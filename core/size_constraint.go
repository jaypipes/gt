package core

import "github.com/jaypipes/gt/types"

// SizeConstraint describes a pair of DimensionConstraints on a Size's width
// and height.
type SizeConstraint struct {
	// width is the SizeConstraint's width DimensionConstraint.
	width types.DimensionConstraint
	// height is the SizeConstraint's height DimensionConstraint.
	height types.DimensionConstraint
}

// Width returns the SizeConstraint's width constraint
func (c SizeConstraint) Width() types.DimensionConstraint {
	return c.width
}

// Height returns the SizeConstraint's height constraint
func (c SizeConstraint) Height() types.DimensionConstraint {
	return c.height
}

// PercentArea returns a SizeConstraint that constrains a Size to a percentage
// of available remaining width and height.
func PercentArea(width uint, height uint) SizeConstraint {
	return SizeConstraint{
		width:  Percent(width),
		height: Percent(height),
	}
}

// PercentWidth returns a SizeConstraint that constrains a Size to a percentage
// of available remaining width but does not constrain the height.
func PercentWidth(width uint) SizeConstraint {
	return SizeConstraint{
		width: Percent(width),
	}
}

// PercentHeight returns a SizeConstraint that constrains a Size to a percentage
// of available remaining height but does not constrain the width.
func PercentHeight(height uint) SizeConstraint {
	return SizeConstraint{
		height: Percent(height),
	}
}

// FixedArea returns a SizeConstraint that constrains a Size to a fixed width
// and height.
func FixedArea(width uint, height uint) SizeConstraint {
	return SizeConstraint{
		width:  Fixed(width),
		height: Fixed(height),
	}
}

// FixedWidth returns a SizeConstraint that constrains a Size to a fixed width
// but does not constrain the height.
func FixedWidth(width uint) SizeConstraint {
	return SizeConstraint{
		width: Fixed(width),
	}
}

// FixedHeight returns a SizeConstraint that constrains a Size to a fixed
// height but does not constrain the width.
func FixedHeight(height uint) SizeConstraint {
	return SizeConstraint{
		height: Fixed(height),
	}
}

var NoSizeConstraint = SizeConstraint{
	width:  NoDimensionConstraint,
	height: NoDimensionConstraint,
}
