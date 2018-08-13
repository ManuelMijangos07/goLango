package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	//"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Sin límetes", 2013, "Desconocido"},
	Movie{"Batman", 1999, "Scorsese"},
	Movie{"A todo gas", 2015, "Juan Antonio"}}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if (err != nil {
		panic(err)
	},

	return session
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOLA MUNDO DESDE GO")
}
func moviesList(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "listado de películas")
	json.NewEncoder(w).Encode(movies)
}
func movieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w, "Detalle de pelicula número %s", movieId)
}

func movieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)
}
