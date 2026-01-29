package element

import "github.com/jaypipes/gt/types"

// WithAlignment sets the Element's alignment mode and returns the Element.
func (e *Element) WithAlignment(alignment types.Alignment) types.Element {
	e.Box.SetAlignment(alignment)
	return e
}
