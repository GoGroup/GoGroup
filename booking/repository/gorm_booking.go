package repository

import (
	"github.com/GoGroup/Movie-and-events/booking"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type BookingGormRepo struct {
	conn *gorm.DB
}

// NewHALLGormRepo returns new object of CommentGormRepo
func NewBookingGormRepo(db *gorm.DB) booking.BookingRepository {
	return &BookingGormRepo{conn: db}
}
func (bkkRepo *BookingGormRepo) Bookings(uid uint) ([]model.Booking, []error) {
	bkk := []model.Booking{}
	errs := bkkRepo.conn.Where("user_id = ?", uid).Find(&bkk).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return bkk, errs
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
func (bkkRepo *BookingGormRepo) StoreBooking(booking *model.Booking) (*model.Booking, []error) {
	bkk := booking
	errs := bkkRepo.conn.Create(bkk).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return bkk, errs
}
