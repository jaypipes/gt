package types

// Bordered describes something contained in a bounding box.
type Bordered interface {
	// SetBounds sets the Bordered's bounding box.
	SetBorder(Border)
	// Bounds returns the bounding box for the Bordered.
	Border() *Border
}
