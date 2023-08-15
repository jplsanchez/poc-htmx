package main

import (
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	log.Println("Starting...")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/index.html", IndexHandler)
	http.HandleFunc("/add-film", AddMovieHandler)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	log.Println("Listening on http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
