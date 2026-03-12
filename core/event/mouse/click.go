package mouse

import "github.com/jaypipes/gt/types"

// NewClickEvent returns a ClickEvent given the current mouse event and whether
// or not it was a double-click.
func NewClickEvent(
	ev types.MouseEvent,
	dclick bool,
) *ClickEvent {
	return &ClickEvent{
		MouseEvent: ev,
		dclick:     dclick,
	}
}

// ClickEvent describes a mouse click or double-click event.
type ClickEvent struct {
	types.MouseEvent
	// dclick is true if the mouse event was a double-click.
	dclick bool
}

// DoubleClicked returns true if the user double-clicked.
func (e *ClickEvent) DoubleClicked() bool {
	return e.dclick
}

var _ types.MouseClickEvent = (*ClickEvent)(nil)
