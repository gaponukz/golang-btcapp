package controller

import (
	"io"
	"net/http"
	"strconv"

	"../entities"
	exporterService "../exporter"
	observerService "../observer"
	storageService "../storage"
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

	} else {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(""))
	}
}

func (controller Controller) SubscribeRouter(responseWriter http.ResponseWriter, request *http.Request) {
	user := entities.User{
		Gmail:           request.URL.Query().Get("gmail"),
		HasSubscription: true,
	}

	err := controller.Storage.Create(user)

	if err == nil {
		responseWriter.Write([]byte("Added"))

	} else {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(""))
	}
}

func (controller Controller) SendEmailsRouter(responseWriter http.ResponseWriter, request *http.Request) {
	err := controller.Observer.Notify(
		controller.Exporter,
		controller.Storage,
	)

	if err == nil {
		responseWriter.Write([]byte("Sended"))

	} else {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(""))
	}
}

func getUserFromRequestBody(request *http.Request) (entities.User, error) {
	defer request.Body.Close()

	body, err := io.ReadAll(request.Body)

	if err != nil {
		return entities.User{}, err
	}

	return entities.UserFromJSON(string(body)), nil
}
