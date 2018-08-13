package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	name       string
	method     string
	pattern    string
	handleFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.handleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"moviesList",
		"GET",
		"/movies_list",
		moviesList,
	},
	Route{
		"movieShow",
		"GET",
		"/movie_show/{id}",
		movieShow,
	},
	Route{
		"movieAdd",
		"POST",
		"/new_movie",
		movieAdd,
	},
}
