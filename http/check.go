package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/cinema/repository"
	"github.com/GoGroup/Movie-and-events/cinema/service"
	usrvim "github.com/GoGroup/Movie-and-events/hall/repository"
	urepim "github.com/GoGroup/Movie-and-events/hall/service"
	"github.com/GoGroup/Movie-and-events/http/handler"
	"github.com/GoGroup/Movie-and-events/model"
	schrep "github.com/GoGroup/Movie-and-events/schedule/repository"
	schser "github.com/GoGroup/Movie-and-events/schedule/service"

	mvrep "github.com/GoGroup/Movie-and-events/movie/repository"
	mvser "github.com/GoGroup/Movie-and-events/movie/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"

	password = "Bangtan123"
	dbname   = "MovieEvent"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.Hall{})
	db.AutoMigrate(&model.Cinema{})
	db.AutoMigrate(&model.Schedule{})
	db.AutoMigrate(&model.Moviem{})

	tmpl := template.Must(template.ParseGlob("./view/template/*"))

	myRouter := httprouter.New()

	scheduleRepo := schrep.NewScheduleGormRepo(db)
	scheduleService := schser.NewScheduleService(scheduleRepo)
	adminScheduleHandler := handler.NewAdminScheduleHandler(scheduleService)

	HallRepo := usrvim.NewHallGormRepo(db)
	Hallsr := urepim.NewHallService(HallRepo)
	HallHandler := handler.NewHallHandler(Hallsr)

	CinemaRepo := repository.NewCinemaGormRepo(db)
	Cinemasr := service.NewCinemaService(CinemaRepo)
	CinemaHandler := handler.NewCinemaHandler(Cinemasr)

	MovieRepo := mvrep.NewMovieGormRepo(db)
	Moviesr := mvser.NewMovieService(MovieRepo)
	MovieHandler := handler.NewMovieHander(Moviesr)

	mh := handler.NewMenuHandler(tmpl, Cinemasr, Hallsr, scheduleService, Moviesr)

	//myRouter.ServeFiles("/assets/css/*filepath", http.Dir("../view/assets"))
	myRouter.ServeFiles("/assets/*filepath", http.Dir("../view/assets"))

	// fs := http.FileServer(http.Dir("../view/assetts"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	myRouter.GET("/adminCinemas", mh.AdminCinema)
	myRouter.GET("/adminCinemas/adminSchedule/:hId", mh.AdminSchedule)
	myRouter.GET("/adminCinemas/adminSchedule/:hId/delete/:sId", mh.AdminScheduleDelete)
	myRouter.GET("/adminCinemas/adminSchedule/:hId/new/", mh.NewAdminSchedule)
	myRouter.POST("/adminCinemas/adminSchedule/:hId/new/", mh.NewAdminSchedulePost)
	myRouter.GET("/", mh.Index)
	myRouter.GET("/movies", mh.Movies)
	myRouter.GET("/movie/:mId", mh.EachMovieHandler)
	myRouter.GET("/theaters", mh.Theaters)
	myRouter.GET("/theater/schedule/:cName/:cId", mh.TheaterSchedule)

	myRouter.GET("/api/schedules", adminScheduleHandler.GetSchedules)
	myRouter.GET("/api/cinemaschedules/:id/:day", adminScheduleHandler.GetSchedulesCinemaDay)
	myRouter.GET("/api/hallschedules/:hid/:day", adminScheduleHandler.GetSchedulesHallDay)
	myRouter.GET("/api/schedule/:id", adminScheduleHandler.GetSingleSchedule)
	myRouter.DELETE("/api/schedule/:id", adminScheduleHandler.DeleteSchedule)
	myRouter.PUT("/api/schedule/:id", adminScheduleHandler.UpdateSchedule)
	myRouter.POST("/api/schedule", adminScheduleHandler.PostSchedule)
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
