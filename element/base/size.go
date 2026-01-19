package base

import (
	"context"
	"strings"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
	"github.com/mitchellh/go-wordwrap"
)

// SetSize constrains the size of the Element's inner bounding box.
func (b *Base) SetSize(constraint types.SizeConstraint) types.Element {
	wc := constraint.Width()
	if wc != nil {
		b.widthConstraint = wc
	}
	hc := constraint.Height()
	if hc != nil {
		b.heightConstraint = hc
	}
	return b
}

// Size returns the width and height of the Element.
func (b *Base) Size() types.Size {
	return types.Size{
		W: int(b.Width()),
		H: int(b.Height()),
	}
}

// SetWidth constrains the width of the Element.
func (b *Base) SetWidth(constraint types.DimensionConstraint) types.Element {
	b.widthConstraint = constraint
	return b
}

// Width returns the Element's width.
//
// If a fixed width has been set and the display mode is `block` or
// `inline-block`, we use the fixed width plus any horizontal padding and
// left-right border width.
//
// If a fixed width has not been set and the display mode is `block` or
// `inline-block`, the width defaults to remaining horizontal space in the
// parent's inner bounding box. If the display mode is `inline`, the width is
// set to the width of the content plus any horizontal padding and left-right
// border width.
func (b *Base) Width() types.Dimension {
	parent := b.Parent()
	if parent == nil {
		return types.Dimension(b.Bounds().Dx())
	}
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())
	borderHoriz := types.Dimension(0)
	border := b.Border()
	if border != nil {
		borderHoriz = render.BorderHorizontalSpace(*border)
	}
	paddingHoriz := b.Padding().HorizontalSpace()

	ctx := context.TODO()
	display := b.Display()
	if display != types.DisplayInline && b.HasFixedWidth() {
		fixedWidth := b.FixedWidth() + paddingHoriz + borderHoriz
		gtlog.Debug(
			ctx,
			"Element.Width[%s]: display=%s padding_horiz=%d border_horiz=%d "+
				"using min(fixed_width=%d, parent_width=%d)",
			b.Tag(), display, paddingHoriz, borderHoriz,
			fixedWidth, parentWidth,
		)
		return types.Dimension(
			min(parentWidth, fixedWidth),
		)
	}
	if display != types.DisplayBlock {
		// For inline or inline-block with no fixed width, we use the lesser of
		// the parent's width or the "natural" width of the content
		contentWidth := b.TextContentWidth()
		contentWidth += paddingHoriz + borderHoriz
		gtlog.Debug(
			ctx,
			"Element.Width[%s]: display=%s padding_horiz=%d border_horiz=%d "+
				"using min(content_width=%d, parent_width=%d)",
			b.Tag(), display, paddingHoriz, borderHoriz,
			contentWidth, parentWidth,
		)
		return types.Dimension(
			min(parentWidth, contentWidth),
		)
	}

	// For block display mode, we start the element at the Min.X coordinate of
	// the parent, which simulates a "newline" on the screen. We then calculate
	// the width based on the whether we have a width constraint.

	wc := b.WidthConstraint()

	// Calculate the width of this based on whether there is a width
	// constraint. If there is no constraint, the element receives the
	// remainder of the horizontal space in the parent's bounding box.
	if wc == nil {
		gtlog.Debug(
			ctx,
			"Element.Width[%s]: display=%s width_constraint=nonb. "+
				"width is parent inner bounds width of %d",
			b.Tag(), display, parentWidth,
		)
		return parentWidth
	}
	remainder := wc.Apply(types.Dimension(parentWidth))
	gtlog.Debug(
		ctx,
		"Element.Width[%s]: display=%s "+
			"width_constraint=%d padding_horiz=%d border_horiz=%d. "+
			"calculated remainder of %d from parent width of %d",
		b.Tag(), display, wc, paddingHoriz, borderHoriz,
		remainder, parentWidth,
	)
	return remainder
}

// FixedWidth returns the Element's fixed width. If the Element does not have a
// fixed width constraint, returns 0.
func (b *Base) FixedWidth() types.Dimension {
	if !b.HasFixedWidth() {
		return types.Dimension(0)
	}
	return types.Dimension(b.widthConstraint.(core.FixedConstraint))
}

// SetMinWidth sets the minimum width of the Element.
func (b *Base) SetMinWidth(w types.Dimension) types.Element {
	b.minWidth = w
	return b
}

// MinWidth returns the Element's minimum width.
func (b *Base) MinWidth() types.Dimension {
	return b.minWidth
}

// SetHeight constrains the height of the Element.
func (b *Base) SetHeight(constraint types.DimensionConstraint) types.Element {
	b.heightConstraint = constraint
	return b
}

