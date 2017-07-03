package handlers

import (
	"fmt"
	"net/http"

	"github.com/sleepfuriously/how-not-to-use-an-http-router/app2/models"
)

func handleProfile(res http.ResponseWriter, id int) {
	if models.KnownUsers[id] {
		fmt.Fprintf(res, "Profile for user %d\n", id)
	} else {
		http.Error(res, fmt.Sprintf("User %d does not exist", id), http.StatusBadRequest)
	}
}
