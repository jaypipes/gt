package focus

import (
	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of an Event.
//
// You can pass zero or more EventWithOptions to optionally set certain
// attributes on the returned Event.
func New(
	opts ...types.FocusEventWithOption,
) *Event {
	e := &Event{
		Event: event.New(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// WithEnabled sets whether the target of the event should receive or lose the
// focus.
func WithEnabled(enabled bool) types.FocusEventWithOption {
	return func(e types.FocusEvent) {
		e.SetEnabled(enabled)
	}
}

// WithProducer modifies the returned Event, setting its producer to the
// supplied value.
func WithProducer(producer any) types.FocusEventWithOption {
	return func(e types.FocusEvent) {
		e.SetProducer(producer)
	}
}
