package main

type Movie struct {
	name     string //`json:"name"`
	year     int
	director string
}

type Movies []Movie
