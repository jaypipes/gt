package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Bordered describes something that has a bounding box.
type Bordered struct {
	// border is the optional Border information for the Bordered.
	border *types.Border
	// fgColor is the border foreground color (i.e the color of the border
	// cell's underlying grapheme).
	fgColor types.Color
	// bgColor is the border background color, i.e. the background color of the
	// border cells.
	bgColor types.Color
}

// SetRect sets the Element's bounding rectangle
func (b *Bordered) SetBorder(border types.Border) {
	b.border = &border
}

// Border returns the Bordered's border, if any.
func (b *Bordered) Border() *types.Border {
	return b.border
}

// SetBorderForegroundColor sets the Bordered's border foreground color
// (i.e the color of the border cell's underlying grapheme).
func (b *Bordered) SetBorderForegroundColor(c types.Color) {
	b.fgColor = c
}

// BorderForegroundColor returns the Bordered's border foreground color.
func (b *Bordered) BorderForegroundColor() types.Color {
	return b.fgColor
}

// SetBorderBackgroundColor sets the Bordered's border background color
// (i.e the background color of the border's cells.
func (b *Bordered) SetBorderBackgroundColor(c types.Color) {
	b.bgColor = c
}

// BorderBackgroundColor returns the Bordered's border background color.
func (b *Bordered) BorderBackgroundColor() types.Color {
	return b.bgColor
}

func (b *Bordered) Draw(screen types.Screen, bounds types.Rectangle) {
	// If we have a border, draw it around the outer bounding box.
	border := b.border
	if border == nil {
		return
	}
	style := types.Style{Fg: b.fgColor, Bg: b.bgColor}
	bb := border.Style(style)
	bb.Draw(screen, bounds)
}

var _ types.Bordered = (*Bordered)(nil)
