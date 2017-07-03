package main

import (
	"net/http"

	"github.com/sleepfuriously/how-not-to-use-an-http-router/app2/handlers"
)
const addr = ":8001"

func serve() {
	http.ListenAndServe(addr, handlers.Main)
}

func main() {
	serve()
}
