package span

import (
	"context"
	"strings"

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
// display mode to `inline-block`.
func (s *Span) SetSize(width, height int) {
	s.Sized.SetSize(width, height)
	s.SetDisplay(types.DisplayInlineBlock)
}

// SetWidth sets the fixed width of the Span and also sets the display mode to
// `inline-block`.
func (s *Span) SetWidth(width int) {
	s.Sized.SetWidth(width)
	s.SetDisplay(types.DisplayInlineBlock)
}

// SetHeight sets the fixed height of the Span and also sets the display mode
// to `inline-block`.
func (s *Span) SetHeight(height int) {
	s.Sized.SetHeight(height)
	s.SetDisplay(types.DisplayInlineBlock)
}

// Height returns the height of the Span.
//
// If a fixed height has been set and the display mode is `block`, we use the
// fixed height.
//
// If a fixed height has not been set or the display mode is not `block`, the
// height defaults to the number of lines of text content, or 1 if there is no
// text content.
func (s *Span) Height() int {
	display := s.Display()
	if display == types.DisplayBlock && s.FixedHeight() {
		return s.Sized.Height()
	}
	return strings.Count(s.Content(), "\n") + 1
}

// Draw renders the Span to the given screen in the specified bounding box.
func (s *Span) Draw(screen types.Screen, bounds types.Rectangle) {
	ctx := context.TODO()
	s.Element.Draw(screen, bounds)
	inner := s.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	content := render.AlignString(
		ctx, s.Content(), inner, s.Alignment(),
	)
	style := s.Style()
	content = style.Styled(content)
	ss := uv.NewStyledString(content)
	ws := s.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}
