package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// HasFocus returns true if the Element has the current focus.
func (e *Element) HasFocus() bool {
	return e.focused
}

// SetFocus marks the Element as either having or not having the focus
// and executes any associated OnFocus or OnLoseFocus callbacks that were
// registered for the Element.
func (e *Element) SetFocus(ctx context.Context, focus bool) {
	stack := e.onFocus
	if !focus {
		stack = e.onLoseFocus
	}
	e.focused = focus
	for _, cb := range stack {
		cb(ctx)
	}
}

// OnFocus registers a callback that will be executed when the Element is
// focused.
func (e *Element) OnFocus(cb types.FocusCallback) {
	e.onFocus = append(e.onFocus, cb)
}

// OnLoseFocus registers a callback that will be executed when the Element
// loses focus.
func (e *Element) OnLoseFocus(cb types.FocusCallback) {
	e.onLoseFocus = append(e.onLoseFocus, cb)
}
