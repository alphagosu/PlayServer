package handlers

import (
	"fmt"
	"net/http"
)

// Handler is a simple handler to requests.
type Handler func(http.ResponseWriter, *http.Request) error

//this is a list of strings
type listOfStrings []string

// RootHandler is the homepage.
func RootHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1>Server is running.</h1>")
	return nil
}

// ErrorHandler is a wrapper for the handlers that returns and logs errors.
func ErrorHandler(f Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
