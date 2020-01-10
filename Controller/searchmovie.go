package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GoGroup/Movie-and-events/model"
)

const a = "https://api.themoviedb.org/3/search/movie?api_key=f4b8e415cb9ab402e5c1d72176cab35b&language=en-US&page=1&include_adult=false&query="

func SearchMovie(query string) (*model.UpcomingMovies, error, error) {

	res, err := http.Get(a + query)
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
