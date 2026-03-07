package types

import (
	"image"
	"image/color"

	"github.com/charmbracelet/x/ansi"
	"github.com/gdamore/tcell/v3"
)

// Convenience wrappers around common tcell and other package structs/funcs
type (
	Color     = color.Color
	Rectangle = image.Rectangle
	Point     = image.Point

	Screen         = tcell.Screen
	Key            = tcell.Key
	CursorStyle    = tcell.CursorStyle
	UnderlineStyle = tcell.UnderlineStyle

	MouseButton = ansi.MouseButton
)

const (
	CursorStyleBar       = tcell.CursorStyleSteadyBar
	CursorStyleSteadyBar = tcell.CursorStyleSteadyBar
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
	UnderlineStyleNone   = tcell.UnderlineStyleNone
	UnderlineStyleSolid  = tcell.UnderlineStyleSolid
	UnderlineStyleDouble = tcell.UnderlineStyleDouble
	UnderlineStyleCurly  = tcell.UnderlineStyleCurly
	UnderlineStyleDotted = tcell.UnderlineStyleDotted
	UnderlineStyleDashed = tcell.UnderlineStyleDashed
)
