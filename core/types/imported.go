package types

import (
	"image"

	uv "github.com/charmbracelet/ultraviolet"
)

// Convenience wrappers around common ultraviolet and core image package
// structs/funcs
type (
	Rectangle    = image.Rectangle
	Point        = image.Point
	Screen       = uv.Screen
	Constraint   = uv.Constraint
	Fixed        = uv.Fixed
	Percent      = uv.Percent
	Border       = uv.Border
	Side         = uv.Side
	StyledString = uv.StyledString
)
