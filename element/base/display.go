package base

import "github.com/jaypipes/gt/types"

// SetDisplayMode sets the display mode of the Element
func (b *Base) SetDisplay(display types.Display) types.Element {
	b.display = display
	return b
}

// DisplayMode returns the display mode of the Element
func (b *Base) Display() types.Display {
	return b.display
}
