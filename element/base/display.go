package base

import "github.com/jaypipes/gt/types"

// WithDisplayMode sets the display mode of the Element and returns the
// Element.
func (b *Base) WithDisplay(display types.Display) types.Element {
	b.Box.SetDisplay(display)
	return b
}
