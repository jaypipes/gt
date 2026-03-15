package scroll

import (
	"fmt"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling scroll-related events.
// Implements [types.ScrollEvent].
type Event struct {
	event.Event
	// pos contains the coordinates of the scroll when the event fired.
	pos types.Point
	// direction is the scroll direction.
	direction types.ScrollDirection
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	return fmt.Sprintf("scroll:%s", e.direction)
}

// Position returns the coordinates of the scroll when the event fired.
func (e *Event) Position() types.Point {
	return e.pos
}

// SetPosition sets the coordinates of the scroll when the event fired.
func (e *Event) SetPosition(pos types.Point) {
	e.pos = pos
}

// Direction returns the scroll direction.
func (e *Event) Direction() types.ScrollDirection {
	return e.direction
}

// SetDirection sets the scroll direction.
func (e *Event) SetDirection(direction types.ScrollDirection) {
	e.direction = direction
}

// scrollDirectionFromTCell translates a tcell ButtonMask to a ScrollDirection.
func scrollDirectionFromTCell(bm tcell.ButtonMask) types.ScrollDirection {
	switch {
	case bm&tcell.WheelUp != 0:
		return types.ScrollDirectionUp
	case bm&tcell.WheelDown != 0:
		return types.ScrollDirectionDown
	case bm&tcell.WheelLeft != 0:
		return types.ScrollDirectionLeft
	case bm&tcell.WheelRight != 0:
		return types.ScrollDirectionRight
	}
	return types.ScrollDirectionNone
}

var _ tcell.Event = (*Event)(nil)
var _ types.ScrollEvent = (*Event)(nil)
