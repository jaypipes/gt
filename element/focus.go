package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// SetDisabled sets whether the Element is disabled. Disabled Elements cannot
// receive the focus.
func (e *Element) SetDisabled(on bool) {
	e.disabled = on
}

// Disabled returns true if the Element cannot get the focus.
func (e *Element) Disabled() bool {
	return e.disabled
}

// WithDisabled sets whether the Element is disabled and returns the Element.
func (e *Element) WithDisabled(on bool) types.Element {
	e.SetDisabled(on)
	return e
}

// SetFocusable sets whether the Element is focusable. Focusable Elements cannot
// receive the focus.
func (e *Element) SetFocusable(on bool) {
	e.focusable = on
}

// Focusable returns true if the Element cannot get the focus.
func (e *Element) Focusable() bool {
	return e.focusable && !e.disabled
}

// WithFocusable sets whether the Element is focusable and returns the Element.
func (e *Element) WithFocusable(on bool) types.Element {
	e.SetFocusable(on)
	return e
}

// NextFocusable returns the next focusable thing, or nil if there is no next
// focusable thing. The Element's children will first be inspected and then the
// next sibling Element.
func (e *Element) NextFocusable(ctx context.Context) types.FocusEventHandler {
	for _, child := range e.children {
		feh, ok := child.(types.FocusEventHandler)
		if ok && feh.Focusable() {
			return feh
		}
	}
	ns := e.NextSibling()
	if ns != nil {
		feh, ok := ns.(types.FocusEventHandler)
		if ok {
			if feh.Focusable() {
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
