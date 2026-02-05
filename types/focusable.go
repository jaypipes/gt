package types

import "context"

// FocusCallback is the function signature for callbacks executed on focus
// events.
type FocusCallback func(context.Context)

// Focusable represents something that can be focused on by the mouse and
// perform some callback.
type Focusable interface {
	// HasFocus returns true if the Focusable has the current focus.
	HasFocus() bool
	// SetFocus marks the Focusable as either having or not having the focus
	// and executes any associated OnFocus or OnLoseFocus callbacks that were
	// registered for the Focusable.
	SetFocus(context.Context, bool)
	// OnFocus registers a callback that will be executed when the Focusable
	// gets the focus.
	OnFocus(FocusCallback)
	// OnLoseFocus registers a callback that will be executed when the
	// Focusable loses focus.
	OnLoseFocus(FocusCallback)
}
