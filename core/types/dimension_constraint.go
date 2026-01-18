package types

// DimensionConstraint describes something that can have both a width or height
// dimension constrained to a given size.
type DimensionConstraint interface {
	// Apply applies the constraint to the given dimension and returns the
	// constrained dimension value.
	Apply(Dimension) Dimension
}
