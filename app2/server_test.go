package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	shutdown := make(chan struct{})
	defer close(shutdown)
	go start(shutdown)

	check(t, http.MethodGet, "", http.StatusNotFound)
	check(t, http.MethodGet, "/", http.StatusNotFound)
	check(t, http.MethodGet, "//", http.StatusNotFound)
	check(t, http.MethodGet, "/user", http.StatusNotFound)
	check(t, http.MethodGet, "/user/a", http.StatusBadRequest)
	check(t, http.MethodGet, "/user/1", http.StatusBadRequest)
	check(t, http.MethodGet, "/user/1/profile", http.StatusBadRequest)
	check(t, http.MethodGet, "/user/1/abc", http.StatusNotFound)

	check(t, http.MethodPut, "", http.StatusNotFound)
	check(t, http.MethodPut, "/user/a", http.StatusBadRequest)
	check(t, http.MethodPut, "/user/1", http.StatusOK)

	check(t, http.MethodGet, "/user/1", http.StatusOK)
	check(t, http.MethodGet, "/user/1/profile", http.StatusOK)

	check(t, http.MethodPost, "/user/1", http.StatusMethodNotAllowed)
}

func check(t *testing.T, method, path string, expectedStatusCode int) {
	req, err := http.NewRequest(method, "http://localhost:8000"+path, nil)
	if err != nil {
		t.Fatal(method, path, err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(method, path, err)
	}
	if res.StatusCode != expectedStatusCode {
		t.Fatal(method, path, fmt.Sprintf("got: %d, expect: %d", res.StatusCode, expectedStatusCode))
	}
}