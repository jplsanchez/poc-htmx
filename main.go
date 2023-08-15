package main

import (
	"fmt"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/index.html", IndexHandler)
	http.HandleFunc("/add-film/", AddMovieHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
