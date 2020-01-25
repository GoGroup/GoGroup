package main

import (
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/http/util"

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
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Session{})
	db.AutoMigrate(&model.Role{})
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
	ah := handler.NewAdminHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr)

	fs := http.FileServer(http.Dir("view/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/adminCinemas", ah.AdminCinema)
	// http.HandleFunc("/adminCinemas/adminSchedule/{hId}", ah.AdminSchedule)
	http.HandleFunc("/adminCinemas/adminSchedule/", ah.AdminSchedule)
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/delete/{sId}", ah.AdminScheduleDelete)
	http.HandleFunc("/adminCinemas/adminSchedule/delete/", ah.AdminScheduleDelete)
	http.HandleFunc("/adminCinemas/adminHalls/edit/", ah.AdminHalls)
	http.HandleFunc("/adminCinemas/adminHalls/new/", ah.AdminHallsNew)
	http.HandleFunc("/adminCinemas/adminHalls/delete/", ah.AdminDeleteHalls)
	http.HandleFunc("/adminCinemas/adminSchedule/new/", ah.NewAdminScheduleHandler)
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/new/", ah.NewAdminSchedule)
	//http.HandleFunc("/adminCinemas/adminSchedule/{hId}/new/", ah.NewAdminSchedulePost)

	http.HandleFunc("/home", mh.Index)
	http.HandleFunc("/movies", mh.Movies)
	//http.HandleFunc("/movie/{mId}", mh.EachMovieHandler)
	http.HandleFunc("/movie/", mh.EachMovieHandler)
	http.HandleFunc("/theaters", mh.Theaters)
	//http.HandleFunc("/theater/schedule/{cName}/{cId}", mh.TheaterSchedule)
	http.HandleFunc("/theater/schedule/", mh.TheaterSchedule)
	http.HandleFunc("/", uh.Login)
	http.Handle("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(ah.AdminCinema))))
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/login", uh.Login)
	http.HandleFunc("/signup", uh.SignUp)

	http.ListenAndServe(":8181", nil)

}
