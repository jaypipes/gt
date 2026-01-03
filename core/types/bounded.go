package types

// Bounded describes something contained in a bounding box.
type Bounded interface {
	// SetBounds sets the Bounded's bounding box.
	SetBounds(Rectangle)
	// Bounds returns the bounding box for the Bounded.
	Bounds() Rectangle
}
