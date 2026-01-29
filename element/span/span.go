package span

import (
	"context"

	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.span"
)

// New returns a new Span instance containing the supplied raw text string.
func New(
	ctx context.Context,
	content string,
) *Span {
	e := element.New(ctx, ElementClass)
	s := &Span{Element: e}
	s.SetDisplay(types.DisplayInline)
	s.SetTextContent(content)
	return s
}

// Span is an Element that uses inline display mode by default.
type Span struct {
	element.Element
}

// SetSize sets the fixed width and height of the Span and also sets the
// display mode to `inline-block`.
func (s *Span) SetSize(constraint types.SizeConstraint) {
	s.Element.SetSize(constraint)
	s.SetDisplay(types.DisplayInlineBlock)
}

// WithSize sets the fixed width and height of the Span and also sets the
// display mode to `inline-block` and returns the Span.
func (s *Span) WithSize(constraint types.SizeConstraint) types.Element {
	s.SetSize(constraint)
	return s
}

// SetWidth sets the fixed width of the Span and also sets the display mode to
// `inline-block`.
func (s *Span) SetWidth(constraint types.DimensionConstraint) {
	s.Element.SetWidth(constraint)
	s.SetDisplay(types.DisplayInlineBlock)
}

// WithWidth sets the fixed width of the Span and also sets the display mode to
// `inline-block` and returns the Span.
func (s *Span) WithWidth(constraint types.DimensionConstraint) types.Element {
	s.SetWidth(constraint)
	return s
}

// SetHeight sets the fixed height of the Span and also sets the display mode
// to `inline-block`.
func (s *Span) SetHeight(constraint types.DimensionConstraint) {
	s.Element.SetHeight(constraint)
	s.SetDisplay(types.DisplayInlineBlock)
}

// WithHeight sets the fixed height of the Span and also sets the display mode
// to `inline-block` and returns the Span.
func (s *Span) WithHeight(constraint types.DimensionConstraint) types.Element {
	s.SetHeight(constraint)
	return s
}
