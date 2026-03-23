package types

import (
	"context"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v3"
)

// EventCallback describes a function that will execute when some Event fires.
type EventCallback func(context.Context)

// Event is the base interface for all gt events.
type Event interface {
	fmt.Stringer
	tcell.Event
	// SetWhen sets the timestamp of the Event.
	SetWhen(time.Time)
	// Source returns the thing that fired the Event.
	Source() any
	// SetSource sets the thing that fired the Event.
	SetSource(any)
}

// EventWithOption describes an optional varg parameter to [core.event.New]
// that modifies the returned Event.
type EventWithOption func(Event)

// ApplicationEvent is used as the message payload for communicating
// Application, Element and Component state changes.
type ApplicationEvent interface {
	Event
}
