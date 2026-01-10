package span

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/element"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
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
	s.SetContent(content)
	return s
}

// Span is an Element that uses the inline display mode by default.
type Span struct {
	*element.Element
	core.Contented
}

// SetSize sets the fixed width and height of the Span and also sets the
// display mode to "inline-block".
func (s *Span) SetSize(width, height int) {
	s.Sized.SetSize(width, height)
	s.SetDisplay(types.DisplayInlineBlock)
}

// Draw renders the Span to the given screen in the specified bounding box.
func (s *Span) Draw(screen types.Screen, bounds types.Rectangle) {
	s.Element.Draw(screen, bounds)
	inner := s.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	ss := uv.NewStyledString(s.Content())
	ws := s.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}
