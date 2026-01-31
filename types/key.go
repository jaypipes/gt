package types

import "context"

// KeyPressCallback describes a function that will execute upon a key press.
type KeyPressCallback func(context.Context)

// KeyPressMap maps key press combination strings to callbacks that will
// execute upon that key press.
type KeyPressMap map[string]KeyPressCallback

// HasKeyPressMap describes something that has a map of key press combinations
// and callbacks associated with it.
type HasKeyPressMap interface {
	// KeyPressMap returns the KeyPressable's map of key press combination
	// strings to callbacks that will execute when that key press combination
	// is entered.
	KeyPressMap() KeyPressMap
}

// KeyPressable describes something that can have key press combinations and
// callbacks associated with it.
type KeyPressable interface {
	HasKeyPressMap
	// OnKeyPress registers a callback to execute upon a key press combination.
	OnKeyPress(string, KeyPressCallback)
}
