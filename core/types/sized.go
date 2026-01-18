package types

// Sized describes something with a width (in cells) and height (in lines).
type Sized interface {
	String() string
	// SetSize constrains the size of the Sized.
	SetSize(SizeConstraint)
	// SetWidth constrains the width of the Sized.
	SetWidth(DimensionConstraint)
	// SetMinWidth sets the minimum width of the Sized.
	SetMinWidth(Dimension)
	// SetHeight constrains the height of the Sized.
	SetHeight(DimensionConstraint)
	// SetMinHeight sets the minimum height of the Sized.
	SetMinHeight(Dimension)
	// HasFixedWidth returns true if the Sized has a fixed width.
	HasFixedWidth() bool
	// Width returns the Sized's width.
	Width() Dimension
	// FixedWidth returns the Sized's fixed width. If the Sized does not have a
	// fixed width constraint, returns 0.
	FixedWidth() Dimension
	// MinWidth returns the Sized's minimum width.
	MinWidth() Dimension
	// WidthConstraint returns any DimensionConstraint set for the Sized's
	// width.
	WidthConstraint() DimensionConstraint
	// HasFixedHeight returns true if the Sized has a fixed height.
	HasFixedHeight() bool
	// Height returns the Sized's height.
	Height() Dimension
	// FixedHeight returns the Sized's fixed height. If the Sized does not have
	// a fixed height constraint, returns 0.
	FixedHeight() Dimension
	// MinHeight returns the Sized's minimum height.
	MinHeight() Dimension
	// HeightConstraint returns any DimensionConstraint set for the Sized's
	// height.
	HeightConstraint() DimensionConstraint
}
