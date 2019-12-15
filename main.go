package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("view/template/*"))

func index(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func display(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Movie.layout", nil)
}
func main() {
	fs := http.FileServer(http.Dir("view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/Movie", display)
	http.ListenAndServe(":8080", nil)
}
