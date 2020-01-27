package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/GoGroup/Movie-and-events/form"

	"github.com/GoGroup/Movie-and-events/hash"
	"github.com/GoGroup/Movie-and-events/rtoken"

	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/controller"
	"github.com/GoGroup/Movie-and-events/hall"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/movie"
	"github.com/GoGroup/Movie-and-events/schedule"
)

type AdminHandler struct {
	tmpl        *template.Template
	csrv        cinema.CinemaService
	hsrv        hall.HallService
	ssrv        schedule.ScheduleService
	msrv        movie.MovieService
	csrfSignKey []byte
}

const nameKey = "name"
const capKey = "cap"
const priceKey = "price"
const vipcapkey = "vipcap"
const vipKey = "vip"
const csrfHKey = "_csrf"

func NewAdminHandler(t *template.Template, cs cinema.CinemaService, hs hall.HallService, ss schedule.ScheduleService, ms movie.MovieService, csKey []byte) *AdminHandler {

	return &AdminHandler{tmpl: t, csrv: cs, hsrv: hs, ssrv: ss, msrv: ms, csrfSignKey: csKey}

}

func (m *AdminHandler) AdminCinema(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminCinemaList.layout", NewCinemaArray))

}

func (m *AdminHandler) AdminScheduleDelete(w http.ResponseWriter, r *http.Request) {

	var HallID int
	var SchedulelID int
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[5])
		code2, err2 := strconv.Atoi(p[6])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil && err2 == nil {
			fmt.Println(".....in first if")
			HallID = code
			SchedulelID = code2
		}
	}

	fmt.Println("In admin schedule*****************")

	fmt.Println("trying to delete*****************")

	uSchID := uint(SchedulelID)
	m.ssrv.DeleteSchedules(uSchID)

	fmt.Println(r.FormValue("hId"))

	var All [][]model.ScheduleWithMovie
	var SWM []model.ScheduleWithMovie
	var err []error
	var sm model.ScheduleWithMovie
	var schedules []model.Schedule

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
	hall, _ := m.hsrv.Hall(uint(HallID))
	cinema, _ := m.csrv.Cinema(uint(hall.CinemaID))
	tempo := struct {
		HallId int
		List   [][]model.ScheduleWithMovie
		Hall   *model.Hall
		Cinema *model.Cinema
	}{
		HallId: HallID,
		List:   All,
		Hall:   hall,
		Cinema: cinema,
	}
	if len(err) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(tempo)

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", tempo))
}

func (m *AdminHandler) AdminDeleteHalls(w http.ResponseWriter, r *http.Request) {
	var CID uint
	var HID uint
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[5])
		code2, err2 := strconv.Atoi(p[6])
		fmt.Println(err)
		fmt.Println(err2)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {
			fmt.Println(".....in first if")
			CID = uint(code)
			HID = uint(code2)
		}
	}
	h, e := m.hsrv.DeleteHall(HID)
	if len(e) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}
	fmt.Println("delteed", h)

	fmt.Println("%%%%%%%%%%%%%%%%%")
	halls, errr := m.hsrv.CinemaHalls(CID)
	if len(errr) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}

	tempo := struct {
		Halls []model.Hall
		Cid   uint
	}{Halls: halls, Cid: CID}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "halls.layout", tempo))

}

func (m *AdminHandler) AdminHallsNew(w http.ResponseWriter, r *http.Request) {
	var CID uint
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[5])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {
			fmt.Println(".....in first if")
			CID = uint(code)
		}
	}
	if r.Method == http.MethodGet {
		CSFRToken, err := rtoken.GenerateCSRFToken(m.csrfSignKey)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewHall.layout", form.Input{
			CSRF: CSFRToken, Cid: CID}))
		return
	}
	if m.isParsableFormPost(w, r) {
		///Validate the form data
		HallNewForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: r.FormValue(csrfKey), Cid: CID}
		fmt.Println(r.FormValue("name"))
		fmt.Println(r.FormValue("cap"))
		fmt.Println(r.FormValue("price"))
		fmt.Println(r.FormValue("name"))
		fmt.Println(CID)
		HallNewForm.ValidateRequiredFields(nameKey, capKey, priceKey, vipcapkey, vipKey)
		HallNewForm.ValidateFieldsInteger(capKey, priceKey, vipcapkey, vipKey)
		HallNewForm.ValidateFieldsRange(capKey, priceKey, vipcapkey, vipKey)

		if !HallNewForm.IsValid() {
			fmt.Println("last")
			err := m.tmpl.ExecuteTemplate(w, "adminNewHall.layout", HallNewForm)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		if m.hsrv.HallExists(r.FormValue(nameKey)) {
			fmt.Println("secondemailexist")
			HallNewForm.VErrors.Add(nameKey, "This HallName is already in use!")
			err := m.tmpl.ExecuteTemplate(w, "adminNewHall.layout", HallNewForm)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		hn := r.FormValue(nameKey)
		c, _ := strconv.Atoi(r.FormValue(capKey))
		pri, _ := strconv.Atoi(r.FormValue(priceKey))
		vp, _ := strconv.Atoi(r.FormValue(vipKey))
		wd, _ := strconv.Atoi(r.FormValue(vipcapkey))
		h := model.Hall{
			HallName:    hn,
			Capacity:    uint(c),
			CinemaID:    CID,
			Price:       uint(pri),
			VIPPrice:    uint(vp),
			VIPCapacity: uint(wd),
		}
		hall, errr := m.hsrv.StoreHall(&h)
		fmt.Println("In ^^^^^^^^^^^^^^^^^^^")
		fmt.Println(hall)
		if len(errr) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}

		// tempo := struct{ Cid uint }{Cid: CID}

		fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewHall.layout", HallNewForm))

	}
}
func (m *AdminHandler) AdminHalls(w http.ResponseWriter, r *http.Request) {
	var CID uint
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[5])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {
			fmt.Println(".....in first if")
			CID = uint(code)
		}
	}

	halls, errr := m.hsrv.CinemaHalls(CID)
	if len(errr) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}

	tempo := struct {
		Halls []model.Hall
		Cid   uint
	}{Halls: halls, Cid: CID}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "halls.layout", tempo))

}

