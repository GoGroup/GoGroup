package cinema

import "github.com/Adona12/GoGroup/Movie-and-events/model"

// CinemaService specifies cinema  related service
type CinemaService interface {
	Cinemas() ([]model.Cinema, []error)
	// 	Cinema(id uint) (*model.Cinema, []error)

	// 	DeleteCinema(id uint) (*model.Cinema, []error)
	StoreCinema(cinema *model.Cinema) (*model.Cinema, []error)
}
