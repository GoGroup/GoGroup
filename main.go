package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/GoGroup/Movie-and-events/controller"
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
func displayTheater(w http.ResponseWriter, r *http.Request) {

	fmt.Println(tmpl.ExecuteTemplate(w, "scheduleDisplay.layout", nil))

}
func admin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in 8080")
	//fmt.Println(controller.GetSchedules())

	fmt.Println(tmpl.ExecuteTemplate(w, "check.layout", nil))

}

func halls(w http.ResponseWriter, r *http.Request) {

	fmt.Println(tmpl.ExecuteTemplate(w, "halls.layout", nil))
}
func adminCinemas(w http.ResponseWriter, r *http.Request) {

	fmt.Println(tmpl.ExecuteTemplate(w, "adminCinemaList.layout", nil))

}
func adminSchedule(w http.ResponseWriter, r *http.Request) {

	fmt.Println(tmpl.ExecuteTemplate(w, "adminScheduleList.layout", nil))

}
func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println(tmpl.ExecuteTemplate(w, "login.layout", nil))

}

func eachmovieHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Println(id)
	trailerKey := controller.GetTrailer(id)
	i, _ := strconv.Atoi(id)
	details, _, _ := controller.GetMovieDetails(i)
	details.Trailer = trailerKey

	fmt.Println(tmpl.ExecuteTemplate(w, "EachMovie.layout", details))

}
func main() {

	fs := http.FileServer(http.Dir("view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/",login)
	http.HandleFunc("/Movie", display)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/adminCinemas", adminCinemas)
	http.HandleFunc("/adminCinemas/adminSchedule", adminSchedule)
	http.HandleFunc("/theater", displayTheater)
	http.HandleFunc("/hall", halls)
	http.HandleFunc("/eachmovie/", eachmovieHandler)
	http.ListenAndServe(":8080", nil)

}
