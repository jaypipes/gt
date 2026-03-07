package event

import "time"

// New returns a new base Event.
func New() Event {
	return Event{
		when: time.Now(),
	}
}
