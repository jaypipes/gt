package element

import "github.com/jaypipes/gt/types"

// WithParent sets the Element's parent and index of the Box within the parent's
// children and returns the Element.
func (e *Element) WithParent(parent types.Element, childIndex int) types.Element {
	e.Box.SetParent(parent, childIndex)
	return e
}
