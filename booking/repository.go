package booking

import "github.com/GoGroup/Movie-and-events/model"

// CommentService specifies  Booking related service
type BookingService interface {
	Bookings(uid uint) ([]model.Booking, []error)

	// Booking(id uint) (*model.Booking, []error)
	// UpdateBooking(Booking *model.Booking) (*model.Booking, []error)
	// DeleteBooking(id uint) (*model.Booking, []error)
	StoreBooking(Booking *model.Booking) (*model.Booking, []error)
}
