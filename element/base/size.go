package base

import (
	"context"
	"strings"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
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
func (b *Base) Width() types.Dimension {
	parent := b.parent
	if parent == nil {
		return types.Dimension(b.Bounds().Dx())
	}
	next := b.NextSibling()
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())
	horizSpace := b.HorizontalSpace()

	ctx := context.TODO()
	display := b.Display()

	if display == types.DisplayInline {
		// For inline, we use the lesser of the parent's width or the "natural"
		// width of the content
		contentWidth := b.TextContentWidth()
		contentWidth += horizSpace
		gtlog.Debug(
			ctx,
			"base.Base.Width[%s]: display=%s horiz_space=%d "+
				"using min(content_width=%d, parent_width=%d)",
			b.Tag(), display, horizSpace, contentWidth, parentWidth,
		)
		return types.Dimension(
			min(parentWidth, contentWidth),
		)
	}

	if b.HasFixedWidth() {
		fixedWidth := b.FixedWidth() + horizSpace
		gtlog.Debug(
			ctx,
			"base.Base.Width[%s]: display=%s horiz_space=%d "+
				"using min(fixed_width=%d, parent_width=%d)",
			b.Tag(), display, horizSpace, fixedWidth, parentWidth,
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
	for _, child := range parent.Children() {
		if child.ChildIndex() == b.index {
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
			"base.Base.Width[%s]: width_constraint=%s. "+
				"calculated width %d "+
				"from total parent available width %d",
			b.Tag(), constraint, percentWidth, parentAvailable,
		)
		if percentWidth != 0 {
			gtlog.Debug(
				ctx,
				"base.Base.Width[%s]: display=%s "+
					"horiz_space=%d width_constraint=%s. "+
					"using min(calc_percent_width=%d, parent_width=%d)",
				b.Tag(), display,
				horizSpace, b.WidthConstraint(),
				percentWidth, parentWidth,
			)
			return types.Dimension(min(parentWidth, percentWidth))
		}
	}

	// No width constraint and not inline display, we consume the remainder of
	// the parent's width if this is the last sibling or the next sibling is
	// block display, otherwise we consume the "natural" width of the content,
	if next == nil || next.Display() == types.DisplayBlock {
		gtlog.Debug(
			ctx,
			"base.Base.Width[%s]: display=%s horiz_space=%d "+
				"last sibling or next sibling is block display. "+
				"using remaining horizontal width in parent %d.",
			b.Tag(), display, horizSpace, parentAvailable,
		)
		return types.Dimension(min(parentWidth, parentAvailable))
	}
	contentWidth := b.TextContentWidth()
	gtlog.Debug(
		ctx,
		"base.Base.Width[%s]: display=%s horiz_space=%d "+
			"not last sibling. using text content width %d.",
		b.Tag(), display, horizSpace, contentWidth,
	)
	return types.Dimension(min(parentWidth, contentWidth))
}

// HasFixedWidth returns true if the Element's inner bounding box has a fixed
// width.
func (b *Base) HasFixedWidth() bool {
	_, ok := b.widthConstraint.(core.FixedConstraint)
	return ok
}

// FixedWidth returns the Element's fixed width. If the Element does not have a
// fixed width constraint, returns 0.
func (b *Base) FixedWidth() types.Dimension {
	if !b.HasFixedWidth() {
		return types.Dimension(0)
	}
	return types.Dimension(b.widthConstraint.(core.FixedConstraint))
}

// HasPercentWidth returns true if the Element's inner bounding box has a percent
// width.
func (b *Base) HasPercentWidth() bool {
	_, ok := b.widthConstraint.(core.PercentConstraint)
	return ok
}

// PercentWidth returns the Element's fixed width. If the Element does not have a
// percent width constraint, returns 0.
func (b *Base) PercentWidth() types.Dimension {
	if !b.HasPercentWidth() {
		return types.Dimension(0)
	}
	return types.Dimension(b.widthConstraint.(core.PercentConstraint))
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

// WidthConstraint returns any optional size constraint for the Element's
// width.  Returns nil when there is no width constraint.
func (b *Base) WidthConstraint() types.DimensionConstraint {
	return b.widthConstraint
}

// SetHeight constrains the height of the Element.
func (b *Base) SetHeight(constraint types.DimensionConstraint) types.Element {
	gtlog.Debug(
		context.TODO(), "base.Base.SetHeight[%s](constraint=%s)",
		b.Tag(), constraint,
	)
	b.heightConstraint = constraint
	return b
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
func (b *Base) Height() types.Dimension {
	parent := b.Parent()
	if parent == nil {
		return types.Dimension(b.Bounds().Dy())
	}
	next := b.NextSibling()
	parentInner := parent.InnerBounds()
	parentWidth := types.Dimension(parentInner.Dx())
	parentHeight := types.Dimension(parentInner.Dy())
	vertSpace := b.VerticalSpace()

	ctx := context.TODO()
	display := b.Display()
	if display != types.DisplayInline && b.HasFixedHeight() {
		fixedHeight := b.FixedHeight() + vertSpace
		gtlog.Debug(
			ctx,
			"base.Base.Height[%s]: display=%s vert_space=%d "+
				"using min(fixed_height=%d, parent_height=%d)",
			b.Tag(), display, vertSpace, fixedHeight, parentHeight,
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
			"base.Base.Height[%s]: height_constraint=%s. "+
				"calculated height %d "+
				"from total parent available height %d",
			b.Tag(), constraint, percentHeight, parentAvailable,
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
				"base.Base.Height[%s]: display=%s "+
					"vert_space=%d height_constraint=%s "+
					"using min(calc_percent_height=%d, parent_height=%d)",
				b.Tag(), display,
				vertSpace, b.HeightConstraint(),
				percentHeight, parentHeight,
			)
			return types.Dimension(min(parentHeight, percentHeight))
		}
	}

	if display == types.DisplayInlineBlock {
		// Default to the height of the parent container
		gtlog.Debug(
			ctx,
			"base.Base.Height[%s]: display=%s "+
				"vert_space=%d height_constraint=%s "+
				"using parent remaining height %d",
			b.Tag(), display,
			vertSpace, b.HeightConstraint(),
			parentAvailable,
		)
		return parentAvailable
	}

	whitespace := b.Whitespace()
	wrapNever := whitespace&types.WhitespaceWrapNever != 0
	if wrapNever && percentHeight == 0 {
		gtlog.Debug(
			ctx,
			"base.Base.Height[%s]: display=%s whitespace=%s "+
				"vert_space=%d height_constraint=none. "+
				"height is always 1 plus padding_vert + border_vert",
			b.Tag(), display, whitespace, vertSpace,
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
	content := b.TextContent()
	contentHeight := b.TextContentHeight()
	contentHeight += vertSpace
	origContentHeight := contentHeight
	contentWidth := b.TextContentWidth()
	if !wrapLine && (contentWidth > parentWidth) {
		wrapped = true
		wrappedContent := wordwrap.WrapString(content, uint(parentWidth))
		contentHeight = types.Dimension(strings.Count(wrappedContent, "\n") + 1)
		contentHeight += vertSpace
		gtlog.Debug(
			ctx,
			"base.Base.Height[%s]: display=%s whitespace=%s "+
				"vert_space=%d original_content_height=%d parent_height=%d "+
				"content_width=%d parent_width=%d wrapped=%t "+
				"calculated new content_height of %d",
			b.Tag(), display, whitespace,
			vertSpace, origContentHeight, parentHeight,
			contentWidth, parentWidth, wrapped,
			contentHeight,
		)
	}
	gtlog.Debug(
		ctx,
		"base.Base.Height[%s]: display=%s whitespace=%s "+
			"vert_space=%d using min(content_height=%d, parent_height=%d)",
		b.Tag(), display, whitespace,
		vertSpace, contentHeight, parentHeight,
	)
	return types.Dimension(min(parentHeight, contentHeight))
}

// HasFixedHeight returns true if the Element's inner bounding box has a fixed
// height.
func (b *Base) HasFixedHeight() bool {
	_, ok := b.heightConstraint.(core.FixedConstraint)
	return ok
}

// FixedHeight returns the Element's fixed height. If the Element does not have
// a fixed height constraint, returns 0.
func (b *Base) FixedHeight() types.Dimension {
	if !b.HasFixedHeight() {
		return types.Dimension(0)
	}
	return types.Dimension(b.heightConstraint.(core.FixedConstraint))
}

// HasPercentHeight returns true if the Element's inner bounding box has a percent
// height.
func (b *Base) HasPercentHeight() bool {
	_, ok := b.heightConstraint.(core.PercentConstraint)
	return ok
}

// PercentHeight returns the Element's percent height. If the Element does not
// have a percent height constraint, returns 0.
func (b *Base) PercentHeight() types.Dimension {
	if !b.HasPercentHeight() {
		return types.Dimension(0)
	}
	return types.Dimension(b.heightConstraint.(core.PercentConstraint))
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

// HeightConstraint returns any optional size constraint for the Element's
// height. Returns nil when there is no height constraint.
func (b *Base) HeightConstraint() types.DimensionConstraint {
	return b.heightConstraint
}
