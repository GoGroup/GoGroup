package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	"net/http"
)

type response struct {
	MovieList []results `json:"results"`
}
type results struct {
	Title       string `json:"title"`
	PosterPath  string `json:"poster_path"`
	Overview    string `json:"overview"`
	Id          int    `json:"id"`
	ReleaseDate string `json:"release_date"`
}

func getUpcomingMovies(body []byte) (*response, error) {
	var s = new(response)

	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	return s, err
}

const b = "https://api.themoviedb.org/3/movie/upcoming?"
const apiKey = "f4b8e415cb9ab402e5c1d72176cab35b"
const image = "http://image.tmdb.org/t/p/w500/pjeMs3yqRmFL3giJy4PMXWZTTPa.jpg"
const path = "/pjeMs3yqRmFL3giJy4PMXWZTTPa.jpg"

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("*.html"))
}

func upcoming_handler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(b + "api_key=" + apiKey)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	s, err := getUpcomingMovies([]byte(body))
	fmt.Println(s)
	temp.ExecuteTemplate(w, "display.html", s)

}

func main() {
	http.HandleFunc("/", upcoming_handler)
	http.ListenAndServe(":8080", nil)

}
