package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/movie"
	"github.com/julienschmidt/httprouter"
)

type MovieHandler struct {
	movieService movie.MovieService
}

func NewMovieHander(mvService movie.MovieService) *MovieHandler {
	return &MovieHandler{movieService: mvService}
}

func (mh *MovieHandler) GetMovies(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	movies, errs := mh.movieService.Movies()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(movies, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (mh *MovieHandler) PostMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	movie := &model.Moviem{}

	err := json.Unmarshal(body, movie)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	movie, errs := mh.movieService.StoreMovie(movie)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//p := fmt.Sprintf("/v1/admin/comments/%d", comment.ID)
	//w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
