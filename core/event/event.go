package event

import (
	"time"

	"github.com/gdamore/tcell/v3"
)

// Event exposes an easy-to-use interface for handling mouse-related events.
// Implements [tcell.Event].
type Event struct {
	// when is the time that the event was generated
	when time.Time
}

// When returns the time the event was generated.
func (e *Event) When() time.Time {
	return e.when
}

// SetWhen sets the time the Event was generated.
func (e *Event) SetWhen(when time.Time) {
	e.when = when
}

var _ tcell.Event = (*Event)(nil)
