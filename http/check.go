package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GoGroup/Movie-and-events/cinema/repository"
	"github.com/GoGroup/Movie-and-events/model"

	"github.com/GoGroup/Movie-and-events/cinema/service"
	usrvim "github.com/GoGroup/Movie-and-events/hall/repository"
	urepim "github.com/GoGroup/Movie-and-events/hall/service"
	"github.com/GoGroup/Movie-and-events/http/handler"
	schrep "github.com/GoGroup/Movie-and-events/schedule/repository"
	schser "github.com/GoGroup/Movie-and-events/schedule/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
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

	tmpl := template.Must(template.ParseGlob("../view/template/*"))

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

	mh := handler.NewMenuHandler(tmpl, Cinemasr)

	myRouter.GET("/adminCinemas", mh.AdminCinema)

	myRouter.GET("/admin/schedules", adminScheduleHandler.GetSchedules)
	myRouter.POST("/admin/schedule", adminScheduleHandler.PostSchedule)
	myRouter.GET("/cinema", CinemaHandler.GetCinemas)
	myRouter.POST("/cinemas", CinemaHandler.PostCinema)
	myRouter.GET("/hall", HallHandler.GetHalls)
	myRouter.POST("/halls", HallHandler.PostHall)
	http.ListenAndServe(":8080", myRouter)

}
