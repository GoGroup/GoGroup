package repository

import (
	"errors"

	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
	"github.com/jinzhu/gorm"
)

type MockScheduleRepo struct {
	conn *gorm.DB
}

func NewMockScheduleRepo(db *gorm.DB) schedule.ScheduleRepository {
	return &ScheduleGormRepo{conn: db}
}

func (scheduleRepo *MockScheduleRepo) Schedules() ([]model.Schedule, []error) {
	schdls := []model.Schedule{model.ScheduleMock}

	return schdls, nil

}
func (scheduleRepo *MockScheduleRepo) HallSchedules(id uint, day string) ([]model.Schedule, []error) {

	schdls := []model.Schedule{model.ScheduleMock}

	return schdls, nil

}

func (scheduleRepo *MockScheduleRepo) StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error) {
	schdl := &model.ScheduleMock

	return schdl, nil
}
func (schRepo *MockScheduleRepo) UpdateSchedules(schedule *model.Schedule) (*model.Schedule, []error) {
	schdl := &model.ScheduleMock

	return schdl, nil
}
func (schRepo *MockScheduleRepo) UpdateSchedulesBooked(schedule *model.Schedule, Amount uint) *model.Schedule {
	schdl := &model.ScheduleMock

	return schdl
}

// DeleteComment deletes a given customer comment from the database
func (schRepo *MockScheduleRepo) DeleteSchedules(id uint) (*model.Schedule, []error) {
	schdl := &model.ScheduleMock

	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return schdl, nil
}
func (schRepo *MockScheduleRepo) Schedule(id uint) (*model.Schedule, []error) {
	schdl := &model.ScheduleMock
	if id == 1 {
		return schdl, nil
	}
	return nil, []error{errors.New("Not found")}
}
func (schRepo *MockScheduleRepo) ScheduleHallDay(hallid uint, day string) ([]model.Schedule, []error) {
	schdl := []model.Schedule{model.ScheduleMock}
	if hallid == 1 {
		return schdl, nil
	}
	return nil, []error{errors.New("Not found")}
}
