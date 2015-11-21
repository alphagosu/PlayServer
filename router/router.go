package main

import (
	"net/http"

	h "github.com/alphagosu/PlayServer/handlers"
	"github.com/gorilla/mux"
)

// type muxRouter func(r *mux.Router) func(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route

const (
	// GET is the method get in a http or https request
	GET = "GET"
)

func main() {
	r := mux.NewRouter()
	get(r, "/", h.RootHandler)
	http.ListenAndServe(":8080", r)
}

func get(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc("/", h.ErrorHandler(h.RootHandler)).Methods(GET)
}
