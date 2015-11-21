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

func main() {
	r := mux.NewRouter()
	get(r, "/", h.RootHandler)

	listenAndServe(":8080", r, newLogger())
}

func get(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc("/", h.ErrorHandler(h.RootHandler)).Methods(GET)
}

func newLogger() *log.Logger {
	logHeader := fmt.Sprintf(time.Now().Format("Jan 2 15:04:05 EST"))
	return log.New(os.Stdout, logHeader, log.LstdFlags)
}

func listenAndServe(addr string, handler http.Handler, logger *log.Logger) error {
	server := &http.Server{Addr: addr, Handler: handler, ErrorLog: logger}
	return server.ListenAndServe()
}
