package controller

import (
	"net/http"
	"strconv"

	"btcapp/src/entities"

	exporterService "btcapp/src/exporter"
	observerService "btcapp/src/observer"
	storageService "btcapp/src/storage"
)

type Controller struct {
	Storage  storageService.IStorage[entities.User]
	Exporter exporterService.IRateExporter
	Observer observerService.IObserverService
}

func (controller Controller) RateRouter(responseWriter http.ResponseWriter, request *http.Request) {
	btcPrice, err := controller.Exporter.GetCurrentBTCPrice()

	if err == nil {
		stringPrice := strconv.FormatFloat(btcPrice, 'f', -1, 64)
		responseWriter.Write([]byte(stringPrice))
		return
	}

	responseWriter.WriteHeader(http.StatusInternalServerError)
	responseWriter.Write([]byte(""))
}

func (controller Controller) SubscribeRouter(responseWriter http.ResponseWriter, request *http.Request) {
	userGmail := request.URL.Query().Get("gmail")

	if userGmail == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte(""))
		return
	}

	user := entities.User{
		Gmail:           userGmail,
		HasSubscription: true,
	}

	err := controller.Storage.Create(user)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(""))
		return
	}

	responseWriter.Write([]byte("Added"))
}

func (controller Controller) SendEmailsRouter(responseWriter http.ResponseWriter, request *http.Request) {
	go controller.Observer.Notify(
		controller.Exporter,
		controller.Storage,
	)

	responseWriter.Write([]byte("Sended"))
}
