package main

import (
	"log"
	"net/http"

	"github.com/sleepfuriously/how-not-to-use-an-http-router/app2/handlers"
)

func start(shutdown chan struct{}) {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: handlers.Main,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	
	<-shutdown
	srv.Shutdown(nil)
}

func main() {
	start(nil)
}
