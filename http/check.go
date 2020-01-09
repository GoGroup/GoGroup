package main

import (
	"fmt"
	"net/http"

	"github.com/GoGroup/Movie-and-events/hall/repository"
	"github.com/GoGroup/Movie-and-events/hall/service"
	"github.com/GoGroup/Movie-and-events/http/handler"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
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
	HallRepo := repository.NewHallGormRepo(db)
	Hallsr := service.NewHallService(HallRepo)
	HallHandler := handler.NewHallHandler(Hallsr)
	myRouter := httprouter.New()
	myRouter.GET("/hall", HallHandler.GetHalls)
	myRouter.POST("/halls", HallHandler.PostHall)
	http.ListenAndServe(":8080", myRouter)

}
