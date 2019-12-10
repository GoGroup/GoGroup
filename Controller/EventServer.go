package main

import (
	"html/template"
	"net/http"
)

var templ = template.Must(template.ParseFiles("slider.html"))

func SliderHandler(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", SliderHandler)
	http.ListenAndServe(":8080", mux)
}
