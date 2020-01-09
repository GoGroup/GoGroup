package cinema

import "gitlab.com/username/excercise/Project-GO/Movie-and-events/model"

// CinemaService specifies cinema  related service
type CinemaService interface {
	Cinemas() ([]model.Hall, []error)
	// 	Cinema(id uint) (*model.Cinema, []error)

	// 	DeleteCinema(id uint) (*model.Cinema, []error)
	StoreCinema(hall *model.Hall) (*model.Cinema, []error)
}
