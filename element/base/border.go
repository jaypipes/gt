package base

import "github.com/jaypipes/gt/types"

// SetRect sets the Element's bounding rectangle
func (b *Base) SetBorder(border types.Border) types.Element {
	b.border = &border
	return b
}

// Border returns the Element's border, if any.
func (b *Base) Border() *types.Border {
	return b.border
}

// SetBorderForegroundColor sets the Element's border foreground color
// (i.e the color of the border cell's underlying grapheme).
func (b *Base) SetBorderForegroundColor(c types.Color) types.Element {
	b.borderFGColor = c
	return b
}

// BorderForegroundColor returns the Element's border foreground color.
func (b *Base) BorderForegroundColor() types.Color {
	return b.borderFGColor
}

// SetBorderBackgroundColor sets the Element's border background color
// (i.e the background color of the border's cells.
func (b *Base) SetBorderBackgroundColor(c types.Color) types.Element {
	b.borderBGColor = c
	return b
}

// BorderBackgroundColor returns the Element's border background color.
func (b *Base) BorderBackgroundColor() types.Color {
	return b.borderBGColor
}

// drawBorder draws the border around the outer bounding box's cells.
func (b *Base) drawBorder(screen types.Screen) {
	// If we have a border, draw it around the outer bounding box.
	border := b.border
	if border == nil {
		return
	}
	style := types.Style{Fg: b.borderFGColor, Bg: b.borderBGColor}
	bb := border.Style(style)
	bb.Draw(screen, b.bounds)
}
