package main
import (
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
func TestAdminCategoriesNew(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	eventRepo := repository.NewMockEventRepo(nil)
	eventServ := service.NewEventService(eventRepo )

	menuEvHandler := handler.MenuHandler(tmpl,eventServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/categories/new", menuEvHandler)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("name", entity.CategoryMock.Name)
	form.Add("Description", entity.CategoryMock.Description)
	form.Add("Image", entity.CategoryMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/categories/new", form)
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

	if !bytes.Contains(body, []byte("Mock Category 01")) {
		t.Errorf("want body to contain %q", body)
	}

}
