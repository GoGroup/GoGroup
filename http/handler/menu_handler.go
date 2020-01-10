package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/GoGroup/Movie-and-events/controller"

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
	fmt.Println("In admin schedule*****************")
	fmt.Println(pm.ByName("hId"))
	HallID := pm.ByName("hId")

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", HallID))

}

func (m *MenuHandler) NewAdminSchedule(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	var MovieTitles *model.UpcomingMovies

	var err error
	var err2 error
	if r.FormValue("movie") != "" {
		Movie := r.FormValue("movie")
		fmt.Println(Movie)
		MovieTitles, err2, err = controller.SearchMovie(Movie)
		if err == nil || err2 == nil {
			fmt.Println("did seatch")
			fmt.Println(MovieTitles)
		}

	} else {
		fmt.Println("empty")
		movie := r.FormValue("movie")
		fmt.Println(movie)

	}

	convid, _ := strconv.Atoi(r.FormValue("id"))
	hallid, _ := strconv.Atoi(pm.ByName("hId"))

	tempo := struct {
		M       *model.UpcomingMovies
		MovieN  string
		MovieID int
		HallID  int
	}{
		M:       MovieTitles,
		MovieN:  r.FormValue("moviename"),
		MovieID: convid,
		HallID:  hallid,
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", tempo))

}
func (m *MenuHandler) NewAdminSchedulePost(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	var a *model.Schedule
	hallid, _ := strconv.Atoi(pm.ByName("hId"))

	MID, _ := strconv.Atoi(r.FormValue("mid"))
	fmt.Println("printing mid", MID)
	Time := r.FormValue("time")
	fmt.Println("printing time", Time)
	DAy := r.FormValue("day")
	fmt.Println("printing day", DAy)
	Dimen := r.FormValue("dimension")
	fmt.Println("printing day", Dimen)
	a = &model.Schedule{MoviemID: MID, StartingTime: Time, Dimension: Dimen, HallID: hallid, Day: DAy}
	if MID != 0 && Time != "" && DAy != "" && hallid != 0 {
		m.ssrv.StoreSchedule(a)
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", nil))

}
