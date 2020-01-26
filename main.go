package main

import (
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/cinev_park/http/util"

	"github.com/GoGroup/Movie-and-events/cinema/repository"
	"github.com/GoGroup/Movie-and-events/cinema/service"
	"github.com/GoGroup/Movie-and-events/cinev_park/http/handler"
	usrvim "github.com/GoGroup/Movie-and-events/hall/repository"
	urepim "github.com/GoGroup/Movie-and-events/hall/service"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/rtoken"

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
	db, err := gorm.Open("postgres", util.DBConnectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Hall{})
	db.AutoMigrate(&model.Cinema{})
	db.AutoMigrate(&model.Schedule{})
	db.AutoMigrate(&model.Moviem{})
	db.AutoMigrate(&model.Session{})
	db.AutoMigrate(&model.User{})
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

	CinemaRepo := repository.NewCinemaGormRepo(db)
	Cinemasr := service.NewCinemaService(CinemaRepo)

	MovieRepo := mvrep.NewMovieGormRepo(db)
	Moviesr := mvser.NewMovieService(MovieRepo)

	uh := handler.NewUserHandler(tmpl, userService, sessionService, roleService, csrfSignKey)

	mh := handler.NewMenuHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr)
	ah := handler.NewAdminHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr, csrfSignKey)

	fs := http.FileServer(http.Dir("view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/adminCinemas", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminCinema))))
	// http.HandleFunc("/adminCinemas/adminSchedule/{hId}", ah.AdminSchedule)
	http.Handle("/adminCinemas/adminSchedule/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminSchedule))))
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/delete/{sId}", ah.AdminScheduleDelete)
	http.Handle("/adminCinemas/adminSchedule/delete/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminScheduleDelete))))
	http.Handle("/adminCinemas/adminHalls/edit/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminHalls))))
	http.Handle("/adminCinemas/adminHalls/new/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminHallsNew))))
	http.Handle("/adminCinemas/adminHalls/delete/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminDeleteHalls))))
	http.Handle("/adminCinemas/adminSchedule/new/", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.NewAdminScheduleHandler))))
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/new/", ah.NewAdminSchedule)
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/new/", ah.NewAdminSchedulePost)
	http.Handle("/home", uh.Authenticated(http.HandlerFunc(mh.Index)))
	http.Handle("/movies", uh.Authenticated(http.HandlerFunc(mh.Movies)))
	//http.HandleFunc("/movie/{mId}", mh.EachMovieHandler)
	http.HandleFunc("/movie/", mh.EachMovieHandler)
	http.HandleFunc("/theaters", mh.Theaters)
	//http.HandleFunc("/theater/schedule/{cName}/{cId}", mh.TheaterSchedule)
	http.HandleFunc("/theater/schedule/", mh.TheaterSchedule)
	//	http.HandleFunc("/", uh.Login)
	http.Handle("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminCinema))))
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/login", uh.Login)
	http.HandleFunc("/signup", uh.SignUp)

	http.ListenAndServe(":8181", nil)

}
