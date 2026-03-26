package types

import "context"

// Focusable represents something that sets and reports whether it can have the
// focus given to it.
type Focusable interface {
	// SetFocusable sets whether the Focusable can receive the focus.
	SetFocusable(bool)
	// Focusable returns true if the Focusable can receive the focus.
	// For Elements that are disabled, this should return false.
	Focusable() bool
}

// FocusEvent describes events received when the focus changes.
type FocusEvent interface {
	Event
	// Focused returns true if the receiver of the event should receive the
	// focus.
	Focused() bool
	// SetFocused sets whether the receiver of the event should receive the
	// focus.
	SetFocused(bool)
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
	Focusable
	// HasFocus returns true if the FocusEventHandler has the current focus.
	HasFocus() bool
	// Focus handles focus events.
	Focus(context.Context, FocusEvent)
	// OnFocus registers a callback that will be executed when a FocusEvent
	// occurs.
	OnFocus(FocusEventCallback)
}
