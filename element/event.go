package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Click executes any OnClick callbacks that were registered for the
// Clickable.
func (e *Element) Click(ctx context.Context, ev types.MouseClickEvent) {
	for _, cb := range e.onClick {
		cb(ctx, ev)
	}
	if ev.Button == types.MouseLeft {
		e.SetFocus(ctx, true)
	}
}

// OnClick registers a callback that will be executed when the Clickable is
// clicked.
func (e *Element) OnClick(cb types.ClickCallback) {
	e.onClick = append(e.onClick, cb)
}

// HasFocus returns true if the Focusable has the current focus.
func (e *Element) HasFocus() bool {
	return e.focused
}

// SetFocus marks the Focusable as either having or not having the focus
// and executes any associated OnFocus or OnLoseFocus callbacks that were
// registered for the Focusable.
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

// OnFocus registers a callback that will be executed when the Focusable is
// focused.
func (e *Element) OnFocus(cb types.FocusCallback) {
	e.onFocus = append(e.onFocus, cb)
}

// OnLoseFocus registers a callback that will be executed when the Focusable
// loses focus.
func (e *Element) OnLoseFocus(cb types.FocusCallback) {
	e.onLoseFocus = append(e.onLoseFocus, cb)
}
