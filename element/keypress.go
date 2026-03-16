package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// KeyPress executes any OnKeyPress callbacks that were registered for the
// Element, returning true if the key press event was consumed/handled.
func (e *Element) KeyPress(ctx context.Context, ev types.KeyPressEvent) bool {
	for _, cb := range e.onKeyPress {
		if cb(ctx, ev) {
			return true
		}
	}
	if len(e.children) > 0 {
		for _, node := range e.children {
			h, ok := node.(types.KeyPressEventHandler)
			if ok {
				if h.KeyPress(ctx, ev) {
					return true
				}
			}
		}
	}
	return false
}

// OnKeyPress registers a callback that will be executed when a keypress
// combination is actuated.
func (e *Element) OnKeyPress(cb types.KeyPressEventCallback) {
	e.onKeyPress = append(e.onKeyPress, cb)
}
