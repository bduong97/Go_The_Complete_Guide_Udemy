package models

import (
	"time"
	"example.com/rest-api/db"
)


type Event struct {
	ID          int64 //optional
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() error {
	// TODO: add  it to a database
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?, ?, ?, ?, ?)` //question mark notation protects against sequel injection attacks
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err 
}

func GetAllEvents() []Event {
	return events
}
