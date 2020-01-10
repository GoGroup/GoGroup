package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/hall"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
	"github.com/julienschmidt/httprouter"
)

type MenuHandler struct {
	tmpl *template.Template
	csrv cinema.CinemaService
	hsrv hall.HallService
	ssrv schedule.ScheduleService
}

func NewMenuHandler(t *template.Template, cs cinema.CinemaService, hs hall.HallService, ss schedule.ScheduleService) *MenuHandler {

	return &MenuHandler{tmpl: t, csrv: cs, hsrv: hs, ssrv: ss}

}

// func (m *MenuHandler) AdminSchedule(w http.ResponseWriter, r *http.Request) {

// 	s := schedule.Schedules()
// 	fmt.Println(tmpl.ExecuteTemplate(w, "adminScheduleList.layout", s))

// }
func (m *MenuHandler) AdminCinema(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var errr []error
	var NewCinemaArray []model.Cinema

	c, err := m.csrv.Cinemas()
	for _, element := range c {
		element.Halls, errr = m.hsrv.CinemaHalls(element.ID)
		NewCinemaArray = append(NewCinemaArray, element)
	}
	fmt.Println(NewCinemaArray)

	if len(err) > 0 || len(errr) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminCinemaList.layout", NewCinemaArray))

}

func (m *MenuHandler) AdminSchedule(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", nil))

}

func (m *MenuHandler) NewAdminSchedule(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", nil))

}
