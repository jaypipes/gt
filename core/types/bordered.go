package types

// Bordered describes something contained in a bounding box.
type Bordered interface {
	// SetBounds sets the Bordered's bounding box.
	SetBorder(Border)
	// Bounds returns the bounding box for the Bordered.
	Border() *Border
	// SetBorderForegroundColor sets the Bordered's border foreground color
	// (i.e the color of the border's cells underlying grapheme).
	SetBorderForegroundColor(Color)
	// BorderForegroundColor returns the Bordered's border foreground color.
	BorderForegroundColor() Color
	// SetBorderBackgroundColor sets the Bordered's border background color
	// (i.e the background color of the border's cells.
	SetBorderBackgroundColor(Color)
	// BorderBackgroundColor returns the Bordered's border background color.
	BorderBackgroundColor() Color
}
