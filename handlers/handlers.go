package handlers

import (
	"fmt"
	"net/http"
)

// handler is a simple handler to requests.
type handler func(http.ResponseWriter, *http.Request) error

// RootHandler is the homepage.
func RootHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		http.Error(w, "This path only supports GET requests", http.StatusNotFound)
		return nil
	}
	fmt.Fprintf(w, "Server is running.")
	return nil
}

// ErrorHandler is a wrapper for the handlers that returns and logs errors.
func ErrorHandler(f handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
