package service

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
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

func (s *ScheduleService) HallSchedules(id uint, day string) ([]model.Schedule, []error) {
	schdls, errs := s.scheduleRepo.HallSchedules(id, day)
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
func (ss *ScheduleService) UpdateSchedules(schedule *model.Schedule) (*model.Schedule, []error) {
	schdls, errs := ss.scheduleRepo.UpdateSchedules(schedule)
	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs
}
func (ss *ScheduleService) DeleteSchedules(id uint) (*model.Schedule, []error) {
	schdls, errs := ss.scheduleRepo.DeleteSchedules(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs
}

func (ss *ScheduleService) Schedule(id uint) (*model.Schedule, []error) {
	schdls, errs := ss.scheduleRepo.Schedule(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs
}
