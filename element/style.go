package element

import (
	"context"
	"image/color"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/motif"
	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/types"
)

// ThemeClass returns the Element's ThemeClass, if any.
func (e *Element) ThemeClass() types.ThemeClass {
	return e.themeClass
}

// SetThemeClass sets the Element's ThemeClass.
func (e *Element) SetThemeClass(class types.ThemeClass) {
	e.themeClass = class
}

// WithThemeClass sets the Element's ThemeClass and returns the Element.
func (e *Element) WithThemeClass(class types.ThemeClass) types.Element {
	e.SetThemeClass(class)
	return e
}

// Theme returns the Element's Theme, if any.
func (e *Element) Theme() types.Theme {
	return e.theme
}

// SetTheme sets the Element's Theme.
func (e *Element) SetTheme(t types.Theme) {
	e.theme = t
}

// WithTheme sets the Element's Theme and returns the Element.
func (e *Element) WithTheme(t types.Theme) types.Element {
	e.SetTheme(t)
	return e
}

// Motif returns the Element's Motif, if any.
func (e *Element) Motif() types.Motif {
	t := e.Theme()
	if t == nil {
		return e.motif
	}
	return nil
}

// SetMotif sets the Element's Motif.
func (e *Element) SetMotif(m types.Motif) {
	e.motif = m
}

// WithMotif sets the Element's Motif and returns the Element.
func (e *Element) WithMotif(m types.Motif) types.Element {
	e.SetMotif(m)
	return e
}

// Unstyled returns true if the Element has no styling.
func (e *Element) Unstyled() bool {
	m := e.Motif()
	if m == nil || m.Unstyled() {
		return true
	}
	return false
}

// Style returns the Element's Style. If the Element has the focus, returns the
// Element's FocusStyle, if set. If the mouse is currently hovering over the
// Element, returns the Element's HoverStyle, if set. Otherwise, returns the
// Element's normal Style or if not set, the nearest parent's Style.
func (e *Element) Style() types.Style {
	m := e.motif
	if e.disabled && m != nil {
		ds := m.DisabledStyle()
		if ds != nil {
			return ds
		}
	}
	if e.focused && m != nil {
		fs := m.FocusedStyle()
		if fs != nil {
			return fs
		}
	}
	if e.hovered && m != nil {
		hs := m.HoveredStyle()
		if hs != nil {
			return hs
		}
	}
	var s types.Style
	if m != nil {
		s = m.NormalStyle()
	}
	if s != nil && !s.Unstyled() {
		return s
	}

	// If there is no style set, inherit from the nearest parent with non-empty
	// style.
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

// SetStyle sets the Element's normal Style. The normal Style is the style of
// the Element when the focusStyle or hoverStyle are not active for the
// Element.
func (e *Element) SetStyle(style types.Style) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetNormalStyle(style)
}

// WithStyle sets the Element's normal Style and returns the Element.
func (e *Element) WithStyle(style types.Style) types.Element {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetNormalStyle(style)
	return e
}

// DisabledStyle returns the Element's Style when it is disabled.
func (e *Element) DisabledStyle() types.Style {
	if e.motif == nil {
		return nil
	}
	return e.motif.DisabledStyle()
}

// SetDisabledStyle sets the Element's Style when it is disabled.
func (e *Element) SetDisabledStyle(style types.Style) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetDisabledStyle(style)
}

// WithDisabledStyle sets the Element's Style when it is disabled and returns
// the Element.
func (e *Element) WithDisabledStyle(style types.Style) types.Element {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetDisabledStyle(style)
	return e
}

// FocusedStyle returns the Element's Style when it has the focus.
func (e *Element) FocusedStyle() types.Style {
	if e.motif == nil {
		return nil
	}
	return e.motif.FocusedStyle()
}

// SetFocusedStyle sets the Element's Style when it has the focus.
func (e *Element) SetFocusedStyle(style types.Style) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetFocusedStyle(style)
}

// WithFocusedStyle sets the Element's Style when it has the focus and returns
// the Element.
func (e *Element) WithFocusedStyle(style types.Style) types.Element {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetFocusedStyle(style)
	return e
}

// HoveredStyle returns the Element's Style when the mouse is hovering over the
// Element.
func (e *Element) HoveredStyle() types.Style {
	if e.motif == nil {
		return nil
	}
	return e.motif.HoveredStyle()
}

// SetHoveredStyle sets the Element's Style when the mouse is hovering over the
// Element.
func (e *Element) SetHoveredStyle(style types.Style) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetHoveredStyle(style)
}

// WithHoveredStyle sets the Element's Style when the mouse is hovering over the
// Element and returns the Element.
func (e *Element) WithHoveredStyle(style types.Style) types.Element {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetHoveredStyle(style)
	return e
}

// Bold returns true if the Element is bolded.
func (e *Element) Bold() bool {
	s := e.Style()
	if s == nil {
		return false
	}
	return s.Bold()
}

