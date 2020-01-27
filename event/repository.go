package event

import (
	"github.com/GoGroup/Movie-and-events/model"
)

// EventRepository specifies food menu event database operations
type EventRepository interface {
	Events() ([]model.Event, []error)
	Event(id uint) (*model.Event, []error)
	UpdateEvent(event *model.Event) (*model.Event, []error)
	DeleteEvent(id uint) (*model.Event, []error)
	StoreEvent(event *model.Event) (*model.Event, []error)
	EventExists(eventName string) bool
}
