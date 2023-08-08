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

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./served"))))
	http.HandleFunc("/", idx)
	http.HandleFunc("/add-film/", addFilm)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func addFilm(w http.ResponseWriter, r *http.Request) {
	log.Print("Request recieved!")
	log.Print(r.Header.Get("HX-Request"))
	log.Print(r.PostForm)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	log.Print(director)
	htmlStr := fmt.Sprintf("<p class='movie-item'> %s - %s</p>", director, title)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}

func idx(w http.ResponseWriter, r *http.Request) {
	films := map[string][]Film{
		"Films": {
			{Title: "TheGodFather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", films)
}
