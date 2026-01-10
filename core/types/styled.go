package types

import "image/color"

// Styled describes something that can have a Style applied to it
type Styled interface {
	// Style returns the thing's Style
	Style() Style
	// SetStyle applies the supplied Style to the Styled.
	SetStyle(Style)
	// SetForegroundColor sets the thing's foreground color
	SetForegroundColor(color.Color)
	// SetBackgroundColor sets the thing's background color
	SetBackgroundColor(color.Color)
}
