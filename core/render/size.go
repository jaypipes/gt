package render

import (
	"context"
	"strings"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
	"github.com/mitchellh/go-wordwrap"
	"github.com/samber/lo"
)

// Width returns the width of the supplied element's outer bounding box.
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
func Width(ctx context.Context, n types.Node) types.Dimension {
	ider, ok := n.(types.Identifiable)
	if !ok {
		return types.Dimension(0)
	}
	id := ider.ID()
	p, ok := n.(types.Plottable)
	if !ok {
		return types.Dimension(0)
	}
	display := p.Display()
	horizSpace := p.HorizontalSpace()

	if p.HasFixedWidth() && display != types.DisplayInline {
		fixedWidth := p.FixedWidth()
		calcWidth := fixedWidth + horizSpace
		gtlog.Debug(
			ctx,
			"render.Width[%s]: display=%s "+
				"fixed_width=%d horiz_space=%d. "+
				"calculated width of %d",
			id, display, fixedWidth, horizSpace, calcWidth,
		)
		return calcWidth
	}

	parentNode := n.Parent()
	if parentNode == nil {
		width := types.Dimension(p.Bounds().Dx())
		gtlog.Debug(
			ctx,
			"render.Width[%s]: no parent. calculated width from bounds: %d",
			id, width,
		)
		return width
	}

	parent, ok := parentNode.(types.Plottable)
	if !ok {
		parentStr := ""
		pid, ok := parentNode.(types.Identifiable)
		if ok {
			parentStr = pid.ID()
		}
		gtlog.Debug(
			ctx, "parent %s is not plottable. it's a %T",
			parentStr, parentNode,
		)
		return 0
	}
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())

	// If this Box is using block display and does not have a fixed width, the
	// Box will start at the left edge of the parent. We calculate the
	// remaining width of the "row" of Boxes by doing a forward pass through
	// this Box's siblings and subtracting any fixed width amounts. If we reach
	// the end of the siblings or we come across a sibling that is using block
	// display, we stop calculating the remaining width.
	if display == types.DisplayBlock {
		nextNode := n.NextSibling()
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
			nextNode = nextNode.NextSibling()
		}
		gtlog.Debug(
			ctx,
			"render.Width[%s]: using block display. "+
				"calculated remaining available width %d "+
				"from original parent width %d",
			id, remainingWidth, parentWidth,
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
	firstChildInRow := n.ChildIndex()
	prevNode := n.PreviousSibling()
	for prevNode != nil {
		prev := prevNode.(types.Plottable)
		if prev.Display() != types.DisplayBlock {
			firstChildInRow = prevNode.ChildIndex()
		}
		prevNode = prevNode.PreviousSibling()
	}

	childIndex := n.ChildIndex()
	children := parentNode.Children()

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

	if p.HasPercentWidth() {
		calcWidth := types.Dimension(0)
		nextNode := n.NextSibling()
		constraint := p.WidthConstraint()
		pw := p.PercentWidth()
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
			id, display,
			horizSpace, constraint, remainingWidth,
			calcWidth,
		)
		return calcWidth
	}

	gtlog.Debug(
		ctx,
		"Box.Width[%s]: display=%s horiz_space=%d. "+
			"using remaining horizontal width in parent %d.",
		id, display, horizSpace, remainingWidth,
	)
	return types.Dimension(min(parentWidth, remainingWidth))
}

