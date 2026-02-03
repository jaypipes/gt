package types

import (
	"image"
	"image/color"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
)

// Convenience wrappers around common ultraviolet and other package
// structs/funcs
type (
	Color     = color.Color
	Rectangle = image.Rectangle
	Point     = image.Point

	Underline = ansi.Underline

	Screen = uv.Screen
	Event  = uv.Event

	Mouse             = uv.Mouse
	MouseMode         = uv.MouseMode
	MouseEvent        = uv.MouseEvent
	MouseClickEvent   = uv.MouseClickEvent
	MouseReleaseEvent = uv.MouseReleaseEvent
	MouseMotionEvent  = uv.MouseMotionEvent
	MouseButton       = ansi.MouseButton

	Border = uv.Border
	Side   = uv.Side
	// Style is semantically slightly different from uv.Style in that a Style
	// can apply to an entire Element not just a single uv.Cell.
	Style        = uv.Style
	StyledString = uv.StyledString
)

const (
	MouseNone       = ansi.MouseNone
	MouseLeft       = ansi.MouseLeft
	MouseMiddle     = ansi.MouseMiddle
	MouseRight      = ansi.MouseRight
	MouseWheelUp    = ansi.MouseWheelUp
	MouseWheelDown  = ansi.MouseWheelDown
	MouseWheelLeft  = ansi.MouseWheelLeft
	MouseWheelRight = ansi.MouseWheelRight
	MouseBackward   = ansi.MouseBackward
	MouseForward    = ansi.MouseForward
)

const (
	UnderlineNone   = ansi.UnderlineNone
	UnderlineSingle = ansi.UnderlineSingle
	UnderlineDouble = ansi.UnderlineDouble
	UnderlineCurly  = ansi.UnderlineCurly
	UnderlineDotted = ansi.UnderlineDotted
	UnderlineDashed = ansi.UnderlineDashed
)
