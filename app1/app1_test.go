package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestApp(t *testing.T) {
	shutdown := make(chan struct{})
	defer close(shutdown)
	go start(shutdown)

	check(t, http.MethodGet, "", http.StatusNotFound)
	check(t, http.MethodGet, "/", http.StatusNotFound)
	check(t, http.MethodGet, "//", http.StatusNotFound)
	check(t, http.MethodGet, "/user", http.StatusBadRequest)
	check(t, http.MethodGet, "/user/a", http.StatusBadRequest)
	check(t, http.MethodGet, "/user/1", http.StatusNotFound)
	check(t, http.MethodGet, "/user/1/profile", http.StatusNotFound)

	check(t, http.MethodPut, "", http.StatusNotFound)
	check(t, http.MethodPut, "/user/a", http.StatusBadRequest)
	check(t, http.MethodPut, "/user/1", http.StatusOK)

	check(t, http.MethodGet, "/user/1", http.StatusOK)
	check(t, http.MethodGet, "/user/1/profile", http.StatusOK)
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
		t.Fatal(method, path, fmt.Sprintf("%d != %d", res.StatusCode, expectedStatusCode))
	}
}
