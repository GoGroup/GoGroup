package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/GoGroup/Movie-and-events/booking"

	"github.com/GoGroup/Movie-and-events/flash"
	"github.com/GoGroup/Movie-and-events/user"

	"github.com/GoGroup/Movie-and-events/comment"
	"github.com/GoGroup/Movie-and-events/controller"
	"github.com/GoGroup/Movie-and-events/event"
	"github.com/GoGroup/Movie-and-events/movie"

	"github.com/GoGroup/Movie-and-events/cinema"
	"github.com/GoGroup/Movie-and-events/hall"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
)

type MenuHandler struct {
	tmpl   *template.Template
	csrv   cinema.CinemaService
	hsrv   hall.HallService
	ssrv   schedule.ScheduleService
	msrv   movie.MovieService
	comsrv comment.CommentService
	evsrv  event.EventService
	usrv   user.UserService
	bsrv   booking.BookingService
}

func NewMenuHandler(t *template.Template, cs cinema.CinemaService, hs hall.HallService, ss schedule.ScheduleService, ms movie.MovieService, comser comment.CommentService, evs event.EventService, u user.UserService, b booking.BookingService) *MenuHandler {

	return &MenuHandler{tmpl: t, csrv: cs, hsrv: hs, ssrv: ss, msrv: ms, comsrv: comser, evsrv: evs, usrv: u, bsrv: b}

}

func (m *MenuHandler) Index(w http.ResponseWriter, r *http.Request) {

	fmt.Println(m.tmpl.ExecuteTemplate(w, "index.layout", nil))

}

func (m *MenuHandler) Bookings(w http.ResponseWriter, r *http.Request) {
	var sm []model.ScheduleWithMovie

	activeSession := r.Context().Value(ctxUserSessionKey).(*model.Session)
	user, errs := m.usrv.User(activeSession.UUID)
	if len(errs) > 0 {

	}
	b, _ := m.bsrv.Bookings(user.ID)
	fmt.Println(b)

	for _, element := range b {
		fmt.Println("in loop")
		s, _ := m.ssrv.Schedule(element.ScheduleID)
		fmt.Println(s, "here is s")
		m, _, _ := controller.GetMovieDetails(s.MoviemID)
		fmt.Println("::::::::::::::::::::::")
		scheduleMovie := model.ScheduleWithMovie{*s, m.Title}
		fmt.Println("<<<<<<<<<")
		sm = append(sm, scheduleMovie)
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "bookings.layout", sm))

}
func (m *MenuHandler) Search(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		if r.FormValue("movie") != "" {

			srch, err, err2 := controller.SearchMovie(r.FormValue("movie"))

			if err != nil || err2 != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			} else {
				fmt.Println(m.tmpl.ExecuteTemplate(w, "search.layout", srch))

			}
		}

	}

}

func (m *MenuHandler) EventList(w http.ResponseWriter, r *http.Request) {
	events, errs := m.evsrv.Events()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}
	tempo := struct{ Events []model.Event }{Events: events}
	fmt.Println(m.tmpl.ExecuteTemplate(w, "eventList.layout", tempo))

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

func (m *MenuHandler) EachNowShowing(w http.ResponseWriter, r *http.Request) {
	var id int
	activeSession := r.Context().Value(ctxUserSessionKey).(*model.Session)
	user, errs := m.usrv.User(activeSession.UUID)
	if len(errs) > 0 {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code, err := strconv.Atoi(p[3])

		if err == nil {
			fmt.Println(".....in first if")
			id = code
		}
	}

	if r.FormValue("comment") != "" {
		c := model.Comment{UserID: user.ID, UserName: user.FullName, MovieID: uint(id), Message: r.FormValue("comment")}

		m.comsrv.StoreComment(&c)
	}

	trailerKey := controller.GetTrailer(strconv.Itoa(id))
	details, _, _ := controller.GetMovieDetails(id)
	details.Trailer = trailerKey
	comments, _ := m.comsrv.RetrieveComments(uint(id))

	tempo := struct {
		Comments    []model.Comment
		MovieDetail *model.MovieDetails
	}{
		Comments:    comments,
		MovieDetail: details,
	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "EachNowShowing.layout", tempo))
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

