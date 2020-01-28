package repository

import (
	"errors"
	"fmt"

	"github.com/GoGroup/Movie-and-events/event"

	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// EventGormRepo implements the event.EventRepository interface
type MockEventRepo struct {
	conn *gorm.DB
}

// NewEventGormRepo will create a new object of EventGormRepo
func NewMockEventRepo(db *gorm.DB) event.EventRepository {
	return &MockEventRepo{conn: db}
}

// Events returns all events stored in the database
func (cRepo *MockEventRepo) Events() ([]model.Event, []error) {
	fmt.Println("(((((((((((in mock))))))))))")

	ctgs := []model.Event{model.EvenMock}
	return ctgs, nil
}

// Event retrieve a event from the database by its id
func (cRepo *MockEventRepo) Event(id uint) (*model.Event, []error) {
	ctg := model.EvenMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateEvent updates a given event in the database
func (cRepo *MockEventRepo) UpdateEvent(event *model.Event) (*model.Event, []error) {
	eve := &model.EvenMock

	return eve, nil
}

// DeleteEvent deletes a given event from the database
func (cRepo *MockEventRepo) DeleteEvent(id uint) (*model.Event, []error) {
	ev := &model.EvenMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}

	return ev, nil
}

// StoreEvent stores a given event in the database
func (cRepo *MockEventRepo) StoreEvent(event *model.Event) (*model.Event, []error) {
	eve := &model.EvenMock

	return eve, nil
}
func (cRepo *MockEventRepo) EventExists(eventName string) bool {
	return true
}
