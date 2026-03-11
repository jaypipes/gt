package border

import (
	"github.com/jaypipes/gt/core/cell"
	"github.com/jaypipes/gt/types"
)

// New returns a new Border
func New(opts ...types.BorderWithOption) types.Border {
	b := &Border{}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// WithT sets the Border's top edge cell. The argument can be either a
// [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithT(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetT(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetT(c)
		}
	}
}

// WithB sets the Border's bottom edge cell. The argument can be either a
// [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithB(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetB(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetB(c)
		}
	}
}

// WithL sets the Border's left edge cell. The argument can be either a
// [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithL(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetL(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetL(c)
		}
	}
}

// WithR sets the Border's right edge cell. The argument can be either a
// [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithR(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetR(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetR(c)
		}
	}
}

// WithTL sets the Border's top left corner cell. The argument can be either a
// [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithTL(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetTL(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetTL(c)
		}
	}
}

// WithTR sets the Border's top right corner cell. The argument can be either a
// [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithTR(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetTR(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetTR(c)
		}
	}
}

// WithBL sets the Border's bottom left corner cell. The argument can be either
// a [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithBL(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetBL(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetBL(c)
		}
	}
}

// WithR sets the Border's bottom right corner cell. The argument can be either
// a [types.Cell] or a string. If a string, we create an unstyled [types.Cell]
// with that string content.
func WithBR(arg any) types.BorderWithOption {
	return func(b types.Border) {
		switch arg := arg.(type) {
		case types.Cell:
			b.SetBR(arg)
		case string:
			c := cell.New(cell.WithContent(arg))
			b.SetBR(c)
		}
	}
}

// WithForegroundColor sets the types.Border's foreground color to the supplied
// value.
func WithForegroundColor(color types.Color) types.BorderWithOption {
	return func(b types.Border) {
		b.SetForegroundColor(color)
	}
}

// WithBackgroundColor sets the types.Border's background color to the supplied
// value.
func WithBackgroundColor(color types.Color) types.BorderWithOption {
	return func(b types.Border) {
		b.SetBackgroundColor(color)
	}
}
