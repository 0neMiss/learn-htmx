package main

import (
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
	log.Fatal(http.ListenAndServe(":8000", nil))
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
