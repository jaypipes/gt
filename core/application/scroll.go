package application

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// handleScrollEvent fires an OnScroll event against the target element, if
// any.
func (a *Application) handleScrollEvent(
	ctx context.Context,
	ev types.ScrollEvent,
) {
	// We determine if the mouse is over a thing that can handle scroll events.
	// If the mouse is over something that can handle mouse events and is not
	// disabled, we will fire the OnScroll event.
	pos := ev.Position()
	v := a.CurrentView()
	node := v.AtPoint(pos)
	if node != nil {
		el, ok := node.(types.Element)
		if ok && !el.Disabled() {
			el.Scroll(ctx, ev)
			a.draw(ctx)
		}
	}
}