// Height returns the height of the Element.
//
// If a fixed height has been set and the display mode is `block`, we use the
// fixed height plus any vertical padding.
//
// If a fixed height has not been set or the display mode is not `block`, the
// height defaults to the number of lines of text content, or 1 if there is no
// text content, plus any vertical padding.
func (b *Base) Height() types.Dimension {
	parent := b.Parent()
	if parent == nil {
		return types.Dimension(b.Bounds().Dy())
	}
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())
	parentHeight := types.Dimension(parentInner.Dy())
	paddingVert := b.Padding().VerticalSpace()
	borderVert := types.Dimension(0)
	border := b.Border()
	if border != nil {
		borderVert = render.BorderVerticalSpace(*border)
	}

	ctx := context.TODO()
	display := b.Display()
	if display == types.DisplayBlock && b.HasFixedHeight() {
		fixedHeight := b.FixedHeight() + paddingVert + borderVert
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: display=%s padding_vert=%d border_vert=%d "+
				"using min(fixed_height=%d, parent_height=%d)",
			b.Tag(), display, paddingVert, borderVert,
			fixedHeight, parentHeight,
		)
		return types.Dimension(min(parentHeight, fixedHeight))
	}

	whitespace := b.Whitespace()
	wrapNever := whitespace&types.WhitespaceWrapNever != 0
	if wrapNever {
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: display=%s whitespace=%s "+
				"padding_vert=%d border_vert=%d "+
				"height is always 1 plus padding_vert + border_vert",
			b.Tag(), display, whitespace, paddingVert, borderVert,
		)
		return paddingVert + borderVert + 1
	}

	// "wrap-line" whitespace mode means don't wrap except on existing
	// newlines...
	wrapLine := whitespace&types.WhitespaceWrapLine != 0
	wrapped := false

	// We use the "natural" height of the content, which is the number of
	// newlines in the content. However, we first need to calculate any
	// wrapping of long text content before returning the number of newlines.
	content := b.TextContent()
	contentHeight := b.TextContentHeight()
	contentHeight += paddingVert + borderVert
	origContentHeight := contentHeight
	contentWidth := b.TextContentWidth()
	if !wrapLine && (contentWidth > parentWidth) {
		wrapped = true
		wrappedContent := wordwrap.WrapString(content, uint(parentWidth))
		contentHeight = types.Dimension(strings.Count(wrappedContent, "\n") + 1)
		contentHeight += paddingVert + borderVert
		gtlog.Debug(
			ctx,
			"Element.Height[%s]: display=%s whitespace=%s "+
				"padding_vert=%d border_vert=%d "+
				"original_content_height=%d parent_height=%d "+
				"content_width=%d parent_width=%d wrapped=%t "+
				"calculated new content_height of %d",
			b.Tag(), display, whitespace,
			paddingVert, borderVert,
			origContentHeight, parentHeight,
			contentWidth, parentWidth, wrapped,
			contentHeight,
		)
	}
	gtlog.Debug(
		ctx,
		"Element.Height[%s]: display=%s whitespace=%s "+
			"padding_vert=%d border_vert=%d "+
			"using min(content_height=%d, parent_height=%d)",
		b.Tag(), display, whitespace,
		paddingVert, borderVert,
		contentHeight, parentHeight,
	)
	return types.Dimension(min(parentHeight, contentHeight))
}

// FixedHeight returns the Element's fixed height. If the Element does not have
// a fixed height constraint, returns 0.
func (b *Base) FixedHeight() types.Dimension {
	if !b.HasFixedHeight() {
		return types.Dimension(0)
	}
	return types.Dimension(b.heightConstraint.(core.FixedConstraint))
}

// SetMinHeight sets the minimum height of the Element.
func (b *Base) SetMinHeight(h types.Dimension) types.Element {
	b.minHeight = h
	return b
}

// MinHeight returns the Element's minimum height.
func (b *Base) MinHeight() types.Dimension {
	return b.minHeight
}

// HasFixedWidth returns true if the Element's inner bounding box has a fixed
// width.
func (b *Base) HasFixedWidth() bool {
	_, ok := b.widthConstraint.(core.FixedConstraint)
	return ok
}

// HasFixedHeight returns true if the Element's inner bounding box has a fixed
// height.
func (b *Base) HasFixedHeight() bool {
	_, ok := b.heightConstraint.(core.FixedConstraint)
	return ok
}

// WidthConstraint returns any optional size constraint for the Element's
// width.  Returns nil when there is no width constraint.
func (b *Base) WidthConstraint() types.DimensionConstraint {
	return b.widthConstraint
}

// HeightConstraint returns any optional size constraint for the Element's
// height. Returns nil when there is no height constraint.
func (b *Base) HeightConstraint() types.DimensionConstraint {
	return b.heightConstraint
}
