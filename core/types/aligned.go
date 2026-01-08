package types

// Aligned things can have their contents horizontally and vertically aligned
type Aligned interface {
	// Alignment returns the Aligned's Alignment
	Alignment() Alignment
	// SetAlignment sets the Aligneds' Alignment
	SetAlignment(Alignment)
}
