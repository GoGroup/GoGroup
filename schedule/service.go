package schedule

import "github.com/GoGroup/Movie-and-events/model"

type ScheduleService interface {
	Schedules() ([]model.Schedule, []error)
	StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error)
	HallSchedules(id uint, day string) ([]model.Schedule, []error)
	Schedule(id uint) (*model.Schedule, []error)
	UpdateSchedules(hall *model.Schedule) (*model.Schedule, []error)
	UpdateSchedulesBooked(user *model.Schedule, Amount uint) *model.Schedule
	DeleteSchedules(id uint) (*model.Schedule, []error)
	ScheduleHallDay(hallid uint, day string) ([]model.Schedule, []error)
}
