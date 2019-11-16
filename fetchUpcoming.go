package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	MovieList []results `json:"results"`
}
type results struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
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

func main() {
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

}
