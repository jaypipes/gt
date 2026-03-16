package application

import (
	"context"
	"io"
	"sync"

	"github.com/jaypipes/gt/core/key"
	"github.com/jaypipes/gt/types"
)

func newController(s types.Screen) *Controller {
	return &Controller{
		screen: s,
	}
}

type Controller struct {
	sync.RWMutex
	screen types.Screen

	// keyMap contains the key combination callbacks managed by the Controller.
	keyMap types.KeyMap
	// inputIntercept points to an io.Writer that receives all keyboard input
	// after TrapKeyPressWithEscape has been called.
	inputIntercept io.Writer
	// keyEscape is the key combination that will trigger the interceptor
	// to be removed.
	keyEscape types.Key
}

// Screen returns the [types.Screen] controlled by the Controller.
func (c *Controller) Screen() types.Screen {
	return c.screen
}

// HandleKeyPress performs the necessary action when the supplied key press
// event is received. Returns a bool indicating whether the event was handled.
func (c *Controller) HandleKeyPress(
	ctx context.Context,
	ev types.KeyPressEvent,
) bool {
	c.Lock()
	defer c.Unlock()

	if c.keyEscape != nil && ev.Matches(c.keyEscape) {
		c.keyEscape = nil
		c.inputIntercept = nil
		return true
	}

	if c.inputIntercept != nil {
		s := string(ev.Key().Code())
		c.inputIntercept.Write([]byte(s))
		return true
	}

	for kp, cb := range c.keyMap {
		if ev.Matches(kp) {
			cb(ctx)
			return true
		}
	}
	return false
}

// SetKeyMap sets the Controller's map of key press combinations to
// callbacks that execute when that key press combination is typed.
func (c *Controller) SetKeyMap(kp types.KeyMap) {
	c.Lock()
	defer c.Unlock()
	c.keyMap = kp
}

// InterceptKey signals the Controller to trap all key press events and
// write all graphemes in key press events to the supplied io.Writer. This
// method allows elements to need to take input from the user when they have
// the focus to prevent keyboard shortcuts from interfering with the input
// stream.
//
// escape can be a string, tcell.Key or types.Key.
func (c *Controller) InterceptKey(escape any, w io.Writer) {
	c.Lock()
	defer c.Unlock()
	c.inputIntercept = w
	c.keyEscape = key.New(escape)
}

// RestoreKey signals the Controller to restore the key press map from
// before it was trapped. This allows elements that lose the focus to release
// any hold they had on the key press events.
func (c *Controller) RestoreKey() {
	c.Lock()
	defer c.Unlock()

	c.keyEscape = nil
	c.inputIntercept = nil
}
