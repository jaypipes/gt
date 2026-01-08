package div

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
	ElementClass = "gt.div"
)

// New returns a new Div instance.
func New[T types.Text](
	ctx context.Context,
	content T,
) *Div[T] {
	e := element.New(ctx, ElementClass)
	d := &Div[T]{
		Element:     e,
		textContent: content,
	}
	d.SetDisplay(types.DisplayBlock)
	return d
}

// Div is an Element that uses the block display mode by default.
type Div[T types.Text] struct {
	*element.Element
	// textContent is the unstyled text content of the Div.
	textContent T
	// wrap indicates the text content should be wrapped.
	wrap bool
}

// SetContent sets the Div's content to the supplied thing. The supplied
// thing can be []byte, string, or *uv.StyledString
func (s *Div[T]) SetContent(content T) {
	s.textContent = content
}

// SetWrap sets the Div's wrapping behaviour.
func (s *Div[T]) SetWrap(enabled bool) {
	s.wrap = enabled
}

// Draw renders the Div to the given buffer at the specified area.
func (s *Div[T]) Prerender(ctx context.Context, buf types.Screen, area types.Rectangle) {
	gtlog.Debug(ctx, "Div(%s).Prerender: bounding box %s\n", s.ID(), area)
	// Draw the border, if any, and clear the inner bounding box of this
	// Element.
	s.Element.Draw(buf, area)
	inner := s.InnerBounds()
	innerClipped := render.Overlapping(area, inner)
	ss := uv.NewStyledString(text.String(s.textContent))
	ss.Wrap = s.wrap
	ss.Draw(buf, innerClipped)
}
