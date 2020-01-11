package service

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/movie"
)

type MovieService struct {
	movieRepo movie.MovieRepository
}

func NewMovieService(mvRepo movie.MovieRepository) movie.MovieService {
	return &MovieService{movieRepo: mvRepo}
}

func (m *MovieService) Movies() ([]model.Moviem, []error) {
	mvs, errs := m.movieRepo.Movies()
	if len(errs) > 0 {
		return nil, errs
	}
	return mvs, errs
}

func (m *MovieService) StoreMovie(movie *model.Moviem) (*model.Moviem, []error) {
	mvs, errs := m.movieRepo.StoreMovie(movie)
	if len(errs) > 0 {
		return nil, errs
	}
	return mvs, errs
}
