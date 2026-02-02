package element

import (
	"context"
	"strings"

	"github.com/mitchellh/go-wordwrap"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// WithSize constrains the size of the Element's outer bounding box and returns
// the Element.
func (e *Element) WithSize(constraint types.SizeConstraint) types.Element {
	e.Box.SetSize(constraint)
	return e
}

// WithWidth constrains the width of the Element and returns the Element.
func (e *Element) WithWidth(constraint types.DimensionConstraint) types.Element {
	e.Box.SetWidth(constraint)
	return e
}

// WithMinWidth sets the minimum width of the Element and returns the Element.
func (e *Element) WithMinWidth(w types.Dimension) types.Element {
	e.Box.SetMinWidth(w)
	return e
}

// WithHeight constrains the height of the Element and returns the Element.
func (e *Element) WithHeight(constraint types.DimensionConstraint) types.Element {
	e.Box.SetHeight(constraint)
	return e
}

// WithMinHeight sets the minimum height of the Element and returns the
// Element.
func (e *Element) WithMinHeight(w types.Dimension) types.Element {
	e.Box.SetMinHeight(w)
	return e
}

// Width returns the Element's width.
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
// parent's inner bounding box. If the display mode is `inline`, the width is
// set to the width of the content plus any horizontal padding and left-right
// border width.
func (e *Element) Width() types.Dimension {
	ctx := context.TODO()
	boxWidth := render.Width(ctx, e)
	display := e.Display()

	// If we're not using inline display mode, we use the box-calculated width.
	// The box-calculated width already includes the horizontal space of the
	// Element.
	if display != types.DisplayInline {
		gtlog.Debug(
			ctx,
			"Element.Width[%s]: display=%s. "+
				"using box width %d.",
			e.Tag(), display, boxWidth,
		)
		return boxWidth
	}

	horizSpace := e.HorizontalSpace()

	var next types.Plottable
	nextNode := e.NextSibling()
	if nextNode != nil {
		next = nextNode.(types.Plottable)
	}

	// No width constraint and not inline display, we consume the remainder of
	// the parent's width if this is the last sibling or the next sibling is
	// block display, otherwise we consume the "natural" width of the content,
	if next == nil || next.Display() == types.DisplayBlock {
		gtlog.Debug(
			ctx,
			"Element.Width[%s]: display=%s horiz_space=%d "+
				"last sibling or next sibling is block display. "+
				"using box width %d.",
			e.Tag(), display, horizSpace, boxWidth,
		)
		return boxWidth
	}

	contentWidth := e.TextContentWidth()
	calcWidth := contentWidth + horizSpace
	gtlog.Debug(
		ctx,
		"Element.Width[%s]: display=%s horiz_space=%d content_width=%d. "+
			"using min(box_width=%d, calc_width=%d).",
		e.Tag(), display, horizSpace, contentWidth,
		boxWidth, calcWidth,
	)
	return types.Dimension(min(boxWidth, calcWidth))
}

// Height returns the height of the Element.
//
// If a fixed height has been set and the display mode is not `inline`, we use
// the fixed height plus any vertical space from padding and border.
//
// If a percent height has been set and the display mode is not `inline`, we
// calculate the height by looking at the set of siblings and determining the
// appropriate percent of the remainder of the parent's height plus any
// vertical space from padding and border.
//
// If neither a fixed or percent height has been set and the display mode is
// `inline-block`, we return the height of the parent.
//
// If a fixed height has not been set or the display mode is inline, the height
// defaults to the number of lines of text content, or 1 if there is no text
// content, plus any vertical space from padding and border.
func (e *Element) Height() types.Dimension {
	ctx := context.TODO()
	display := e.Display()
	boxHeight := render.Height(ctx, e)

	// If we're not using inline display mode and there is a fixed height, we
	// use the box-calculated height.
	if display != types.DisplayInline && e.HasFixedHeight() {
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: display=%s. "+
				"using fixed box height %d.",
			e.ID(), display, boxHeight,
		)
		return boxHeight
	}

	parentNode := e.Parent()
	if parentNode == nil {
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: no parent. using box height %d.",
			e.ID(), boxHeight,
		)
		return boxHeight
	}

	parent := parentNode.(types.Plottable)
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())
	parentHeight := types.Dimension(parentInner.Dy())
	vertSpace := e.VerticalSpace()

	if display != types.DisplayInline && e.HasPercentHeight() {
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: display=%s "+
				"vert_space=%d height_constraint=%s. "+
				"using box height %d.",
			e.Tag(), display,
			vertSpace, e.HeightConstraint(),
			boxHeight,
		)
		return boxHeight
	}

	whitespace := e.Whitespace()
	wrapNever := whitespace&types.WhitespaceWrapNever != 0
	if wrapNever && boxHeight == 0 {
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: display=%s whitespace=%s "+
				"vert_space=%d height_constraint=none. "+
				"height is always 1 plus padding_vert + border_vert",
			e.Tag(), display, whitespace, vertSpace,
		)
		return vertSpace + 1
	}

	// "wrap-line" whitespace mode means don't wrap except on existing
	// newlines...
	wrapLine := whitespace&types.WhitespaceWrapLine != 0
	wrapped := false

	// We use the "natural" height of the content, which is the number of
	// newlines in the content. However, we first need to calculate any
	// wrapping of long text content before returning the number of newlines.
	content := e.TextContent()
	contentHeight := e.TextContentHeight()
	contentHeight += vertSpace
	origContentHeight := contentHeight
	contentWidth := e.TextContentWidth()
	if !wrapLine && (contentWidth > parentWidth) {
		wrapped = true
		wrappedContent := wordwrap.WrapString(content, uint(parentWidth))
		contentHeight = types.Dimension(strings.Count(wrappedContent, "\n") + 1)
		contentHeight += vertSpace
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: display=%s whitespace=%s "+
				"vert_space=%d original_content_height=%d parent_height=%d "+
				"content_width=%d parent_width=%d wrapped=%t "+
				"calculated new content_height of %d",
			e.Tag(), display, whitespace,
			vertSpace, origContentHeight, parentHeight,
			contentWidth, parentWidth, wrapped,
			contentHeight,
		)
	}
	gtlog.Debug(
		ctx,
		"Element.Height[%s]: display=%s whitespace=%s "+
			"vert_space=%d using min(content_height=%d, parent_height=%d)",
		e.Tag(), display, whitespace,
		vertSpace, contentHeight, parentHeight,
	)
	return types.Dimension(min(parentHeight, contentHeight))
}
