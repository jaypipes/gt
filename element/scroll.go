package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Scroll executes any OnScroll callbacks that were registered for the Element.
func (e *Element) Scroll(ctx context.Context, ev types.ScrollEvent) {
	for _, cb := range e.onScroll {
		cb(ctx, ev)
	}
}

// OnScroll registers a callback that will be executed when a mouse
// wheel/scroll occurs over the Element.
func (e *Element) OnScroll(cb types.ScrollEventCallback) {
	e.onScroll = append(e.onScroll, cb)
}
