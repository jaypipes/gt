package box

import "github.com/jaypipes/gt/types"

// SetAlignment sets the Box's alignment modb.
func (b *Box) SetAlignment(alignment types.Alignment) {
	b.alignment = alignment
}

// Alignment returns the Box's Alignment.
func (b *Box) Alignment() types.Alignment {
	return b.alignment
}
