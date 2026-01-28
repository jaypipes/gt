package render

import (
	"context"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// Plotting describes the process of determining an outer bounding box and
// positioning coordinates for something on a Screen.
//
// # Determining size and bounds.
//
// An Element has an outer bounding box and an inner bounding box. The outer
// bounding box represents the outer edge of the Element. The outer bounding
// box's cells contain the Element's border, if any. The inner bounding box
// represents the outer edge of the Element's *content* after accounting for
// any padding. The Element's content may be some text or it may be the content
// of child elements of the Element.
//
// ## No border, no padding, no content
//
// An empty Element has no size, no border, and no padding. It can be
// considered simply a Point at (0,0) on the Screen.
//
// ## No border, no padding, some content
//
// The inner bounding box will be equal to the outer bounding box when the
// Element has no border or padding.
//
// Consider an Element that has no border or padding and its content is the
// string "Hello". You can envision the Element laid out on the Screen like so:
//
// |     0   1   2   3   4   5   6   7   8
// |
// | 0   H   e   l   l   o
// |
// | 1
// |
// | 2
// |
// | 3
// |
// | 4
//
// Here's what the Element's methods would return:
//
// * OuterBounds(): (0,0)-(0,4)
// * Width(): 5
// * Height(): 1
// * InnerBounds(): (0,0)-(0,4)
// * InnerWidth(): 5
// * InnerHeight(): 1
//
// ## A border, no padding, some content
//
// Consider an Element that has a border but no padding and its content is the
// string "Hello". You can envision the Element laid out on the Screen like so:
//
// |     0   1   2   3   4   5   6   7   8
// |
// | 0   B   B   B   B   B   B   B
// |
// | 1   B   H   e   l   l   o   B
// |
// | 2   B   B   B   B   B   B   B
// |
// | 3
// |
// | 4
//
// In the above diagram, the letter "B" has been placed in the cells where the
// border will be drawn. The letters for the content "Hello" have been placed
// in their appropriate cells.
//
// Here's what the Element's methods would return:
//
// * OuterBounds(): (0,0)-(3,6)
// * Width(): 7
// * Height(): 3
// * InnerBounds(): (1,1)-(1,5)
// * InnerWidth(): 5
// * InnerHeight(): 1
//
// ## A border, some padding, some content
//
// Consider an Element has a border, a padding of 1 on all sides and its
// content is the string "Hello".  You can envision the Element's two bounding
// boxes like so:
//
// |     0   1   2   3   4   5   6   7   8
// |
// | 0   B   B   B   B   B   B   B   B   B
// |
// | 1   B   P   P   P   P   P   P   P   B
// |
// | 2   B   P   H   e   l   l   o   P   B
// |
// | 3   B   P   P   P   P   P   P   P   B
// |
// | 4   B   B   B   B   B   B   B   B   B
//
// In the above diagram, the letter "B" has been placed in the cells where the
// border will be drawn. The letter "P" has been placed in the cells where the
// padding takes up some width and height. And the letters for the content
// "Hello" have been placed in their appropriate cells.
//
// Here's what the Element's methods would return:
//
// * OuterBounds(): (0,0)-(4,8)
// * Width(): 9
// * Height(): 5
// * InnerBounds(): (2,2)-(2,6)
// * InnerWidth(): 5
// * InnerHeight(): 1
//
// # Impact of fixed width or height
//
// When the SetSize(), SetWidth() and SetHeight() methods are called on a
// Element, the user is saying that they want the Element's *inner bounding
// box* to be a specific fixed width and/or height.
//
// In other words, if the Element with a fixed width and/or height has a border
// or padding, that border and padding will cause the Element's Height()
// and Width() to be more than the specified fixed width and height.

// Plot calculates the bounds and positioning coordinates of the supplied
// element.
//
// If the supplied element's bounds have already been set, Plot does nothing.
//
// It traverses the tree of elements rooted at this element and calculates the
// top left coordinates for the element.
//
// To calculate the top left (anchor point) coordinates of the element's
// bounding box, we use the following algorithm:
//
// If the element is using absolute positioning, its bounding box is anchored
// at the absolute coordinates. If the element is using relative positioning,
// the anchor point is calculated boxd on the element's Display property and
// is relative to the previous sibling or, if no previous sibling, the parent.
func Plot(
	ctx context.Context,
	p types.Plottable,
	containerBounds types.Rectangle,
) {
	// If bounds has already been set, no need to plot.
	bounds := p.Bounds()
	if !bounds.Empty() {
		gtlog.Debug(
			ctx, "render.Plot[%s]: bounds already set: %s",
			p.ID(), bounds,
		)
	} else {
		bounds = calculateBounds(ctx, p, containerBounds)
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: calculated bounds %s",
			p.ID(), bounds,
		)
		p.SetBounds(bounds)
	}
	for _, child := range p.Children() {
		cp, ok := child.(types.Plottable)
		if ok {
			Plot(ctx, cp, containerBounds)
		}
	}
}

