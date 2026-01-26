package base

import "github.com/jaypipes/gt/types"

// WithParent sets the Element's parent and index of the Box within the parent's
// children and returns the Element.
func (b *Base) WithParent(parent types.Element, childIndex int) types.Element {
	b.Box.SetParent(parent, childIndex)
	return b
}
