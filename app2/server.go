package main

import (
	"net/http"

	"github.com/sleepfuriously/how-not-to-use-an-http-router/app2/handlers"
)

func start() {
	http.ListenAndServe(":8000", handlers.Main)
}

func main() {
	start()
}
