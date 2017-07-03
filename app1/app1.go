package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

var knownUsers = map[int]bool{}

type App struct {
	// We could use http.Handler as a type here; using the specific type has
	// the advantage that static analysis tools can link directly from
	// h.UserHandler.ServeHTTP to the correct definition. The disadvantage is
	// that we have slightly stronger coupling. Do the tradeoff yourself.
	UserHandler *UserHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "user" {
		h.UserHandler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}

type UserHandler struct {
	ProfileHandler *ProfileHandler
}

func (h *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
		return
	}
	if req.URL.Path != "/" {
		head, _ := ShiftPath(req.URL.Path)
		switch head {
		case "profile":
			// We can't just make ProfileHandler an http.Handler; it needs the
			// user id. Let's instead…
			h.ProfileHandler.Handler(id).ServeHTTP(res, req)
		default:
			http.Error(res, "Not Found", http.StatusNotFound)
		}
		return
	}
	switch req.Method {
	case "GET":
		h.handleGet(res, id)
	case "PUT":
		h.handlePut(res, id)
	default:
		http.Error(res, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) handleGet(res http.ResponseWriter, id int) {
	if knownUsers[id] {
		fmt.Fprintf(res, "User %d\n", id)
	} else {
		http.Error(res, fmt.Sprintf("User %d not found", id), http.StatusNotFound)
	}
}

func (h *UserHandler) handlePut(res http.ResponseWriter, id int) {
	if knownUsers[id] {
		http.Error(res, fmt.Sprintf("User %d already exists", id), http.StatusBadRequest)
	} else {
		knownUsers[id] = true
		fmt.Fprintf(res, "User %d added\n", id)
	}
}

type ProfileHandler struct {
}

func (h *ProfileHandler) Handler(id int) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if knownUsers[id] {
			fmt.Fprintf(res, "Profile for user %d\n", id)
		} else {
			http.Error(res, fmt.Sprintf("User %d not found", id), http.StatusNotFound)
		}
	})
}

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func start(shutdown chan struct{}) {
	srv := &http.Server{
		Addr: ":8000",
		Handler: &App{
			UserHandler: new(UserHandler),
		},
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
