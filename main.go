package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/hannasamuel20/Movie-and-events/controller"
)

var tmpl = template.Must(template.ParseGlob("view/template/*"))



func index(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func display(w http.ResponseWriter, r *http.Request) {

	upcomingmovies, _, _ := controller.GetUpcomingMovies()
	fmt.Println(upcomingmovies)
	fmt.Println(tmpl.ExecuteTemplate(w, "Movie.layout", upcomingmovies))


}
func main() {
	fs := http.FileServer(http.Dir("view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/Movie", display)
	http.ListenAndServe(":8080", nil)
}
