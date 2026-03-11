package style

import (
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a Style.
//
// You can pass zero or more StyleWithOptions to optionally set certain
// attributes on the returned Styled.
func New(opts ...types.StyleWithOption) *Style {
	s := &Style{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Empty returns a new Style that has no attributes set.
func Empty() *Style {
	return &Style{}
}

// WithBold enables the bold attribute in the Style.
func WithBold() types.StyleWithOption {
	return func(s types.Style) {
		s.SetBold(true)
	}
}

// WithItalic enables the italic attribute in the Style.
func WithItalic() types.StyleWithOption {
	return func(s types.Style) {
		s.SetItalic(true)
	}
}

// WithDim enables the dim attribute in the Style.
func WithDim() types.StyleWithOption {
	return func(s types.Style) {
		s.SetDim(true)
	}
}

// WithStrikethrough enables the strikethrough attribute in the Style.
func WithStrikethrough() types.StyleWithOption {
	return func(s types.Style) {
		s.SetStrikethrough(true)
	}
}

// WithForegroundColor sets the types.Style's foreground color to the supplied
// value.
func WithForegroundColor(color types.Color) types.StyleWithOption {
	return func(s types.Style) {
		s.SetForegroundColor(color)
	}
}

// WithBackgroundColor sets the types.Style's background color to the supplied
// value.
func WithBackgroundColor(color types.Color) types.StyleWithOption {
	return func(s types.Style) {
		s.SetBackgroundColor(color)
	}
}

// WithUnderlineColor sets the types.Style's underline color to the supplied
// value.
func WithUnderlineColor(color types.Color) types.StyleWithOption {
	return func(s types.Style) {
		s.SetUnderlineColor(color)
	}
}

// WithUnderlineStyle sets the types.Style's underline shape to the supplied
// value.
func WithUnderlineStyle(
	ulStyle types.UnderlineStyle,
) types.StyleWithOption {
	return func(s types.Style) {
		s.SetUnderlineStyle(ulStyle)
	}
}
