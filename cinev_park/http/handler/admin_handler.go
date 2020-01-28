package handler

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/GoGroup/Movie-and-events/event"
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
	evrv        event.EventService
	csrfSignKey []byte
}

const nameKey = "name"
const locationKey = "location"
const descriptionKey = "description"
const datetimekey = "datetime"
const fileKey = "file"

const capKey = "cap"
const priceKey = "price"
const vipcapkey = "vipcap"
const vipKey = "vip"
const csrfHKey = "_csrf"

func NewAdminHandler(t *template.Template, cs cinema.CinemaService, hs hall.HallService, ss schedule.ScheduleService, ms movie.MovieService, ev event.EventService, csKey []byte) *AdminHandler {

	return &AdminHandler{tmpl: t, csrv: cs, hsrv: hs, ssrv: ss, msrv: ms, evrv: ev, csrfSignKey: csKey}

}

func (m *AdminHandler) AdminCinema(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		if r.FormValue("cinemaName") != "" {
			fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
			c := model.Cinema{
				CinemaName: r.FormValue("cinemaName"),
			}
			cc, ee := m.csrv.StoreCinema(&c)
			fmt.Println(cc)
			fmt.Println(ee)
		}

	}
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
func (m *AdminHandler) AdminDeleteEvents(w http.ResponseWriter, r *http.Request) {

	var EID uint
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		fmt.Println("in first if")
		//return defaultCode, p[0]
	} else if len(p) > 1 {
		fmt.Println("..in first if")
		code2, err2 := strconv.Atoi(p[4])

		fmt.Println(err2)
		fmt.Println(p)

		if err2 == nil {
			fmt.Println(".....in first if")
			EID = uint(code2)
		}
	}
	h, e := m.evrv.DeleteEvent(EID)
	if len(e) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}
	fmt.Println("delteed", h)

	fmt.Println("%%%%%%%%%%%%%%%%%")
	events, errr := m.evrv.Events()
	if len(errr) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminEventList.layout", events))

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
func (m *AdminHandler) AdminEventsNew(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	CSFRToken, err := rtoken.GenerateCSRFToken(m.csrfSignKey)
	if r.Method == http.MethodGet {
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewEvent.layout", form.Input{
			CSRF: CSFRToken}))
		return
	}
	if m.isParsableFormPost(w, r) {

		EventNewForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: r.FormValue(csrfHKey)}
		fmt.Println("name")
		fmt.Println(r.FormValue(nameKey))
		fmt.Println("name")
		// _, s, _ := r.FormFile(fileKey)
		// fmt.Println("FILE")
		// fmt.Println(s.Filename)
		// fmt.Println("FILE")
		EventNewForm.ValidateRequiredFields(nameKey, locationKey, descriptionKey, datetimekey)
		EventNewForm.MinLength(descriptionKey, 20)
		EventNewForm.Date(datetimekey)
		mh, _, err := r.FormFile(fileKey)
		if err != nil || mh == nil {
			EventNewForm.VErrors.Add(fileKey, "File error")
		}
		if !EventNewForm.IsValid() {
			fmt.Println("last")
			err := m.tmpl.ExecuteTemplate(w, "adminNewEvent.layout", EventNewForm)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		if m.evrv.EventExists(r.FormValue(nameKey)) {

			EventNewForm.VErrors.Add(nameKey, "This Event exists!")
			err := m.tmpl.ExecuteTemplate(w, "adminNewEvent.layout", EventNewForm)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		en := r.FormValue(nameKey)
		c := r.FormValue(descriptionKey)
		pri := r.FormValue(datetimekey)
		vp := r.FormValue(locationKey)
		mf, fh, _ := r.FormFile(fileKey)
		defer mf.Close()
		fname := fh.Filename
		wd, err := os.Getwd()
		path := filepath.Join(wd, "view", "assets", "images", fname)
		image, err := os.Create(path)
		fmt.Println(path)
		if err != nil {
			fmt.Println("error")
		}

		defer image.Close()
		io.Copy(image, mf)
		h := model.Event{

			Name:        en,
			Description: c,
			Location:    vp,
			Time:        pri,
			Image:       fname,
		}
		event, errr := m.evrv.StoreEvent(&h)
		fmt.Println("In ^^^^^^^^^^^^^^^^^^^")
		fmt.Println(event)
		if len(errr) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}

		// tempo := struct{ Cid uint }{Cid: CID}

		fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewEvent.layout", form.Input{CSRF: CSFRToken}))

	}

}
func (m *AdminHandler) AdminHallUpdateList(w http.ResponseWriter, r *http.Request) {
fmt.Println("check")
	halls1, _ := m.hsrv.Halls()
	if m.isParsableFormPost(w, r) {
		var EID uint
		p := strings.Split(r.URL.Path, "/")
		if len(p) == 1 {
			fmt.Println("in first if")
			fmt.Println("in go if")
			//return defaultCode, p[0]
		} else if len(p) > 1 {
			fmt.Println("..in first if")
			fmt.Println("please")
			code2, err2 := strconv.Atoi(p[4])

			fmt.Println(err2)
			fmt.Println(p)

			if err2 == nil {
				fmt.Println(".....in first if")
				EID = uint(code2)
			}
		}
		myhall, errr := m.hsrv.Hall(EID)
		if errr != nil {
			fmt.Println("this part")
		}
		HallNewForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: r.FormValue(csrfHKey)}
		fmt.Println("name")
		fmt.Println(r.FormValue(nameKey))
		fmt.Println("name")
		HallNewForm.ValidateRequiredFields(nameKey, capKey, priceKey, vipcapkey, vipKey)
		HallNewForm.ValidateFieldsInteger(capKey, priceKey, vipcapkey, vipKey)
		HallNewForm.ValidateFieldsRange(capKey, priceKey, vipcapkey, vipKey)

		tempo1 := struct {
			Halls []model.Hall
			From  form.Input
			ID    uint
		}{Halls: halls1, From: HallNewForm, ID: EID}
		if !HallNewForm.IsValid() {
			fmt.Println("last")
			err := m.tmpl.ExecuteTemplate(w, "halls.layout", tempo1)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		if m.hsrv.HallExists(r.FormValue(nameKey)) {

			HallNewForm.VErrors.Add(nameKey, "This Event exists!")
			tempo2 := struct {
				Halls []model.Hall
				From  form.Input
				ID    uint
			}{Halls: halls1, From: HallNewForm, ID: EID}

			err := m.tmpl.ExecuteTemplate(w, "halls.layout", tempo2)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		hall, errrr := m.hsrv.UpdateHall(myhall)
		fmt.Println("In ^^^^^^^^^^^^^^^^^^^")
		fmt.Println(hall)
		if len(errrr) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		halls, _ := m.hsrv.Halls()
		tempo := struct {
			Halls []model.Hall
			From  form.Input
			ID    uint
		}{Halls: halls, From: HallNewForm, ID: EID}

		fmt.Println(m.tmpl.ExecuteTemplate(w, "halls.layout", tempo))
	}
}
func (m *AdminHandler) AdminEventUpdateList(w http.ResponseWriter, r *http.Request) {
	events1, _ := m.evrv.Events()
	r.ParseMultipartForm(0)
	if m.isParsableFormPost(w, r) {
		var EID uint
		p := strings.Split(r.URL.Path, "/")
		if len(p) == 1 {
			fmt.Println("in first if")
			//return defaultCode, p[0]
		} else if len(p) > 1 {
			fmt.Println("..in first if")
			code2, err2 := strconv.Atoi(p[4])

			fmt.Println(err2)
			fmt.Println(p)

			if err2 == nil {
				fmt.Println(".....in first if")
				EID = uint(code2)
			}
		}
		Myevent, errr := m.evrv.Event(EID)
		if errr != nil {
			fmt.Println("this part")
		}
		EventNewForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: r.FormValue(csrfHKey)}
		fmt.Println("name")
		fmt.Println(r.FormValue(nameKey))
		fmt.Println("name")
		// _, s, _ := r.FormFile(fileKey)
		// fmt.Println("FILE")
		// fmt.Println(s.Filename)
		// fmt.Println("FILE")
		EventNewForm.ValidateRequiredFields(nameKey, locationKey, descriptionKey, datetimekey)
		EventNewForm.MinLength(descriptionKey, 20)
		EventNewForm.Date(datetimekey)

		tempo1 := struct {
			Events []model.Event
			From   form.Input
			ID     uint
		}{Events: events1, From: EventNewForm, ID: EID}

		if !EventNewForm.IsValid() {
			fmt.Println("last")
			err := m.tmpl.ExecuteTemplate(w, "adminEventList.layout", tempo1)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		if m.evrv.EventExists(r.FormValue(nameKey)) && r.FormValue(nameKey) != Myevent.Name {

			EventNewForm.VErrors.Add(nameKey, "This Event exists!")
			tempo2 := struct {
				Events []model.Event
				From   form.Input
				ID     uint
			}{Events: events1, From: EventNewForm, ID: EID}

			err := m.tmpl.ExecuteTemplate(w, "adminEventList.layout", tempo2)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}

		mf, fh, _ := r.FormFile(fileKey)
		defer mf.Close()
		fname := fh.Filename
		if fname != Myevent.Image {
			wd, err := os.Getwd()
			path := filepath.Join(wd, "view", "assets", "images", fname)
			image, err := os.Create(path)
			fmt.Println(path)
			if err != nil {
				fmt.Println("error")
			}

			defer image.Close()
			io.Copy(image, mf)
		}

		event, errrr := m.evrv.UpdateEvent(Myevent)
		fmt.Println("In ^^^^^^^^^^^^^^^^^^^")
		fmt.Println(event)
		if len(errrr) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		events, _ := m.evrv.Events()
		tempo := struct {
			Events []model.Event
			From   form.Input
			ID     uint
		}{Events: events, From: EventNewForm, ID: EID}

		fmt.Println(m.tmpl.ExecuteTemplate(w, "adminEventList.layout", tempo))

	}

}

func (m *AdminHandler) AdminEventList(w http.ResponseWriter, r *http.Request) {
	events, errr := m.evrv.Events()
	CSFRToken, _ := rtoken.GenerateCSRFToken(m.csrfSignKey)
	tempo := struct {
		Events []model.Event
		From   form.Input
	}{Events: events, From: form.Input{CSRF: CSFRToken}}

	if len(errr) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	}

	fmt.Println(m.tmpl.ExecuteTemplate(w, "adminEventList.layout", tempo))
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
		HallNewForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: r.FormValue(csrfHKey), Cid: CID}
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
	CSFRToken, _ := rtoken.GenerateCSRFToken(m.csrfSignKey)
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
		From  form.Input
	}{Halls: halls, Cid: CID, From: form.Input{CSRF: CSFRToken}}

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
	CSFRToken, err := rtoken.GenerateCSRFToken(m.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

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
		From    form.Input
		Cinema  *model.Cinema
	}{
		M:       MovieTitles,
		MovieN:  r.FormValue("moviename"),
		MovieID: convid,
		HallID:  hallid,
		Hall:    hall,
		From:    form.Input{CSRF: CSFRToken},
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
	if m.isParsableFormPost(w, r) {
		var a *model.Schedule
		var movie *model.Moviem
		ScheduleNewForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}, CSRF: r.FormValue(csrfHKey)}
		hall, _ := m.hsrv.Hall(uint(hallid))
		cinema, _ := m.csrv.Cinema(uint(hall.CinemaID))
		MID, _ := strconv.Atoi(r.FormValue("mid"))
		fmt.Println("printing mid", MID)
		Time := r.FormValue("time")
		fmt.Println("printing time", Time)
		DAy := r.FormValue("day")
		fmt.Println("printing day", DAy)
		Dimen := r.FormValue("3or2d")
		fmt.Println("printing day", Dimen)
		ScheduleNewForm.ValidateRequiredFields("time", "day", "3or2d")
		ScheduleNewForm.Date("time")
		tempo := struct {
			M       *model.UpcomingMovies
			MovieN  string
			MovieID int
			HallID  int
			Hall    *model.Hall
			From    form.Input
			Cinema  *model.Cinema
		}{
			M:       nil,
			MovieN:  "",
			MovieID: 0,
			HallID:  hallid,
			Hall:    hall,
			From:    ScheduleNewForm,
			Cinema:  cinema,
		}
		if !ScheduleNewForm.IsValid() {
			fmt.Println("last")
			err := m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", tempo)
			if err != nil {
				fmt.Println("hiiii")
				fmt.Println(err)
			}
			return
		}
		a = &model.Schedule{MoviemID: MID, StartingTime: Time, Dimension: Dimen, HallID: hallid, Day: DAy}
		movie = &model.Moviem{TmdbID: MID}
		if MID != 0 && Time != "" && DAy != "" && hallid != 0 {
			m.ssrv.StoreSchedule(a)
			m.msrv.StoreMovie(movie)
		}

		fmt.Println(m.tmpl.ExecuteTemplate(w, "adminNewSchedule.layout", tempo))
	}
	fmt.Println("Error")
}

func (m *AdminHandler) isParsableFormPost(w http.ResponseWriter, r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false
	}
	fmt.Println("parse")
	fmt.Println(r.Method == http.MethodPost)
	fmt.Println(hash.ParseForm(w, r))
	fmt.Println(r.FormValue(csrfHKey))
	fmt.Println(rtoken.IsCSRFValid(r.FormValue(csrfHKey), m.csrfSignKey))

	return r.Method == http.MethodPost &&
		hash.ParseForm(w, r) &&
		rtoken.IsCSRFValid(r.FormValue(csrfHKey), m.csrfSignKey)
}