// SetBold sets the Element's bold attribute.
func (e *Element) SetBold(on bool) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetBold(on)
	e.motif.SetNormalStyle(s)
}

// WithBold sets the Element's bold attribute and returns the Element
func (e *Element) WithBold(on bool) types.Element {
	e.SetBold(on)
	return e
}

// Italic returns true if the Element is italicized.
func (e *Element) Italic() bool {
	s := e.Style()
	if s == nil {
		return false
	}
	return s.Italic()
}

// SetItalic sets the Element's italic attribute.
func (e *Element) SetItalic(on bool) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetItalic(on)
	e.motif.SetNormalStyle(s)
}

// WithItalic sets the Element's italic attribute and returns the Element
func (e *Element) WithItalic(on bool) types.Element {
	e.SetItalic(on)
	return e
}

// Dim returns true if the Element is dimmed.
func (e *Element) Dim() bool {
	s := e.Style()
	if s == nil {
		return false
	}
	return s.Dim()
}

// SetDim sets the Element's dim attribute.
func (e *Element) SetDim(on bool) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetDim(on)
	e.motif.SetNormalStyle(s)
}

// WithDim sets the Element's dim attribute and returns the Element
func (e *Element) WithDim(on bool) types.Element {
	e.SetDim(on)
	return e
}

// Strikethrough returns true if the Element is struckthrough.
func (e *Element) Strikethrough() bool {
	s := e.Style()
	if s == nil {
		return false
	}
	return s.Strikethrough()
}

// SetStrikethrough sets the Element's strikethrough attribute.
func (e *Element) SetStrikethrough(on bool) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetStrikethrough(on)
	e.motif.SetNormalStyle(s)
}

// WithStrikethrough sets the Element's strikethrough attribute and returns the
// Element
func (e *Element) WithStrikethrough(on bool) types.Element {
	e.SetStrikethrough(on)
	return e
}

// Blink returns true if the Element is blinked.
func (e *Element) Blink() bool {
	s := e.Style()
	if s == nil {
		return false
	}
	return s.Blink()
}

// SetBlink sets the Element's blink attribute.
func (e *Element) SetBlink(on bool) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetBlink(on)
	e.motif.SetNormalStyle(s)
}

// WithBlink sets the Element's blink attribute and returns the Element
func (e *Element) WithBlink(on bool) types.Element {
	e.SetBlink(on)
	return e
}

// Underline returns true if the Element is underlined.
func (e *Element) Underline() bool {
	s := e.Style()
	if s == nil {
		return false
	}
	return s.Underline()
}

// UnderlineStyle returns the Element's underline style.
func (e *Element) UnderlineStyle() types.UnderlineStyle {
	s := e.Style()
	if s == nil {
		return types.UnderlineStyleNone
	}
	return s.UnderlineStyle()
}

// SetUnderlineStyle sets the Element's underline style.
func (e *Element) SetUnderlineStyle(us types.UnderlineStyle) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetUnderlineStyle(us)
	e.motif.SetNormalStyle(s)
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
	s := e.Style()
	if s == nil {
		return color.Transparent
	}
	return s.ForegroundColor()
}

// SetForegroundColor sets the Style's foreground color.
func (e *Element) SetForegroundColor(color types.Color) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetForegroundColor(color)
	e.motif.SetNormalStyle(s)
}

// WithForegroundColor sets the Element's foreground color and returns the
// Element.
func (e *Element) WithForegroundColor(color types.Color) types.Element {
	e.SetForegroundColor(color)
	return e
}

// BackgroundColor returns the Element's background color.
func (e *Element) BackgroundColor() types.Color {
	s := e.Style()
	if s == nil {
		return color.Transparent
	}
	return s.BackgroundColor()
}

// SetBackgroundColor sets the Style's background color.
func (e *Element) SetBackgroundColor(color types.Color) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetBackgroundColor(color)
	e.motif.SetNormalStyle(s)
}

// WithBackgroundColor sets the Element's background color and returns the
// Element.
func (e *Element) WithBackgroundColor(color types.Color) types.Element {
	e.SetBackgroundColor(color)
	return e
}

// UnderlineColor returns the Element's underline color.
func (e *Element) UnderlineColor() types.Color {
	s := e.Style()
	if s == nil {
		return color.Transparent
	}
	return s.UnderlineColor()
}

// SetUnderlineColor sets the Style's underline color.
func (e *Element) SetUnderlineColor(color types.Color) {
	s := e.Style()
	if s == nil {
		e.motif = motif.Empty()
	}
	s = e.motif.NormalStyle()
	if s == nil {
		s = style.Empty()
	}
	s.SetUnderlineColor(color)
	e.motif.SetNormalStyle(s)
}

// WithUnderlineColor sets the Element's underline color and returns the
// Element.
func (e *Element) WithUnderlineColor(color types.Color) types.Element {
	e.SetUnderlineColor(color)
	return e
}

var _ types.Style = (*Element)(nil)
