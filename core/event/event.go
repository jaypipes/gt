package event

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling mouse-related events.
// Implements [types.Event].
type Event struct {
	// when is the time that the event was generated
	when time.Time
	// source contains the thing that triggered the event.
	source any
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	return fmt.Sprintf(
		"event@%s", e.when,
	)
}

// When returns the time the event was generated.
func (e *Event) When() time.Time {
	return e.when
}

// SetWhen sets the time the Event was generated.
func (e *Event) SetWhen(when time.Time) {
	e.when = when
}

// Source returns the thing that triggered/created the event.
func (e *Event) Source() any {
	return e.source
}

// SetSource sets the thing that triggered/created the event.
func (e *Event) SetSource(source any) {
	e.source = source
}

var _ tcell.Event = (*Event)(nil)
var _ types.Event = (*Event)(nil)
