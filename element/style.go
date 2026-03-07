package element

import (
	"context"
	"image/color"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/types"
)

// Unstyled returns true if the Element has no styling.
func (e *Element) Unstyled() bool {
	return e.style == nil || e.style.Unstyled()
}

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
	// If there is no style set, inherit from the nearest parent with non-empty
	// style.
	if !e.Unstyled() {
		return e.style
	}
	ctx := context.TODO()
	parentNode := e.Parent()
	parent, ok := parentNode.(types.Element)
	if ok {
		if !parent.Unstyled() {
			gtlog.Debug(
				ctx, "Element.Style[%s]: inheriting parent %s style",
				e.Tag(), core.ID(parentNode),
			)
			return parent.Style()
		}
	}
	return nil
}

// Bold returns true if the Element is bolded.
func (e *Element) Bold() bool {
	if e.style == nil {
		return false
	}
	return e.style.Bold()
}

// SetBold sets the Element's bold attribute.
func (e *Element) SetBold(on bool) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetBold(on)
}

// WithBold sets the Element's bold attribute and returns the Element
func (e *Element) WithBold(on bool) types.Element {
	e.SetBold(on)
	return e
}

// Italic returns true if the Element is italicized.
func (e *Element) Italic() bool {
	if e.style == nil {
		return false
	}
	return e.style.Italic()
}

// SetItalic sets the Element's italic attribute.
func (e *Element) SetItalic(on bool) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetItalic(on)
}

// WithItalic sets the Element's italic attribute and returns the Element
func (e *Element) WithItalic(on bool) types.Element {
	e.SetItalic(on)
	return e
}

// Dim returns true if the Element is dimmed.
func (e *Element) Dim() bool {
	if e.style == nil {
		return false
	}
	return e.style.Dim()
}

// SetDim sets the Element's dim attribute.
func (e *Element) SetDim(on bool) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetDim(on)
}

// WithDim sets the Element's dim attribute and returns the Element
func (e *Element) WithDim(on bool) types.Element {
	e.SetDim(on)
	return e
}

// Strikethrough returns true if the Element is struckthrough.
func (e *Element) Strikethrough() bool {
	if e.style == nil {
		return false
	}
	return e.style.Strikethrough()
}

// SetStrikethrough sets the Element's strikethrough attribute.
func (e *Element) SetStrikethrough(on bool) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetStrikethrough(on)
}

// WithStrikethrough sets the Element's strikethrough attribute and returns the
// Element
func (e *Element) WithStrikethrough(on bool) types.Element {
	e.SetStrikethrough(on)
	return e
}

// Blink returns true if the Element is blinked.
func (e *Element) Blink() bool {
	if e.style == nil {
		return false
	}
	return e.style.Blink()
}

// SetBlink sets the Element's blink attribute.
func (e *Element) SetBlink(on bool) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetBlink(on)
}

// WithBlink sets the Element's blink attribute and returns the Element
func (e *Element) WithBlink(on bool) types.Element {
	e.SetBlink(on)
	return e
}

// Underline returns true if the Element is underlined.
func (e *Element) Underline() bool {
	if e.style == nil {
		return false
	}
	return e.style.Underline()
}

// UnderlineStyle returns the Element's underline style.
func (e *Element) UnderlineStyle() types.UnderlineStyle {
	if e.style == nil {
		return types.UnderlineStyleNone
	}
	return e.style.UnderlineStyle()
}

// SetUnderlineStyle sets the Element's underline style.
func (e *Element) SetUnderlineStyle(us types.UnderlineStyle) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetUnderlineStyle(us)
}

// WithUnderlineStyle sets the Element's underline style and returns the
// Element.
func (e *Element) WithUnderlineStyle(
	ulStyle types.UnderlineStyle,
) types.Element {
	e.SetUnderlineStyle(ulStyle)
	return e
}

// ForegroundColor returns the Element's underline color.
func (e *Element) ForegroundColor() types.Color {
	if e.style == nil {
		return color.Transparent
	}
	return e.style.ForegroundColor()
}

// SetForegroundColor sets the Style's foreground color.
func (e *Element) SetForegroundColor(color types.Color) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetForegroundColor(color)
}

// WithForegroundColor sets the Element's foreground color and returns the
// Element.
func (e *Element) WithForegroundColor(color types.Color) types.Element {
	e.SetForegroundColor(color)
	return e
}

// BackgroundColor returns the Element's background color.
func (e *Element) BackgroundColor() types.Color {
	if e.style == nil {
		return color.Transparent
	}
	return e.style.BackgroundColor()
}

// SetBackgroundColor sets the Style's background color.
func (e *Element) SetBackgroundColor(color types.Color) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetBackgroundColor(color)
}

// WithBackgroundColor sets the Element's background color and returns the
// Element.
func (e *Element) WithBackgroundColor(color types.Color) types.Element {
	e.SetBackgroundColor(color)
	return e
}

// UnderlineColor returns the Element's underline color.
func (e *Element) UnderlineColor() types.Color {
	if e.style == nil {
		return color.Transparent
	}
	return e.style.UnderlineColor()
}

// SetUnderlineColor sets the Style's underline color.
func (e *Element) SetUnderlineColor(color types.Color) {
	if e.style == nil {
		e.style = style.Empty()
	}
	e.style.SetUnderlineColor(color)
}

// WithUnderlineColor sets the Element's underline color and returns the
// Element.
func (e *Element) WithUnderlineColor(color types.Color) types.Element {
	e.SetUnderlineColor(color)
	return e
}

var _ types.Style = (*Element)(nil)
