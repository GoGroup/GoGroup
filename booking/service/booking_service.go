package service

import (
	"github.com/GoGroup/Movie-and-events/booking"
	"github.com/GoGroup/Movie-and-events/model"
)

type BookingService struct {
	bookingRepo booking.BookingRepository
}

func NewBookingService(BookingRepos booking.BookingRepository) booking.BookingService {
	return &BookingService{bookingRepo: BookingRepos}
}

// bookings returns all stored comments
func (bk *BookingService) Bookings(uid uint) ([]model.Booking, []error) {
	bkk, errs := bk.bookingRepo.Bookings(uid)
	if len(errs) > 0 {
		return nil, errs
	}
	return bkk, errs
}

func (bk *BookingService) StoreBooking(booking *model.Booking) (*model.Booking, []error) {
	boks, errs := bk.bookingRepo.StoreBooking(booking)
	if len(errs) > 0 {
		return nil, errs
	}
	return boks, errs
}
