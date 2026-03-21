package application

import (
	"context"

	"github.com/samber/lo"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/key"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// SetExitKey configures the Application to exit on any of the supplied
// keypress combinations. These keypress combinations are always evaluated
// first when a keypress event is received by the Application.
//
// The keypress combinations can be strings -- e.g. "Ctrl+C", "Esc" -- or
// [tcell.Key] codes -- e.g. tcell.KeyCtrlC, KeyEscape.
//
// If no exit keypress combinations are set for the Application, it defaults to
// "Ctrl+C".
func (a *Application) SetExitKey(subject ...any) {
	k := key.New(subject)
	for _, ek := range a.exitKeys {
		if !k.Equal(ek) {
			a.exitKeys = append(a.exitKeys, k)
		}
	}
}

// InterceptKeyPressEvents signals the Application to trap all key press events
// and route all key press events to the supplied KeyPressEventHandler. This
// method allows elements to need to take input from the user when they have
// the focus to prevent keyboard shortcuts from interfering with the input
// stream.
func (a *Application) InterceptKeyPressEvents(
	ctx context.Context,
	escapeKey types.Key,
	handler types.KeyPressEventHandler,
) {
	a.Lock()
	defer a.Unlock()
	if handler == nil || escapeKey == nil {
		return
	}
	gtlog.Debug(
		ctx,
		"Application.InterceptKeyPressEvents: to=%s escape=%q",
		core.ID(handler), escapeKey,
	)
	a.keyInterceptor = handler
	a.keyInterceptEscape = escapeKey
}

// StopInterceptKey signals the Application to restore the key map from
// before it was trapped. This allows elements that lose the focus to
// release any hold they had on the key press events.
func (a *Application) StopInterceptKeyPressEvents(
	ctx context.Context,
) {
	a.Lock()
	defer a.Unlock()
	if a.keyInterceptor == nil {
		return
	}
	gtlog.Debug(ctx, "Application.StopInterceptKeyPressEvents")
	a.keyInterceptEscape = nil
	a.keyInterceptor = nil
}

// exitKeyPressed returns true if the supplied KeyPressEvent matches any of the
// exit keys registered for the Application.
func (a *Application) exitKeyPressed(ev types.KeyPressEvent) bool {
	for _, ek := range a.exitKeys {
		if ek.Equal(ev.Key()) {
			return true
		}
	}
	return false
}

// handleKeyPressEvent passes a KeyPressEvent to any handlers that are
// listening for KeyPressEvents.
func (a *Application) handleKeyPressEvent(
	ctx context.Context,
	ev types.KeyPressEvent,
) {
	a.RLock()
	focused := a.focused
	interceptor := a.keyInterceptor
	escapeKey := a.keyInterceptEscape
	a.RUnlock()

	k := ev.Key()

	// If we have an intercepting handler, just route all KeyPressEvents to
	// that handler until the escape key is seen, at which point we stop the
	// intercepting.
	if interceptor != nil {
		if escapeKey.Equal(k) {
			if focused != nil {
				if core.ID(focused) == core.ID(interceptor) {
					// release the focus on an element that has stopped
					// intercepting key press events.
					a.removeFocus(ctx)
				}
			}
			a.StopInterceptKeyPressEvents(ctx)
			a.draw(ctx)
			return
		}
		interceptor.KeyPress(ctx, ev)
		a.draw(ctx)
		return
	}

	// If there is an element that has the focus, we send the key press event
	// to that element. That element can return false, meaning it did not
	// consume/handle the event. If that is the case, or there was no element
	// with the focus, we send the key press event to all elements in the
	// current view, stopping when any element returns a true value.
	handled := false
	if focused != nil {
		handler, ok := focused.(types.KeyPressEventHandler)
		if ok {
			handled = handler.KeyPress(ctx, ev)
		}
		if handled {
			a.draw(ctx)
			return
		}
	}
	v := a.CurrentView()
	if v.KeyPress(ctx, ev) {
		a.draw(ctx)
	}
}

// KeyMap returns the Application's *global* map of key press
// combination strings to callbacks that will execute when that key press
// combination is entered.
func (a *Application) KeyMap() types.KeyMap {
	return a.keyMap
}

// OnKeyPress registers an Application-level (global)  callback to execute
// upon a key press combination.
//
// The keypress combination can be a string -- e.g. "Ctrl+C", "Esc" -- or a
// [tcell.Key] code -- e.g. tcell.KeyCtrlC, KeyEscape.
func (a *Application) OnKeyPress(subject any, cb types.EventCallback) {
	kp := key.New(subject)
	a.keyMap[kp] = cb
}

// buildKeyMap builds the Application's outermost map of key press combinations
// to callback functions to execute when those key press combinations are
// entered.
//
// The outermost map will always be the "current view" key press combinations
// that the Application's registered Views have along with any key press
// combinations registered with the Application itself and any key press
// combinations that the *current* View contains.
func (a *Application) buildKeyMap(
	ctx context.Context,
) types.KeyMap {
	res := types.KeyMap{}

	// copy in our global key press callbacks
	for k, cb := range a.keyMap {
		res[k] = cb
	}
	globalKPs := lo.Keys(a.keyMap)

	// now add our "current view" key press callbacks
	for viewID, view := range a.views {
		currentViewKP := view.CurrentViewKey()
		if currentViewKP != nil {
			if lo.Contains(globalKPs, currentViewKP) {
				gtlog.Warn(
					ctx,
					"current view key press combination %q for view %q "+
						"shadows global key press combination",
					currentViewKP, viewID,
				)
			}
			res[currentViewKP] = func(_ context.Context) {
				a.SetCurrentView(viewID)
			}
		}
	}

	// finally, add all the current View's key press callbacks
	curView := a.views[a.curView]
	curViewKPMap := curView.KeyMap()
	if len(curViewKPMap) > 0 {
		appKPs := lo.Keys(res)
		for kp, cb := range curViewKPMap {
			if lo.Contains(appKPs, kp) {
				gtlog.Warn(
					ctx,
					"view key press combination %q for view %q "+
						"shadows application key press combination",
					kp, curView.ID(),
				)
			}
			res[kp] = cb
		}
	}

	gtlog.Debug(
		ctx,
		"Application.buildKeyMap: built map for combinations %v",
		lo.Keys(res),
	)
	return res
}
