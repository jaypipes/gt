package types

import (
	"image"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
)

// Convenience wrappers around common ultraviolet and other package
// structs/funcs
type (
	Rectangle = image.Rectangle
	Point     = image.Point

	Underline = ansi.Underline

	Screen         = uv.Screen
	SizeConstraint = uv.Constraint
	Fixed          = uv.Fixed
	Percent        = uv.Percent
	Border         = uv.Border
	Side           = uv.Side
	// Style is semantically slightly different from uv.Style in that a Style
	// can apply to an entire Element not just a single uv.Cell.
	Style        = uv.Style
	StyledString = uv.StyledString
)

const (
	UnderlineNone   = ansi.UnderlineNone
	UnderlineSingle = ansi.UnderlineSingle
	UnderlineDouble = ansi.UnderlineDouble
	UnderlineCurly  = ansi.UnderlineCurly
	UnderlineDotted = ansi.UnderlineDotted
	UnderlineDashed = ansi.UnderlineDashed
)
