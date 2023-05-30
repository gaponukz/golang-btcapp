package main

import (
	"net/http"

	"btcapp/src/controller"
	"btcapp/src/exporter"
	"btcapp/src/observer"
	"btcapp/src/storage"
)

func main() {
	routerService := controller.Controller{
		Exporter: &exporter.CoingeckoExporter{},
		Storage: &storage.JsonFileUserStorage{
			Filename: "users.json",
		},
		Observer: &observer.ObserverService{
			Strategy: observer.ConsoleLogStrategy,
		},
	}

	httpRoute := http.NewServeMux()

	httpRoute.HandleFunc("/rate", RequiredMethod(routerService.RateRouter, http.MethodGet))
	httpRoute.HandleFunc("/subscribe", RequiredMethod(routerService.SubscribeRouter, http.MethodPost))
	httpRoute.HandleFunc("/sendEmails", RequiredMethod(routerService.SendEmailsRouter, http.MethodPost))

	loggedRouter := LoggingMiddleware(httpRoute)

	server := http.Server{
		Addr:    ":8080",
		Handler: loggedRouter,
	}

	server.ListenAndServe()
}
