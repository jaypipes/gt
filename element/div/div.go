package div

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/element"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
)

const (
	ElementClass = "gt.div"
)

// New returns a new Div instance containing the supplied raw string content.
func New(
	ctx context.Context,
	content string,
) *Div {
	e := element.New(ctx, ElementClass)
	d := &Div{Element: e}
	d.SetDisplay(types.DisplayBlock)
	d.SetContent(content)
	return d
}

// Div is an Element that uses the block display mode by default.
type Div struct {
	*element.Element
	core.Contented
}

// Draw renders the Div to the given screen in the specified bounding box.
func (d *Div) Draw(screen types.Screen, bounds types.Rectangle) {
	d.Element.Draw(screen, bounds)
	inner := d.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	ss := uv.NewStyledString(d.Content())
	ws := d.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}
