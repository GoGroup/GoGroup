package hall

import "github.com/GoGroup/Movie-and-events/model"

// CommentService specifies cinema hall related service
type HallService interface {
	Halls() ([]model.Hall, []error)
	CinemaHalls(id uint) ([]model.Hall, []error)
	Hall(id uint) (*model.Hall, []error)
	// 	UpdateHall(hall *model.Hall) (*model.Hall, []error)
	// 	DeleteHall(id uint) (*model.Hall, []error)
	StoreHall(hall *model.Hall) (*model.Hall, []error)
}
