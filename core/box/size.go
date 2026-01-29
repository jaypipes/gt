package box

import (
	"context"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
	"github.com/samber/lo"
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
func (b *Box) Width() types.Dimension {

	ctx := context.TODO()
	display := b.Display()
	horizSpace := b.HorizontalSpace()

	if b.HasFixedWidth() && display != types.DisplayInline {
		fixedWidth := b.FixedWidth()
		calcWidth := fixedWidth + horizSpace
		gtlog.Debug(
			ctx,
			"Box.Width[%s]: display=%s "+
				"fixed_width=%d horiz_space=%d. "+
				"calculated width of %d",
			b.ID(), display, fixedWidth, horizSpace, calcWidth,
		)
		return calcWidth
	}

	parentNode := b.Parent()
	if parentNode == nil {
		width := types.Dimension(b.Bounds().Dx())
		gtlog.Debug(
			ctx,
			"Box.Width[%s]: no parent. calculated width from bounds: %d",
			b.ID(), width,
		)
		return width
	}

	parent := parentNode.(types.Plottable)
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())

	// If this Box is using block display and does not have a fixed width, the
	// Box will start at the left edge of the parent. We calculate the
	// remaining width of the "row" of Boxes by doing a forward pass through
	// this Box's siblings and subtracting any fixed width amounts. If we reach
	// the end of the siblings or we come across a sibling that is using block
	// display, we stop calculating the remaining width.
	if display == types.DisplayBlock {
		nextNode := b.NextSibling()
		remainingWidth := parentWidth
		for nextNode != nil {
			next := nextNode.(types.Plottable)
			if next.Display() == types.DisplayBlock {
				break
			}
			remainingWidth -= next.HorizontalSpace()
			if next.HasFixedWidth() {
				remainingWidth -= next.FixedWidth()
			}
			nextNode = next.NextSibling()
		}
		gtlog.Debug(
			ctx,
			"Box.Width[%s]: using block display. "+
				"calculated remaining available width %d "+
				"from original parent width %d",
			b.ID(), remainingWidth, parentWidth,
		)
		return types.Dimension(
			min(parentWidth, remainingWidth),
		)
	}

	// If the Box is NOT using block display and does not have a fixed width,
	// we calculate the remainder of the available width by determining the set
	// of siblings IN THIS ROW and subtracting any fixed width values and
	// horizontal space.
	remainingWidth := parentWidth
	firstChildInRow := b.ChildIndex()
	prevNode := b.PreviousSibling()
	for prevNode != nil {
		prev := prevNode.(types.Plottable)
		if prev.Display() != types.DisplayBlock {
			firstChildInRow = prev.ChildIndex()
		}
		prevNode = prev.PreviousSibling()
	}

	childIndex := b.ChildIndex()
	children := parent.Children()

	for x, childNode := range children[firstChildInRow:] {
		if x == childIndex {
			// ignore THIS element for the purposes of calculating
			// remaining width...
			continue
		}
		child := childNode.(types.Plottable)
		if child.Display() == types.DisplayBlock {
			// Sibling starts a new row at the parent's inner bounds left edge
			// and therefore we are done calculating the available width.
			break
		}
		remainingWidth -= child.HorizontalSpace()
		if child.HasFixedWidth() {
			remainingWidth -= child.FixedWidth()
		}
	}

	if b.HasPercentWidth() {
		calcWidth := types.Dimension(0)
		nextNode := b.NextSibling()
		constraint := b.WidthConstraint()
		pw := b.PercentWidth()
		calcWidth = remainingWidth * pw / 100
		calcWidth += horizSpace
		if nextNode == nil {
			// If we're the last child in the row to use a percentage width
			// constraint, we need to decrease the calculated width by a single
			// cell in order to snug to the parent's inner bounds.
			calcWidth -= 1
		}
		gtlog.Debug(
			ctx,
			"Box.Width[%s]: display=%s "+
				"horiz_space=%d width_constraint=%s remaining_width=%d. "+
				"calculated width of %d.",
			b.ID(), display,
			horizSpace, constraint, remainingWidth,
			calcWidth,
		)
		return calcWidth
	}

	gtlog.Debug(
		ctx,
		"Box.Width[%s]: display=%s horiz_space=%d. "+
			"using remaining horizontal width in parent %d.",
		b.ID(), display, horizSpace, remainingWidth,
	)
	return types.Dimension(min(parentWidth, remainingWidth))
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
func (b *Box) Height() types.Dimension {
	ctx := context.TODO()
	display := b.Display()
	vertSpace := b.VerticalSpace()

	if display != types.DisplayInline && b.HasFixedHeight() {
		fixedHeight := b.FixedHeight()
		calcHeight := fixedHeight + vertSpace
		gtlog.Debug(
			ctx,
			"Box.Height[%s]: "+
				"display=%s vert_space=%d fixed_height=%d. "+
				"calculated height of %d",
			b.ID(), display, vertSpace, fixedHeight, calcHeight,
		)
		return calcHeight
	}

	parentNode := b.Parent()
	if parentNode == nil {
		height := types.Dimension(b.Bounds().Dy())
		gtlog.Debug(
			ctx,
			"Box.Height[%s]: no parent. calculated height from bounds: %d",
			b.ID(), height,
		)
		// Note that the bounds already includes the vertical space, so no need
		// to add it again here.
		return height
	}

	parent := parentNode.(types.Plottable)
	parentInner := parent.InnerBounds()
	parentHeight := types.Dimension(parentInner.Dy())

	// To determine the remaining available height from which we might
	// calculate a percentage height, we determine the max fixed height of
	// previous "rows" and subtract those max-fixed-height values from the
	// parent's inner height.
	remainingHeight := parentHeight
	childIndex := b.ChildIndex()
	children := parent.Children()
	rowMaxHeights := map[int]types.Dimension{}
	rowMaxHeight := types.Dimension(0)
	curRow := 0
	for x, childNode := range children {
		if x == childIndex {
			continue
		}
		child := childNode.(types.Plottable)
		cellVertSpace := child.VerticalSpace()
		cellFixedHeight := child.FixedHeight()
		childDisplay := child.Display()
		if childDisplay != types.DisplayInline {
			rowMaxHeight = max(rowMaxHeight, cellFixedHeight+cellVertSpace)
		}
		if childDisplay == types.DisplayBlock {
			rowMaxHeights[curRow] = rowMaxHeight
			curRow++
			rowMaxHeight = 0
			continue
		}
	}
	rowMaxHeights[curRow] = rowMaxHeight
	remainingHeight -= lo.Sum(lo.Values(rowMaxHeights))
	gtlog.Debug(
		ctx,
		"Box.Height[%s]: parent_height=%d row_max_fixed_heights=%v. "+
			"calculated remaining height %d",
		b.ID(), parentHeight, rowMaxHeights, remainingHeight,
	)

	next := b.NextSibling()
	if display != types.DisplayInline && b.HasPercentHeight() {
		calcHeight := types.Dimension(0)
		constraint := b.HeightConstraint()
		ph := b.PercentHeight()
		calcHeight = remainingHeight * ph / 100
		if next == nil {
			// If we're the last child in the column to use a percentage height
			// constraint, we expand the height by a single line to consume the
			// remainder of the available parent's height.
			calcHeight += 1
		}
		calcHeight += vertSpace
		gtlog.Debug(
			ctx,
			"Box.Height[%s]: display=%s "+
				"vert_space=%d height_constraint=%s remaining_height=%d. "+
				"calculated height of %d",
			b.ID(), display,
			vertSpace, constraint, remainingHeight,
			calcHeight,
		)
		return calcHeight
	}

	if display != types.DisplayBlock {
		remainingHeight += vertSpace
	}

	gtlog.Debug(
		ctx,
		"Box.Height[%s]: display=%s vert_space=%d height_constraint=%s. "+
			"returning remaining height %d",
		b.ID(), display,
		vertSpace, b.HeightConstraint(),
		remainingHeight,
	)
	return remainingHeight
}
