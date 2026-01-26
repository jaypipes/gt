package box

import "github.com/jaypipes/gt/types"

// SetDisplayMode sets the display mode of the Box
func (b *Box) SetDisplay(display types.Display) {
	b.display = display
}

// DisplayMode returns the display mode of the Box
func (b *Box) Display() types.Display {
	return b.display
}
