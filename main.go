package main

import (
	"log"
	"net/http"

	"btcapp/src/controller"
	"btcapp/src/exporter"
	"btcapp/src/observer"
	"btcapp/src/storage"
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

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s?%s", request.Method, request.URL.Path, request.URL.RawQuery)
		next.ServeHTTP(responseWriter, request)
	})
}

func main() {
	routerService := controller.Controller{
		Exporter: &exporter.MemoryExporter{},
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
