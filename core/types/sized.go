package types

// Sized describes something with a width (in cells) and height (in lines).
type Sized interface {
	String() string
	// SetSize sets the Sized's size and marks the Sized as having a fixed
	// width and height.
	SetSize(int, int)
	// SetWidth sets the Sized's width in cells and marks the Sized as having a
	// fixed width.
	SetWidth(int)
	// SetHeight sets the Sized's height in lines and marks the Sized as having
	// a fixed height.
	SetHeight(int)
	// SetWidthConstraint sets the Sized's width size constraint.
	SetWidthConstraint(SizeConstraint)
	// SetHeightConstraint sets the Sized's height size constraint.
	SetHeightConstraint(SizeConstraint)
	// Width returns the current width of the Sized.
	Width() int
	// Height returns the current height of the Sized.
	Height() int
	// Size returns the current width and height for the Sized.
	Size() Size
	// FixedWidth returns true if the Sized is using a fixed width.
	FixedWidth() bool
	// FixedHeight returns true if the Sized is using a fixed height.
	FixedHeight() bool
	// WidthConstraint returns the width size constraint of the Sized, if any.
	WidthConstraint() SizeConstraint
	// HeightConstraint returns the height size constraint of the Sized, if
	// any.
	HeightConstraint() SizeConstraint
}