// Height returns the height of the supplied element's outer bounding box.
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
func Height(ctx context.Context, n types.Node) types.Dimension {
	ider, ok := n.(types.Identifiable)
	if !ok {
		return types.Dimension(0)
	}
	id := ider.ID()
	p, ok := n.(types.Plottable)
	if !ok {
		return types.Dimension(0)
	}
	display := p.Display()
	vertSpace := p.VerticalSpace()

	if display != types.DisplayInline && p.HasFixedHeight() {
		fixedHeight := p.FixedHeight()
		calcHeight := fixedHeight + vertSpace
		gtlog.Debug(
			ctx,
			"render.Height[%s]: "+
				"display=%s vert_space=%d fixed_height=%d. "+
				"calculated height of %d",
			id, display, vertSpace, fixedHeight, calcHeight,
		)
		return calcHeight
	}

	parentNode := n.Parent()
	if parentNode == nil {
		height := types.Dimension(p.Bounds().Dy())
		gtlog.Debug(
			ctx,
			"render.Height[%s]: no parent. calculated height from bounds: %d",
			id, height,
		)
		// Note that the bounds already includes the vertical space, so no need
		// to add it again here.
		return height
	}

	parent, ok := parentNode.(types.Plottable)
	if !ok {
		parentStr := ""
		pid, ok := parentNode.(types.Identifiable)
		if ok {
			parentStr = pid.ID()
		}
		gtlog.Debug(
			ctx, "parent %s is not plottable. it's a %T",
			parentStr, parentNode,
		)
		return 0
	}
	parentInner := parent.InnerBounds()
	parentHeight := types.Dimension(parentInner.Dy())

	// To determine the remaining available height from which we might
	// calculate a percentage height, we determine the max fixed height of
	// previous "rows" and subtract those max-fixed-height values from the
	// parent's inner height.
	remainingHeight := parentHeight
	childIndex := n.ChildIndex()
	children := parentNode.Children()
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
		"render.Height[%s]: parent_height=%d row_max_fixed_heights=%v. "+
			"calculated remaining height %d",
		id, parentHeight, rowMaxHeights, remainingHeight,
	)

	next := n.NextSibling()
	if display != types.DisplayInline && p.HasPercentHeight() {
		calcHeight := types.Dimension(0)
		constraint := p.HeightConstraint()
		ph := p.PercentHeight()
		calcHeight = remainingHeight * ph / 100
		if next == nil && ph != 100 {
			// If we're the last child in the column to use a percentage height
			// constraint, we expand the height by a single line to consume the
			// remainder of the available parent's height.
			calcHeight += 1
		}
		calcHeight += vertSpace
		gtlog.Debug(
			ctx,
			"render.Height[%s]: display=%s "+
				"vert_space=%d height_constraint=%s remaining_height=%d. "+
				"calculated height of %d",
			id, display,
			vertSpace, constraint, remainingHeight,
			calcHeight,
		)
		return calcHeight
	}

	// If we get here, we're either inline display mode or no height constraint
	// was specified. In this case, if we're plotting a types.Element, we
	// determine the minimum number of lines that the element's content would
	// consume given the container's width and return the minimum of that or
	// the previously-calculated remaining height.
	e, ok := n.(types.Element)
	if !ok {
		gtlog.Debug(
			ctx,
			"render.Height[%s]: display=%s vert_space=%d height_constraint=%s. "+
				"node was not an element. using remaining height %d",
			id, display,
			vertSpace, p.HeightConstraint(),
			remainingHeight,
		)
		return remainingHeight
	}

	parentWidth := types.Dimension(parentInner.Dx())

	whitespace := p.Whitespace()
	wrapNever := whitespace&types.WhitespaceWrapNever != 0
	if wrapNever && remainingHeight == 0 {
		gtlog.Debug(
			ctx,
			"render.Height[%s]: display=%s whitespace=%s "+
				"vert_space=%d height_constraint=none. "+
				"height is always 1 plus padding_vert + border_vert",
			id, display, whitespace, vertSpace,
		)
		return vertSpace + 1
	}

	// "wrap-line" whitespace mode means don't wrap EXCEPT on existing
	// newlines.
	wrapLine := whitespace&types.WhitespaceWrapLine != 0
	wrapped := false

	// We use the "natural" height of the content, which is the number of
	// newlines in the content. However, we first need to calculate any
	// wrapping of long text content before returning the number of newlines.
	horizSpace := p.HorizontalSpace()
	content := e.TextContent()
	contentHeight := e.TextContentHeight()
	contentHeight += vertSpace
	origContentHeight := contentHeight
	contentWidth := e.TextContentWidth()
	if !wrapLine && ((contentWidth + horizSpace) > parentWidth) {
		wrapped = true
		wrapWidth := uint(parentWidth - horizSpace)
		wrappedContent := wordwrap.WrapString(content, wrapWidth)
		contentHeight = types.Dimension(strings.Count(wrappedContent, "\n") + 1)
		contentHeight += vertSpace
		gtlog.Debug(
			ctx,
			"render.Height[%s]: display=%s whitespace=%s "+
				"vert_space=%d horiz_space=%d wrap_width=%d "+
				"original_content_height=%d parent_height=%d "+
				"content_width=%d parent_width=%d wrapped=%t "+
				"calculated new content_height of %d",
			e.Tag(), display, whitespace,
			vertSpace, horizSpace, wrapWidth,
			origContentHeight, parentHeight,
			contentWidth, parentWidth, wrapped,
			contentHeight,
		)
		e.SetTextContent(wrappedContent)
	}

	gtlog.Debug(
		ctx,
		"render.Height[%s]: display=%s whitespace=%s "+
			"vert_space=%d using min(content_height=%d, remaining_height=%d)",
		e.Tag(), display, whitespace,
		vertSpace, contentHeight, parentHeight,
	)
	return types.Dimension(min(parentHeight, contentHeight))
}
