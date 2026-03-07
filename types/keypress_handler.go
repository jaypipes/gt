package types

import "context"

// KeyPressHandler describes something that can have key press combinations and
// callbacks associated with it.
type KeyPressHandler interface {
	// HandleKeyPress handles key press events.
	HandleKeyPress(context.Context, KeyPressEvent)
	// OnKeyPress registers a callback to execute upon a key press combination.
	OnKeyPress(string, EventCallback)
}
