package focus

import (
	"fmt"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling focus-related events.
// Implements [types.FocusEvent].
type Event struct {
	*event.Event
	// on indicates whether the focus should be gained or lost.
	on bool
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	return fmt.Sprintf(
		"focus:%t", e.on,
	)
}

// Focused returns whether the focus should be gained or lost by the target of
// the event.
func (e *Event) Focused() bool {
	return e.on
}

// SetFocused sets whether the focus should be gained or lost by the target of
// the event.
func (e *Event) SetFocused(on bool) {
	e.on = on
}

var _ tcell.Event = (*Event)(nil)
var _ types.FocusEvent = (*Event)(nil)
