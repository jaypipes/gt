package key

import (
	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/core/key"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of an Event.
//
// You can pass zero or more EventWithOptions to optionally set certain
// attributes on the returned Event.
func New(
	opts ...types.KeyPressEventWithOption,
) *Event {
	e := &Event{
		Event: event.New(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// WithTCell modifies the returned Event to base on the supplied
// [tcell.EventKey]
func WithTCell(
	te *tcell.EventKey,
) types.KeyPressEventWithOption {
	return func(e types.KeyPressEvent) {
		k := key.New(te)
		e.SetKey(k)
		e.SetWhen(te.When())
	}
}
