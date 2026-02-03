package types

import "context"

// ClickCallback is the function signature for callbacks executed on mouse
// click events.
type ClickCallback func(context.Context, MouseClickEvent)

// Clickable represents something that can be clicked on by the mouse and
// perform some callback.
type Clickable interface {
	// Click executes any OnClick callbacks that were registered for the
	// Clickable.
	Click(context.Context, MouseClickEvent)
	// OnClick registers a callback that will be executed when the Clickable is
	// clicked.
	OnClick(ClickCallback)
}
