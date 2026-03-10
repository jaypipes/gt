package style

import (
	"fmt"
	"strings"

	"github.com/jaypipes/gt/types"
)

type attrs uint8

const (
	attrNone          attrs = 0
	attrBold                = 1
	attrBlink               = 1 << 1
	attrReverse             = 1 << 2
	attrDim                 = 1 << 3
	attrItalic              = 1 << 4
	attrStrikethrough       = 1 << 5
)

// Style represents the style of a [Cell] being displayed in a [Screen].
type Style struct {
	// attrs is a bitmap storing style attributes
	attrs attrs
	// fgColor is the foreground color of the Style, if any
	fgColor types.Color
	// bgColor is the background color or the Style, if any
	bgColor types.Color
	// ulStyle is the style of the underline, if any
	ulStyle types.UnderlineStyle
	// ulColor is the style of the underline, if any
	ulColor types.Color
}

// String returns a short string description of the Style.
func (s *Style) String() string {
	if s.Unstyled() {
		return "none"
	}
	parts := []string{}
	attrsOn := []string{}
	if s.Bold() {
		attrsOn = append(attrsOn, "bold")
	}
	if s.Italic() {
		attrsOn = append(attrsOn, "italic")
	}
	if s.Dim() {
		attrsOn = append(attrsOn, "dim")
	}
	if s.Strikethrough() {
		attrsOn = append(attrsOn, "strikethrough")
	}
	if s.Blink() {
		attrsOn = append(attrsOn, "blink")
	}
	if s.Underline() {
		attrsOn = append(attrsOn, "underline")
	}
	if len(attrsOn) > 0 {
		parts = append(
			parts,
			fmt.Sprintf("attrs:%s", strings.Join(attrsOn, ",")),
		)
	}
	if s.fgColor != nil {
		parts = append(parts, fmt.Sprintf("fg:%s", colorRGBHex(s.fgColor)))
	}
	if s.bgColor != nil {
		parts = append(parts, fmt.Sprintf("bg:%s", colorRGBHex(s.bgColor)))
	}
	return strings.Join(parts, " ")
}

// Unstyled returns true if the Style hasn't had any attributes set.
func (s *Style) Unstyled() bool {
	return s.attrs == attrNone &&
		s.ulStyle == types.UnderlineStyleNone &&
		s.fgColor == nil &&
		s.bgColor == nil
}

// Bold returns true if the Style is bolded.
func (s *Style) Bold() bool {
	return s.attrs&attrBold != 0
}

// SetBold sets the Style's bold attribute.
func (s *Style) SetBold(on bool) {
	if on {
		s.attrs |= attrBold
	} else {
		s.attrs &^= attrBold
	}
}

// WithBold sets the Style's bold attribute and returns the Style
func (s *Style) WithBold(on bool) *Style {
	s.SetBold(on)
	return s
}

// Italic returns true if the Style is bolded.
func (s *Style) Italic() bool {
	return s.attrs&attrItalic != 0
}

// SetItalic sets the Style's italic attribute.
func (s *Style) SetItalic(on bool) {
	if on {
		s.attrs |= attrItalic
	} else {
		s.attrs &^= attrItalic
	}
}

// WithItalic sets the Style's italic attribute and returns the Style
func (s *Style) WithItalic(on bool) *Style {
	s.SetItalic(on)
	return s
}

// Dim returns true if the Style is dimmed.
func (s *Style) Dim() bool {
	return s.attrs&attrDim != 0
}

// SetDim sets the Style's dim attribute.
func (s *Style) SetDim(on bool) {
	if on {
		s.attrs |= attrDim
	} else {
		s.attrs &^= attrDim
	}
}

// WithDim sets the Style's dim attribute and returns the Style
func (s *Style) WithDim(on bool) *Style {
	s.SetDim(on)
	return s
}

// Strikethrough returns true if the Style is bolded.
func (s *Style) Strikethrough() bool {
	return s.attrs&attrStrikethrough != 0
}

// SetStrikethrough sets the Style's strikethrough attribute.
func (s *Style) SetStrikethrough(on bool) {
	if on {
		s.attrs |= attrStrikethrough
	} else {
		s.attrs &^= attrStrikethrough
	}
}

// WithStrikethrough sets the Style's strikethrough attribute and returns the
// Style
func (s *Style) WithStrikethrough(on bool) *Style {
	s.SetStrikethrough(on)
	return s
}

// Blink returns true if the Style has blinking on.
func (s *Style) Blink() bool {
	return s.attrs&attrBlink != 0
}

// SetBlink sets the Style's blink attribute.
func (s *Style) SetBlink(on bool) {
	if on {
		s.attrs |= attrBlink
	} else {
		s.attrs &^= attrBlink
	}
}

// WithBlink sets the Style's blink attribute and returns the Style
func (s *Style) WithBlink(on bool) *Style {
	s.SetBlink(on)
	return s
}

// Underline returns true if the Style is underlined.
func (s *Style) Underline() bool {
	return s.ulStyle != types.UnderlineStyleNone
}

// UnderlineStyle returns the Style's underline style.
func (s *Style) UnderlineStyle() types.UnderlineStyle {
	return s.ulStyle
}

// SetUnderlineStyle sets the Style's underline style.
func (s *Style) SetUnderlineStyle(us types.UnderlineStyle) {
	s.ulStyle = us
}

// WithUnderlineStyle sets the Style's underline style and returns the Style.
func (s *Style) WithUnderlineStyle(ulStyle types.UnderlineStyle) *Style {
	s.SetUnderlineStyle(ulStyle)
	return s
}

// UnderlineColor returns the Style's underline color.
func (s *Style) UnderlineColor() types.Color {
	return s.ulColor
}

// SetUnderlineColor sets the Style's underline color.
func (s *Style) SetUnderlineColor(color types.Color) {
	s.ulColor = color
}

// WithUnderlineColor sets the Style's underline color and returns the Style.
func (s *Style) WithUnderlineColor(color types.Color) *Style {
	s.SetUnderlineColor(color)
	return s
}

// ForegroundColor returns the Style's foreground color.
func (s *Style) ForegroundColor() types.Color {
	return s.fgColor
}

// SetForegroundColor sets the Style's foreground color.
func (s *Style) SetForegroundColor(color types.Color) {
	s.fgColor = color
}

// WithForegroundColor sets the Style's foreground color and returns the Style.
func (s *Style) WithForegroundColor(color types.Color) *Style {
	s.SetForegroundColor(color)
	return s
}

// BackgroundColor returns the Style's background color.
func (s *Style) BackgroundColor() types.Color {
	return s.bgColor
}

// SetBackgroundColor sets the Style's background color.
func (s *Style) SetBackgroundColor(color types.Color) {
	s.bgColor = color
}

// WithBackgroundColor sets the Style's background color and returns the Style.
func (s *Style) WithBackgroundColor(color types.Color) *Style {
	s.SetBackgroundColor(color)
	return s
}

// colorRGBHex returns the supplied color's 6-character (RRGGBB) hex string.
func colorRGBHex(c types.Color) string {
	cr, cg, cb, _ := c.RGBA()
	r := uint8(cr >> 8)
	g := uint8(cg >> 8)
	b := uint8(cb >> 8)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

var _ types.Style = (*Style)(nil)
