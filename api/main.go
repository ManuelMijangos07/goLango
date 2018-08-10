package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/movies_list", moviesList)
	router.HandleFunc("/movie_show/{id}", movieShow)

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
