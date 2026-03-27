package types

import (
	"context"
)

// MouseButton indicates the mouse button that was activated.
type MouseButton int16

const (
	MouseButtonNone MouseButton = iota
	MouseButtonPrimary
	MouseButtonSecondary
	MouseButtonMiddle
	MouseButtonForward
	MouseButtonBackward
	MouseButtonLeft  = MouseButtonPrimary
	MouseButtonRight = MouseButtonSecondary
)

var (
	mouseButtonNames = []string{
		"none",
		"primary",
		"secondary",
		"middle",
		"forward",
		"backward",
	}
)

func (b MouseButton) String() string {
	return mouseButtonNames[int(b)]
}

// Pressable returns true if the button is pressable (i.e. clickable)
func (b MouseButton) Pressable() bool {
	return b != MouseButtonNone
}

// MouseEvent describes events received when a mouse moved, clicked or
// released.
type MouseEvent interface {
	Event
	KeyModifiable
	// Button returns the mouse button that was depressed, if any.
	Button() MouseButton
	// SetButton sets the mouse button that was depressed when the event fired.
	SetButton(MouseButton)
	// Position returns where the mouse was when the MouseEvent was triggered.
	Position() Point
	// SetPosition sets the coordinates of the mouse when the event fired.
	SetPosition(Point)
}

// MouseHoverEvent describes events received when the mouse hovers (or stops
// hovering) over something.
type MouseHoverEvent interface {
	MouseEvent
	// Hovered returns true if the receiver of the event has the mouse hovering
	// over it.
	Hovered() bool
}

// MouseClickEvent describes a mouse event for when the user clicked or
// double-clicked a mouse button.
type MouseClickEvent interface {
	MouseEvent
	// DoubleClicked returns true if the user double-clicked.
	DoubleClicked() bool
}

// MouseDragEvent describes a mouse event for when the user held a mouse button
// down and moved the mouse.
type MouseDragEvent interface {
	MouseEvent
	// Start returns the MouseDragStartEvent associated with the start of the
	// drag action.
	Start() MouseEvent
}

// MouseEventWithOption describes an optional varg parameter to
// [core.event.mouse.New] that modifies the returned MouseEvent.
type MouseEventWithOption func(MouseEvent)

// MouseEventCallback is the function signature for callbacks executed on mouse
// events.
type MouseEventCallback func(context.Context, MouseEvent)

// MouseHoverEventCallback is the function signature for callbacks executed on
// mouse hover events.
type MouseHoverEventCallback func(context.Context, MouseHoverEvent)

// MouseClickEventCallback is the function signature for callbacks executed on
// mouse click and double-click events.
type MouseClickEventCallback func(context.Context, MouseClickEvent)

// MouseDragStopEventCallback is the function signature for callbacks executed
// on mouse drag move or stop events.
type MouseDragEventCallback func(context.Context, MouseDragEvent)

// MouseEventHandler represents something that can handle mouse events.
type MouseEventHandler interface {
	// MouseHover executes any OnMouseHover callbacks that were registered for
	// the MouseEventHandler.
	MouseHover(context.Context, MouseHoverEvent)
	// OnMouseHover registers a callback that will be executed when the mouse
	// is over top of an element but the element does *not* have the focus.
	OnMouseHover(MouseHoverEventCallback)
	// MouseClick executes any OnMouseClick callbacks that were registered for
	// the MouseEventHandler.
	MouseClick(context.Context, MouseClickEvent)
	// OnMouseClick registers a callback that will be executed when the
	// MouseEventHandler is clicked.
	OnMouseClick(MouseClickEventCallback)
	// MouseDoubleClick executes any OnMouseDoubleClick callbacks that were
	// registered for the MouseEventHandler.
	MouseDoubleClick(context.Context, MouseClickEvent)
	// OnMouseDoubleClick registers a callback that will be executed when the
	// MouseEventHandler is double-clicked.
	OnMouseDoubleClick(MouseClickEventCallback)
	// MouseScroll executes any OnMouseScroll callbacks that were registered for
	// the MouseEventHandler.
	MouseScroll(context.Context, MouseEvent)
	// OnMouseScroll registers a callback that will be executed when the mouse
	// wheel is scrolled and the MouseEventHandler has the focus.
	OnMouseScroll(MouseEventCallback)
	// MouseDragMove executes any OnMouseDragMove callbacks that were
	// registered for the MouseEventHandler.
	MouseDragMove(context.Context, MouseDragEvent)
	// OnMouseDragMove registers a callback that will be executed when a mouse
	// button is held down and the mouse is moved.
	OnMouseDragMove(MouseDragEventCallback)
	// MouseDragStop executes any OnMouseDragStop callbacks that were
	// registered for the MouseEventHandler.
	MouseDragStop(context.Context, MouseDragEvent)
	// OnMouseDragStop registers a callback that will be executed when the
	// mouse button is released after dragging the mouse.
	OnMouseDragStop(MouseDragEventCallback)
}
