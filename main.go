package main

import (
	"net/http"

	"./src/controller"
	"./src/exporter"
	"./src/observer"
	"./src/storage"
)

func main() {
	routerService := controller.Controller{
		Storage:  &storage.UserMemoryStorage{},
		Exporter: &exporter.MemoryExporter{},
		Observer: &observer.ObserverService{
			Strategy: observer.ConsoleLogStrategy,
		},
	}

	httpRoute := http.NewServeMux()

	httpRoute.HandleFunc("/rate", routerService.RateRouter)
	httpRoute.HandleFunc("/subscribe", routerService.SubscribeRouter)
	httpRoute.HandleFunc("/sendEmails", routerService.SendEmailsRouter)

	server := http.Server{
		Addr:    ":8080",
		Handler: httpRoute,
	}

	server.ListenAndServe()
}
