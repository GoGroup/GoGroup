package repository

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
	"github.com/jinzhu/gorm"
)

type ScheduleGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo returns new object of CommentGormRepo
func NewScheduleGormRepo(db *gorm.DB) schedule.ScheduleRepository {
	return &ScheduleGormRepo{conn: db}
}

// Comments returns all customer comments stored in the database
func (scheduleRepo *ScheduleGormRepo) Schedules() ([]model.Schedule, []error) {
	schdls := []model.Schedule{}
	errs := scheduleRepo.conn.Find(&schdls).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs
}

// StoreComment stores a given customer comment in the database
func (scheduleRepo *ScheduleGormRepo) StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error) {
	schdl := schedule
	errs := scheduleRepo.conn.Create(schdl).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return schdl, errs
}
