package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GoGroup/Movie-and-events/booking"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/julienschmidt/httprouter"
)

type BookingHandler struct {
	bookingService booking.BookingService
}

func NewBookingHandler(BkkService booking.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: BkkService}

}

//Gets all halls
func (bkk *BookingHandler) GetBookings(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	Bookings, errs := bkk.bookingService.Bookings(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Bookings, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//adds a hall
func (ach *BookingHandler) PostBooking(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	booking := &model.Booking{}

	err := json.Unmarshal(body, booking)

	if err != nil {
fmt.Println("first")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	booking, errs := ach.bookingService.StoreBooking(booking)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/bookings/%d", booking.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
