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
	for _, opt := range opts {
		opt(h)
	}
	// An HR always starts on a new line and is one line high.
	h.SetDisplay(types.DisplayBlock)
	h.SetHeight(core.Fixed(1))
	h.SetAlignment(types.AlignmentCenter)
	h.SetWhitespace(types.WhitespaceWrapNever)
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
	defStyle := h.Style()
	startX := inner.Min.X
	startY := inner.Min.Y
	for x := range line {
		screen.Put(startX+x, startY, string(line[x]), style.TCell(defStyle))
	}
}
