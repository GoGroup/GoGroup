package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/kalkidm19/Event/entity"
	"github.com/kalkidm19/Event/menu"
)

// MockEventRepo implements the menu.EventRepository interface
type MockEventRepo struct {
	conn *gorm.DB
}

// NewMockEventRepo will create a new object of MockEventRepo
func NewMockEventRepo(db *gorm.DB) menu.EventRepository {
	return &MockEventRepo{conn: db}
}

// Events returns all fake events
func (mEveRepo *MockEventRepo) Events() ([]entity.Event, []error) {
	ctgs := []entity.Event{entity.EventMock}
	return ctgs, nil
}

// Event retrieve a fake eventy with id 1
func (mEveRepo *MockEventRepo) Event(id uint) (*entity.Event, []error) {
	ctg := entity.EventMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateEvent updates a given fake event
func (mEveRepo *MockEventRepo) UpdateEvent(event *entity.Event) (*entity.Event, []error) {
	eve := entity.EventMock
	return &eve, nil
}

// DeleteEvent deletes a given event from the database
func (mEveRepo *MockEventRepo) DeleteEvent(id uint) (*entity.Event, []error) {
	eve := entity.EventMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &eve, nil
}

// StoreEvent stores a given mock event
func (mEveRepo *MockEventRepo) StoreEvent(event *entity.Event) (*entity.Event, []error) {
	eve := event
	return eve, nil
}
