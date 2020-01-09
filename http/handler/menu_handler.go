package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/julienschmidt/httprouter"
)

type MenuHandler struct {
	tmpl *template.Template
	csrv cinema.CinemaService
}

func NewMenuHandler(t *template.Template, cs cinema.CinemaService) *MenuHandler {

	return &MenuHandler{tmpl: t, csrv: cs}

}

// func (m *MenuHandler) AdminSchedule(w http.ResponseWriter, r *http.Request) {

// 	s := schedule.Schedules()
// 	fmt.Println(tmpl.ExecuteTemplate(w, "adminScheduleList.layout", s))

// }
func (m *MenuHandler) AdminCinema(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	c, err := m.csrv.Cinemas()
	fmt.Println(c)

	if len(err) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminCinemaList.layout", c))

}
