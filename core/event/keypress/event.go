package keypress

import (
	"fmt"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling KeyPress events.
// Implements [types.KeyPressEvent].
type Event struct {
	*event.Event
	// key is the key combination that was pressed.
	key types.Key
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	return fmt.Sprintf(
		"keypress:%s",
		e.key,
	)
}

// SetKey sets the underlying key combination that was pressed.
func (e *Event) SetKey(k types.Key) {
	e.key = k
}

// Key returns the underlying key combination that was pressed.
func (e *Event) Key() types.Key {
	return e.key
}

// Matches returns true if the event matches for any of the supplied key
// combination strings or [types.Key] objects.
func (e *Event) Matches(subjects ...types.Key) bool {
	for _, subject := range subjects {
		if subject.Equal(e.key) {
			return true
		}
	}
	return false
}

var _ tcell.Event = (*Event)(nil)
var _ types.KeyPressEvent = (*Event)(nil)