// calculateBounds determines the outer bounding box for the supplied Plottable.
func calculateBounds(
	ctx context.Context,
	p types.Plottable,
	containerBounds types.Rectangle,
) types.Rectangle {
	var parent types.Plottable
	var prevSibling types.Plottable

	parentNode := p.Parent()
	if parent != nil {
		parent = parentNode.(types.Plottable)
		containerBounds = parent.InnerBounds()
	}

	prevSiblingNode := p.PreviousSibling()
	if prevSiblingNode != nil {
		prevSibling = prevSiblingNode.(types.Plottable)
	}

	containerTL := containerBounds.Min
	display := p.Display()

	gtlog.Debug(
		ctx, "render.Plot[%s].start: container_bounds=%s display=%s",
		p.ID(), containerBounds, display,
	)
	// First we calculate the anchoring coordinates (top-left of our bounding
	// box)
	var anchor types.Point
	if p.HasAbsolutePosition() {
		anchor = p.TL()
		gtlog.Debug(
			ctx, "render.Plot[%s]: using absolute positioning. anchor to %s",
			p.ID(), anchor,
		)
	} else {
		// We place our anchor position depending on the display mode of the
		// current element. If the display mode is inline or inline-block, we
		// place our element directly to the right of the previous sibling or,
		// if no previous sibling, the left edge of the parent.
		//
		// If the display mode of the current element is block, we anchor our
		// element on the left margin of the parent and the bottom margin of
		// the previous sibling.
		if prevSibling == nil || display == types.DisplayBlock {
			anchor = containerTL
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: using relative positioning. "+
					"anchor to container top left %s",
				p.ID(), anchor,
			)
		} else {
			anchor = prevSibling.TR()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: using relative positioning. "+
					"anchor to prev sibling outer top right %s",
				p.ID(), anchor,
			)
		}

		// For elements with inline or inline-block display mode, we set the
		// anchor's top (min.Y) to the previous sibling's anchor point's Y
		// coordinatb. If no previous sibling, we use the parent's top margin.
		if display != types.DisplayBlock {
			if prevSibling != nil {
				psy := prevSibling.MinY()
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: using inline display. "+
						"anchor y to min.y of previous sibling %d",
					p.ID(), psy,
				)
				anchor.Y = psy
			}
		} else {
			// For elements with block display mode, we need to start this
			// element on the next line after the tallest previous sibling, or
			// if none, the parent's inner bounds top left coordinates.
			nextY := NextLineY(p)
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: using block display. anchor y to "+
					"max.y of previous siblings or "+
					"min.y of container bounds %d",
				p.ID(), nextY,
			)
			anchor.Y = nextY
		}
	}

	bounds := types.Rectangle{}
	// Set the top left corner of our bounding box to the anchor point.
	bounds.Min = anchor

	// Set the bottom right corner of our bounding box to the anchor
	// point plus the element's outer width and height.
	width := p.Width()
	height := p.Height()

	maxWidth := types.Dimension(containerBounds.Dx())
	maxHeight := types.Dimension(containerBounds.Dy())

	if width > maxWidth {
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: calculated width %d exceeds container "+
				"width %d. constraining to container width.",
			p.ID(), width, maxWidth,
		)
		width = maxWidth
	}

	if height > maxHeight {
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: calculated height %d exceeds container "+
				"height %d. constraining to container height.",
			p.ID(), height, maxHeight,
		)
		height = maxHeight
	}

	gtlog.Debug(
		ctx,
		"render.Plot[%s]: expanding bounds by "+
			"adding width %d and height %d to anchor %s",
		p.ID(), width, height, anchor,
	)
	bounds.Max.X = anchor.X + int(width)
	bounds.Max.Y = anchor.Y + int(height)

	// Make sure that the parent bounds is never exceeded by a child.
	if !bounds.In(containerBounds) {
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: plotted bounds %s exceeds container "+
				"bounds %s. constraining to container bounds",
			p.ID(), bounds, containerBounds,
		)
		if bounds.Dx() > containerBounds.Dx() {
			bounds.Min.X = containerBounds.Min.X
			bounds.Max.X = containerBounds.Max.X
		}
		if bounds.Dy() > containerBounds.Dy() {
			bounds.Min.Y = containerBounds.Min.Y
			bounds.Max.Y = containerBounds.Max.Y
		}
	}
	return bounds
}

// NextLineY returns the maximum Y value of any previous sibling, or if
// no siblings, the parent inner bounds top-left coordinate's Y valub.
func NextLineY(p types.Plottable) int {
	parentNode := p.Parent()
	if parentNode == nil {
		return 0
	}
	parent := parentNode.(types.Plottable)
	y := parent.InnerBounds().Min.Y
	for _, prevSiblingNode := range p.PreviousSiblings() {
		prevSibling := prevSiblingNode.(types.Plottable)
		y = max(y, prevSibling.MaxY())
	}
	return y
}
