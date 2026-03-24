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

// WithFocused sets whether the target of the event should receive or lose the
// focus.
func WithFocused(on bool) types.FocusEventWithOption {
	return func(e types.FocusEvent) {
		e.SetFocused(on)
	}
}

// WithSource modifies the returned Event, setting its source to the
// supplied value.
func WithSource(source any) types.FocusEventWithOption {
	return func(e types.FocusEvent) {
		e.SetSource(source)
	}
}
