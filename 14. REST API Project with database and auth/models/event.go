package models

import (
	"time"

	"e.com/events-rest-api/db"
)

type Event struct {
	ID          int64     // sql ID is int64 so this field's datatype was changed to int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	// ? is special character that prevents sql injections
	// ? is understood by mostly all SQL packages by go
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	// Prepare() put the sql query in RAM so its very fast
	// mostly used when sql query is very large or complex
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// close the statement at last
	defer stmt.Close()

	// execute query to create a entry in databse and provide values for all '?' from event instance
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	// take the id of the event entry created in database
	id, err := result.LastInsertId()

	// in the event struct instance add the ID which came from database
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// get back all all event rows from db
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// all database rows will be put into this slice
	var events []Event

	// iterate over rows and append the populated event struct instance
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}
