package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/kalkidm19/Event/delivery/http/handler"
	"github.com/kalkidm19/Event/entity"
	mrepim "github.com/kalkidm19/Event/menu/repository"
	msrvim "github.com/kalkidm19/Event/menu/service"
	"github.com/kalkidm19/Event/rtoken"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	urepimp "github.com/kalkidm19/Event/user/repository"
	usrvimp "github.com/kalkidm19/Event/user/service"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.User{}, &entity.Role{}, &entity.Session{}, &entity.Event{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {

	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	tmpl := template.Must(template.ParseGlob("ui/templates/*"))

	dbconn, err := gorm.Open("postgres", "postgres://postgres:1234567890@localhost/restaurantdb?sslmode=disable")
	//createTables(dbconn)
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	sessionRepo := urepimp.NewSessionGormRepo(dbconn)
	sessionSrv := usrvimp.NewSessionService(sessionRepo)

	eventRepo := mrepim.NewEventGormRepo(dbconn)
	eventServ := msrvim.NewEventService(eventRepo)

	userRepo := urepimp.NewUserGormRepo(dbconn)
	userServ := usrvimp.NewUserService(userRepo)

	roleRepo := urepimp.NewRoleGormRepo(dbconn)
	roleServ := usrvimp.NewRoleService(roleRepo)

	ach := handler.NewAdminEventHandler(tmpl, eventServ, csrfSignKey)
	mh := handler.NewMenuHandler(tmpl, eventServ, csrfSignKey)

	sess := configSess()
	uh := handler.NewUserHandler(tmpl, userServ, sessionSrv, roleServ, sess, csrfSignKey)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", mh.Index)
	http.HandleFunc("/detail/", mh.Detail)
	http.Handle("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(mh.Admin))))

	http.Handle("/admin/events", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminEvents))))
	http.Handle("/admin/events/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminEventsNew))))
	http.Handle("/admin/events/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminEventsUpdate))))
	http.Handle("/admin/events/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminEventsDelete))))

	http.HandleFunc("/login", uh.Login)
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/signup", uh.Signup)

	http.ListenAndServe(":8282", nil)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
