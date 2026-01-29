package element

import "github.com/jaypipes/gt/types"

// SetStyle sets the Element's style.
func (e *Element) SetStyle(style types.Style) {
	e.style = style
}

// WithStyle sets the Element's style and returns the Element.
func (e *Element) WithStyle(style types.Style) types.Element {
	e.style = style
	return e
}

// Style returns the Element's Style.
func (e *Element) Style() types.Style {
	return e.style
}

// SetForegroundColor sets the Element's foreground color.
func (e *Element) SetForegroundColor(c types.Color) {
	e.style.Fg = c
}

// WithForegroundColor sets the Element's foreground color and returns the Element.
func (e *Element) WithForegroundColor(c types.Color) types.Element {
	e.style.Fg = c
	return e
}

// ForegroundColor returns the Element's foreground color.
func (e *Element) ForegroundColor() types.Color {
	return e.style.Fg
}

// SetBackgroundColor sets the Element's background color.
func (e *Element) SetBackgroundColor(c types.Color) {
	e.style.Bg = c
}

// WithBackgroundColor sets the Element's background color and returns the Element.
func (e *Element) WithBackgroundColor(c types.Color) types.Element {
	e.style.Bg = c
	return e
}

// BackgroundColor returns the Element's background color.
func (e *Element) BackgroundColor() types.Color {
	return e.style.Bg
}
