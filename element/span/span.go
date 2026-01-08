package span

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/element"
	gtlog "github.com/jaypipes/gt/core/log"
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
	// wrap indicates the text content should be wrapped.
	wrap bool
}

// SetContent sets the Span's content to the supplied thing. The supplied
// thing can be []byte, string, or *uv.StyledString
func (s *Span[T]) SetContent(content T) {
	s.textContent = content
}

// SetWrap sets the Span's wrapping behaviour.
func (s *Span[T]) SetWrap(enabled bool) {
	s.wrap = enabled
}

// Draw renders the Span to the given buffer at the specified area.
func (s *Span[T]) Prerender(ctx context.Context, buf types.Screen, area types.Rectangle) {
	gtlog.Debug(ctx, "Span(%s).Prerender: bounding box %s\n", s.ID(), area)
	// Draw the border, if any, and clear the inner bounding box of this
	// Element.
	s.Element.Draw(buf, area)
	inner := s.InnerBounds()
	innerClipped := render.Overlapping(area, inner)
	ss := uv.NewStyledString(text.String(s.textContent))
	ss.Wrap = s.wrap
	ss.Draw(buf, innerClipped)
}
