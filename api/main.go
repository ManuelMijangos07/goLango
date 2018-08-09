package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HOLA MUNDO DESDE GO")
	})

	server := http.ListenAndServe(":8080", nil)

	log.Fatal(server)
	fmt.Println("El servidor está corriendo en http://localhost:8080")
}
