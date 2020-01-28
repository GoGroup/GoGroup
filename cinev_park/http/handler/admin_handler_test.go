package handler

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/rtoken"

	"github.com/GoGroup/Movie-and-events/event/repository"
	"github.com/GoGroup/Movie-and-events/event/service"
	schrep "github.com/GoGroup/Movie-and-events/schedule/repository"
	schser "github.com/GoGroup/Movie-and-events/schedule/service"

	cinrep "github.com/GoGroup/Movie-and-events/cinema/repository"
	cinser "github.com/GoGroup/Movie-and-events/cinema/service"
	hallrepo "github.com/GoGroup/Movie-and-events/hall/repository"
	hallser "github.com/GoGroup/Movie-and-events/hall/service"
)

func TestAdminEventNew(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

	eventRepo := repository.NewMockEventRepo(nil)
	eventServ := service.NewEventService(eventRepo)

	adminEvHandler := NewAdminHandler(tmpl, nil, nil, nil, nil, eventServ, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cinemas/events/new/", adminEvHandler.AdminEventsNew)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	URL := ts.URL
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	form := url.Values{}
	form.Add("name", model.EvenMock.Name)
	form.Add("description", model.EvenMock.Description)
	form.Add("time", model.EvenMock.Time)
	form.Add("image", model.EvenMock.Image)
	form.Add("location", model.EvenMock.Location)
	CSFRToken, _ := rtoken.GenerateCSRFToken(csrfSignKey)
	form.Add(csrfHKey, CSFRToken)

	resp, err := tc.PostForm(URL+"/admin/cinemas/events/new/", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

}

func TestAdminScheduleDelete(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

	schrepo := schrep.NewMockScheduleRepo(nil)
	schserv := schser.NewScheduleService(schrepo)

	hrep := hallrepo.NewMockHallRepo(nil)
	hser := hallser.NewHallService(hrep)

	cinr := cinrep.NewMockCinemaRepo(nil)
	cins := cinser.NewCinemaService(cinr)

	adminSchHandler := NewAdminHandler(tmpl, cins, hser, schserv, nil, nil, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cinemas/schedule/delete/", adminSchHandler.AdminScheduleDelete)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	URL := ts.URL

	resp, err := tc.Get(URL + "/admin/cinemas/schedule/delete/1/1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

}

func TestAdminEvent(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

	eventRepo := repository.NewMockEventRepo(nil)
	eventServ := service.NewEventService(eventRepo)

	adminEvHandler := NewAdminHandler(tmpl, nil, nil, nil, nil, eventServ, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/", adminEvHandler.AdminEventList)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	URL := ts.URL

	resp, err := tc.Get(URL + "/admin/events/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("MockName 1")) {
		t.Errorf("want body to contain %q", body)
	}

}



func TestAdminEventDelete(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

	eventRepo := repository.NewMockEventRepo(nil)
	eventServ := service.NewEventService(eventRepo)

	adminEvHandler := NewAdminHandler(tmpl, nil, nil, nil, nil, eventServ, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/delete/", adminEvHandler.AdminDeleteEvents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	URL := ts.URL

	resp, err := tc.Get(URL + "/admin/events/delete/1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

}

func TestAdminEventUpdate(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

	eventRepo := repository.NewMockEventRepo(nil)
	eventServ := service.NewEventService(eventRepo)

	adminEvHandler := NewAdminHandler(tmpl, nil, nil, nil, nil, eventServ, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/update/", adminEvHandler.AdminEventUpdateList)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	URL := ts.URL
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	form := url.Values{}
	form.Add("name", model.EvenMock.Name)
	form.Add("description", model.EvenMock.Description)
	form.Add("time", model.EvenMock.Time)
	form.Add("image", model.EvenMock.Image)
	form.Add("location", model.EvenMock.Location)
	CSFRToken, _ := rtoken.GenerateCSRFToken(csrfSignKey)
	form.Add(csrfHKey, CSFRToken)

	resp, err := tc.PostForm(URL+"/admin/events/update/1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

}
