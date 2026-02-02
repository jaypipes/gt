package types

import (
	"image/color"
)

// Styleable represents something that has a foreground and background color.
type Styleable interface {
	// SetStyle sets the Styleable's style.
	SetStyle(Style)
	// Style returns the thing's Style
	Style() Style
	// SetForegroundColor sets the Styleable's foreground color.
	SetForegroundColor(color.Color)
	// ForegroundColor returns the Styleable's foreground color.
	ForegroundColor() Color
	// SetBackgroundColor sets the Styleable's background color.
	SetBackgroundColor(color.Color)
	// BackgroundColor returns the Styleable's background color.
	BackgroundColor() Color
}
