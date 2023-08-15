package main

import (
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	ctx, err := NewContext()
	if err != nil {
		log.Println(err)
		return
	}
	defer ctx.Close()

	filmSlice, err := ctx.GetAllFilms()
	if err != nil {
		log.Println(err)
		return
	}

	temp := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": *filmSlice,
	}

	if err := temp.Execute(w, films); err != nil {
		log.Println(err)
	}
}

func AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	film := Film{
		Title:    r.PostFormValue("title"),
		Director: r.PostFormValue("director"),
	}

	ctx, err := NewContext()
	if err != nil {
		log.Println(err)
		return
	}
	defer ctx.Close()

	if err := ctx.InsertFilm(film); err != nil {
		log.Println(err)
		return
	}

	temp := template.Must(template.ParseFiles("index.html"))
	temp.ExecuteTemplate(w, "film-list-item", film)
}
