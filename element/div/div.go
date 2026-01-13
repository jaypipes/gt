package div

import (
	"context"
	"strings"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/mitchellh/go-wordwrap"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/element"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
)

const (
	ElementClass = "gt.div"
)

// New returns a new Div instance containing the supplied raw string content.
func New(
	ctx context.Context,
	content string,
) *Div {
	e := element.New(ctx, ElementClass)
	d := &Div{Element: e}
	d.SetDisplay(types.DisplayBlock)
	d.SetContent(content)
	return d
}

// Div is an Element that uses the block display mode by default.
type Div struct {
	*element.Element
	core.Contented
}

// Width returns the width of the Div.
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
func (d *Div) Width() int {
	parent := d.Parent()
	if parent == nil {
		return d.Bounds().Dx()
	}
	parentInner := parent.InnerBounds()
	parentWidth := parentInner.Dx()
	borderHoriz := 0
	border := d.Border()
	if border != nil {
		borderHoriz = render.BorderHorizontalWidth(*border)
	}
	paddingHoriz := d.Padding().Horizontal()

	ctx := context.TODO()
	display := d.Display()
	if display != types.DisplayInline && d.FixedWidth() {
		fixedWidth := d.Sized.Width() + paddingHoriz + borderHoriz
		gtlog.Debug(
			ctx,
			"Div.Width[%s]: display=%s padding_horiz=%d border_horiz=%d "+
				"using min(fixed_width=%d, parent_width=%d)",
			d.Tag(), display, paddingHoriz, borderHoriz,
			fixedWidth, parentWidth,
		)
		return min(parentWidth, fixedWidth)
	}
	if display != types.DisplayBlock {
		// For inline or inline-block with no fixed width, we use the lesser of
		// the parent's width or the "natural" width of the content
		contentWidth := ansi.StringWidth(d.Content())
		contentWidth += paddingHoriz + borderHoriz
		gtlog.Debug(
			ctx,
			"Div.Width[%s]: display=%s padding_horiz=%d border_horiz=%d "+
				"using min(content_width=%d, parent_width=%d)",
			d.Tag(), display, paddingHoriz, borderHoriz,
			contentWidth, parentWidth,
		)
		return min(parentWidth, contentWidth)
	}

	// For block display mode, we start the element at the Min.X coordinate of
	// the parent, which simulates a "newline" on the screen. We then calculate
	// the width based on the whether we have a width constraint.

	wc := d.WidthConstraint()

	// Calculate the width of this based on whether there is a width
	// constraint. If there is no constraint, the element receives the
	// remainder of the horizontal space in the parent's bounding box.
	if wc == nil {
		gtlog.Debug(
			ctx,
			"Div.Width[%s]: display=%s "+
				"width_constraint=none padding_horiz=%d border_horiz=%d. "+
				"remainder is parent inner bounds "+
				"width of %d minus padding_horiz + border_horiz",
			d.Tag(), display, paddingHoriz, borderHoriz, parentWidth,
		)
		return parentWidth - paddingHoriz
	}
	remainder := wc.Apply(parentWidth)
	gtlog.Debug(
		ctx,
		"Div.Width[%s]: display=%s "+
			"width_constraint=%d padding_horiz=%d border_horiz=%d. "+
			"calculated remainder of %d from parent width of %d",
		d.Tag(), display, wc, paddingHoriz, borderHoriz,
		remainder, parentWidth,
	)
	return remainder
}

// Height returns the height of the Div.
//
// If a fixed height has been set and the display mode is `block`, we use the
// fixed height plus any vertical padding.
//
// If a fixed height has not been set or the display mode is not `block`, the
// height defaults to the number of lines of text content, or 1 if there is no
// text content, plus any vertical padding.
func (d *Div) Height() int {
	parent := d.Parent()
	if parent == nil {
		return d.Bounds().Dx()
	}
	parentInner := parent.InnerBounds()
	parentWidth := parentInner.Dx()
	parentHeight := parentInner.Dy()
	paddingVert := d.Padding().Vertical()
	borderVert := 0
	border := d.Border()
	if border != nil {
		borderVert = render.BorderVerticalHeight(*border)
	}

	ctx := context.TODO()
	display := d.Display()
	if display == types.DisplayBlock && d.FixedHeight() {
		fixedHeight := d.Sized.Height() + paddingVert + borderVert
		gtlog.Debug(
			ctx,
			"Div.Height[%s]: display=%s padding_vert=%d border_vert=%d "+
				"using min(fixed_height=%d, parent_height=%d)",
			d.Tag(), display, paddingVert, borderVert,
			fixedHeight, parentHeight,
		)
		return min(parentHeight, fixedHeight)
	}

	whitespace := d.Whitespace()
	wrapNever := whitespace&types.WhitespaceWrapNever != 0
	if wrapNever {
		gtlog.Debug(
			ctx,
			"Div.Height[%s]: display=%s whitespace=%s "+
				"padding_vert=%d border_vert=%d "+
				"height is always 1 plus padding_vert + border_vert",
			d.Tag(), display, whitespace, paddingVert, borderVert,
		)
		return 1
	}

	// "wrap-line" whitespace mode means don't wrap except on existing
	// newlines...
	wrapLine := whitespace&types.WhitespaceWrapLine != 0
	wrapped := false

	//We use the "natural" height of the content, which is the number of
	//newlines in the content. However, we first need to calculate any
	//wrapping of long text content before returning the number of
	//newlines.
	content := d.Content()
	contentHeight := strings.Count(content, "\n") + 1
	contentHeight += paddingVert + borderVert
	origContentHeight := contentHeight
	contentWidth := ansi.StringWidth(content)
	if !wrapLine && (contentWidth > parentWidth) {
		wrapped = true
		wrappedContent := wordwrap.WrapString(content, uint(parentWidth))
		contentHeight = strings.Count(wrappedContent, "\n")
		contentHeight += paddingVert + borderVert
		gtlog.Debug(
			ctx,
			"Div.Height[%s]: display=%s whitespace=%s "+
				"padding_vert=%d border_vert=%d "+
				"original_content_height=%d parent_height=%d "+
				"content_width=%d parent_width=%d wrapped=%t "+
				"calculated new content_height of %d",
			d.Tag(), display, whitespace,
			paddingVert, borderVert,
			origContentHeight, parentHeight,
			contentWidth, parentWidth, wrapped,
			contentHeight,
		)
	}
	gtlog.Debug(
		ctx,
		"Div.Height[%s]: display=%s whitespace=%s "+
			"padding_vert=%d border_vert =%d "+
			"using min(content_height=%d, parent_height=%d)",
		d.Tag(), display, whitespace,
		paddingVert, borderVert,
		contentHeight, parentHeight,
	)
	return min(parentHeight, contentHeight)
}

// Draw renders the Div to the given screen in the specified bounding box.
func (d *Div) Draw(screen types.Screen, bounds types.Rectangle) {
	ctx := context.TODO()
	d.Element.Draw(screen, bounds)
	inner := d.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	content := render.AlignString(
		ctx, d.Content(), inner, d.Alignment(),
	)
	style := d.Style()
	content = style.Styled(content)
	ss := uv.NewStyledString(content)
	ws := d.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}
