package core

import (
	"fmt"

	"github.com/jaypipes/gt/core/types"
)

// Aligned describes something that can be horizontally, vertically and/or
// item-aligned.
type Aligned struct {
	// alignment is the alignment mode of the Aligned
	alignment types.Alignment
}

func (a *Aligned) String() string {
	return fmt.Sprintf("align=%s", a.alignment)
}

// SetAlignment sets the Aligned's alignment mode.
func (a *Aligned) SetAlignment(alignment types.Alignment) {
	a.alignment = alignment
}

// Alignment returns the Aligned's Alignment.
func (a *Aligned) Alignment() types.Alignment {
	return a.alignment
}

var _ types.Aligned = (*Aligned)(nil)
