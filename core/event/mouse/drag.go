package mouse

import "github.com/jaypipes/gt/types"

// NewDragMoveEvent returns a DragMoveEvent given the current and starting
// mouse event.
func NewDragEvent(
	cur types.MouseEvent,
	start types.MouseEvent,
) *DragEvent {
	return &DragEvent{
		MouseEvent: cur,
		start:      start,
	}
}

// DragEvent describes an ongoing mouse drag operation or the stop of a mouse
// drag operation.
type DragEvent struct {
	types.MouseEvent
	// start is the starting event for the drag operation
	start types.MouseEvent
}

// Start returns the DragStartEvent for the start of the drag operation.
func (e *DragEvent) Start() types.MouseEvent {
	return e.start
}

var _ types.MouseDragEvent = (*DragEvent)(nil)
