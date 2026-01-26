package base

import (
	"context"
	"strings"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
	"github.com/mitchellh/go-wordwrap"
)

// WithSize constrains the size of the Element's outer bounding box and returns
// the Element.
func (b *Base) WithSize(constraint types.SizeConstraint) types.Element {
	b.Box.SetSize(constraint)
	return b
}

// WithWidth constrains the width of the Element and returns the Element.
func (b *Base) WithWidth(constraint types.DimensionConstraint) types.Element {
	b.Box.SetWidth(constraint)
	return b
}

// WithMinWidth sets the minimum width of the Element and returns the Element.
func (b *Base) WithMinWidth(w types.Dimension) types.Element {
	b.Box.SetMinWidth(w)
	return b
}

// WithHeight constrains the height of the Element and returns the Element.
func (b *Base) WithHeight(constraint types.DimensionConstraint) types.Element {
	b.Box.SetHeight(constraint)
	return b
}

// WithMinHeight sets the minimum height of the Element and returns the
// Element.
func (b *Base) WithMinHeight(w types.Dimension) types.Element {
	b.Box.SetMinHeight(w)
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
	parent := b.Parent()
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
	childIndex := b.ChildIndex()
	for _, child := range parent.Children() {
		if child.ChildIndex() == childIndex {
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
