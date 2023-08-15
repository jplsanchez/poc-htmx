package main

import (
	handle "htmx/src/service"
	"log"
	"net/http"
)

func main() {
	log.Println("info: Starting...")
	http.HandleFunc("/", handle.IndexHandler)
	http.HandleFunc("/index.html", handle.IndexHandler)
	http.HandleFunc("/add-film", handle.AddMovieHandler)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	log.Println("info: Listening on http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
