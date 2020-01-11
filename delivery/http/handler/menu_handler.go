package handler

import (
	"html/template"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Betehemgebresilasse/Event/entity"
	"github.com/Betehemgebresilasse/Event/form"
	"github.com/Betehemgebresilasse/Event/menu"
	"github.com/Betehemgebresilasse/Event/rtoken"
)

// MenuHandler handles menu related requests
type MenuHandler struct {
	tmpl        *template.Template
	eventSrv    menu.EventService
	csrfSignKey []byte
}

// NewMenuHandler initializes and returns new MenuHandler
func NewMenuHandler(T *template.Template, CS menu.EventService, csKey []byte) *MenuHandler {
	return &MenuHandler{tmpl: T, eventSrv: CS, csrfSignKey: csKey}
}

// Index handles request on route /
func (mh *MenuHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	events, errs := mh.eventSrv.Events()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		Events  []entity.Event
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		Events:  events,
		CSRF:    token,
	}

	mh.tmpl.ExecuteTemplate(w, "index.layout", tmplData)
}

// Detail handles request on route /
func (mh *MenuHandler) Detail(w http.ResponseWriter, r *http.Request) {

	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		eve, errs := mh.eventSrv.Event(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		values := url.Values{}
		values.Add("eveid", idRaw)
		values.Add("evename", eve.Name)
		values.Add("evedesc", eve.Description)
		values.Add("eveloc", eve.Location)
		values.Add("evetime", eve.Time)
		values.Add("eveimg", eve.Image)
		upEveForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Event   *entity.Event
			CSRF    string
		}{
			Values:  values,
			VErrors: form.ValidationErrors{},
			Event:   eve,
			CSRF:    token,
		}
		mh.tmpl.ExecuteTemplate(w, "detail.layout", upEveForm)
		return

	}

}

// About handles requests on route /about
func (mh *MenuHandler) About(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "about.layout", tmplData)
}

// Menu handle request on route /menu
func (mh *MenuHandler) Menu(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "menu.layout", tmplData)
}

// Contact handle request on route /Contact
func (mh *MenuHandler) Contact(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "contact.layout", tmplData)
}

// Admin handle request on route /admin
func (mh *MenuHandler) Admin(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "admin.index.layout", tmplData)
}
