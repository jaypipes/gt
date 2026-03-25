package application

import (
	"context"

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

// SetFocusNextKey configures the keypress combinations that tell the
// Application to move the focus to the next focusable element.
//
// The keypress combinations can be strings ("Tab"), [tcell.Key] codes
// (tcell.KeyTab), [types.KeyCode] values (types.KeyCodeTab) or [types.Key]
// objects (core.key.KeyTab)
//
// If no move focus keypress combinations are set for the Application, it defaults to
// "Tab".
func (a *Application) SetFocusNextKey(subject ...any) {
	k := key.New(subject)
	for _, ek := range a.focusNextKeys {
		if !k.Equal(ek) {
			a.focusNextKeys = append(a.focusNextKeys, k)
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

// SetKeyShortcut registers an Application-level KeyShortcut that will execute
// upon a key press combination.
func (a *Application) SetKeyShortcut(shortcut types.KeyShortcut) {
	k := shortcut.Key()
	for _, ks := range a.keyShortcuts {
		ksk := ks.Key()
		if ksk.Equal(k) {
			gtlog.Warn(
				context.TODO(),
				"key shortcut %q shadows previously-registered "+
					"application-level key shortcut",
				k,
			)
		}
	}
	a.keyShortcuts = append(a.keyShortcuts, shortcut)
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

// focusNextKeyPressed returns true if the supplied KeyPressEvent matches any
// of the focusNext keys registered for the Application.
func (a *Application) focusNextKeyPressed(ev types.KeyPressEvent) bool {
	for _, fnk := range a.focusNextKeys {
		if fnk.Equal(ev.Key()) {
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
	keyShortcuts := a.keyShortcuts
	views := a.views
	activeView := a.ActiveView()
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

	handled := false

	// If our "move focus to next focusable" key press combination was pressed,
	// let's move our focus.
	if a.focusNextKeyPressed(ev) {
		handled = a.FocusNext(ctx)
		if handled {
			a.draw(ctx)
			return
		}
	}

	// Next, we handle our Application-level global key shortcuts.
	for _, ks := range keyShortcuts {
		ksk := ks.Key()
		if ksk.Equal(k) {
			cb := ks.Callback()
			cb(ctx)
			return
		}
	}

	// Then we check if the key press combination is a "switch active view"
	// key, and if so, set the active view.
	activeViewID := activeView.ID()
	for viewID, v := range views {
		if viewID == activeViewID {
			continue
		}
		vk := v.ActiveKey()
		if vk != nil && vk.Equal(k) {
			a.SetActiveView(viewID)
			a.draw(ctx)
			return
		}
	}

	// If there is an element that has the focus, we send the key press event
	// to that element. That element can return false, meaning it did not
	// consume/handle the event. If that is the case, or there was no element
	// with the focus, we send the key press event to all elements in the
	// active view, stopping when any element returns a true value.
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

	// Finally, if nothing has handled the KeyPressEvent, we ask the active
	// view to handle it.
	if activeView.KeyPress(ctx, ev) {
		a.draw(ctx)
	}
}
