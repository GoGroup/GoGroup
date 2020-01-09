package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hannasamuel20/Movie-and-events/model"
)

const allschedulesurl = "http://localhost:8181/admin/schedules"

func GetSchedules() ([]model.Schedule, error, error) {

	res, err := http.Get(allschedulesurl)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var schdls []model.Schedule
	err2 := json.Unmarshal([]byte(body), &schdls)
	if err2 != nil {
		fmt.Println("whoops:", err2)
	}

	return schdls, err, err2

}
