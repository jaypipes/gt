package palette

import (
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a Palette.
//
// You can pass zero or more PaletteWithOptions to optionally set certain
// attributes on the returned Palette.
func New(opts ...types.PaletteWithOption) *Palette {
	p := &Palette{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// WithColors sets the Palette's set of colors.
func WithColors(colors types.PaletteColors) types.PaletteWithOption {
	return func(p types.Palette) {
		p.SetColors(colors)
	}
}
