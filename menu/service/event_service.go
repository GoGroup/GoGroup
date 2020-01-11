package service

import (
	"github.com/kalkidm19/Event/entity"
	"github.com/kalkidm19/Event/menu"
)

// EventService implements menu.EventService interface
type EventService struct {
	eventRepo menu.EventRepository
}

// NewEventService will create new EventService object
func NewEventService(EveRepo menu.EventRepository) menu.EventService {
	return &EventService{eventRepo: EveRepo}
}

// Events returns list of events
func (cs *EventService) Events() ([]entity.Event, []error) {

	events, errs := cs.eventRepo.Events()

	if len(errs) > 0 {
		return nil, errs
	}

	return events, nil
}

// StoreEvent persists new event information
func (cs *EventService) StoreEvent(event *entity.Event) (*entity.Event, []error) {

	eve, errs := cs.eventRepo.StoreEvent(event)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}

// Event returns a event object with a given id
func (cs *EventService) Event(id uint) (*entity.Event, []error) {

	c, errs := cs.eventRepo.Event(id)

	if len(errs) > 0 {
		return c, errs
	}

	return c, nil
}

// UpdateEvent updates a event with new data
func (cs *EventService) UpdateEvent(event *entity.Event) (*entity.Event, []error) {

	eve, errs := cs.eventRepo.UpdateEvent(event)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}

// DeleteEvent delete a event by its id
func (cs *EventService) DeleteEvent(id uint) (*entity.Event, []error) {

	eve, errs := cs.eventRepo.DeleteEvent(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}
