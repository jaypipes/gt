package types

// Displayed describes how a renderable box is displayed.
type Displayed interface {
	// DisplayMode returns the display mode of the Displayed
	Display() Display
	// SetDisplayMode sets the display mode of the Displayed
	SetDisplay(Display)
}
