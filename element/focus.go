package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// NextFocusable returns the next focusable thing, or nil if there is no next
// focusable thing. The Element's children will first be inspected and then the
// next sibling Element.
func (e *Element) NextFocusable(ctx context.Context) types.FocusEventHandler {
	for _, child := range e.children {
		feh, ok := child.(types.FocusEventHandler)
		if ok && feh.CanFocus() {
			return feh
		}
	}
	ns := e.NextSibling()
	if ns != nil {
		feh, ok := ns.(types.FocusEventHandler)
		if ok {
			if feh.CanFocus() {
				return feh
			}
			nse, ok := ns.(types.Element)
			if ok {
				// Continue looking for focusable siblings...
				feh = nse.NextFocusable(ctx)
				if feh != nil {
					return feh
				}
			}
		}
	}
	return nil
}

// CanFocus returns true if the Element can receive the focus. Disabled
// Elements cannot receive the focus.
func (e *Element) CanFocus() bool {
	return !e.disabled
}

// HasFocus returns true if the Element has the current focus.
func (e *Element) HasFocus() bool {
	return e.focused
}

// Focus handles focus events.
func (e *Element) Focus(ctx context.Context, ev types.FocusEvent) {
	e.focused = ev.Focused()
	for _, cb := range e.onFocus {
		cb(ctx, ev)
	}
}

// OnFocus registers a callback that will be executed when the Element receives
// or loses the focus.
func (e *Element) OnFocus(cb types.FocusEventCallback) {
	e.onFocus = append(e.onFocus, cb)
}
