package service

import (
	"github.com/hannasamuel20/Movie-and-events/model"
	"github.com/hannasamuel20/Movie-and-events/schedule"
)

type ScheduleService struct {
	scheduleRepo schedule.ScheduleRepository
}

func NewScheduleService(schRepo schedule.ScheduleRepository) schedule.ScheduleService {
	return &ScheduleService{scheduleRepo: schRepo}
}

func (s *ScheduleService) Schedules() ([]model.Schedule, []error) {
	schdls, errs := s.scheduleRepo.Schedules()
	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs
}

func (s *ScheduleService) StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error) {
	schdls, errs := s.scheduleRepo.StoreSchedule(schedule)
	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs
}
