package types

// Plotted describes something that can be plotted on a Screen.
type Plotted interface {
	Aligned
	Bounded
	Displayed
	Sized
	// Anchor sets the Plotted's anchor point (i.e. its top-left grid
	// coordinates) and marks the Plotted as using absolute positioning.
	Anchor(Point)
	// AbsolutePositioned returns true if the Plotted used absolute
	// positioning.
	AbsolutePositioned() bool
	// InnerBounds returns the inner bounding box for the Plotted, which
	// accounts for any border and padding.
	InnerBounds() Rectangle
	// LeftMargin returns the distance, in cells, from the left edge of the
	// outer bounds to left edge of the inner bounds.
	LeftMargin() int
	// RightMargin returns the distance, in cells, from the right edge of the
	// outer bounds to right edge of the inner bounds.
	RightMargin() int
	// TopMargin returns the distance, in cells, from the top edge of the
	// outer bounds to top edge of the inner bounds.
	TopMargin() int
	// BottomMargin returns the distance, in cells, from the bottom edge of the
	// outer bounds to bottom edge of the inner bounds.
	BottomMargin() int
}
