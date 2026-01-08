package core

import (
	"fmt"

	"github.com/jaypipes/gt/core/types"
)

// Displayed describes how a renderable box is displayed.
type Displayed struct {
	display types.Display
}

func (d *Displayed) String() string {
	return fmt.Sprintf("display=%s", d.display)
}

// DisplayMode returns the display mode of the Displayed
func (d *Displayed) Display() types.Display {
	return d.display
}

// SetDisplayMode sets the display mode of the Displayed
func (d *Displayed) SetDisplay(display types.Display) {
	d.display = display
}
