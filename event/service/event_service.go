package service

import (
	"github.com/GoGroup/Movie-and-events/event"
	"github.com/GoGroup/Movie-and-events/model"
)

// EventService implements event.EventService interface
type EventService struct {
	eventRepo event.EventRepository
}

// NewEventService will create new EventService object
func NewEventService(EveRepo event.EventRepository) event.EventService {
	return &EventService{eventRepo: EveRepo}
}

// Events returns list of events
func (cs *EventService) Events() ([]model.Event, []error) {

	events, errs := cs.eventRepo.Events()

	if len(errs) > 0 {
		return nil, errs
	}

	return events, nil
}

// StoreEvent persists new event information
func (cs *EventService) StoreEvent(event *model.Event) (*model.Event, []error) {

	eve, errs := cs.eventRepo.StoreEvent(event)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}

// Event returns a event object with a given id
func (cs *EventService) Event(id uint) (*model.Event, []error) {

	c, errs := cs.eventRepo.Event(id)

	if len(errs) > 0 {
		return c, errs
	}

	return c, nil
}

// UpdateEvent updates a event with new data
func (cs *EventService) UpdateEvent(event *model.Event) (*model.Event, []error) {

	eve, errs := cs.eventRepo.UpdateEvent(event)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}

// DeleteEvent delete a event by its id
func (cs *EventService) DeleteEvent(id uint) (*model.Event, []error) {

	eve, errs := cs.eventRepo.DeleteEvent(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return eve, nil
}
func (cs *EventService) EventExists(eventName string) bool {
	exists := cs.eventRepo.EventExists(eventName)
	return exists
}
