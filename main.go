package main

import (
	"net/http"

	"./src/controller"
	"./src/exporter"
	"./src/observer"
	"./src/storage"
)

type RouterFunc = func(rw http.ResponseWriter, r *http.Request)

func RequiredMethod(router RouterFunc, required string) RouterFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == required {
			router(responseWriter, request)

		} else {
			http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	routerService := controller.Controller{
		Storage:  &storage.UserMemoryStorage{},
		Exporter: &exporter.MemoryExporter{},
		Observer: &observer.ObserverService{
			Strategy: observer.ConsoleLogStrategy,
		},
	}

	httpRoute := http.NewServeMux()

	httpRoute.HandleFunc("/rate", RequiredMethod(routerService.RateRouter, http.MethodGet))
	httpRoute.HandleFunc("/subscribe", RequiredMethod(routerService.SubscribeRouter, http.MethodPost))
	httpRoute.HandleFunc("/sendEmails", RequiredMethod(routerService.SendEmailsRouter, http.MethodPost))

	server := http.Server{
		Addr:    ":8080",
		Handler: httpRoute,
	}

	server.ListenAndServe()
}
