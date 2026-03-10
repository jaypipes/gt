package cell

import (
	"fmt"
	"image/color"

	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/types"
)

// Cell represents a single cell within a [types.Screen]
type Cell struct {
	// style contains the Cell's Style, if any
	style types.Style
	// content is the Cell's string content
	content string
}

func (c *Cell) String() string {
	s := "none"
	if c.style != nil {
		s = c.style.String()
	}
	return fmt.Sprintf("cell style=%s", s)
}

// Unstyled returns true if the Cell has no style applied.
func (c *Cell) Unstyled() bool {
	return (c.style == nil || c.style.Unstyled())
}

// Empty returns true if the Cell has no content.
func (c *Cell) Empty() bool {
	return len(c.content) == 0 &&
		(c.style == nil || c.style.Unstyled())
}

// Content returns the Cell's string content.
func (c *Cell) Content() string {
	return c.content
}

// SetContent sets the Cell's string content.
func (c *Cell) SetContent(content string) {
	c.content = content
}

// Style returns the Cell's Style, if any.
func (c *Cell) Style() types.Style {
	return c.style
}

// SetStyle sets the Cell's Style.
func (c *Cell) SetStyle(s types.Style) {
	c.style = s
}

// WithStyle sets the Cell's Style and returns the Cell.
func (c *Cell) WithStyle(s types.Style) types.Cell {
	c.SetStyle(s)
	return c
}

// Bold returns true if the Cell is bolded.
func (c *Cell) Bold() bool {
	if c.style == nil {
		return false
	}
	return c.style.Bold()
}

// SetBold sets the Cell's bold attribute.
func (c *Cell) SetBold(on bool) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetBold(on)
}

// WithBold sets the Cell's bold attribute and returns the Cell
func (c *Cell) WithBold(on bool) types.Cell {
	c.SetBold(on)
	return c
}

// Italic returns true if the Cell is bolded.
func (c *Cell) Italic() bool {
	if c.style == nil {
		return false
	}
	return c.style.Italic()
}

// SetItalic sets the Cell's italic attribute.
func (c *Cell) SetItalic(on bool) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetItalic(on)
}

// WithItalic sets the Cell's italic attribute and returns the Cell
func (c *Cell) WithItalic(on bool) types.Cell {
	c.SetItalic(on)
	return c
}

// Dim returns true if the Cell is dimmed.
func (c *Cell) Dim() bool {
	if c.style == nil {
		return false
	}
	return c.style.Dim()
}

// SetDim sets the Cell's dim attribute.
func (c *Cell) SetDim(on bool) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetDim(on)
}

// WithDim sets the Cell's dim attribute and returns the Cell
func (c *Cell) WithDim(on bool) types.Cell {
	c.SetDim(on)
	return c
}

// Strikethrough returns true if the Cell is bolded.
func (c *Cell) Strikethrough() bool {
	if c.style == nil {
		return false
	}
	return c.style.Strikethrough()
}

// SetStrikethrough sets the Cell's strikethrough attribute.
func (c *Cell) SetStrikethrough(on bool) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetStrikethrough(on)
}

// WithStrikethrough sets the Cell's strikethrough attribute and returns the
// Cell
func (c *Cell) WithStrikethrough(on bool) types.Cell {
	c.SetStrikethrough(on)
	return c
}

// Blink returns true if the Cell has blinking on.
func (c *Cell) Blink() bool {
	if c.style == nil {
		return false
	}
	return c.style.Blink()
}

// SetBlink sets the Cell's blink attribute.
func (c *Cell) SetBlink(on bool) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetBlink(on)
}

// WithBlink sets the Cell's blink attribute and returns the Cell
func (c *Cell) WithBlink(on bool) types.Cell {
	c.SetBlink(on)
	return c
}

// Underline returns true if the Cell is underlined.
func (c *Cell) Underline() bool {
	if c.style == nil {
		return false
	}
	return c.style.Underline()
}

// UnderlineCell returns the Cell's underline style.
func (c *Cell) UnderlineStyle() types.UnderlineStyle {
	if c.style == nil {
		return types.UnderlineStyleNone
	}
	return c.style.UnderlineStyle()
}

// SetUnderlineCell sets the Cell's underline style.
func (c *Cell) SetUnderlineStyle(ulStyle types.UnderlineStyle) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetUnderlineStyle(ulStyle)
}

// WithUnderlineCell sets the Cell's underline style and returns the Cell.
func (c *Cell) WithUnderlineStyle(ulStyle types.UnderlineStyle) types.Cell {
	c.SetUnderlineStyle(ulStyle)
	return c
}

// UnderlineColor returns the Cell's foreground color.
func (c *Cell) UnderlineColor() types.Color {
	if c.style == nil {
		return color.Transparent
	}
	return c.style.UnderlineColor()
}

// SetUnderlineColor sets the Cell's underline color.
func (c *Cell) SetUnderlineColor(color types.Color) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetUnderlineColor(color)
}

// WithUnderlineColor sets the Cell's underline color and returns the Cell.
func (c *Cell) WithUnderlineColor(color types.Color) types.Cell {
	c.SetUnderlineColor(color)
	return c
}

// ForegroundColor returns the Cell's foreground color.
func (c *Cell) ForegroundColor() types.Color {
	if c.style == nil {
		return color.Transparent
	}
	return c.style.ForegroundColor()
}

// SetForegroundColor sets the Cell's foreground color.
func (c *Cell) SetForegroundColor(color types.Color) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetForegroundColor(color)
}

// WithForegroundColor sets the Cell's foreground color and returns the Cell.
func (c *Cell) WithForegroundColor(color types.Color) types.Cell {
	c.SetForegroundColor(color)
	return c
}

// BackgroundColor returns the Cell's background color.
func (c *Cell) BackgroundColor() types.Color {
	if c.style == nil {
		return color.Transparent
	}
	return c.style.BackgroundColor()
}

// SetBackgroundColor sets the Cell's background color.
func (c *Cell) SetBackgroundColor(color types.Color) {
	if c.style == nil {
		c.style = style.Empty()
	}
	c.style.SetBackgroundColor(color)
}

// WithBackgroundColor sets the Cell's background color and returns the Cell.
func (c *Cell) WithBackgroundColor(color types.Color) types.Cell {
	c.SetBackgroundColor(color)
	return c
}

var _ types.Cell = (*Cell)(nil)
