package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Padded describes something that has a bounding box.
type Padded struct {
	// padding is any padding applied to the Padded.
	padding types.Padding
}

// Padding returns the padding for the Padded.
func (p *Padded) Padding() types.Padding {
	return p.padding
}

// SetPadding sets the Padded's padding.
func (p *Padded) SetPadding(padding types.Padding) {
	p.padding = padding
}

var _ types.Padded = (*Padded)(nil)
