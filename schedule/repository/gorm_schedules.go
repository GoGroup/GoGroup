package repository

import (
	"fmt"

	asvv "github.com/GoGroup/Movie-and-events/hall/repository"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/schedule"
	"github.com/jinzhu/gorm"
)

type ScheduleGormRepo struct {
	conn *gorm.DB
}

func NewScheduleGormRepo(db *gorm.DB) schedule.ScheduleRepository {
	return &ScheduleGormRepo{conn: db}
}

func (scheduleRepo *ScheduleGormRepo) Schedules() ([]model.Schedule, []error) {
	schdls := []model.Schedule{}
	errs := scheduleRepo.conn.Find(&schdls).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs

}
func (scheduleRepo *ScheduleGormRepo) HallSchedules(id uint, day string) ([]model.Schedule, []error) {
	ids := []uint{}
	HallRepo := asvv.NewHallGormRepo(scheduleRepo.conn)
	hll := []model.Hall{}
	hll, err := HallRepo.CinemaHalls(id)
	if len(err) > 0 {
		return nil, err
	}
	for _, ar := range hll {
		ids = append(ids, ar.ID)
	}

	schdls := []model.Schedule{}
	fmt.Printf(day)
	errs := scheduleRepo.conn.Where("hall_id in (?) And Day=?", ids, day).Find(&schdls).GetErrors()
	//errs := scheduleRepo.conn.Joins("JOIN schedules on hall_id=halls.id AND day = ?", day).Joins("Join halls on halls.id=cinemas.id").Where("cinemas.id=?", id).Find(&schdls).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs

}

func (scheduleRepo *ScheduleGormRepo) StoreSchedule(schedule *model.Schedule) (*model.Schedule, []error) {
	schdl := schedule
	errs := scheduleRepo.conn.Create(schdl).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return schdl, errs
}
func (schRepo *ScheduleGormRepo) UpdateSchedules(schedule *model.Schedule) (*model.Schedule, []error) {
	schdl := schedule
	errs := schRepo.conn.Save(schdl).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return schdl, errs
}
func (schRepo *ScheduleGormRepo) UpdateSchedulesBooked(schedule *model.Schedule, Amount uint) *model.Schedule {
	schdl := schedule
	schRepo.conn.Model(&schdl).UpdateColumn("booked", schdl.Booked+1)

	return schdl
}

// DeleteComment deletes a given customer comment from the database
func (schRepo *ScheduleGormRepo) DeleteSchedules(id uint) (*model.Schedule, []error) {
	fmt.Println("(((((((((((((((((((((((((in delete gorm))))))))))))))))))))))")

	schdl, errs := schRepo.Schedule(id)
	fmt.Println("(((((((((((((((((((((((((in delete gorm))))))))))))))))))))))")
	if len(errs) > 0 {
		return nil, errs
	}

	errs = schRepo.conn.Delete(schdl, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return schdl, errs
}
func (schRepo *ScheduleGormRepo) Schedule(id uint) (*model.Schedule, []error) {
	schdl := model.Schedule{}
	errs := schRepo.conn.First(&schdl, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &schdl, errs
}
func (schRepo *ScheduleGormRepo) ScheduleHallDay(hallid uint, day string) ([]model.Schedule, []error) {
	schdls := []model.Schedule{}
	errs := schRepo.conn.Where("hall_id in (?) And Day=?", hallid, day).Find(&schdls).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return schdls, errs
}
