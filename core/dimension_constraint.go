package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Percent returns a PercentConstraint representing a percentage of an
// available remaining amount of a dimension.
func Percent(p uint) PercentConstraint {
	return PercentConstraint(types.Dimension(p))
}

// PercentConstraint implements DimensionConstraint a percentage of an
// available remaining amount of the dimension.
type PercentConstraint uint

// Apply applies the percentage constraint to the given dimension.
func (p PercentConstraint) Apply(d types.Dimension) types.Dimension {
	if p > 100 {
		return d
	}
	return types.Dimension(int(d) * int(p) / 100)
}

// Fixed returns a FixedConstraint representing a fixed amount of a dimension.
func Fixed(p uint) FixedConstraint {
	return FixedConstraint(types.Dimension(p))
}

// FixedConstraint implements DimensionConstraint and represents a fixed amount
// of a dimension.
type FixedConstraint uint

// Apply applies the fixed size constraint to the given dimension. If the fixed
// size is greater than the supplied dimension, returns the supplied dimension.
func (f FixedConstraint) Apply(d types.Dimension) types.Dimension {
	if uint(f) > uint(d) {
		return d
	}
	return types.Dimension(f)
}

type noDimensionConstraint uint

func (f noDimensionConstraint) Apply(d types.Dimension) types.Dimension {
	return d
}

const NoDimensionConstraint = noDimensionConstraint(0)
