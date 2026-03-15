package types

import "context"

// ScrollDirection describes the direction of a scroll event
type ScrollDirection int8

const (
	ScrollDirectionNone ScrollDirection = iota
	ScrollDirectionUp
	ScrollDirectionDown
	ScrollDirectionLeft
	ScrollDirectionRight
)

var (
	scrollDirectionNames = []string{
		"none",
		"up",
		"down",
		"left",
		"right",
	}
)

func (d ScrollDirection) String() string {
	return scrollDirectionNames[int(d)]
}

// ScrollEvent describes an event when the user wants to scroll some content.
// This event is created when the mouse wheel is engaged or can be created
// manually by a component when, for instance, an arrow key is pressed and the
// cursor is at the edge of the component's content.
type ScrollEvent interface {
	Event
	// Position returns where the mouse was when the ScrollEvent was triggered.
	// Note that this will be the zero position if the ScrollEvent was manually
	// created by a component in response to, for instance, a key press event.
	Position() Point
	// SetPosition sets the coordinates of the mouse when the ScrollEvent was
	// triggered.
	SetPosition(Point)
	// Direction returns the ScrollDirection.
	Direction() ScrollDirection
	// SetDirection sets the ScrollDirection.
	SetDirection(ScrollDirection)
}

// ScrollEventWithOption describes an optional varg parameter to
// [core.event.scroll.New] that modifies the returned ScrollEvent.
type ScrollEventWithOption func(ScrollEvent)

// ScrollEventCallback is the function signature for callbacks executed on
// scroll events.
type ScrollEventCallback func(context.Context, ScrollEvent)

// ScrollEventHandler represents something that can handle scroll events.
type ScrollEventHandler interface {
	// Scroll executes any OnScroll callbacks that were registered for
	// the ScrollEventHandler.
	Scroll(context.Context, ScrollEvent)
	// OnScrollHover registers a callback that will be executed when the mouse
	// is over top of an element and a wheel/scroll action occurs.
	OnScroll(ScrollEventCallback)
}
