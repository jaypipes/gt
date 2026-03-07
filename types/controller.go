package types

import (
	"context"
	"io"

	"github.com/gdamore/tcell/v3"
)

// Controller represents something that controls screen and cursor display.
type Controller interface {
	tcell.Screen
	// HandleFocus sets the focus on the supplied Focusable and releases the
	// focus on any previously-focused Focusable.
	HandleFocus(context.Context, Focusable)
	// HandleKeyPress performs the necessary action when the supplied key press
	// event is received. Returns a bool indicating whether the event was handled.
	HandleKeyPress(context.Context, KeyPressEvent) bool
	// SetKeyPressMap sets the Controller's map of key press combinations to
	// callbacks that execute when that key press combination is typed.
	SetKeyPressMap(KeyPressMap)
	// InterceptKeyPress signals the Controller to trap all key press events
	// and write all graphemes in key press events to the supplied io.Writer.
	// This method allows elements to need to take input from the user when
	// they have the focus to prevent keyboard shortcuts from interfering with
	// the input stream.
	InterceptKeyPress(string, io.Writer)
	// RestoreKeyPress signals the Controller to restore the key press map from
	// before it was trapped. This allows elements that lose the focus to
	// release any hold they had on the key press events.
	RestoreKeyPress()
}

// Controllable has a Controller
type Controllable interface {
	// Controller returns the screen controller.
	Controller() Controller
	// SetController sets the screen controller.
	SetController(Controller)
}
