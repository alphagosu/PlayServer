package main

import (
	"net/http"

	"github.com/alphagosu/PlayServer/handlers"
)

// TODO: I need to write my own router using a http.mux
// to allow for specifying the method a path can support.
func main() {
	m := http.NewServeMux()
	r := router{
		mux:    m,
		routes: []route{},
	}
	r.routes = append(r.routes, route{
		path: "/",
		h:    handlers.ErrorHandler(handlers.RootHandler),
	})

	r.registerRoutes()

	//validRootMethods := []string{"GET", "POST"}
	//http.HandleFunc("/", handlers.ErrorHandler(handlers.RootHandler, validRootMethods))

	http.ListenAndServe(":8080", r.mux)
}

//paths, routes, handler fucntions, router, mux

// router
// 	a route is composed of path, handler, methods(get, post.. etc)
// 	mux
//

type router struct {
	mux    *http.ServeMux
	routes []route
}

type route struct {
	path string
	h    func(http.ResponseWriter, *http.Request)
}

func (rtr router) registerRoutes() {
	for _, r := range rtr.routes {
		rtr.mux.HandleFunc(r.path, r.h)
	}
}
