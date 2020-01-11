package repository

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/movie"
	"github.com/jinzhu/gorm"
)

type MovieGormRepo struct {
	conn *gorm.DB
}

func NewMovieGormRepo(db *gorm.DB) movie.MovieRepository {
	return &MovieGormRepo{conn: db}

}

func (movieRepo *MovieGormRepo) Movies() ([]model.Moviem, []error) {
	mvs := []model.Moviem{}
	errs := movieRepo.conn.Find(&mvs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mvs, errs

}

func (movieRepo *MovieGormRepo) StoreMovie(movie *model.Moviem) (*model.Moviem, []error) {
	mv := movie
	errs := movieRepo.conn.Create(mv).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mv, errs
}
