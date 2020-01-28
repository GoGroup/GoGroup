package repository

import (
	"errors"

	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type MockCinemaRepo struct {
	conn *gorm.DB
}

// NewHALLGormRepo returns new object of CommentGormRepo
func NewMockCinemaRepo(db *gorm.DB) cinema.CinemaRepository {
	return &MockCinemaRepo{conn: db}
}
func (cllRepo *MockCinemaRepo) Cinemas() ([]model.Cinema, []error) {
	cll := []model.Cinema{model.CinemaMock}
	return cll, nil
}

//Cinema retrieves a cinema from the database by its id
func (cllRepo *MockCinemaRepo) Cinema(id uint) (*model.Cinema, []error) {
	cll := &model.CinemaMock

	if id == 1 {
		return cll, nil
	}
	return nil, []error{errors.New("Not found")}
}

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
func (cllRepo *MockCinemaRepo) StoreCinema(cinema *model.Cinema) (*model.Cinema, []error) {
	cll := &model.CinemaMock

	return cll, nil
}
