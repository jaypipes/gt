package div

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/element"
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
}

// SetContent sets the Div's content to the supplied thing. The supplied
// thing can be []byte, string, or *uv.StyledString
func (s *Div[T]) SetContent(content T) {
	s.textContent = content
}

// Draw renders the Div to the given screen in the specified bounding box.
func (s *Div[T]) Draw(screen types.Screen, bounds types.Rectangle) {
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
