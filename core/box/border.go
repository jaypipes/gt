package box

import (
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// SetBorder sets the Box's border.
func (b *Box) SetBorder(border types.Border) {
	b.border = &border
}

// Border returns the Box's border, if any.
func (b *Box) Border() *types.Border {
	return b.border
}

// SetBorderForegroundColor sets the Box's border foreground color
// (i.e the color of the border cell's underlying grapheme).
func (b *Box) SetBorderForegroundColor(c types.Color) {
	b.borderFGColor = c
}

// BorderForegroundColor returns the Box's border foreground color.
func (b *Box) BorderForegroundColor() types.Color {
	return b.borderFGColor
}

// SetBorderBackgroundColor sets the Box's border background color
// (i.e the background color of the border's cells.
func (b *Box) SetBorderBackgroundColor(c types.Color) {
	b.borderBGColor = c
}

// BorderBackgroundColor returns the Box's border background color.
func (b *Box) BorderBackgroundColor() types.Color {
	return b.borderBGColor
}

// DrawBorder draws the border around the outer bounding box's cells.
func (b *Box) DrawBorder(screen types.Screen) {
	// If we have a border, draw it around the outer bounding box.
	border := b.border
	if border == nil {
		return
	}
	style := types.Style{Fg: b.borderFGColor, Bg: b.borderBGColor}
	bb := border.Style(style)
	bb.Draw(screen, b.bounds)
}

// SetPadding sets the Box's padding.
func (b *Box) SetPadding(padding types.Padding) {
	b.padding = padding
}

// Padding returns the padding for the Box.
func (b *Box) Padding() types.Padding {
	return b.padding
}

// HorizontalSpace returns the number of cells consumed by the Box's
// left-right padding and border.
func (b *Box) HorizontalSpace() types.Dimension {
	space := b.padding.HorizontalSpace()
	if b.border != nil {
		space += render.BorderHorizontalSpace(*b.border)
	}
	return space
}

// VerticalSpace returns the number of lines consumed by the Box's
// top-bottom padding and border
func (b *Box) VerticalSpace() types.Dimension {
	space := b.padding.VerticalSpace()
	if b.border != nil {
		space += render.BorderVerticalSpace(*b.border)
	}
	return space
}
