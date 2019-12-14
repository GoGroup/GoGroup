package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hannasamuel20/Movie-and-events/model"
)

const b = "https://api.themoviedb.org/3/movie/upcoming?"
const apiKey = "f4b8e415cb9ab402e5c1d72176cab35b"
const image = "http://image.tmdb.org/t/p/w500/pjeMs3yqRmFL3giJy4PMXWZTTPa.jpg"

func GetUpcomingMovies() (*model.UpcomingMovies, error, error) {

	res, err := http.Get(b + "api_key=" + apiKey)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var upcoming = new(model.UpcomingMovies)
	err2 := json.Unmarshal([]byte(body), &upcoming)
	if err2 != nil {
		fmt.Println("whoops:", err2)
	}

	return upcoming, err, err2

}

//const path = "/pjeMs3yqRmFL3giJy4PMXWZTTPa.jpg"

// func getUpcomingMovies(body []byte) (*UpcomingMovies, error) {
// 	var upcoming = new(UpcomingMovies)

// 	err := json.Unmarshal(body, &upcoming)
// 	if err != nil {
// 		fmt.Println("whoops:", err)
// 	}

// 	return upcoming, err
// }

// func upcoming_handler(w http.ResponseWriter, r *http.Request) {
// 	res, err := http.Get(b + "api_key=" + apiKey)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	s, err := getUpcomingMovies([]byte(body))
// 	fmt.Println(s)
// 	temp.ExecuteTemplate(w, "display.html", s)

// }
