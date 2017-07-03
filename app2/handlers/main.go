package handlers

import (
	"net/http"

	"github.com/sleepfuriously/how-not-to-use-an-http-router/app2/path"
)

type request struct {
	httpRequest  *http.Request
	pathIterator *path.Iterator
}

var Main http.HandlerFunc = main

func main(res http.ResponseWriter, req *http.Request) {
	request := &request{
		httpRequest:  req,
		pathIterator: path.NewIterator(req.URL.Path),
	}
	segment, _ := request.pathIterator.Next()
	switch {
	case segment == "user":
		user(res, request)
	default:
		notFound(res)
	}
}

func notFound(res http.ResponseWriter) {
	http.NotFound(res, nil)
}
