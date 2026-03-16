package application

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// setFocus sets the currently-focused thing and calls Focus(false) on the
// previously-focused thing, returning whether there was a change in focus.
func (a *Application) setFocus(ctx context.Context, f types.Focusable) bool {
	if a.focused != nil {
		if f == nil {
			a.focused.SetFocus(ctx, false)
			a.focused = nil
			return true
		}
		if f.HasFocus() {
			// already has the focus, no need to do anything...
			return false
		}
		a.focused.SetFocus(ctx, false)
	} else {
		if f == nil {
			return false
		}
	}
	if f != nil {
		f.SetFocus(ctx, true)
	}
	a.focused = f
	return true
}
