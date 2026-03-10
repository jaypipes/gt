package types

import "fmt"

// Style represents the style of a something being displayed in a [Screen].
type Style interface {
	fmt.Stringer
	// Unstyled returns true if there is no styling applied.
	Unstyled() bool
	// SetBold returns whether the Style's bold attribute is set.
	Bold() bool
	// SetBold sets the Style's bold attribute.
	SetBold(bool)
	// Italic returns whether the Style's italic attribute is set.
	Italic() bool
	// SetItalic sets the Style's italic attribute.
	SetItalic(bool)
	// Dim returns whether the Style's dim attribute is set.
	Dim() bool
	// SetDim sets the Style's dim attribute.
	SetDim(bool)
	// Strikethrough returns whether theStyle's strikethrough attribute is
	// set.
	Strikethrough() bool
	// SetStrikethrough sets the Style's strikethrough attribute.
	SetStrikethrough(bool)
	// Blink returns whether the Style's blink attribute is set.
	Blink() bool
	// SetBlink sets the Style's blink attribute.
	SetBlink(bool)
	// Underline returns whether the Style's underline style is not none.
	Underline() bool
	// UnderlineStyle sets the Style's underline style.
	UnderlineStyle() UnderlineStyle
	// SetUnderlineStyle sets the Style's underline style.
	SetUnderlineStyle(UnderlineStyle)
	// SetForegroundColor returns the Style's foreground color.
	ForegroundColor() Color
	// SetForegroundColor sets the Style's foreground color.
	SetForegroundColor(Color)
	// SetBackgroundColor returns the Style's background color.
	BackgroundColor() Color
	// SetBackgroundColor sets the Style's background color.
	SetBackgroundColor(Color)
	// SetUnderlineColor returns the Style's underline color.
	UnderlineColor() Color
	// SetUnderlineColor sets the Style's underline color.
	SetUnderlineColor(Color)
}

// StyleWithOption describes an optional varg parameter to [style.New] that
// modifies the returned Style.
type StyleWithOption func(Style)
