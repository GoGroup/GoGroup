package schedule

import "github.com/GoGroup/Movie-and-events/model"

type ScheduleRepository interface {
	Schedules() ([]model.Schedule, []error)
	StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error)
	HallSchedules(id uint, day string) ([]model.Schedule, []error)
	Schedule(id uint) (*model.Schedule, []error)
	UpdateSchedules(hall *model.Schedule) (*model.Schedule, []error)
	DeleteSchedules(id uint) (*model.Schedule, []error)
}
