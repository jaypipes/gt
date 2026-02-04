package element

import (
	"context"
	"strings"

	"github.com/charmbracelet/x/ansi"

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

// ScrollWidth returns the minimum number of cells (width) that the Element
// would consume in order to fit all of its content on the screen without
// using a horizontal scrollbar.
func (e *Element) ScrollWidth() types.Dimension {
	return types.Dimension(ansi.StringWidth(e.textContent))
}

// ScrollHeight returns the minimum number of lines (height) that the
// Element would consume in order to fit all of its content on the screen
// without using a vertical scrollbar.
func (e *Element) ScrollHeight() types.Dimension {
	return types.Dimension(strings.Count(e.textContent, "\n")) + 1
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
