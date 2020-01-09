package main

import (
	"fmt"
	"net/http"

	"gitlab.com/username/excercise/Project-GO/Movie-and-events/hall/service"

	"gitlab.com/username/excercise/Project-GO/Movie-and-events/hall/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/username/excercise/Project-GO/Movie-and-events/http/handler"
	"gitlab.com/username/excercise/Project-GO/Movie-and-events/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Bangtan123"
	dbname   = "movieevent"
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
	HallRepo := repository.NewHallGormRepo(db)
	Hallsr := service.NewHallService(HallRepo)
	HallHandler := handler.NewHallHandler(Hallsr)
	myRouter := httprouter.New()
	myRouter.GET("/hall", HallHandler.GetHalls)
	myRouter.POST("/halls", HallHandler.PostHall)
	http.ListenAndServe(":8080", myRouter)

}
