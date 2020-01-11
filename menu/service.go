package menu

import "github.com/kalkidm19/Event/entity"

// EventService specifies food menu event services
type EventService interface {
	Events() ([]entity.Event, []error)
	Event(id uint) (*entity.Event, []error)
	UpdateEvent(event *entity.Event) (*entity.Event, []error)
	DeleteEvent(id uint) (*entity.Event, []error)
	StoreEvent(event *entity.Event) (*entity.Event, []error)
}
