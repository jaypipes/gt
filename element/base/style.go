package base

import "github.com/jaypipes/gt/types"

// SetStyle sets the Element's style.
func (b *Base) SetStyle(style types.Style) types.Element {
	b.style = style
	return b
}

// Style returns the Element's Stylb.
func (b *Base) Style() types.Style {
	return b.style
}

// SetForegroundColor sets the Element's foreground color.
func (b *Base) SetForegroundColor(c types.Color) types.Element {
	b.style.Fg = c
	return b
}

// ForegroundColor returns the Element's foreground color.
func (b *Base) ForegroundColor() types.Color {
	return b.style.Fg
}

// SetBackgroundColor sets the Element's foreground color.
func (b *Base) SetBackgroundColor(c types.Color) types.Element {
	b.style.Bg = c
	return b
}

// BackgroundColor returns the Element's background color.
func (b *Base) BackgroundColor() types.Color {
	return b.style.Bg
}
