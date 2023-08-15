package main

import (
	"htmx/src/data"
	"htmx/src/service"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	h := resolveDependencyInjection()
	defer h.Repository.Close()

	log.Println("info: Starting...")
	http.HandleFunc("/", h.IndexHandler)
	http.HandleFunc("/index.html", h.IndexHandler)
	http.HandleFunc("/add-film/", h.AddMovieHandler)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	log.Println("info: Listening on http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func resolveDependencyInjection() *service.Handler {
	godotenv.Load()
	doNotUseDatabase := os.Getenv("DO_NOT_USE_DATABASE")
	if strings.ToLower(doNotUseDatabase) == "true" {
		ctx, err := data.NewInMemoryContext()
		if err != nil {
			log.Fatal(err)
		}
		return &service.Handler{Repository: ctx}

	}

	ctx, err := data.NewMySqlContext()
	if err != nil {
		log.Fatal(err)
	}
	return &service.Handler{Repository: ctx}
}
