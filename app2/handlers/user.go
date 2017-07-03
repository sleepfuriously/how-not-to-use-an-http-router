package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sleepfuriously/how-not-to-use-an-http-router/app2/models"
)

func user(res http.ResponseWriter, req *request) {
	idString, ok := req.pathIterator.Next()
	if !ok {
		notFound(res)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid user id %q", id), http.StatusBadRequest)
		return
	}
	if next, ok := req.pathIterator.Next(); ok {
		if next == "profile" {
			handleProfile(res, id)
			return
		} else {
			notFound(res)
			return
		}
	}
	switch req.httpRequest.Method {
	case http.MethodGet:
		userGet(res, id)
	case http.MethodPut:
		userPut(res, id)
	default:
		http.Error(res, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}
}

func userGet(res http.ResponseWriter, id int) {
	if models.KnownUsers[id] {
		fmt.Fprintf(res, "User %d\n", id)
	} else {
		http.Error(res, fmt.Sprintf("User %d does not exist", id), http.StatusBadRequest)
	}
}

func userPut(res http.ResponseWriter, id int) {
	if models.KnownUsers[id] {
		http.Error(res, fmt.Sprintf("User %d already exists", id), http.StatusBadRequest)
	} else {
		models.KnownUsers[id] = true
		fmt.Fprintf(res, "User %d added\n", id)
	}
}
