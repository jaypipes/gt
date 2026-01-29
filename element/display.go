package element

import "github.com/jaypipes/gt/types"

// WithDisplayMode sets the display mode of the Element and returns the
// Element.
func (e *Element) WithDisplay(display types.Display) types.Element {
	e.Box.SetDisplay(display)
	return e
}
