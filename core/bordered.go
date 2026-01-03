package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Bordered describes something that has a bounding box.
type Bordered struct {
	// border is the optional Border information for the Bordered.
	border *types.Border
}

// SetRect sets the Element's bounding rectangle
func (b *Bordered) SetBorder(border types.Border) {
	b.border = &border
}

// Border returns the Bordered's border, if any.
func (b *Bordered) Border() *types.Border {
	return b.border
}

var _ types.Bordered = (*Bordered)(nil)
