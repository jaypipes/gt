package palette

import (
	"github.com/jaypipes/gt/types"
)

// Palette represents a related set of 7 grayscale and 8 colors.
type Palette struct {
	// colors contains the Palette's colors.
	colors types.PaletteColors
}

// Colors returns the Palette's set of colors.
func (p *Palette) Colors() types.PaletteColors {
	return p.colors
}

// SetColors sets the Palette's set of colors.
func (p *Palette) SetColors(colors types.PaletteColors) {
	p.colors = colors
}

// Grayscale returns the color of the Palette's grayscale color set at the
// supplied index, with 0 being the darkest color and 6 being the lightest
// color.
func (p *Palette) Grayscale(index int) types.Color {
	index = min(max(0, index), 6)
	return p.colors[index]
}

// Color returns the color at the specified index.
func (p *Palette) Color(index int) types.Color {
	index = min(max(0, index), 15)
	return p.colors[index]
}

var _ types.Palette = (*Palette)(nil)
