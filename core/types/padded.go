package types

// Padded describes something contained in a bounding box.
type Padded interface {
	// SetPadding sets the Padded's padding.
	SetPadding(Padding)
	// Padding returns the padding for the Padded.
	Padding() Padding
}