func (m *AdminHandler) AdminSchedule(w http.ResponseWriter, r *http.Request) {
	var HallID int

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
			HallID = code
		}
	}

	var All [][]model.ScheduleWithMovie
	var SWM []model.ScheduleWithMovie
	var err []error
	var sm model.ScheduleWithMovie
	var schedules []model.Schedule
	//HallID, _ := strconv.Atoi(r.FormValue("hId"))
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
	hall, _ := m.hsrv.Hall(uint(HallID))
	cinema, _ := m.csrv.Cinema(uint(hall.CinemaID))
	tempo := struct {
		HallId int
		List   [][]model.ScheduleWithMovie
		Hall   *model.Hall
		Cinema *model.Cinema
	}{
		HallId: HallID,
		List:   All,
		Hall:   hall,
		Cinema: cinema,
	}
	if len(err) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println(tempo)

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminScheduleList.layout", tempo))

}

func (m *AdminHandler) NewAdminScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "POST" {
		m.NewAdminSchedulePost(w, r)
	} else if r.Method == "GET" {
		m.NewAdminSchedule(w, r)
	}

	if err != nil {
		fmt.Println("hi")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

}

func (m *AdminHandler) NewAdminSchedule(w http.ResponseWriter, r *http.Request) {
	var MovieTitles *model.UpcomingMovies
	var err error
	var err2 error
	var hallid int
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[5])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {
			fmt.Println(".....in first if")
			hallid = code
		}
	}

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
	//hallid, _ := strconv.Atoi(r.FormValue("hId"))
	hall, _ := m.hsrv.Hall(uint(hallid))
	cinema, _ := m.csrv.Cinema(uint(hall.CinemaID))

	tempo := struct {
		M       *model.UpcomingMovies
		MovieN  string
		MovieID int
		HallID  int
		Hall    *model.Hall
		Cinema  *model.Cinema
	}{
		M:       MovieTitles,
		MovieN:  r.FormValue("moviename"),
		MovieID: convid,
		HallID:  hallid,
		Hall:    hall,
		Cinema:  cinema,
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", tempo))

}
func (m *AdminHandler) NewAdminSchedulePost(w http.ResponseWriter, r *http.Request) {
	var hallid int

	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[5])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {
			fmt.Println(".....in first if")
			hallid = code
		}
	}

	var a *model.Schedule
	var movie *model.Moviem

	hall, _ := m.hsrv.Hall(uint(hallid))
	cinema, _ := m.csrv.Cinema(uint(hall.CinemaID))
	tempo := struct {
		M       *model.UpcomingMovies
		MovieN  string
		MovieID int
		HallID  int
		Hall    *model.Hall
		Cinema  *model.Cinema
	}{
		M:       nil,
		MovieN:  "",
		MovieID: 0,
		HallID:  hallid,
		Hall:    hall,
		Cinema:  cinema,
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

func (m *AdminHandler) isParsableFormPost(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("parse")
	fmt.Println(r.Method == http.MethodPost)
	fmt.Println(hash.ParseForm(w, r))
	fmt.Println(rtoken.IsCSRFValid(r.FormValue(csrfKey), m.csrfSignKey))

	return r.Method == http.MethodPost &&
		hash.ParseForm(w, r) &&
		rtoken.IsCSRFValid(r.FormValue(csrfKey), m.csrfSignKey)
}
