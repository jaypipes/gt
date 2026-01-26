package box

import (
	"context"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
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
		W: int(b.OuterWidth()),
		H: int(b.OuterHeight()),
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

// OuterWidth returns the width of the Box's outer bounding box.
//
// If a fixed width has been set and the display mode is not "inline", we use
// the fixed width plus any horizontal padding and left-right border width.
//
// If a percent width has been set and the display mode is not "inline", we
// calculate the width by looking at the siblings and subtracting any fixed
// width siblings from the parent's available width.
//
// If a fixed width has not been set and the display mode is `block` or
// `inline-block`, the width defaults to remaining horizontal space in the
// parent's inner bounding box.
func (b *Box) OuterWidth() types.Dimension {
	parent := b.Parent()
	if parent == nil {
		return types.Dimension(b.Bounds().Dx())
	}
	next := b.NextSibling()
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())
	horizSpace := b.HorizontalSpace()

	ctx := context.TODO()
	display := b.Display()

	if b.HasFixedWidth() {
		fixedWidth := b.FixedWidth() + horizSpace
		gtlog.Debug(
			ctx,
			"box.Box.OuterWidth[%s]: display=%s horiz_space=%d "+
				"using min(fixed_width=%d, parent_width=%d)",
			b.ID(), display, horizSpace, fixedWidth, parentWidth,
		)
		return types.Dimension(
			min(parentWidth, fixedWidth),
		)
	}

	percentWidth := types.Dimension(0)
	parentAvailable := parentWidth
	// Calculate the remainder of the parent's available width by examining the
	// set of siblings and subtracting any fixed width values and horizontal
	// space.
	childIndex := b.ChildIndex()
	for _, child := range parent.Children() {
		if child.ChildIndex() == childIndex {
			continue
		}
		parentAvailable -= child.HorizontalSpace()
		childDisplay := child.Display()
		if childDisplay != types.DisplayInline && child.HasFixedWidth() {
			parentAvailable -= child.FixedWidth()
		}
	}
	if b.HasPercentWidth() {
		constraint := b.WidthConstraint()
		pw := b.PercentWidth()
		percentWidth = parentAvailable * pw / 100
		percentWidth += horizSpace
		if next == nil {
			// If we're the last child in the row to use a percentage width
			// constraint, we need to reduce the calculated width by a single
			// cell in order to not exceed the parent inner bounds.
			percentWidth -= 1
		}
		gtlog.Debug(
			ctx,
			"box.Box.OuterWidth[%s]: width_constraint=%s. "+
				"calculated width %d "+
				"from total parent available width %d",
			b.ID(), constraint, percentWidth, parentAvailable,
		)
		if percentWidth != 0 {
			gtlog.Debug(
				ctx,
				"box.Box.OuterWidth[%s]: display=%s "+
					"horiz_space=%d width_constraint=%s. "+
					"using min(calc_percent_width=%d, parent_width=%d)",
				b.ID(), display,
				horizSpace, b.WidthConstraint(),
				percentWidth, parentWidth,
			)
			return types.Dimension(min(parentWidth, percentWidth))
		}
	}

	gtlog.Debug(
		ctx,
		"box.Box.OuterWidth[%s]: display=%s horiz_space=%d "+
			"last sibling or next sibling is block display. "+
			"using remaining horizontal width in parent %d.",
		b.ID(), display, horizSpace, parentAvailable,
	)
	return types.Dimension(min(parentWidth, parentAvailable))
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

// OuterHeight returns the height of the Box's outer bounding box.
//
// If a fixed height has been set and the display mode is not `inline`, we use
// the fixed height plus any vertical space from padding and border.
//
// If a percent height has been set and the display mode is not `inline`, we
// calculate the height by looking at the set of siblings and determining the
// appropriate percent of the remainder of the parent's height plus any
// vertical space from padding and border.
//
// If neither a fixed or percent height has been set, we return the remaining
// available height of the parent.
func (b *Box) OuterHeight() types.Dimension {
	parent := b.Parent()
	if parent == nil {
		return types.Dimension(b.Bounds().Dy())
	}
	next := b.NextSibling()
	parentInner := parent.InnerBounds()
	parentHeight := types.Dimension(parentInner.Dy())
	vertSpace := b.VerticalSpace()

	ctx := context.TODO()
	display := b.Display()
	if display != types.DisplayInline && b.HasFixedHeight() {
		fixedHeight := b.FixedHeight() + vertSpace
		gtlog.Debug(
			ctx,
			"box.Box.OuterHeight[%s]: display=%s vert_space=%d "+
				"using min(fixed_height=%d, parent_height=%d)",
			b.ID(), display, vertSpace, fixedHeight, parentHeight,
		)
		return types.Dimension(min(parentHeight, fixedHeight))
	}

	percentHeight := types.Dimension(0)
	parentAvailable := parentHeight
	if display != types.DisplayInline && b.HasPercentHeight() {
		// Calculate the remainder of the parent's available height by
		// examining the set of siblings and subtracting any fixed height
		// values.
		for _, child := range parent.Children() {
			childDisplay := child.Display()
			if childDisplay != types.DisplayInline && child.HasFixedHeight() {
				parentAvailable -= child.FixedHeight()
			}
		}
		constraint := b.HeightConstraint()
		ph := b.PercentHeight()
		percentHeight = parentAvailable * ph / 100
		gtlog.Debug(
			ctx,
			"box.Box.OuterHeight[%s]: height_constraint=%s. "+
				"calculated height %d "+
				"from total parent available height %d",
			b.ID(), constraint, percentHeight, parentAvailable,
		)
		if next == nil {
			// If we're the last child in the column to use a percentage height
			// constraint, we expand the height by a single line to consume the
			// remainder of the available parent's height.
			percentHeight += 1
		}
		if percentHeight != 0 {
			gtlog.Debug(
				ctx,
				"box.Box.OuterHeight[%s]: display=%s "+
					"vert_space=%d height_constraint=%s "+
					"using min(calc_percent_height=%d, parent_height=%d)",
				b.ID(), display,
				vertSpace, b.HeightConstraint(),
				percentHeight, parentHeight,
			)
			return types.Dimension(min(parentHeight, percentHeight))
		}
	}

	// Default to the remaining height of the parent container
	gtlog.Debug(
		ctx,
		"box.Box.OuterHeight[%s]: display=%s "+
			"vert_space=%d height_constraint=%s "+
			"using parent remaining height %d",
		b.ID(), display,
		vertSpace, b.HeightConstraint(),
		parentAvailable,
	)
	return parentAvailable
}
