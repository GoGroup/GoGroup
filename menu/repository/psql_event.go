package repository

import (
	"database/sql"
	"errors"

	"github.com/kalkidm19/Event/entity"
)

// EventRepositoryImpl implements the menu.EventRepository interface
type EventRepositoryImpl struct {
	conn *sql.DB
}

// NewEventRepositoryImpl will create an object of PsqlEventRepository
func NewEventRepositoryImpl(Conn *sql.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{conn: Conn}
}

//Events returns all events from the database
func (cri *EventRepositoryImpl) Events() ([]entity.Event, error) {

	rows, err := cri.conn.Query("SELECT * FROM events;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Event{}

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Time, &event.Image)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, event)
	}

	return ctgs, nil
}

// Event returns a event with a given id
func (cri *EventRepositoryImpl) Event(id uint) (entity.Event, error) {

	row := cri.conn.QueryRow("SELECT * FROM events WHERE id = $1", id)

	c := entity.Event{}

	err := row.Scan(&c.ID, &c.Name, &c.Description, &c.Location, &c.Time, &c.Image)
	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateEvent updates a given object with a new data
func (cri *EventRepositoryImpl) UpdateEvent(c entity.Event) error {

	_, err := cri.conn.Exec("UPDATE events SET name=$1,description=$2,location=$3,time=$4, image=$5 WHERE id=$6", c.Name, c.Description, c.Location, c.Time, c.Image, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteEvent removes a event from a database by its id
func (cri *EventRepositoryImpl) DeleteEvent(id uint) error {

	_, err := cri.conn.Exec("DELETE FROM events WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreEvent stores new event information to database
func (cri *EventRepositoryImpl) StoreEvent(c entity.Event) error {

	_, err := cri.conn.Exec("INSERT INTO events (name,description,image) values($1, $2, $3)", c.Name, c.Description, c.Location, c.Time, c.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
