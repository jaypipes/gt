package types

import "context"

// KeyPressEvent describes events received when a key press occurs.
type KeyPressEvent interface {
	Event
	// Key returns the key combination that was pressed
	Key() Key
	// SetKey sets the key combination that was pressed.
	SetKey(Key)
	// Matches returns true if the KeyPressEvent matches for any of the
	// supplied Keys.
	Matches(...Key) bool
}

// KeyEventWithOption describes an optional varg parameter to
// [core.event.key.New] that modifies the returned KeyEvent.
type KeyPressEventWithOption func(KeyPressEvent)

// KeyEventCallback is the function signature for callbacks executed on key
// combination actuation. The callback returns whether the event was
// consumed/handled.
type KeyPressEventCallback func(context.Context, KeyPressEvent) bool

// KeyPressHandler describes something that can have key press combinations and
// callbacks associated with it.
type KeyPressEventHandler interface {
	// KeyPress handles key press events. It returns true if the handler
	// consumed/handled the event, false if not.
	KeyPress(context.Context, KeyPressEvent) bool
	// OnKeyPress registers a callback to execute upon a key press combination.
	OnKeyPress(KeyPressEventCallback)
}
