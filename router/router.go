package main

import (
	"net/http"

	"github.com/alphagosu/PlayServer/handlers"
)

// TODO: I need to write my own router using a http.mux
// to allow for specifying the method a path can support.
func main() {
	http.HandleFunc("/", handlers.ErrorHandler(handlers.RootHandler, "GET"))
	http.ListenAndServe(":8080", nil)
}
