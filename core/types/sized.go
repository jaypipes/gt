package types

// Sized describes something with a width (in cells) and height (in lines).
type Sized interface {
	// SetSize sets the Sized's size.
	SetSize(int, int)
	// SetWidth sets the Sized's width in cells.
	SetWidth(int)
	// SetHeight sets the Sized's height in lines.
	SetHeight(int)
	// SetWidthConstraint sets the Sized's width size constraint.
	SetWidthConstraint(SizeConstraint)
	// SetHeightConstraint sets the Sized's height size constraint.
	SetHeightConstraint(SizeConstraint)
	// Height returns the height of the Sized.
	Height() int
	// Width returns the width of the Sized.
	Width() int
	// Bounds returns the width and height of the Sized.
	Size() Size
	// WidthConstraint returns the width size constraint of the Sized, if any.
	WidthConstraint() *SizeConstraint
	// HeightConstraint returns the height size constraint of the Sized, if
	// any.
	HeightConstraint() *SizeConstraint
}
