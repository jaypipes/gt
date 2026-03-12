package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// MouseClick executes any OnMouseClick callbacks that were registered for the
// Element.
func (e *Element) MouseClick(ctx context.Context, ev types.MouseClickEvent) {
	for _, cb := range e.onMouseClick {
		cb(ctx, ev)
	}
}

// OnMouseClick registers a callback that will be executed when the Element is
// clicked.
func (e *Element) OnMouseClick(cb types.MouseClickEventCallback) {
	e.onMouseClick = append(e.onMouseClick, cb)
}

// MouseDoubleClick executes any OnMouseDoubleClick callbacks that were
// registered for the Element.
func (e *Element) MouseDoubleClick(ctx context.Context, ev types.MouseClickEvent) {
	for _, cb := range e.onMouseDoubleClick {
		cb(ctx, ev)
	}
}

// OnMouseDoubleClick registers a callback that will be executed when the
// Element is double-clicked.
func (e *Element) OnMouseDoubleClick(cb types.MouseClickEventCallback) {
	e.onMouseDoubleClick = append(e.onMouseDoubleClick, cb)
}

// MouseScroll executes any OnMouseScroll callbacks that were registered for
// the Element.
func (e *Element) MouseScroll(ctx context.Context, ev types.MouseEvent) {
	for _, cb := range e.onMouseScroll {
		cb(ctx, ev)
	}
}

// OnMouseScroll registers a callback that will be executed when the mouse
// wheel scrolls and the Element has the focus.
func (e *Element) OnMouseScroll(cb types.MouseEventCallback) {
	e.onMouseScroll = append(e.onMouseScroll, cb)
}

// MouseDragMove executes any OnMouseDragMove callbacks that were registered for
// the Element.
func (e *Element) MouseDragMove(
	ctx context.Context,
	ev types.MouseDragMoveEvent,
) {
	for _, cb := range e.onMouseDragMove {
		cb(ctx, ev)
	}
}

// OnMouseDragMove registers a callback that will be executed when the user
// performs a mouse drag action.
func (e *Element) OnMouseDragMove(cb types.MouseDragMoveEventCallback) {
	e.onMouseDragMove = append(e.onMouseDragMove, cb)
}

// MouseDragStop executes any OnMouseDragStop callbacks that were registered for
// the Element.
func (e *Element) MouseDragStop(
	ctx context.Context,
	ev types.MouseDragStopEvent,
) {
	for _, cb := range e.onMouseDragStop {
		cb(ctx, ev)
	}
}

// OnMouseDragStop registers a callback that will be executed when the user
// ends a mouse drag action.
func (e *Element) OnMouseDragStop(cb types.MouseDragStopEventCallback) {
	e.onMouseDragStop = append(e.onMouseDragStop, cb)
}
