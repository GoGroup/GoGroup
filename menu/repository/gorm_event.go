package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kalkidm19/Event/entity"
	"github.com/kalkidm19/Event/menu"
)

// EventGormRepo implements the menu.EventRepository interface
type EventGormRepo struct {
	conn *gorm.DB
}

// NewEventGormRepo will create a new object of EventGormRepo
func NewEventGormRepo(db *gorm.DB) menu.EventRepository {
	return &EventGormRepo{conn: db}
}

// Events returns all events stored in the database
func (cRepo *EventGormRepo) Events() ([]entity.Event, []error) {
	ctgs := []entity.Event{}
	errs := cRepo.conn.Find(&ctgs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ctgs, errs
}

// Event retrieve a event from the database by its id
func (cRepo *EventGormRepo) Event(id uint) (*entity.Event, []error) {
	ctg := entity.Event{}
	errs := cRepo.conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

// UpdateEvent updates a given event in the database
func (cRepo *EventGormRepo) UpdateEvent(event *entity.Event) (*entity.Event, []error) {
	eve := event
	errs := cRepo.conn.Save(eve).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return eve, errs
}

// DeleteEvent deletes a given event from the database
func (cRepo *EventGormRepo) DeleteEvent(id uint) (*entity.Event, []error) {
	eve, errs := cRepo.Event(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = cRepo.conn.Delete(eve, eve.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return eve, errs
}

// StoreEvent stores a given event in the database
func (cRepo *EventGormRepo) StoreEvent(event *entity.Event) (*entity.Event, []error) {
	eve := event
	errs := cRepo.conn.Create(eve).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return eve, errs
}
