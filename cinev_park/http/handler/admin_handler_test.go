package handler

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoGroup/Movie-and-events/event/repository"
	"github.com/GoGroup/Movie-and-events/event/service"
	schrep "github.com/GoGroup/Movie-and-events/schedule/repository"
	schser "github.com/GoGroup/Movie-and-events/schedule/service"
)

// func TestAdminEventNew(t *testing.T) {

// 	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

// 	eventRepo := repository.NewMockEventRepo(nil)
// 	eventServ := service.NewEventService(eventRepo)

// 	adminEvHandler := NewAdminHandler(tmpl,nil,nil,nil,nil ,eventServ,nil)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/admin/cinemas/events/new/", adminEvHandler.AdminEventsNew)
// 	ts := httptest.NewTLSServer(mux)
// 	defer ts.Close()

// 	tc := ts.Client()
// 	sURL := ts.URL

// 	form := url.Values{}
// 	form.Add("name", model.EvenMock.Name)
// 	form.Add("description", model.EvenMock.Description)
// 	form.Add("time", model.EvenMock.Time)
// 	form.Add("image", model.EvenMock.Image)
// 	form.Add("location", model.EvenMock.Location)
// 	resp, err := tc.PostForm(sURL+"/admin/cinemas/events/new/", form)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !bytes.Contains(body, []byte("Mock Event 01")) {
// 		t.Errorf("want body to contain %q", body)
// 	}

// }

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
func TestAdminSchedule(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../../view/template/*"))

	schdlRepo := schrep.NewMockScheduleRepo(nil)
	schdlServ := schser.NewScheduleService(schdlRepo)

	adminEvHandler := NewAdminHandler(tmpl, nil, nil, schdlServ, nil, nil, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cinemas/schedule/1", adminEvHandler.AdminSchedule)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	URL := ts.URL

	resp, err := tc.Get(URL + "/admin/cinemas/schedule/")
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

	if !bytes.Contains(body, []byte("Monday")) {
		t.Errorf("want body to contain %q", body)
	}

}
