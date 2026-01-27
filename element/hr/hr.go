package hr

import (
	"context"
	"strings"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element/base"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass   = "gt.hr"
	thinHorizontal = "─"
)

// New returns a new HR instance.
func New(ctx context.Context) *HR {
	b := base.New(ctx, ElementClass)
	h := &HR{Base: b}
	h.SetHeight(core.Fixed(1))
	// An HR always starts on a new line.
	h.SetDisplay(types.DisplayBlock)
	h.SetAlignment(types.AlignmentCenter)
	return h
}

// HR is an Element that renders a single horizontal rule line on the screen.
//
// An HR defaults to the width of the parent container. An HR's height defaults
// to 1. By default, the HR is centered within the parent container.
type HR struct {
	base.Base
}

// Render draws the HR to the supplied Screen.
func (h *HR) Render(ctx context.Context, screen types.Screen) {
	render.Plot(ctx, h)
	numCellsWide := h.Width()
	inner := h.InnerBounds()
	if numCellsWide == 0 {
		numCellsWide = types.Dimension(inner.Dx())
	}
	line := strings.Repeat(thinHorizontal, int(numCellsWide))
	line = render.AlignString(
		ctx, line, inner, h.Alignment(),
	)
	style := h.Style()
	line = style.Styled(line)
	ss := uv.NewStyledString(line)
	ss.Draw(screen, inner)
}
