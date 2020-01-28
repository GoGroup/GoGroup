package repository

import (
	"errors"

	"github.com/GoGroup/Movie-and-events/hall"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type MockHallRepo struct {
	conn *gorm.DB
}

// NewHALLGormRepo returns new object of CommentGormRepo
func NewMockHallRepo(db *gorm.DB) hall.HallRepository {
	return &MockHallRepo{conn: db}
}
func (hllRepo *MockHallRepo) Halls() ([]model.Hall, []error) {
	hll := []model.Hall{model.HallMock}

	return hll, nil
}

func (hllRepo *MockHallRepo) CinemaHalls(id uint) ([]model.Hall, []error) {
	hlls := []model.Hall{model.HallMock}

	return hlls, nil
}

//Hall retrieves a Hall from the database by its id
func (hllRepo *MockHallRepo) Hall(id uint) (*model.Hall, []error) {
	hll := model.HallMock
	if id == 1 {
		return &hll, nil
	}
	return nil, []error{errors.New("Not found")}
}

// // UpdateHall
func (hllRepo *MockHallRepo) UpdateHall(hall *model.Hall) (*model.Hall, []error) {
	hll := &model.HallMock

	return hll, nil
}

// DeleteHall
func (hllRepo *MockHallRepo) DeleteHall(id uint) (*model.Hall, []error) {
	hll := &model.HallMock

	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return hll, nil
}

// StoreComment stores a given customer comment in the database
func (hllRepo *MockHallRepo) StoreHall(hall *model.Hall) (*model.Hall, []error) {
	hll := &model.HallMock

	return hll, nil
}
func (hllRepo *MockHallRepo) HallExists(hallName string) bool {

	return true
}
