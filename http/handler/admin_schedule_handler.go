package handler

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
	"github.com/julienschmidt/httprouter"
)

type AdminScheduleHandler struct {
	scheduleService schedule.ScheduleService
}

func NewAdminScheduleHandler(schdlService schedule.ScheduleService) *AdminScheduleHandler {
	fmt.Println("admin schedule handler created")
	return &AdminScheduleHandler{scheduleService: schdlService}
}

func (as *AdminScheduleHandler) GetSchedules(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	schedules, errs := as.scheduleService.Schedules()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(schedules, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (as *AdminScheduleHandler) PostSchedule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("in post schedule 1")

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	schedule := &model.Schedule{}
	fmt.Println("in post schedule 2")

	err := json.Unmarshal(body, schedule)
	fmt.Println(schedule)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	schedule, errs := as.scheduleService.StoreSchedule(schedule)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//p := fmt.Sprintf("/v1/admin/comments/%d", comment.ID)
	//w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
