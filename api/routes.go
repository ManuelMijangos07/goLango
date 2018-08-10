package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	name       string
	method     string
	pattern    string
	handleFunc http.HandleFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index},
}
