package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	h "github.com/alphagosu/PlayServer/handlers"
	"github.com/gorilla/mux"
)

const (
	// GET is the method get in a http or https request
	GET = "GET"
)

// sets up a router for incoming requests. blocks on listen and serve.
func main() {
	r := mux.NewRouter()
	get(r, "/", h.RootHandler)

	listenAndServe(":8080", r, newLogger())
}

// get specifies that only the GET method can work for that route.
func get(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc("/", h.ErrorHandler(h.RootHandler)).Methods(GET)
}

// newLogger returns a custom logger to log events.
// TODO: find out if this log header is redundant with the flags supplied.
func newLogger() *log.Logger {
	logHeader := fmt.Sprintf(time.Now().Format("Jan 2 15:04:05 EST"))
	return log.New(os.Stdout, logHeader, log.LstdFlags)
}

// listenAndServe creates a new SERVER and uses it to listen for incoming
// requests.
// TODO: Find out how to log ALL requests that hit the server.
// TODO: Find out how to throttle requests to 1 request a second. (brute forcing prevention)
// TODO: Find out how to block ip ranges for incoming requests. (DDoS prevention)
func listenAndServe(addr string, handler http.Handler, logger *log.Logger) error {
	server := &http.Server{Addr: addr, Handler: handler, ErrorLog: logger}
	return server.ListenAndServe()
}
