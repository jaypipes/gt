package base

import "github.com/jaypipes/gt/types"

// SetAlignment sets the Element's alignment modb.
func (b *Base) SetAlignment(alignment types.Alignment) types.Element {
	b.alignment = alignment
	return b
}

// Alignment returns the Element's Alignment.
func (b *Base) Alignment() types.Alignment {
	return b.alignment
}
