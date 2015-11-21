package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Handler is a simple handler to requests.
type Handler func(http.ResponseWriter, *http.Request) error

// RootHandler is the homepage.
func RootHandler(w http.ResponseWriter, r *http.Request) error {
	ServeContent(w, r, "<h1>Server is running.</h1>", http.StatusOK)
	return nil
}

// ErrorHandler is a wrapper for the handlers that returns and logs errors.
func ErrorHandler(f Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			ServeContent(w, r, err.Error(), http.StatusInternalServerError)
		}
	}
}

// ServeContent is a wrapper for responding to clients. It can log the
func ServeContent(w http.ResponseWriter, r *http.Request, body string, statusCode int) {
	logHeader := fmt.Sprintf(time.Now().Format("Jan 2 15:04:05 EST"))
	path := fmt.Sprintf(", request on path '%s' ", r.URL.Path)
	status := fmt.Sprintf("returned status code: %d", statusCode)

	log.Println(logHeader + path + status)
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, body)
}
