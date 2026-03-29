package cell

import (
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a Cell.
//
// You can pass zero or more CellWithOptions to optionally set certain
// attributes on the returned Cell.
func New(
	opts ...types.CellWithOption,
) *Cell {
	c := &Cell{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Empty returns a new Cell with no content or style.
func Empty() *Cell {
	return &Cell{}
}

// WithContent sets the Cell's content to the supplied string or rune.
func WithContent[T string | rune](content T) types.CellWithOption {
	return func(c types.Cell) {
		c.SetContent(string(content))
	}
}

// WithBold enables the bold attribute in the Cell.
func WithBold() types.CellWithOption {
	return func(c types.Cell) {
		c.SetBold(true)
	}
}

// WithItalic enables the italic attribute in the Cell.
func WithItalic() types.CellWithOption {
	return func(c types.Cell) {
		c.SetItalic(true)
	}
}

// WithDim enables the dim attribute in the Cell.
func WithDim() types.CellWithOption {
	return func(c types.Cell) {
		c.SetDim(true)
	}
}

// WithStrikethrough enables the strikethrough attribute in the Cell.
func WithStrikethrough() types.CellWithOption {
	return func(c types.Cell) {
		c.SetStrikethrough(true)
	}
}

// WithBlink enables the blink attribute in the Cell.
func WithBlink() types.CellWithOption {
	return func(c types.Cell) {
		c.SetBlink(true)
	}
}

// WithUnderlineStyle sets the types.Cell's underline style to the supplied
// value.
func WithUnderlineStyle(ulStyle types.UnderlineStyle) types.CellWithOption {
	return func(c types.Cell) {
		c.SetUnderlineStyle(ulStyle)
	}
}

// WithForegroundColor sets the types.Cell's foreground color to the supplied
// value.
func WithForegroundColor(color types.Color) types.CellWithOption {
	return func(c types.Cell) {
		c.SetForegroundColor(color)
	}
}

// WithBackgroundColor sets the types.Cell's background color to the supplied
// value.
func WithBackgroundColor(color types.Color) types.CellWithOption {
	return func(c types.Cell) {
		c.SetBackgroundColor(color)
	}
}

// WithUnderlineColor sets the types.Cell's underline color to the supplied
// value.
func WithUnderlineColor(color types.Color) types.CellWithOption {
	return func(c types.Cell) {
		c.SetUnderlineColor(color)
	}
}
