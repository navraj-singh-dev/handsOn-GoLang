package models

import "time"

type Event struct {
	// struct tags are supported by GIN
	// here i am using tags to make some of the fields as must 'required',
	// which a user must send in POST request other-wise error will be sent as response
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
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
