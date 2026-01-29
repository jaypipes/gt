package types

// Bounded has a bounding box.
type Bounded interface {
	// Bounds returns the Bounded's outer bounding box.
	Bounds() Rectangle
}
