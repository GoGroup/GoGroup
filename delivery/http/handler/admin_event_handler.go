package handler

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Betehemgebresilasse/Event/entity"
	"github.com/Betehemgebresilasse/Event/form"
	"github.com/Betehemgebresilasse/Event/menu"
	"github.com/Betehemgebresilasse/Event/rtoken"
)

// AdminEventHandler handles event handler admin requests
type AdminEventHandler struct {
	tmpl        *template.Template
	eventSrv    menu.EventService
	csrfSignKey []byte
}

// NewAdminEventHandler initializes and returns new AdminEventHandler
func NewAdminEventHandler(t *template.Template, cs menu.EventService, csKey []byte) *AdminEventHandler {
	return &AdminEventHandler{tmpl: t, eventSrv: cs, csrfSignKey: csKey}
}

// AdminEvents handle requests on route /admin/events
func (ach *AdminEventHandler) AdminEvents(w http.ResponseWriter, r *http.Request) {
	events, errs := ach.eventSrv.Events()
	if errs != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
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
	ach.tmpl.ExecuteTemplate(w, "admin.eveeg.layout", tmplData)
}

// AdminEventsNew hanlde requests on route /admin/events/new
func (ach *AdminEventHandler) AdminEventsNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		newEveForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.eveeg.new.layout", newEveForm)
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		newEveForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		newEveForm.Required("evename", "evedesc", "eveloc", "evetime")
		newEveForm.MinLength("evedesc", 10)
		newEveForm.MinLength("eveloc", 10)
		newEveForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !newEveForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "admin.eveeg.new.layout", newEveForm)
			return
		}
		mf, fh, err := r.FormFile("eveimg")
		if err != nil {
			newEveForm.VErrors.Add("eveimg", "File error")
			ach.tmpl.ExecuteTemplate(w, "admin.eveeg.new.layout", newEveForm)
			return
		}
		defer mf.Close()
		ctg := &entity.Event{
			Name:        r.FormValue("evename"),
			Description: r.FormValue("evedesc"),
			Location:    r.FormValue("eveloc"),
			Time:        r.FormValue("evetime"),
			Image:       fh.Filename,
		}
		writeFile(&mf, fh.Filename)
		_, errs := ach.eventSrv.StoreEvent(ctg)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/admin/events", http.StatusSeeOther)
	}
}

// AdminEventsUpdate handle requests on /admin/events/update
func (ach *AdminEventHandler) AdminEventsUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		eve, errs := ach.eventSrv.Event(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		values := url.Values{}
		values.Add("eveid", idRaw)
		values.Add("evename", eve.Name)
		values.Add("evedesc", eve.Description)
		values.Add("eveimg", eve.Image)
		values.Add("eveloc", eve.Location)
		values.Add("evetime", eve.Time)
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
		ach.tmpl.ExecuteTemplate(w, "admin.eveeg.update.layout", upEveForm)
		return
	}
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		updateEveForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		updateEveForm.Required("evename", "evedesc", "eveloc", "evetime")
		updateEveForm.MinLength("evedesc", 10)
		updateEveForm.MinLength("eveloc", 10)
		updateEveForm.CSRF = token

		eveID, err := strconv.Atoi(r.FormValue("eveid"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		ctg := &entity.Event{
			ID:          uint(eveID),
			Name:        r.FormValue("evename"),
			Description: r.FormValue("evedesc"),
			Image:       r.FormValue("imgname"),
			Location:    r.FormValue("eveloc"),
			Time:        r.FormValue("evetime"),
		}
		mf, fh, err := r.FormFile("eveimg")
		if err == nil {
			ctg.Image = fh.Filename
			err = writeFile(&mf, ctg.Image)
		}
		if mf != nil {
			defer mf.Close()
		}
		_, errs := ach.eventSrv.UpdateEvent(ctg)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/events", http.StatusSeeOther)
		return
	}
}

// AdminEventsDelete handle requests on route /admin/events/delete
func (ach *AdminEventHandler) AdminEventsDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := ach.eventSrv.DeleteEvent(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/events", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)
	if err != nil {
		return err
	}
	defer image.Close()
	io.Copy(image, *mf)
	return nil
}
