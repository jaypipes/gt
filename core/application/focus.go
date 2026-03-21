package application

import (
	"context"

	"github.com/jaypipes/gt/core"
	fevent "github.com/jaypipes/gt/core/event/focus"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// setFocus sets the currently-focused thing and removes the focus from the
// previously-focused thing, returning whether there was a change in focus.
func (a *Application) setFocus(
	ctx context.Context,
	target types.FocusEventHandler,
) bool {
	if a.focused != nil {
		a.removeFocus(ctx)
		if target == nil {
			return true
		}
		if target.HasFocus() {
			// already has the focus, no need to do anything...
			return false
		}
	} else {
		if target == nil {
			return false
		}
	}
	gtlog.Debug(ctx, "Application.setFocus on %s", core.ID(target))
	ev := fevent.New(
		fevent.WithEnabled(true), fevent.WithProducer(a),
	)
	target.Focus(ctx, ev)
	a.focused = target
	return true
}

// removeFocus removes the Application's focus and fires a FocusEvent to the
// currently focused element to remove its focus.
func (a *Application) removeFocus(ctx context.Context) {
	if a.focused == nil {
		return
	}
	gtlog.Debug(ctx, "Application.removeFocus on %s", core.ID(a.focused))
	ev := fevent.New(
		fevent.WithEnabled(false), fevent.WithProducer(a),
	)
	a.focused.Focus(ctx, ev)
	a.focused = nil
}
