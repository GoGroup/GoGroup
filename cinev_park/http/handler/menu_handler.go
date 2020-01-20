package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/GoGroup/Movie-and-events/controller"
	"github.com/GoGroup/Movie-and-events/movie"

	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/hall"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
)

type MenuHandler struct {
	tmpl *template.Template
	csrv cinema.CinemaService
	hsrv hall.HallService
	ssrv schedule.ScheduleService
	msrv movie.MovieService
}

func NewMenuHandler(t *template.Template, cs cinema.CinemaService, hs hall.HallService, ss schedule.ScheduleService, ms movie.MovieService) *MenuHandler {

	return &MenuHandler{tmpl: t, csrv: cs, hsrv: hs, ssrv: ss, msrv: ms}

}

// func (m *MenuHandler) AdminSchedule(w http.ResponseWriter, r *http.Request) {

// 	s := schedule.Schedules()
// 	fmt.Println(tmpl.ExecuteTemplate(w, "adminScheduleList.layout", s))

// }

// func (m *MenuHandler) AdminCinema(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

// 	var errr []error
// 	var NewCinemaArray []model.Cinema

// 	c, err := m.csrv.Cinemas()
// 	for _, element := range c {
// 		element.Halls, errr = m.hsrv.CinemaHalls(element.ID)
// 		NewCinemaArray = append(NewCinemaArray, element)
// 	}
// 	fmt.Println(NewCinemaArray)

// 	if len(err) > 0 || len(errr) > 0 {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	}
// 	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminCinemaList.layout", NewCinemaArray))

// }

// func (m *MenuHandler) AdminScheduleDelete(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
// 	fmt.Println("In admin schedule*****************")

// 	fmt.Println("trying to delete*****************")
// 	SchedulelID, _ := strconv.Atoi(pm.ByName("sId"))
// 	uSchID := uint(SchedulelID)
// 	m.ssrv.DeleteSchedules(uSchID)

// 	fmt.Println(pm.ByName("hId"))

// 	var All [][]model.Schedule
// 	var err []error
// 	var schedules []model.Schedule
// 	HallID, _ := strconv.Atoi(pm.ByName("hId"))
// 	uHallID := uint(HallID)
// 	Days := []string{"Monday", "Tuesday", "Wednsday", "Thursday", "Friday", "Saturday", "Sunday"}
// 	for _, d := range Days {
// 		schedules, err = m.ssrv.ScheduleHallDay(uHallID, d)
// 		All = append(All, schedules)

// 	}
// 	if len(err) > 0 {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	}

// 	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", All))

// }

// func (m *MenuHandler) AdminSchedule(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
// 	fmt.Println("In admin schedule*****************")

// 	fmt.Println(pm.ByName("hId"))

// 	var All [][]model.ScheduleWithMovie
// 	var SWM []model.ScheduleWithMovie
// 	var err []error
// 	var sm model.ScheduleWithMovie
// 	var schedules []model.Schedule
// 	HallID, _ := strconv.Atoi(pm.ByName("hId"))
// 	uHallID := uint(HallID)
// 	Days := []string{"Monday", "Tuesday", "Wednsday", "Thursday", "Friday", "Saturday", "Sunday"}
// 	for _, d := range Days {
// 		schedules, err = m.ssrv.ScheduleHallDay(uHallID, d)
// 		for _, s := range schedules {
// 			m, _, _ := controller.GetMovieDetails(s.MoviemID)
// 			sm = model.ScheduleWithMovie{s, m.Title}
// 			SWM = append(SWM, sm)

// 		}
// 		All = append(All, SWM)

// 	}
// 	tempo := struct {
// 		HallId int
// 		List   [][]model.ScheduleWithMovie
// 	}{
// 		HallId: HallID,
// 		List:   All,
// 	}
// 	if len(err) > 0 {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	}
// 	fmt.Println(tempo)

// 	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", tempo))

// }

// func (m *MenuHandler) NewAdminSchedule(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
// 	var MovieTitles *model.UpcomingMovies
// 	var err error
// 	var err2 error

// 	if r.FormValue("movie") != "" {
// 		Movie := r.FormValue("movie")
// 		fmt.Println(Movie)
// 		MovieTitles, err2, err = controller.SearchMovie(Movie)
// 		if err == nil || err2 == nil {
// 			fmt.Println("did seatch")
// 			fmt.Println(MovieTitles)
// 		}

// 	} else {
// 		fmt.Println("empty")
// 		movie := r.FormValue("movie")
// 		fmt.Println(movie)

// 	}

// 	convid, _ := strconv.Atoi(r.FormValue("id"))
// 	hallid, _ := strconv.Atoi(pm.ByName("hId"))

// 	tempo := struct {
// 		M       *model.UpcomingMovies
// 		MovieN  string
// 		MovieID int
// 		HallID  int
// 	}{
// 		M:       MovieTitles,
// 		MovieN:  r.FormValue("moviename"),
// 		MovieID: convid,
// 		HallID:  hallid,
// 	}

// 	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", tempo))

