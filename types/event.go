package types

import (
	"context"
	"fmt"

	"github.com/gdamore/tcell/v3"
)

// EventCallback describes a function that will execute when some Event fires.
type EventCallback func(context.Context)

// Event is the base interface for all gt events.
type Event interface {
	fmt.Stringer
	tcell.Event
}

// KeyPressEvent describes events received when a key press occurs.
type KeyPressEvent interface {
	Event
	KeyModifiable
	// Key returns the virtual key code
	Key() Key
	// Printable returns the printable character(s) associated with the key
	// press event.
	Printable() string
	// MatchAny returns true if the KeyPressEvent matches for any of the
	// keypress strings or key codes supplied.
	MatchAny(...any) bool
}

// MouseEvent describes events received when a mouse moved, clicked or
// released.
type MouseEvent interface {
	Event
	KeyModifiable
	// Button returns the mouse button that was depressed, if any.
	Button() MouseButton
	// Position returns where the mouse was when the MouseEvent was triggered.
	Position() Point
}
