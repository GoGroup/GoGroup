package movie

import "github.com/GoGroup/Movie-and-events/model"

type MovieService interface {
	Movies() ([]model.Moviem, []error)
	StoreMovie(movie *model.Moviem) (*model.Moviem, []error)
}
