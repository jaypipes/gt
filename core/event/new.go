package event

import (
	"time"

	"github.com/jaypipes/gt/types"
)

// New returns a new base Event.
func New(
	opts ...types.EventWithOption,
) *Event {
	e := &Event{
		when: time.Now(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// WithSource modifies the returned Event, setting its source to the
// supplied value.
func WithSource(source any) types.EventWithOption {
	return func(e types.Event) {
		e.SetSource(source)
	}
}
