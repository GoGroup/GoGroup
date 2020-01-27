package main

import (
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/cinema/repository"
	"github.com/GoGroup/Movie-and-events/cinema/service"
	"github.com/GoGroup/Movie-and-events/cinev_park/http/handler"
	usrvim "github.com/GoGroup/Movie-and-events/hall/repository"
	urepim "github.com/GoGroup/Movie-and-events/hall/service"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/rtoken"

	cmrep "github.com/GoGroup/Movie-and-events/comment/repository"
	cmser "github.com/GoGroup/Movie-and-events/comment/service"

	evrep "github.com/GoGroup/Movie-and-events/event/repository"
	evser "github.com/GoGroup/Movie-and-events/event/service"

	schrep "github.com/GoGroup/Movie-and-events/schedule/repository"
	schser "github.com/GoGroup/Movie-and-events/schedule/service"

	mvrep "github.com/GoGroup/Movie-and-events/movie/repository"
	mvser "github.com/GoGroup/Movie-and-events/movie/service"

	usrep "github.com/GoGroup/Movie-and-events/user/repository"
	usser "github.com/GoGroup/Movie-and-events/user/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:Bangtan123@localhost/MovieEvent?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Hall{})
	db.AutoMigrate(&model.Cinema{})
	db.AutoMigrate(&model.Schedule{})
	db.AutoMigrate(&model.Moviem{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Session{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Booking{})
	db.AutoMigrate(&model.Event{})
	db.AutoMigrate(&model.Role{ID: 1, Name: "USER"})
	db.AutoMigrate(&model.Role{ID: 2, Name: "ADMIN"})
	tmpl := template.Must(template.ParseGlob("view/template/*"))

	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	userRepo := usrep.NewUserGormRepo(db)
	userService := usser.NewUserService(userRepo)

	sessionRepo := usrep.NewSessionGormRepo(db)
	sessionService := usser.NewSessionService(sessionRepo)

	roleRepo := usrep.NewRoleGormRepo(db)
	roleService := usser.NewRoleService(roleRepo)

	scheduleRepo := schrep.NewScheduleGormRepo(db)
	scheduleService := schser.NewScheduleService(scheduleRepo)

	HallRepo := usrvim.NewHallGormRepo(db)
	Hallsr := urepim.NewHallService(HallRepo)

	EventRepo := evrep.NewEventGormRepo(db)
	EventSer := evser.NewEventService(EventRepo)

	CommentRepo := cmrep.NewCommentGormRepo(db)
	CommentSer := cmser.NewCommentService(CommentRepo)

	CinemaRepo := repository.NewCinemaGormRepo(db)
	Cinemasr := service.NewCinemaService(CinemaRepo)

	MovieRepo := mvrep.NewMovieGormRepo(db)
	Moviesr := mvser.NewMovieService(MovieRepo)

	uh := handler.NewUserHandler(tmpl, userService, sessionService, roleService, csrfSignKey)

	mh := handler.NewMenuHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr, CommentSer, EventSer)
	ah := handler.NewAdminHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr, EventSer, csrfSignKey)

	fs := http.FileServer(http.Dir("view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/admin/cinemas", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminCinema))))
	// http.HandleFunc("/adminCinemas/adminSchedule/{hId}", ah.AdminSchedule)
	http.Handle("/admin/cinemas/schedule/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminSchedule))))
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/delete/{sId}", ah.AdminScheduleDelete)
	http.Handle("/admin/cinemas/schedule/delete/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminScheduleDelete))))
	http.Handle("/admin/cinemas/halls/edit/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminHalls))))
	http.Handle("/admin/cinemas/halls/new/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminHallsNew))))

	http.Handle("/admin/cinemas/halls/delete/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminDeleteHalls))))
	http.Handle("/admin/cinemas/schedule/new/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.NewAdminScheduleHandler))))
	http.Handle("/admin/cinemas/events/new/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminEventsNew))))
	http.Handle("/admin/cinemas/events/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminEventList))))
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/new/", ah.NewAdminSchedule)
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/new/", ah.NewAdminSchedulePost)
	http.Handle("/home", uh.Authenticated(http.HandlerFunc(mh.Index)))
	http.Handle("/movies", uh.Authenticated(http.HandlerFunc(mh.Movies)))
	//http.HandleFunc("/movie/{mId}", mh.EachMovieHandler)
	http.Handle("/movie/", uh.Authenticated(http.HandlerFunc(mh.EachMovieHandler)))
	http.Handle("/movie/nowshowing/", uh.Authenticated(http.HandlerFunc(mh.EachNowShowing)))
	http.Handle("/theaters", uh.Authenticated(http.HandlerFunc(mh.Theaters)))
	http.Handle("/events", uh.Authenticated(http.HandlerFunc(mh.EventList)))

	//http.HandleFunc("/theater/schedule/{cName}/{cId}", mh.TheaterSchedule)
	http.Handle("/theater/schedule/", uh.Authenticated(http.HandlerFunc(mh.TheaterSchedule)))
	//	http.HandleFunc("/", uh.Login)
	http.Handle("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminCinema))))
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/login", uh.Login)
	http.HandleFunc("/signup", uh.SignUp)

	http.ListenAndServe(":8181", nil)

}
