package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Casablanca", Director: "Michael Curtiz"},
				{Title: "Cool Hand Luke", Director: "Stuart Rosenberg"},
				{Title: "Bullitt", Director: "Peter Yates"},
			},
		}
		err := temp.Execute(w, films)
		if err != nil {
			log.Println(err)
		}
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(2 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// temp, _ := template.New("add-film").Parse(htmlStr)
		// temp.Execute(w, nil)

		temp := template.Must(template.ParseFiles("index.html"))
		temp.ExecuteTemplate(w, "film-list-item", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/index.html", h1)
	http.HandleFunc("/add-film/", h2)
	log.Fatal(http.ListenAndServe(":80", nil))
}
