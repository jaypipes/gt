package types

import (
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

// MouseAction indicates one of the actions the mouse is logically doing.
type MouseAction int16

const (
	MouseActionNone MouseAction = iota
	MouseActionMove
	MouseActionLeftDown
	MouseActionLeftUp
	MouseActionLeftClick
	MouseActionLeftDoubleClick
	MouseActionMiddleDown
	MouseActionMiddleUp
	MouseActionMiddleClick
	MouseActionMiddleDoubleClick
	MouseActionRightDown
	MouseActionRightUp
	MouseActionRightClick
	MouseActionRightDoubleClick
	MouseActionScrollUp
	MouseActionScrollDown
	MouseActionScrollLeft
	MouseActionScrollRight
)

var (
	mouseActionNames = []string{
		"none",
		"move",
		"left-down",
		"left-up",
		"left-click",
		"left-double-click",
		"middle-down",
		"middle-up",
		"middle-click",
		"middle-double-click",
		"right-down",
		"right-up",
		"right-click",
		"right-double-click",
		"scroll-up",
		"scroll-down",
		"scroll-left",
		"scroll-right",
	}
)

func (a MouseAction) String() string {
	return mouseActionNames[int(a)]
}

func (a MouseAction) MouseDown() bool {
	return a == MouseActionLeftDown ||
		a == MouseActionRightDown ||
		a == MouseActionMiddleDown
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
	// Action returns the semantic action that was taken by the user.
	Action() MouseAction
	// SetAction sets the semantic action that was taken by the user.
	SetAction(MouseAction)
}

// MouseEventWithOption describes an optional varg parameter to
// [core.event.mouse.New] that modifies the returned MouseEvent.
type MouseEventWithOption func(MouseEvent)
