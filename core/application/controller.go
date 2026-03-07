package application

import (
	"context"
	"io"
	"sync"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/types"
)

func newController(s tcell.Screen) *Controller {
	return &Controller{
		Screen: s,
	}
}

type Controller struct {
	sync.RWMutex
	tcell.Screen

	// prevFocused is the Focusable that previously had the focus.
	prevFocused types.Focusable
	// focused is the Focusable that currently has the focus.
	focused types.Focusable

	// keyPressMap contains the key press combination callbacks managed by the
	// Controller.
	keyPressMap types.KeyPressMap
	// keyIntercept points to an io.Writer that receives all keyboard input
	// after TrapKeyPressWithEscape has been called.
	keyIntercept io.Writer
	// keyEscape is the key press combination that will trigger the interceptor
	// to be removed.
	keyEscape string
}

// HandleFocus sets the focus on the supplied Focusable and releases the focus
// on any previously-focused Focusable.
func (c *Controller) HandleFocus(ctx context.Context, f types.Focusable) {
	if c.prevFocused != nil {
		c.prevFocused.SetFocus(ctx, false)
	}
	c.prevFocused = f
	f.SetFocus(ctx, true)
	c.focused = f
}

// HandleKeyPress performs the necessary action when the supplied key press
// event is received. Returns a bool indicating whether the event was handled.
func (c *Controller) HandleKeyPress(
	ctx context.Context,
	ev types.KeyPressEvent,
) bool {
	c.Lock()
	defer c.Unlock()

	if c.keyEscape != "" && ev.MatchAny(c.keyEscape) {
		c.keyEscape = ""
		c.keyIntercept = nil
		return true
	}

	if c.keyIntercept != nil {
		c.keyIntercept.Write([]byte(ev.Printable()))
		return true
	}

	for kp, cb := range c.keyPressMap {
		if ev.MatchAny(kp) {
			cb(ctx)
			return true
		}
	}
	return false
}

// SetKeyPressMap sets the Controller's map of key press combinations to
// callbacks that execute when that key press combination is typed.
func (c *Controller) SetKeyPressMap(kp types.KeyPressMap) {
	c.Lock()
	defer c.Unlock()
	c.keyPressMap = kp
}

// InterceptKeyPress signals the Controller to trap all key press events and
// write all graphemes in key press events to the supplied io.Writer. This
// method allows elements to need to take input from the user when they have
// the focus to prevent keyboard shortcuts from interfering with the input
// stream.
func (c *Controller) InterceptKeyPress(escape string, w io.Writer) {
	c.Lock()
	defer c.Unlock()
	c.keyIntercept = w
	c.keyEscape = escape
}

// RestoreKeyPress signals the Controller to restore the key press map from
// before it was trapped. This allows elements that lose the focus to release
// any hold they had on the key press events.
func (c *Controller) RestoreKeyPress() {
	c.Lock()
	defer c.Unlock()

	c.keyEscape = ""
	c.keyIntercept = nil
}
