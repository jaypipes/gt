package types

import (
	"context"
	"time"

	"github.com/gdamore/tcell/v3"
)

// NOTE(jaypipes): Some of this code adapted from:
// https://github.com/rivo/tview/blob/f39b95c73dbb30877f4b5145b835333002afb2a8/application.go

// DefaultMouseDoubleClickInterval specifies the default maximum time between
// clicks to register a double click rather than click.
var DefaultMouseDoubleClickInterval = 500 * time.Millisecond

// MouseButton indicates the mouse button that was activated.
type MouseButton int16

const (
	MouseButtonNone      = MouseButton(tcell.ButtonNone)
	MouseButtonLeft      = MouseButton(tcell.ButtonPrimary)
	MouseButtonPrimary   = MouseButton(tcell.ButtonPrimary)
	MouseButtonMiddle    = MouseButton(tcell.ButtonMiddle)
	MouseButtonSecondary = MouseButton(tcell.ButtonSecondary)
	MouseButtonRight     = MouseButton(tcell.ButtonSecondary)
	MouseWheelUp         = MouseButton(tcell.WheelUp)
	MouseWheelDown       = MouseButton(tcell.WheelDown)
	MouseWheelLeft       = MouseButton(tcell.WheelLeft)
	MouseWheelRight      = MouseButton(tcell.WheelRight)
	MouseButtonBackward  = MouseButton(tcell.Button4)
	MouseButtonForward   = MouseButton(tcell.Button5)
)

var (
	mouseButtonNames = []string{
		"none",
		"primary",
		"primary",
		"middle",
		"secondary",
		"secondary",
		"wheel-up",
		"wheel-down",
		"wheel-left",
		"wheel-right",
		"backward",
		"forward",
	}
)

func (b MouseButton) String() string {
	return mouseButtonNames[int(b)]
}

// Pressable returns true if the button is pressable (i.e. clickable)
func (b MouseButton) Pressable() bool {
	return b == MouseButtonPrimary ||
		b == MouseButtonMiddle ||
		b == MouseButtonSecondary
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

// MouseClickEvent describes a mouse event for when the user clicked or
// double-clicked a mouse button.
type MouseClickEvent interface {
	MouseEvent
	// DoubleClicked returns true if the user double-clicked.
	DoubleClicked() bool
}

// MouseDragMoveEvent describes a mouse event for when the user held a mouse
// button down and moved the mouse.
type MouseDragMoveEvent interface {
	MouseEvent
	// Start returns the MouseDragStartEvent associated with the start of the
	// drag action.
	Start() MouseEvent
}

// MouseDragStopEvent describes a mouse event for when the user released the mouse
// button after dragging the mouse.
type MouseDragStopEvent interface {
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

// MouseClickEventCallback is the function signature for callbacks executed on
// mouse click and double-click events.
type MouseClickEventCallback func(context.Context, MouseClickEvent)

// MouseDragMoveEventCallback is the function signature for callbacks executed
// on mouse drag move events.
type MouseDragMoveEventCallback func(context.Context, MouseDragMoveEvent)

// MouseDragStopEventCallback is the function signature for callbacks executed
// on mouse drag stop events.
type MouseDragStopEventCallback func(context.Context, MouseDragStopEvent)

// MouseEventHandler represents something that can handle mouse events.
type MouseEventHandler interface {
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
	MouseDragMove(context.Context, MouseDragMoveEvent)
	// OnMouseDragMove registers a callback that will be executed when a mouse
	// button is held down and the mouse is moved.
	OnMouseDragMove(MouseDragMoveEventCallback)
	// MouseDragStop executes any OnMouseDragStop callbacks that were
	// registered for the MouseEventHandler.
	MouseDragStop(context.Context, MouseDragStopEvent)
	// OnMouseDragStop registers a callback that will be executed when the
	// mouse button is released after dragging the mouse.
	OnMouseDragStop(MouseDragStopEventCallback)
}
