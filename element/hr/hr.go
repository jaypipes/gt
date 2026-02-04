package hr

import (
	"context"
	"strings"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/render"
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

// Draw draws the HR to the supplied Screen.
func (h *HR) Draw(screen types.Screen, bounds types.Rectangle) {
	ctx := context.TODO()
	numCellsWide := h.Width()
	inner := h.InnerBounds()
	if numCellsWide == 0 {
		numCellsWide = types.Dimension(inner.Dx())
	}
	line := strings.Repeat(thinHorizontal, int(numCellsWide))
	line = render.Align(
		ctx, line, inner, h.Alignment(), h.Whitespace(),
	)
	style := h.Style()
	line = style.Styled(line)
	ss := uv.NewStyledString(line)
	ss.Draw(screen, inner)
}
