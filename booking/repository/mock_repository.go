package repository

import (
	"errors"

	"github.com/GoGroup/Movie-and-events/booking"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type MockBookingepo struct {
	conn *gorm.DB
}

// NewHALLGormRepo returns new object of CommentGormRepo
func NewMockBookingepo(db *gorm.DB) booking.BookingRepository {
	return &BookingGormRepo{conn: db}
}
func (bkkRepo *MockBookingepo) Bookings(uid uint) ([]model.Booking, []error) {

	ctgs := []model.Booking{model.BookingMock}
	if uid == 1 {
		return ctgs, nil
	}
	return nil, []error{errors.New("Not found")}
}

//Hall retrieves a Hall from the database by its id
// func (hllRepo *HallGormRepo) Hall(id uint) (*model.Hall, []error) {
// 	hll := model.Hall{}
// 	errs := hllRepo.conn.First(&hll, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return &hll, errs
// }

// // UpdateHall
// func (hllRepo *HallGormRepo) UpdateHall(hall *model.Hall) (*model.Hall, []error) {
// 	hll := hall
// 	errs := hllRepo.conn.Save(hll).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return hll, errs
// }

// // DeleteHall
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
func (bkkRepo *MockBookingepo) StoreBooking(booking *model.Booking) (*model.Booking, []error) {
	eve := &model.BookingMock

	return eve, []error{}
}
