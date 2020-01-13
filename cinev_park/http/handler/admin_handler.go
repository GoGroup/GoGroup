package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/controller"
	"github.com/GoGroup/Movie-and-events/hall"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/movie"
	"github.com/GoGroup/Movie-and-events/schedule"
	"github.com/julienschmidt/httprouter"
)

type AdminHandler struct {
	tmpl *template.Template
	csrv cinema.CinemaService
	hsrv hall.HallService
	ssrv schedule.ScheduleService
	msrv movie.MovieService
}

func NewAdminHandler(t *template.Template, cs cinema.CinemaService, hs hall.HallService, ss schedule.ScheduleService, ms movie.MovieService) *AdminHandler {

	return &AdminHandler{tmpl: t, csrv: cs, hsrv: hs, ssrv: ss, msrv: ms}

}

func (m *AdminHandler) AdminCinema(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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

func (m *AdminHandler) AdminScheduleDelete(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	fmt.Println("In admin schedule*****************")

	fmt.Println("trying to delete*****************")
	SchedulelID, _ := strconv.Atoi(pm.ByName("sId"))
	uSchID := uint(SchedulelID)
	m.ssrv.DeleteSchedules(uSchID)

	fmt.Println(pm.ByName("hId"))

	var All [][]model.Schedule
	var err []error
	var schedules []model.Schedule
	HallID, _ := strconv.Atoi(pm.ByName("hId"))
	uHallID := uint(HallID)
	Days := []string{"Monday", "Tuesday", "Wednsday", "Thursday", "Friday", "Saturday", "Sunday"}
	for _, d := range Days {
		schedules, err = m.ssrv.ScheduleHallDay(uHallID, d)
		All = append(All, schedules)

	}
	if len(err) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", All))

}

func (m *AdminHandler) AdminSchedule(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	fmt.Println("In admin schedule*****************")

	fmt.Println(pm.ByName("hId"))

	var All [][]model.ScheduleWithMovie
	var SWM []model.ScheduleWithMovie
	var err []error
	var sm model.ScheduleWithMovie
	var schedules []model.Schedule
	HallID, _ := strconv.Atoi(pm.ByName("hId"))
	uHallID := uint(HallID)
	Days := []string{"Monday", "Tuesday", "Wednsday", "Thursday", "Friday", "Saturday", "Sunday"}
	for _, d := range Days {
		schedules, err = m.ssrv.ScheduleHallDay(uHallID, d)
		SWM = nil
		for _, s := range schedules {
			m, _, _ := controller.GetMovieDetails(s.MoviemID)
			sm = model.ScheduleWithMovie{s, m.Title}
			SWM = append(SWM, sm)

		}
		All = append(All, SWM)

	}
	tempo := struct {
		HallId int
		List   [][]model.ScheduleWithMovie
	}{
		HallId: HallID,
		List:   All,
	}
	if len(err) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(tempo)

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", tempo))

}

func (m *AdminHandler) NewAdminSchedule(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
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
func (m *AdminHandler) NewAdminSchedulePost(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
	var a *model.Schedule
	var movie *model.Moviem
	hallid, _ := strconv.Atoi(pm.ByName("hId"))
	tempo := struct {
		M       *model.UpcomingMovies
		MovieN  string
		MovieID int
		HallID  int
	}{
		M:       nil,
		MovieN:  "",
		MovieID: 0,
		HallID:  hallid,
	}
	MID, _ := strconv.Atoi(r.FormValue("mid"))
	fmt.Println("printing mid", MID)
	Time := r.FormValue("time")
	fmt.Println("printing time", Time)
	DAy := r.FormValue("day")
	fmt.Println("printing day", DAy)
	Dimen := r.FormValue("3or2d")
	fmt.Println("printing day", Dimen)
	a = &model.Schedule{MoviemID: MID, StartingTime: Time, Dimension: Dimen, HallID: hallid, Day: DAy}
	movie = &model.Moviem{TmdbID: MID}
	if MID != 0 && Time != "" && DAy != "" && hallid != 0 {
		m.ssrv.StoreSchedule(a)
		m.msrv.StoreMovie(movie)
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", tempo))

}
