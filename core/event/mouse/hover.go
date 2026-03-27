package mouse

import "github.com/jaypipes/gt/types"

// NewHoverEvent returns a HoverEvent given the current mouse event and whether
// or not the mouse was hovering over the element.
func NewHoverEvent(
	ev types.MouseEvent,
	hovered bool,
) *HoverEvent {
	return &HoverEvent{
		MouseEvent: ev,
		hovered:    hovered,
	}
}

// HoverEvent describes a mouse Hover event.
type HoverEvent struct {
	types.MouseEvent
	// hovered is true if the mouse was hovering over the receiver of the
	// event.
	hovered bool
}

// Hovered returns true if the user is currently hovering over the receiver of
// the event.
func (e *HoverEvent) Hovered() bool {
	return e.hovered
}

var _ types.MouseHoverEvent = (*HoverEvent)(nil)
