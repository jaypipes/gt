package render

import (
	"context"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/types"
)

// Plot calculates the bounding box of a supplied element. It traverses the
// tree of elements rooted at the supplied element and calculates the top left
// and bottom right coordinates for the element.
//
// To calculate the top left (anchor point) coordinates of the element's
// bounding box, we use the following algorithm:
//
// If the element is using absolute positioning, its bounding box is anchored
// at the absolute coordinates. If the element is using relative positioning,
// the anchor point is calculated based on the element's Display property and
// is relative to the previous sibling or, if no previous sibling, the parent.
//
// To calculate the bottom right coordinates of the bounding box, we add the
// element's width to its anchor point's X value and add the element's height
// to its anchor point's Y value. As such, each element class is responsible
// for properly calculating its own width and height depending on, for example,
// whether there is a fixed width/height, what the element's Display property
// is, and whether there are width or height constraints.
func Plot(
	ctx context.Context,
	el types.Element,
) {
	parent := el.Parent()
	if parent == nil {
		for _, child := range el.Children() {
			Plot(ctx, child)
		}
		return
	}

	parentInner := parent.InnerBounds()
	parentTL := parentInner.Min
	prevSibling := el.PreviousSibling()
	display := el.Display()
	bounds := types.Rectangle{}

	// First we calculate the anchoring coordinates (top-left of our bounding
	// box)
	var anchor types.Point
	if el.AbsolutePositioned() {
		anchor = el.TL()
		gtlog.Debug(
			ctx, "render.Plot[%s]: anchor to absolute position %s",
			el.Tag(), anchor,
		)
	} else {
		// anchor at our parent's top left and add our relative offset

		// We place our anchor position depending on the display mode of the
		// current element. If the display mode is inline or inline-block, we
		// place our element directly to the right of the previous sibling or,
		// if no previous sibling, the left margin of the parent.
		//
		// If the display mode of the current element is block, we anchor our
		// element on the left margin of the parent and the bottom margin of
		// the previous sibling.
		if prevSibling == nil || display == types.DisplayBlock {
			anchor = parentTL
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: anchor to parent inner top left %s",
				el.Tag(), anchor,
			)
		} else {
			anchor = prevSibling.TR()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: anchor to prev sibling outer top right %s",
				el.Tag(), anchor,
			)
		}

		// For elements with inline or inline-block display mode, we set the
		// anchor's top (min.Y) to the previous sibling's anchor point's Y
		// coordinate. If no previous sibling, we use the parent's top margin.
		if display != types.DisplayBlock {
			if prevSibling != nil {
				psy := prevSibling.MinY()
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: inline display, setting anchor y to %d "+
						"(min.y of previous sibling)",
					el.Tag(), psy,
				)
				anchor.Y = psy
			}
		} else {
			// For elements with block display mode, we need to start this
			// element on the next line after the tallest previous sibling, or
			// if none, the parent's inner bounds top left coordinates.
			nextY := nextLineY(el)
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: block display, setting anchor y to %d "+
					"(max.y of previous siblings + 1 or parent inner bounds)",
				el.Tag(), nextY,
			)
			anchor.Y = nextY
		}

		offset := el.TL()
		if offset.X > 0 || offset.Y > 0 {
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: adding offset %s to anchor %s",
				el.Tag(), offset, anchor,
			)
			anchor.Add(offset)
		}

		gtlog.Debug(
			ctx,
			"render.Plot[%s]: calculated anchor position %s",
			el.Tag(), anchor,
		)
	}
	bounds.Min.X = anchor.X
	bounds.Min.Y = anchor.Y

	// Set the bottom right corner of the bounding box to the anchor point plus
	// the element's width and height plus any padding and border.
	width := el.Width()
	height := el.Height()
	gtlog.Debug(
		ctx,
		"render.Plot[%s]: expanding bounds by adding width %d and height %d to anchor point",
		el.Tag(), width, height,
	)
	bounds.Max.X = anchor.X + width
	bounds.Max.Y = anchor.Y + height

	/*
		// Then we calculate the width and height, which will inform us what our
		// bottom-right coordinates will be.
		if display != types.DisplayInline && el.FixedWidth() {
			w := el.Width()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: using fixed width %d",
				el.Tag(), w,
			)
			bounds.Max.X += w
		} else if display == types.DisplayBlock {
			// Calculate the width of this Plotted based on whether there is a
			// width constraint. If there is no constraint, the element receives
			// the remainder of the horizontal space in the parent's bounding box.
			pw := parent.InnerBounds().Dx()
			remainder := pw - bounds.Dx()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: calculated width remainder of %d "+
					"from parent width of %d",
				el.Tag(), remainder, pw,
			)
			wc := el.WidthConstraint()
			if wc != nil {
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: calculating width constraint %v",
					el.Tag(), wc,
				)
			} else {
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: width defaulting to remainder %d",
					el.Tag(), remainder,
				)
				bounds.Max.X += remainder
			}
		}

		if el.FixedHeight() {
			h := el.Height()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: using fixed height %d",
				el.Tag(), h,
			)
			bounds.Max.Y += h
		} else if display == types.DisplayBlock {
			// Calculate the height of this Plotted based on whether there is a
			// height constraint. If there is no constraint, the element receives
			// the remainder of the vertical space in the parent's bounding box.
			// The remainder of the vertical space in the bounding box can be
			// calculated by subtracting the previous sibling's Max.Y from the
			// parent's inner bounds Max.Y.
			parentMaxY := parent.InnerBounds().Max.Y
			prevSibMaxY := parent.InnerBounds().Min.Y
			if prevSibling != nil {
				prevSibMaxY = prevSibling.MaxY()
			}
			remainder := parentMaxY - prevSibMaxY
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: calculated height remainder of %d "+
					"from parent max.y of %d and prevsib max.y of %d",
				el.Tag(), remainder, parentMaxY, prevSibMaxY,
			)
			hc := el.HeightConstraint()
			if hc != nil {
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: calculating height constraint %v",
					el.Tag(), hc,
				)
			} else {
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: height defaulting to remainder %d",
					el.Tag(), remainder,
				)
				bounds.Max.Y += remainder
			}
		} else {
			// el.Height() returns the "natural" height of the element. For things
			// like a Span, the natural height will be the number of newlines in
			// the Span's text content.
			h := el.Height()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: using natural height %d",
				el.Tag(), h,
			)
			bounds.Max.Y += h
		}
	*/

	// Make sure that the parent bounds is never exceeded by a child.
	if !bounds.In(parentInner) {
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: plotted bounds %s exceeds parent inner "+
				"bounds %s. constraining to parent inner bounds.",
			el.Tag(), bounds, parentInner,
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
		"render.Plot[%s]: final plotted bounds %s",
		el.Tag(), bounds,
	)
	el.SetBounds(bounds)
	for _, child := range el.Children() {
		Plot(ctx, child)
	}
}

// nextLineY returns the maximum Y value of any previous sibling plus 1, or if
// no siblings, the parent inner bounds top-left coordinate's Y value.
func nextLineY(el types.Element) int {
	y := 0
	parent := el.Parent()
	if parent != nil {
		y = parent.InnerBounds().Min.Y
	}
	prevSibling := el.PreviousSibling()
	for prevSibling != nil {
		y = max(y, prevSibling.BL().Y+1)
		prevSibling = prevSibling.PreviousSibling()
	}
	return y
}
