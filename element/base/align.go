package base

import "github.com/jaypipes/gt/types"

// WithAlignment sets the Element's alignment mode and returns the Element.
func (b *Base) WithAlignment(alignment types.Alignment) types.Element {
	b.Box.SetAlignment(alignment)
	return b
}
