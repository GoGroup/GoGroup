package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Adona12/GoGroup/Movie-and-events/hall"
	"github.com/Adona12/GoGroup/Movie-and-events/model"
	"github.com/julienschmidt/httprouter"
)

type HallHandler struct {
	hallService hall.HallService
}

func NewHallHandler(HllService hall.HallService) *HallHandler {
	return &HallHandler{hallService: HllService}

}
func (hh *HallHandler) GetHalls(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	Halls, errs := hh.hallService.Halls()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Halls, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (ach *HallHandler) PostHall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	hall := &model.Hall{}

	err := json.Unmarshal(body, hall)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	hall, errs := ach.hallService.StoreHall(hall)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/halls/%d", hall.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
