package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOLA MUNDO DESDE GO")
}
func moviesList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Sin límetes", 2013, "Desconocido"},
		Movie{"Batman", 1999, "Scorsese"},
		Movie{"A todo gas", 2015, "Juan Antonio"}}
	//fmt.Fprintf(w, "listado de películas")
	json.NewEncoder(w).Encode(movies)
}
func movieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w, "Detalle de pelicula número %s", movieId)
}
