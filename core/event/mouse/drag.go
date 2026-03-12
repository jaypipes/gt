package mouse

import "github.com/jaypipes/gt/types"

// NewDragMoveEvent returns a DragMoveEvent given the current and starting
// mouse event.
func NewDragMoveEvent(
	cur types.MouseEvent,
	start types.MouseEvent,
) *DragMoveEvent {
	return &DragMoveEvent{
		MouseEvent: cur,
		start:      start,
	}
}

// DragMoveEvent describes an ongoing mouse drag operation
type DragMoveEvent struct {
	types.MouseEvent
	// start is the starting event for the drag operation
	start types.MouseEvent
}

// Start returns the DragStartEvent for the start of the drag operation.
func (e *DragMoveEvent) Start() types.MouseEvent {
	return e.start
}

// NewDragStopEvent returns a DragStopEvent given the stopping and starting
// mouse event.
func NewDragStopEvent(
	stop types.MouseEvent,
	start types.MouseEvent,
) *DragStopEvent {
	return &DragStopEvent{
		MouseEvent: stop,
		start:      start,
	}
}

// DragStopEvent describes the stop of a mouse drag operation
type DragStopEvent struct {
	types.MouseEvent
	// start is the starting event for the drag operation
	start types.MouseEvent
}

// Start returns the DragStartEvent for the start of the drag operation.
func (e *DragStopEvent) Start() types.MouseEvent {
	return e.start
}

var _ types.MouseDragMoveEvent = (*DragMoveEvent)(nil)
var _ types.MouseDragStopEvent = (*DragStopEvent)(nil)
