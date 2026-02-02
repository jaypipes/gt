package box

import (
	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/types"
)

// SetSize constrains the size of the Box's inner bounding box.
func (b *Box) SetSize(constraint types.SizeConstraint) {
	wc := constraint.Width()
	if wc != nil {
		b.widthConstraint = wc
	}
	hc := constraint.Height()
	if hc != nil {
		b.heightConstraint = hc
	}
}

// Size returns the width and height of the Box.
func (b *Box) Size() types.Size {
	return types.Size{
		W: int(b.Width()),
		H: int(b.Height()),
	}
}

// SetWidth constrains the width of the Box.
func (b *Box) SetWidth(constraint types.DimensionConstraint) {
	b.widthConstraint = constraint
}

// HasFixedWidth returns true if the Box's inner bounding box has a fixed
// width.
func (b *Box) HasFixedWidth() bool {
	_, ok := b.widthConstraint.(core.FixedConstraint)
	return ok
}

// FixedWidth returns the Box's fixed width. If the Box does not have a
// fixed width constraint, returns 0.
func (b *Box) FixedWidth() types.Dimension {
	if !b.HasFixedWidth() {
		return types.Dimension(0)
	}
	return types.Dimension(b.widthConstraint.(core.FixedConstraint))
}

// HasPercentWidth returns true if the Box's inner bounding box has a percent
// width.
func (b *Box) HasPercentWidth() bool {
	_, ok := b.widthConstraint.(core.PercentConstraint)
	return ok
}

// PercentWidth returns the Box's fixed width. If the Box does not have a
// percent width constraint, returns 0.
func (b *Box) PercentWidth() types.Dimension {
	if !b.HasPercentWidth() {
		return types.Dimension(0)
	}
	return types.Dimension(b.widthConstraint.(core.PercentConstraint))
}

// SetMinWidth sets the minimum width of the Box.
func (b *Box) SetMinWidth(w types.Dimension) {
	b.minWidth = w
}

// MinWidth returns the Box's minimum width.
func (b *Box) MinWidth() types.Dimension {
	return b.minWidth
}

// WidthConstraint returns any optional size constraint for the Box's
// width.  Returns nil when there is no width constraint.
func (b *Box) WidthConstraint() types.DimensionConstraint {
	return b.widthConstraint
}

// Width returns the width of the Box's outer bounding box.
func (b *Box) Width() types.Dimension {
	return types.Dimension(b.Bounds().Dx())
}

// SetHeight constrains the height of the Box.
func (b *Box) SetHeight(constraint types.DimensionConstraint) {
	b.heightConstraint = constraint
}

// HasFixedHeight returns true if the Box's inner bounding box has a fixed
// height.
func (b *Box) HasFixedHeight() bool {
	_, ok := b.heightConstraint.(core.FixedConstraint)
	return ok
}

// FixedHeight returns the Box's fixed height. If the Box does not have
// a fixed height constraint, returns 0.
func (b *Box) FixedHeight() types.Dimension {
	if !b.HasFixedHeight() {
		return types.Dimension(0)
	}
	return types.Dimension(b.heightConstraint.(core.FixedConstraint))
}

// HasPercentHeight returns true if the Box's inner bounding box has a percent
// height.
func (b *Box) HasPercentHeight() bool {
	_, ok := b.heightConstraint.(core.PercentConstraint)
	return ok
}

// PercentHeight returns the Box's percent height. If the Box does not
// have a percent height constraint, returns 0.
func (b *Box) PercentHeight() types.Dimension {
	if !b.HasPercentHeight() {
		return types.Dimension(0)
	}
	return types.Dimension(b.heightConstraint.(core.PercentConstraint))
}

// SetMinHeight sets the minimum height of the Box.
func (b *Box) SetMinHeight(h types.Dimension) {
	b.minHeight = h
}

// MinHeight returns the Box's minimum height.
func (b *Box) MinHeight() types.Dimension {
	return b.minHeight
}

// HeightConstraint returns any optional size constraint for the Box's
// height. Returns nil when there is no height constraint.
func (b *Box) HeightConstraint() types.DimensionConstraint {
	return b.heightConstraint
}

// Height returns the height of the Box's outer bounding box.
func (b *Box) Height() types.Dimension {
	return types.Dimension(b.Bounds().Dy())
}
