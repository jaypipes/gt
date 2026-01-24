package span

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element/base"
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
	b := base.New(ctx, ElementClass)
	s := &Span{Base: b}
	s.SetDisplay(types.DisplayInline)
	s.SetTextContent(content)
	return s
}

// Span is an Element that uses the inline display mode by default.
type Span struct {
	base.Base
}

// SetSize sets the fixed width and height of the Span and also sets the
// display mode to `inline-block`.
func (s *Span) SetSize(constraint types.SizeConstraint) types.Element {
	s.Base.SetSize(constraint)
	s.SetDisplay(types.DisplayInlineBlock)
	return s
}

// SetWidth sets the fixed width of the Span and also sets the display mode to
// `inline-block`.
func (s *Span) SetWidth(constraint types.DimensionConstraint) types.Element {
	s.Base.SetWidth(constraint)
	s.SetDisplay(types.DisplayInlineBlock)
	return s
}

// SetHeight sets the fixed height of the Span and also sets the display mode
// to `inline-block`.
func (s *Span) SetHeight(constraint types.DimensionConstraint) types.Element {
	s.Base.SetHeight(constraint)
	s.SetDisplay(types.DisplayInlineBlock)
	return s
}

// Render draws the Span to the given screen.
func (s *Span) Render(ctx context.Context, screen types.Screen) {
	s.Base.Render(ctx, screen)
	gtlog.Debug(ctx, "span.Span.Render[%s]", s)
	bounds := s.Bounds()
	inner := s.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	content := render.AlignString(
		ctx, s.TextContent(), inner, s.Alignment(),
	)
	style := s.Style()
	content = style.Styled(content)
	ss := uv.NewStyledString(content)
	ws := s.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
	for _, child := range s.Children() {
		child.Render(ctx, screen)
	}
}
