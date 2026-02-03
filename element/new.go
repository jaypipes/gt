package element

import (
	"context"
	"sync"

	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a [element.Element] with the specified
// type/class.
//
// You can pass zero or more WithOptions to optionally set certain attributes
// on the returned Element.
func New(
	ctx context.Context,
	class string,
	opts ...types.ElementWithOption,
) Element {
	e := Element{
		RWMutex: new(sync.RWMutex),
		class:   class,
	}
	for _, opt := range opts {
		opt(&e)
	}
	return e
}

// WithID sets the types.Element's ID to the supplied value.
func WithID(id string) types.ElementWithOption {
	return func(e types.Element) {
		e.SetID(id)
	}
}

// WithBounds sets the types.Element's bounds to the supplied value.
func WithBounds(bounds types.Rectangle) types.ElementWithOption {
	return func(e types.Element) {
		e.SetBounds(bounds)
	}
}

// WithAbsolutePosition sets the types.Element's absolute position to the supplied
// value and marks the types.Element as using fixed positioning.
func WithAbsolutePosition(pt types.Point) types.ElementWithOption {
	return func(e types.Element) {
		e.SetAbsolutePosition(pt)
	}
}

// WithSize constrains the size of the types.Element.
func WithSize(constraint types.SizeConstraint) types.ElementWithOption {
	return func(e types.Element) {
		e.SetSize(constraint)
	}
}

// WithWidth constrains the width of the types.Element.
func WithWidth(constraint types.DimensionConstraint) types.ElementWithOption {
	return func(e types.Element) {
		e.SetWidth(constraint)
	}
}

// WithMinWidth sets the minimum width of the types.Element.
func WithMinWidth(width types.Dimension) types.ElementWithOption {
	return func(e types.Element) {
		e.SetMinWidth(width)
	}
}

// WithHeight constrains the height of the types.Element.
func WithHeight(constraint types.DimensionConstraint) types.ElementWithOption {
	return func(e types.Element) {
		e.SetHeight(constraint)
	}
}

// WithMinHeight sets the minimum height of the types.Element.
func WithMinHeight(height types.Dimension) types.ElementWithOption {
	return func(e types.Element) {
		e.SetMinHeight(height)
	}
}

// WithDisplay sets the types.Element's display mode to the supplied value.
func WithDisplay(display types.Display) types.ElementWithOption {
	return func(e types.Element) {
		e.SetDisplay(display)
	}
}

// WithAlignment sets the types.Element's alignment mode to the supplied value.
func WithAlignment(align types.Alignment) types.ElementWithOption {
	return func(e types.Element) {
		e.SetAlignment(align)
	}
}

// WithWhitespace sets the types.Element's whitespace mode to the supplied value.
func WithWhitespace(whitespace types.Whitespace) types.ElementWithOption {
	return func(e types.Element) {
		e.SetWhitespace(whitespace)
	}
}

// WithPadding sets the types.Element's padding to the supplied value.
func WithPadding(padding types.Padding) types.ElementWithOption {
	return func(e types.Element) {
		e.SetPadding(padding)
	}
}

// WithBorder sets the types.Element's border to the supplied value.
func WithBorder(border types.Border) types.ElementWithOption {
	return func(e types.Element) {
		e.SetBorder(border)
	}
}

// WithBorderForegroundColor sets the types.Element's border foreground color to the
// supplied value.
func WithBorderForegroundColor(color types.Color) types.ElementWithOption {
	return func(e types.Element) {
		e.SetBorderForegroundColor(color)
	}
}

// WithBorderBackgroundColor sets the types.Element's border background color to the
// supplied value.
func WithBorderBackgroundColor(color types.Color) types.ElementWithOption {
	return func(e types.Element) {
		e.SetBorderBackgroundColor(color)
	}
}

// WithStyle sets the types.Element's style to the supplied value.
func WithStyle(style types.Style) types.ElementWithOption {
	return func(e types.Element) {
		e.SetStyle(style)
	}
}

// WithForegroundColor sets the types.Element's foreground color to the supplied
// value.
func WithForegroundColor(color types.Color) types.ElementWithOption {
	return func(e types.Element) {
		e.SetForegroundColor(color)
	}
}

// WithBackgroundColor sets the types.Element's background color to the supplied
// value.
func WithBackgroundColor(color types.Color) types.ElementWithOption {
	return func(e types.Element) {
		e.SetBackgroundColor(color)
	}
}

// WithTextContent sets the types.Element's text content to the supplied value.
func WithTextContent(content string) types.ElementWithOption {
	return func(e types.Element) {
		e.SetTextContent(content)
	}
}
