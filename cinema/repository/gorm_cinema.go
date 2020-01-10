package repository

import (
	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type CinemaGormRepo struct {
	conn *gorm.DB
}

// NewHALLGormRepo returns new object of CommentGormRepo
func NewCinemaGormRepo(db *gorm.DB) cinema.CinemaRepository {
	return &CinemaGormRepo{conn: db}
}
func (cllRepo *CinemaGormRepo) Cinemas() ([]model.Cinema, []error) {
	cll := []model.Cinema{}
	errs := cllRepo.conn.Find(&cll).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cll, errs
}

// Comment retrieves a customer comment from the database by its id
// func (hllRepo *HallGormRepo) Hall(id uint) (*model.Hall, []error) {
// 	hll := model.Hall{}
// 	errs := hllRepo.conn.First(&hll, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return &hll, errs
// }

// // UpdateComment updates a given customer comment in the database
// func (hllRepo *HallGormRepo) UpdateHall(hall *model.Hall) (*model.Hall, []error) {
// 	hll := hall
// 	errs := hllRepo.conn.Save(hll).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return hll, errs
// }

// // DeleteComment deletes a given customer comment from the database
// func (hllRepo *HallGormRepo) DeleteHall(id uint) (*model.Hall, []error) {
// 	hll, errs := hllRepo.Hall(id)

// 	if len(errs) > 0 {
// 		return nil, errs
// 	}

// 	errs = hllRepo.conn.Delete(hll, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return hll, errs
// }

// StoreComment stores a given customer comment in the database
func (cllRepo *CinemaGormRepo) StoreCinema(cinema *model.Cinema) (*model.Cinema, []error) {
	cll := cinema
	errs := cllRepo.conn.Create(cll).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cll, errs
}
