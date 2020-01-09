package main

import (
	"fmt"
	"net/http"

	"github.com/GoGroup/Movie-and-events/http/handler"
	"github.com/GoGroup/Movie-and-events/model"
	schrep "github.com/GoGroup/Movie-and-events/schedule/repository"
	schser "github.com/GoGroup/Movie-and-events/schedule/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:admin@localhost/MovieEvent?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()
	dbconn.AutoMigrate(&model.Schedule{})

	// roleRepo := urepim.NewRoleGormRepo(dbconn)
	// roleSrv := usrvim.NewRoleService(roleRepo)
	// adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)
	scheduleRepo := schrep.NewScheduleGormRepo(dbconn)
	scheduleService := schser.NewScheduleService(scheduleRepo)
	adminScheduleHandler := handler.NewAdminScheduleHandler(scheduleService)
	router := httprouter.New()
	router.GET("/admin/schedules", adminScheduleHandler.GetSchedules)
	// router.GET("/v1/admin/comments/:id", adminCommentHandler.GetSingleComment)
	// router.GET("/v1/admin/comments", adminCommentHandler.GetComments)
	// router.PUT("/v1/admin/comments/:id", adminCommentHandler.PutComment)
	router.POST("/admin/schedule", adminScheduleHandler.PostSchedule)
	// router.DELETE("/v1/admin/comments/:id", adminCommentHandler.DeleteComment)
	fmt.Println("called from main2.go")

	http.ListenAndServe(":8181", router)
}