// }
// func (m *MenuHandler) NewAdminSchedulePost(w http.ResponseWriter, r *http.Request, pm httprouter.Params) {
// 	var a *model.Schedule
// 	var movie *model.Moviem
// 	hallid, _ := strconv.Atoi(pm.ByName("hId"))
// 	tempo := struct {
// 		M       *model.UpcomingMovies
// 		MovieN  string
// 		MovieID int
// 		HallID  int
// 	}{
// 		M:       nil,
// 		MovieN:  "",
// 		MovieID: 0,
// 		HallID:  hallid,
// 	}
// 	MID, _ := strconv.Atoi(r.FormValue("mid"))
// 	fmt.Println("printing mid", MID)
// 	Time := r.FormValue("time")
// 	fmt.Println("printing time", Time)
// 	DAy := r.FormValue("day")
// 	fmt.Println("printing day", DAy)
// 	Dimen := r.FormValue("3or2d")
// 	fmt.Println("printing day", Dimen)
// 	a = &model.Schedule{MoviemID: MID, StartingTime: Time, Dimension: Dimen, HallID: hallid, Day: DAy}
// 	movie = &model.Moviem{TmdbID: MID}
// 	if MID != 0 && Time != "" && DAy != "" && hallid != 0 {
// 		m.ssrv.StoreSchedule(a)
// 		m.msrv.StoreMovie(movie)
// 	}

// 	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", tempo))

// }

func (m *MenuHandler) Index(w http.ResponseWriter, r *http.Request) {

	fmt.Println(m.tmpl.ExecuteTemplate(w, "index.layout", nil))

}
func getCode(r *http.Request, defaultCode int) (int, string) {
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[2])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {
			fmt.Println(".....in first if")
			return code, p[2]
		} else {
			fmt.Println("...........in first if")
			fmt.Println(p)
			return defaultCode, p[1]
		}
	} else {
		fmt.Println("...........in not if")
		return defaultCode, ""
	}

}

func (m *MenuHandler) EachMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sdfdasasdcsdgvasfgsfkljsdfklsjdlfmslk")
	d, stringid := getCode(r, 0)
	fmt.Println(d)
	fmt.Println(stringid)
	trailerKey := controller.GetTrailer(stringid)
	id, _ := strconv.Atoi(stringid)

	details, _, _ := controller.GetMovieDetails(id)
	details.Trailer = trailerKey

	fmt.Println(m.tmpl.ExecuteTemplate(w, "EachMovie.layout", details))
	//	fmt.Println(m.tmpl.ExecuteTemplate(w, "index.layout", nil))

}
func (m *MenuHandler) Movies(w http.ResponseWriter, r *http.Request) {
	var nowshowingdetails []model.MovieDetails
	upcomingmovies, err1, err2 := controller.GetUpcomingMovies()
	fmt.Println(upcomingmovies)
	nowshowingmovies, err := m.msrv.Movies()
	for _, a := range nowshowingmovies {

		movie, _, _ := controller.GetMovieDetails(a.TmdbID)

		nowshowingdetails = append(nowshowingdetails, *movie)
	}

	if len(err) > 0 || err1 != nil || err2 != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tempo := struct {
		Upc *model.UpcomingMovies
		Nws []model.MovieDetails
	}{
		Upc: upcomingmovies,
		Nws: nowshowingdetails,
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "Movie.layout", tempo))

}
func (m *MenuHandler) Theaters(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(m.tmpl.ExecuteTemplate(w, "theatersList.layout", NewCinemaArray))

}

func (m *MenuHandler) TheaterSchedule(w http.ResponseWriter, r *http.Request) {
	var CName string
	var CId string
	fmt.Println("In theatesrs sdlkfsdjf")
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[4])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {
			fmt.Println(".....in first if")
			CName = p[3]
			CId = p[4]
			//return code, p[2]
		} else {
			fmt.Println("...........in first if")
			fmt.Println(p)
			//return defaultCode, p[1]
		}
	} else {
		fmt.Println("...........in not if")
		//return defaultCode, ""
	}

	//CName := r.FormValue("cName")
	//CId, _ := strconv.Atoi(r.FormValue("cId"))
	CcId, _ := strconv.Atoi(CId)
	uCId := uint(CcId)
	H := model.HallSchedule{}
	B := model.BindedSchedule{}

	Days := []string{"Monday", "Tuesday", "Wednsday", "Thursday", "Friday", "Saturday", "Sunday"}

	for _, d := range Days {
		fmt.Println(d)
		schdls, _ := m.ssrv.HallSchedules(uCId, d)
		fmt.Println(schdls)
		for _, s := range schdls {
			mo, _, _ := controller.GetMovieDetails(s.MoviemID)
			fmt.Println(mo)

			B.PosterPath = mo.PosterPath
			B.MovieName = mo.Title
			B.Runtime = mo.RunTime
			hall, _ := m.hsrv.Hall(uint(s.HallID))
			fmt.Println("hall is", hall)
			B.HallName = hall.HallName
			B.StartTime = s.StartingTime
			B.Day = d

			H.All = append(H.All, B)
		}

	}
	H.CinemaName = CName
	fmt.Println("************************************")
	fmt.Println(H)
	// if len(err) > 0 {
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// }

	fmt.Println(m.tmpl.ExecuteTemplate(w, "scheduleDisplay.layout", H))

}
