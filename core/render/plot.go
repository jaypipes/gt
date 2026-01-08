package render

import (
	"context"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/types"
)

// Plot calculates the bounding box of a supplied element.  It traverses the
// tree of elements rooted at the supplied element and calculates the anchor
// position, width and height of the element.
//
// If the element is using absolute positioning and a fixed size, returns the
// bounding box anchored at the absolute coordinates and the bottom-right
// coordinates calculated by adding the fixed width and height to the anchor
// point. If the element is using absolute positioning but relative sizing, we
// calculate the bottom-right coordinates relative to any supplied parent or
// previous sibling.
//
// If the element is not using absolute positioning, the anchor point is
// calculated relative to the supplied parent or previous sibling. If the
// element is using relative sizing, the bottom-right coordinates are
// calculated relative to any supplied parent or previous sibling.
func Plot(
	ctx context.Context,
	el types.Element,
) {
	parent := el.Parent()
	if parent == nil {
		// We don't plot the Document...
		return
	}
	prevSibling := el.PreviousSibling()
	display := el.Display()
	bounds := types.Rectangle{}

	// First we calculate the anchoring coordinates (top-left of our bounding
	// box)
	anchor := el.TL()
	if el.AbsolutePositioned() {
		gtlog.Debug(
			ctx, "render.Plot[%s]: anchor to absolute position %s",
			el.Tag(), anchor,
		)
	} else {
		// start with our relative offset
		anchor = el.TL()
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: anchor to relative position with offset %s",
			el.Tag(), anchor,
		)

		// We place our anchor position depending on the display mode of the
		// current element. If the display mode is inline or inline-block, we
		// place our element directly to the right of the previous sibling or,
		// if no previous sibling, the left margin of the parent.
		//
		// If the display mode of the current element is block, we anchor our
		// element on the left margin of the parent and the bottom margin of
		// the previous sibling.
		if prevSibling == nil || display == types.DisplayBlock {
			leftMargin := parent.LeftMargin()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: moving x by %d "+
					"(parent left margin)",
				el.Tag(), leftMargin,
			)
			anchor.X += leftMargin
		} else {
			psWidth := prevSibling.Width()
			gtlog.Debug(
				ctx,
				"render.Plot[%s]: moving x by %d "+
					"(width of previous sibling)",
				el.Tag(), psWidth,
			)
			anchor.X += psWidth
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
			} else {
				topMargin := parent.TopMargin()
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: inline display, moving y by %d "+
						"(parent top margin)",
					el.Tag(), topMargin,
				)
				anchor.Y += topMargin
			}
		} else {
			// For elements with block display mode, we need to start this
			// element on the next line after the previous sibling, of if none,
			// the parent's top margin.
			if prevSibling != nil {
				psy := prevSibling.MaxY()
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: block display, setting anchor y to %d "+
						"(max.y of previous sibling)",
					el.Tag(), psy,
				)
				anchor.Y = psy
			} else {
				topMargin := parent.TopMargin()
				gtlog.Debug(
					ctx,
					"render.Plot[%s]: block display, moving y by %d "+
						"(parent top margin)",
					el.Tag(), topMargin,
				)
				anchor.Y += topMargin
			}
		}
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: calculated anchor position %s",
			el.Tag(), anchor,
		)
	}
	bounds.Min.X = anchor.X
	bounds.Min.Y = anchor.Y

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
	} else {
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

	if display != types.DisplayInline && el.FixedHeight() {
		h := el.Height()
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: using fixed height %d",
			el.Tag(), h,
		)
		bounds.Max.Y += h
	} else {
		// Calculate the height of this Plotted based on whether there is a
		// height constraint. If there is no constraint, the element receives
		// the remainder of the vertical space in the parent's bounding box.
		ph := parent.InnerBounds().Dy()
		remainder := ph - bounds.Dy()
		gtlog.Debug(
			ctx,
			"render.Plot[%s]: calculated height remainder of %d "+
				"from parent height of %d",
			el.Tag(), remainder, ph,
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
