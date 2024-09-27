package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

// method to save a event struct instance in a slice
// just testing out simply, SQL will be included later
var events = []Event{}

func (e Event) Save() {
	// just temporary testing, SQL will be added later
	events = append(events, e)
}

// method to return all events (which are in slice)
func GetAllEvents() []Event {
	return events
}
