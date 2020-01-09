package schedule

import "github.com/hannasamuel20/Movie-and-events/model"

type ScheduleService interface {
	Schedules() ([]model.Schedule, []error)
	StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error)
}