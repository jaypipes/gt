package core

import (
	"image"

	"github.com/jaypipes/gt/core/types"
)

// Bounded describes something that has a bounding box.
type Bounded struct {
	// bounds is the outermost bounding box of the Bounded.
	bounds image.Rectangle
}

// SetBounds sets the Bounded's bounding box.
func (b *Bounded) SetBounds(r image.Rectangle) {
	b.bounds = r
}

// Bounds returns the bounding box for the Bounded
func (b *Bounded) Bounds() image.Rectangle {
	return b.bounds
}

var _ types.Bounded = (*Bounded)(nil)
