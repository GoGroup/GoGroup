package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hannasamuel20/Movie-and-events/model"
)

const base = "https://api.themoviedb.org/3/movie/"
const upcomingQuery = "upcoming?"
const apiKey = "f4b8e415cb9ab402e5c1d72176cab35b"

//const videoBase = "http://api.themoviedb.org/3/movie/157336/videos?api_key=###"
const videoQuery = "/videos?"

func GetMovieDetails(id int) (*model.MovieDetails, error, error) {

	res, err := http.Get(base + strconv.Itoa(id) + "?api_key=" + apiKey)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var details = new(model.MovieDetails)
	err2 := json.Unmarshal([]byte(body), &details)
	if err2 != nil {
		fmt.Println("whoops:", err2)
	}

	return details, err, err2

}

func GetUpcomingMovies() (*model.UpcomingMovies, error, error) {

	res, err := http.Get(base + upcomingQuery + "api_key=" + apiKey)
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

func GetTrailer(i string) string {
	res, err := http.Get(base + i + videoQuery + "api_key=" + apiKey)

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var vid = new(model.VideoLists)
	err2 := json.Unmarshal([]byte(body), &vid)
	//m.Trailer=vid.VList[0].Key

	if err2 != nil {
		fmt.Println("whoops:", err2)
	}

	return vid.VList[0].Key

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
