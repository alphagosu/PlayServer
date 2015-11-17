package handlers

import (
	"fmt"
	"net/http"
)

// handler is a simple handler to requests.
type handler func(http.ResponseWriter, *http.Request) error

//this is a list of strings
type listOfStrings []string

// RootHandler is the homepage.
func RootHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1>Server is running.</h1>")
	return nil
}

// ErrorHandler is a wrapper for the handlers that returns and logs errors.
func ErrorHandler(f handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if !methods.contains(r.Method) {
		// 	err := fmt.Sprintf("This path only supports the following requests: %s", strings.Join(methods, ", "))
		// 	http.Error(w, err, http.StatusNotFound)
		// } else {
		if err := f(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// checks to see if a value is contained in an array.
func (vals listOfStrings) contains(s string) bool {
	for _, val := range vals {
		if s == val {
			return true
		}
	}
	return false
}
