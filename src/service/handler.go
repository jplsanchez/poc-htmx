package service

import (
	"html/template"
	"htmx/src/data"
	"htmx/src/model"
	"log"
	"net/http"
)

type Handler struct {
	Repository data.Repository
}

func (h Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	filmSlice, err := h.Repository.GetAllFilms()
	if err != nil {
		log.Println(err)
		return
	}

	temp := template.Must(template.ParseFiles("index.html"))
	films := map[string][]model.Film{
		"Films": *filmSlice,
	}

	if err := temp.Execute(w, films); err != nil {
		log.Println(err)
	}
}

func (h Handler) AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	film := model.Film{
		Title:    r.PostFormValue("title"),
		Director: r.PostFormValue("director"),
	}
	if err := h.Repository.InsertFilm(film); err != nil {
		log.Println(err)
		return
	}

	temp := template.Must(template.ParseFiles("index.html"))
	temp.ExecuteTemplate(w, "film-list-item", film)
}
