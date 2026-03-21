package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// HasFocus returns true if the Element has the current focus.
func (e *Element) HasFocus() bool {
	return e.focused
}

// Focus handles focus events.
func (e *Element) Focus(ctx context.Context, ev types.FocusEvent) {
	e.focused = ev.Enabled()
	for _, cb := range e.onFocus {
		cb(ctx, ev)
	}
}

// OnFocus registers a callback that will be executed when the Element receives
// or loses the focus.
func (e *Element) OnFocus(cb types.FocusEventCallback) {
	e.onFocus = append(e.onFocus, cb)
}
