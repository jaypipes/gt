package types

// SizeConstraint describes something that can have both a width and height
// dimension constrained to a given size.
type SizeConstraint interface {
	// WidthConstraint returns the SizeConstraint's width DimensionConstraint.
	Width() DimensionConstraint
	// HeightConstraint returns the SizeConstraint's height
	// DimensionConstraint.
	Height() DimensionConstraint
}
