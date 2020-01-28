package handler

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	schrep "github.com/GoGroup/Movie-and-events/schedule/repository"
	schser "github.com/GoGroup/Movie-and-events/schedule/service"
)

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
