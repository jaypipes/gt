package types

import "image/color"

type PaletteColors [16]color.Color

// Palette represents a related set of 7 grayscale and 8 colors.
type Palette interface {
	// Colors returns the Palette's set of colors.
	Colors() PaletteColors
	// SetColors sets the Palette's set of colors.
	SetColors(PaletteColors)
	// Grayscale returns the color of the Palette's grayscale color set at the
	// supplied index, with 0 being the darkest color and 6 being the lightest
	// color.
	Grayscale(int) Color
	// Color returns the color at the specified index.
	Color(int) Color
}

// PaletteWithOption describes an optional varg parameter to [palette.New] that
// modifies the returned Palette.
type PaletteWithOption func(Palette)