func (m *MenuHandler) TheaterScheduleBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Int theradkafsldkfjas Error next to this")

	var CName string
	var CId string
	var SId string
	fmt.Println(SId)

	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {

	} else if len(p) > 1 {

		fmt.Println("Error next to this")

		code, err := strconv.Atoi(p[5])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {

			CName = p[4]
			CId = p[5]
			SId = p[6]

		} else {

			fmt.Println(p)

		}
	} else {

	}
	Sidint, _ := strconv.Atoi(SId)

	fmt.Println(r.FormValue("seat"))
	givenPrice, _ := strconv.Atoi(r.FormValue("seat"))
	activeSession := r.Context().Value(ctxUserSessionKey).(*model.Session)
	user, errs := m.usrv.User(activeSession.UUID)
	fmt.Println(user)
	fmt.Println(errs)
	schdl, _ := m.ssrv.Schedule(uint(Sidint))

	if r.FormValue("seat") != "" {
		if givenPrice >= int(user.Amount) {
			flash.SetFlash(w, "error", []byte("You dont have enough money in your account"))

		} else {

			b := model.Booking{
				UserID:     user.ID,
				ScheduleID: uint(Sidint),
			}
			m.bsrv.StoreBooking(&b)
			m.ssrv.UpdateSchedulesBooked(schdl, 1)
			m.usrv.UpdateUserAmount(user, user.Amount-uint(givenPrice))
			flash.SetFlash(w, "success", []byte("You Have successfully booked"))

		}

	}
	url := "/theater/schedule/" + CName + "/" + CId
	http.Redirect(w, r, url, 303)

	// if c == nil {
	// 	fmt.Fprint(w, "No flash messages")
	// 	fmt.Fprintf(w, "%s", r.FormValue("seat"))
	// } else {
	// 	fmt.Fprintf(w, "%s", c)
	// 	fmt.Fprintf(w, "%s", r.FormValue("seat"))

	// }

}

func (m *MenuHandler) TheaterSchedule(w http.ResponseWriter, r *http.Request) {
	var CName string
	var CId string
	fmt.Println("In >>>>>>>>>>>>>>>theater schedule")

	c, er := flash.GetFlash(w, r, "error")
	suc, er := flash.GetFlash(w, r, "success")
	fmt.Println(er)
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {

	} else if len(p) > 1 {
		fmt.Println("In >>>>>>>>>>>>>>>theater schedule")

		code, err := strconv.Atoi(p[4])
		fmt.Println(err)
		fmt.Println(p)
		fmt.Println(code)
		if err == nil {

			CName = p[3]
			CId = p[4]

		} else {

			fmt.Println(p)

		}
	} else {
		fmt.Println("...........in not if")

	}

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
			B.ScheduleID = s.ID
			B.Booked = s.Booked

			hall, _ := m.hsrv.Hall(uint(s.HallID))
			fmt.Println("hall is", hall)
			B.HallName = hall.HallName
			B.VIPPrice = hall.VIPPrice
			B.VIPCapacity = hall.VIPCapacity
			B.Price = hall.Price
			B.Capacity = hall.Capacity
			B.StartTime = s.StartingTime
			B.Day = d
			B.Dimension = s.Dimension
			B.Available = B.Capacity - B.Booked

			H.All = append(H.All, B)
		}

	}
	H.CinemaName = CName
	tempo := struct {
		HallS    model.HallSchedule
		CinemaID uint
		Err      string
		Succ     string
	}{
		HallS:    H,
		CinemaID: uint(CcId),
		Err:      string(c[:]),
		Succ:     string(suc[:]),
	}
	// if len(err) > 0 {
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// }

	fmt.Println(m.tmpl.ExecuteTemplate(w, "scheduleDisplay.layout", tempo))

}
