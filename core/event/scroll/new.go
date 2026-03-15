package scroll

import (
	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of an Event.
//
// You can pass zero or more EventWithOptions to optionally set certain
// attributes on the returned Event.
func New(
	opts ...types.ScrollEventWithOption,
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
// [tcell.EventScroll]
func WithTCell(
	te *tcell.EventMouse,
) types.ScrollEventWithOption {
	return func(e types.ScrollEvent) {
		x, y := te.Position()
		pos := types.Point{
			X: x, Y: y,
		}
		e.SetWhen(te.When())
		e.SetPosition(pos)
		e.SetDirection(scrollDirectionFromTCell(te.Buttons()))
	}
}
