package main
import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/mine/Movie-and-events/http/handler"
	"github.com/GoGroup/Movie-and-events/event/service"
	"github.com/GoGroup/Movie-and-events/event/repository"
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"


)
func TestAdminEventNew(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../view/template/*"))

	eventRepo := repository.NewMockEventRepo(nil)
	eventServ := service.NewEventService(eventRepo )

	adminEvHandler:=handler.

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cinemas/events/new/", adminEvHandler)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("name", model.EvenMock.Name)
	form.Add("description", model.EvenMock.Description)
	form.Add("time", model.EvenMock.Time)
	form.Add("image", model.EvenMock.Image)
	form.Add("location", model.EvenMock.Location)
	resp, err := tc.PostForm(sURL+"/admin/cinemas/events/new/", form)
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

	if !bytes.Contains(body, []byte("Mock Event 01")) {
		t.Errorf("want body to contain %q", body)
	}

}
func TestAdminEvent(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../view/template/*"))

	eventRepo := repository.NewMockEventRepo(nil)
	eventServ := service.NewEventService(eventRepo )

	adminEvHandler:=handler.

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/", adminEvHandler)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(url + "/admin/categories")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Category 01")) {
		t.Errorf("want body to contain %q", body)
	}

}