package models

import "time"

type Event struct {
	ID          int       `json:"ID,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Location    string    `json:"location,omitempty"`
	DateTime    time.Time `json:"dateTime,omitempty"`
	UserID      int       `json:"userID,omitempty"`
}

var events []Event

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
