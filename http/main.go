package main

import (
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/http/util"

	"github.com/GoGroup/Movie-and-events/cinema/repository"
	"github.com/GoGroup/Movie-and-events/cinema/service"
	usrvim "github.com/GoGroup/Movie-and-events/hall/repository"
	urepim "github.com/GoGroup/Movie-and-events/hall/service"
	"github.com/GoGroup/Movie-and-events/http/handler"
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
	"github.com/julienschmidt/httprouter"
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
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.Role{ID: 1, Name: "USER"})
	db.AutoMigrate(&model.Role{ID: 2, Name: "ADMIN"})
	tmpl := template.Must(template.ParseGlob("../view/template/*"))

	myRouter := httprouter.New()
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	userRepo := usrep.NewUserGormRepo(db)
	userService := usser.NewUserService(userRepo)

	sessionRepo := usrep.NewSessionGormRepo(db)
	sessionService := usser.NewSessionService(sessionRepo)

	roleRepo := usrep.NewRoleGormRepo(db)
	roleService := usser.NewRoleService(roleRepo)

	scheduleRepo := schrep.NewScheduleGormRepo(db)
	scheduleService := schser.NewScheduleService(scheduleRepo)
	sheduleHandler := handler.NewScheduleHandler(scheduleService)

	HallRepo := usrvim.NewHallGormRepo(db)
	Hallsr := urepim.NewHallService(HallRepo)
	HallHandler := handler.NewHallHandler(Hallsr)

	CinemaRepo := repository.NewCinemaGormRepo(db)
	Cinemasr := service.NewCinemaService(CinemaRepo)
	CinemaHandler := handler.NewCinemaHandler(Cinemasr)

	MovieRepo := mvrep.NewMovieGormRepo(db)
	Moviesr := mvser.NewMovieService(MovieRepo)
	MovieHandler := handler.NewMovieHander(Moviesr)

	uh := handler.NewUserHandler(tmpl, userService, sessionService, roleService, csrfSignKey)

	mh := handler.NewMenuHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr)
	ah := handler.NewAdminHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr)

	myRouter.ServeFiles("/assets/*filepath", http.Dir("../view/assets"))

	myRouter.GET("/adminCinemas", ah.AdminCinema)
	myRouter.GET("/adminCinemas/adminSchedule/:hId", ah.AdminSchedule)
	myRouter.GET("/adminCinemas/adminSchedule/:hId/delete/:sId", ah.AdminScheduleDelete)
	myRouter.GET("/adminCinemas/adminSchedule/:hId/new/", ah.NewAdminSchedule)
	myRouter.POST("/adminCinemas/adminSchedule/:hId/new/", ah.NewAdminSchedulePost)
	myRouter.GET("/home", mh.Index)
	myRouter.GET("/movies", mh.Movies)
	myRouter.GET("/movie/:mId", mh.EachMovieHandler)
	myRouter.GET("/theaters", mh.Theaters)
	myRouter.GET("/theater/schedule/:cName/:cId", mh.TheaterSchedule)

	myRouter.GET("/", uh.Login)
	myRouter.GET("/login", uh.Login)
	myRouter.GET("/signup", uh.SignUp)
	myRouter.POST("/login", uh.Login)
	myRouter.POST("/signup", uh.SignUp)

	myRouter.GET("/api/schedules", sheduleHandler.GetSchedules)
	myRouter.GET("/api/cinemaschedules/:id/:day", sheduleHandler.GetSchedulesCinemaDay)
	myRouter.GET("/api/hallschedules/:hid/:day", sheduleHandler.GetSchedulesHallDay)
	myRouter.GET("/api/schedule/:id", sheduleHandler.GetSingleSchedule)
	myRouter.DELETE("/api/schedule/:id", sheduleHandler.DeleteSchedule)
	myRouter.PUT("/api/schedule/:id", sheduleHandler.UpdateSchedule)
	myRouter.POST("/api/schedule", sheduleHandler.PostSchedule)
	myRouter.GET("/api/cinemas", CinemaHandler.GetCinemas)
	myRouter.POST("/api/cinema", CinemaHandler.PostCinema)
	myRouter.GET("/api/cinema/:id", CinemaHandler.GetSingleCinema)
	myRouter.GET("/api/hallcinema/:id", HallHandler.GetCinemaHalls)
	myRouter.GET("/api/halls", HallHandler.GetHalls)
	myRouter.GET("/api/hall/:id", HallHandler.GetSingleHall)
	myRouter.POST("/api/hall", HallHandler.PostHall)
	myRouter.GET("/api/movies", MovieHandler.GetMovies)
	myRouter.POST("/api/movie", MovieHandler.PostMovie)
	http.ListenAndServe(":8080", myRouter)

}
