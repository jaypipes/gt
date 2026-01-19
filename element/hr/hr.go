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
	return h
}

// HR is an Element that renders a single horizontal rule line on the screen.
//
// An HR defaults to the width of the parent container. An HR's height defaults
// to 1. By default, the HR is centered within the parent container.
type HR struct {
	base.Base
}

// Display always returns DisplayBlock since an HR always starts on a new line.
func (h *HR) Display() types.Display {
	return types.DisplayBlock
}

// InnerBounds returns the HR's inner bounding box. The bounding box within
// which a horizontal rule is aligned is always as wide as their parent
// container and the fixed height of the horizontal rule.
func (h *HR) InnerBounds() types.Rectangle {
	parent := h.Parent()
	inner := h.Base.InnerBounds()
	if parent != nil {
		parentInner := parent.InnerBounds()
		inner.Min.X = parentInner.Min.X
		inner.Max.X = parentInner.Max.X
	}
	return inner
}

// Draw renders the HR to the given screen in the specified bounding box.
func (h *HR) Draw(screen types.Screen, bounds types.Rectangle) {
	ctx := context.TODO()

	numCellsWide := h.Width()
	inner := h.InnerBounds()
	if numCellsWide == 0 {
		numCellsWide = types.Dimension(inner.Dx())
	}

	numLinesHigh := h.Height()
	if numLinesHigh == 0 {

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
