package box

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
// * OuterWidth(): 5
// * OuterHeight(): 1
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
// * OuterWidth(): 7
// * OuterHeight(): 3
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
// * OuterWidth(): 9
// * OuterHeight(): 5
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
// or padding, that border and padding will cause the Element's OuterHeight()
// and OuterWidth() to be more than the specified fixed width and height.

// SetBounds sets the Box's outer bounding box.
func (b *Box) SetBounds(bounds types.Rectangle) {
	b.bounds = bounds
}

// Bounds returns the Box's outer bounding box.
func (b *Box) Bounds() types.Rectangle {
	return b.bounds
}

// TL returns the Box's outer bounding box's top-left coordinates.
func (b *Box) TL() types.Point {
	return b.bounds.Min
}

// TR returns the Box's outer bounding box's top-right coordinates.
func (b *Box) TR() types.Point {
	return types.Point{
		X: b.bounds.Max.X,
		Y: b.bounds.Min.Y,
	}
}

// MinY returns the Min Y (top) of the Box's outer bounding box.
func (b *Box) MinY() int {
	return b.bounds.Min.Y
}

// MaxY returns the Max Y (bottom) of the Box's outer bounding box.
func (b *Box) MaxY() int {
	return b.bounds.Max.Y
}

// InnerBounds returns the inner bounding box for the Box, which is the
// outer bounding box adjusted for any border and padding.
func (b *Box) InnerBounds() types.Rectangle {
	bounds := b.Bounds()
	border := b.Border()
	if border != nil {
		bounds.Min.X++
		bounds.Min.Y++
		bounds.Max.X--
		bounds.Max.Y--
	}
	return b.Padding().AdjustBounds(bounds)
}

// SetAbsolutePosition sets the Box's outer bounding box's top-left
// coordinates and marks the Box as using absolute positioning.
func (b *Box) SetAbsolutePosition(pt types.Point) {
	b.bounds.Min = pt
	b.absolute = true
}

// HasAbsolutePosition returns true if the Box used absolute positioning.
func (b *Box) HasAbsolutePosition() bool {
	return b.absolute
}

// Plot calculates the anchoring positioning coordinates of the element.
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
func (b *Box) Plot(ctx context.Context) {
	// If bounds has already been set, no need to plot.
	origBounds := b.Bounds()
	if !origBounds.Empty() {
		return
	}

	var parentInner types.Rectangle
	var parentTL types.Point
	parent := b.Parent()
	if parent != nil {
		parentInner = parent.InnerBounds()
		parentTL = parentInner.Min
	}
	prevSibling := b.PreviousSibling()
	display := b.display
	bounds := types.Rectangle{}

	// First we calculate the anchoring coordinates (top-left of our bounding
	// box)
	var anchor types.Point
	if b.HasAbsolutePosition() {
		anchor = b.TL()
		gtlog.Debug(
			ctx, "box.Box.Plot[%s]: anchor to absolute position %s",
			b.ID(), anchor,
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
			anchor = parentTL
			gtlog.Debug(
				ctx,
				"box.Box.Plot[%s]: anchor to parent inner top left %s",
				b.ID(), anchor,
			)
		} else {
			anchor = prevSibling.TR()
			gtlog.Debug(
				ctx,
				"box.Box.Plot[%s]: anchor to prev sibling outer top right %s",
				b.ID(), anchor,
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
					"box.Box.Plot[%s]: inline display, setting anchor y to %d "+
						"(min.y of previous sibling)",
					b.ID(), psy,
				)
				anchor.Y = psy
			}
		} else {
			// For elements with block display mode, we need to start this
			// element on the next line after the tallest previous sibling, or
			// if none, the parent's inner bounds top left coordinates.
			nextY := NextLineY(b)
			gtlog.Debug(
				ctx,
				"box.Box.Plot[%s]: block display, setting y to next line y %d "+
					"(max.y of previous siblings or parent inner bounds)",
				b.ID(), nextY,
			)
			anchor.Y = nextY
		}
	}
	// Set the top left corner of our bounding box to the anchor point.
	bounds.Min = anchor

	// Set the bottom right corner of our bounding box to the anchor
	// point plus the element's outer width and height.
	width := b.OuterWidth()
	height := b.OuterHeight()
	gtlog.Debug(
		ctx,
		"box.Box.Plot[%s]: expanding bounds by adding width %d and height %d to anchor point",
		b.ID(), width, height,
	)
	bounds.Max.X = anchor.X + int(width)
	bounds.Max.Y = anchor.Y + int(height)

	// Make sure that the parent bounds is never exceeded by a child.
	if !bounds.In(parentInner) {
		gtlog.Debug(
			ctx,
			"box.Box.Plot[%s]: plotted bounds %s exceeds parent inner "+
				"bounds %s. constraining to parent inner bounds.",
			b.ID(), bounds, parentInner,
		)
		if bounds.Dx() > parentInner.Dx() {
			bounds.Min.X = parentInner.Min.X
			bounds.Max.X = parentInner.Max.X
		}
		if bounds.Dy() > parentInner.Dy() {
			bounds.Min.Y = parentInner.Min.Y
			bounds.Max.Y = parentInner.Max.Y
		}
	}

	gtlog.Debug(
		ctx,
		"box.Box.Plot[%s]: final plotted bounds %s",
		b.ID(), bounds,
	)
	b.SetBounds(bounds)
	for _, child := range b.Children() {
		child.Plot(ctx)
	}
}

// NextLineY returns the maximum Y value of any previous sibling, or if
// no siblings, the parent inner bounds top-left coordinate's Y valub.
func NextLineY(b types.Plottable) int {
	parent := b.Parent()
	if parent == nil {
		return 0
	}
	y := parent.InnerBounds().Min.Y
	for _, prevSibling := range b.PreviousSiblings() {
		y = max(y, prevSibling.MaxY())
	}
	return y
}
