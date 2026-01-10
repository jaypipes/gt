package span

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/element"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/text"
	"github.com/jaypipes/gt/core/types"
)

const (
	ElementClass = "gt.span"
)

// New returns a new Span instance.
func New[T types.Text](
	ctx context.Context,
	content T,
) *Span[T] {
	e := element.New(ctx, ElementClass)
	d := &Span[T]{
		Element: e, textContent: content,
	}
	d.SetDisplay(types.DisplayInline)
	return d
}

// Span is an Element that uses the inline display mode by default.
type Span[T types.Text] struct {
	*element.Element
	// textContent is the unstyled text content of the Span.
	textContent T
}

// SetSize sets the fixed width and height of the Span and also sets the
// display mode to "inline-block".
func (s *Span[T]) SetSize(width, height int) {
	s.Sized.SetSize(width, height)
	s.SetDisplay(types.DisplayInlineBlock)
}

// SetContent sets the Span's content to the supplied thing. The supplied
// thing can be []byte, string, or *uv.StyledString
func (s *Span[T]) SetContent(content T) {
	s.textContent = content
}

// Draw renders the Span to the given screen in the specified bounding box.
func (s *Span[T]) Draw(screen types.Screen, bounds types.Rectangle) {
	s.Element.Draw(screen, bounds)
	inner := s.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	ss := uv.NewStyledString(text.String(s.textContent))
	ws := s.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}
