package hr

import (
	"context"
	"strings"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass   = "gt.hr"
	thinHorizontal = "─"
)

// New returns a new HR instance with the given options.
func New(
	ctx context.Context,
	opts ...types.ElementWithOption,
) *HR {
	e := element.New(ctx, ElementClass)
	h := &HR{Element: e}
	// An HR defaults to starting on a new line and is one line high and is
	// centered within the width of the parent container.
	h.SetDisplay(types.DisplayBlock)
	h.SetHeight(core.Fixed(1))
	h.SetAlignment(types.AlignmentCenter)
	h.SetWhitespace(types.WhitespaceWrapNever)
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// HR is an Element that renders a single horizontal rule line on the screen.
//
// An HR defaults to the width of the parent container. An HR's height defaults
// to 1. By default, the HR is centered within the parent container.
type HR struct {
	element.Element
}

// Render implements the types.Renderable interface
func (h *HR) Render(ctx context.Context, screen types.Screen) {
	bounds := h.Bounds()
	gtlog.Debug(ctx, "HR.Render[%s]: bounds=%s", h.Tag(), bounds)
	numCellsWide := h.Width()
	inner := h.InnerBounds()
	if numCellsWide == 0 {
		numCellsWide = types.Dimension(inner.Dx())
	}
	line := strings.Repeat(thinHorizontal, int(numCellsWide))
	line = render.Align(
		ctx, line, inner, h.Alignment(), h.Whitespace(),
	)
	s := h.Style()
	startX := inner.Min.X
	startY := inner.Min.Y
	screen.PutStrStyled(startX, startY, line, style.TCell(s))
}
