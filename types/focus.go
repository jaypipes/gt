package types

import "context"

// FocusEvent describes events received when the focus changes.
type FocusEvent interface {
	Event
	// Enabled returns true if the receiver of the event should receive the
	// focus.
	Enabled() bool
	// SetEnabled sets whether the receiver of the event should receive the
	// focus.
	SetEnabled(bool)
}

// FocusEventWithOption describes an optional varg parameter to
// [core.event.focus.New] that modifies the returned FocusEvent.
type FocusEventWithOption func(FocusEvent)

// FocusEventCallback is the function signature for callbacks executed on focus
// events.
type FocusEventCallback func(context.Context, FocusEvent)

// FocusEventHandler represents something that can be focused on and perform
// some callback.
type FocusEventHandler interface {
	// HasFocus returns true if the FocusEventHandler has the current focus.
	HasFocus() bool
	// Focus handles focus events.
	Focus(context.Context, FocusEvent)
	// OnFocus registers a callback that will be executed when a FocusEvent
	// occurs.
	OnFocus(FocusEventCallback)
}
