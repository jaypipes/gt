package types

import (
	"image"
	"image/color"

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
)

var (
	Rect = image.Rect
)

const (
	CursorStyleBar       = tcell.CursorStyleSteadyBar
	CursorStyleSteadyBar = tcell.CursorStyleSteadyBar
)

const (
	UnderlineStyleNone   = tcell.UnderlineStyleNone
	UnderlineStyleSolid  = tcell.UnderlineStyleSolid
	UnderlineStyleDouble = tcell.UnderlineStyleDouble
	UnderlineStyleCurly  = tcell.UnderlineStyleCurly
	UnderlineStyleDotted = tcell.UnderlineStyleDotted
	UnderlineStyleDashed = tcell.UnderlineStyleDashed
)
