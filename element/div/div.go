package div

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element/base"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.div"
)

// New returns a new Div instance containing the supplied raw string content.
func New(
	ctx context.Context,
	content string,
) *Div {
	b := base.New(ctx, ElementClass)
	d := &Div{Base: b}
	d.SetDisplay(types.DisplayBlock)
	d.SetTextContent(content)
	return d
}

// Div is an Element that uses the block display mode by default.
type Div struct {
	base.Base
}

// Render renders the Div to the given screen in the specified bounding box.
func (d *Div) Render(ctx context.Context, screen types.Screen) {
	gtlog.Debug(ctx, "div.Div.Render[%s]", d)
	bounds := d.Bounds()
	d.Base.Render(ctx, screen)
	inner := d.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	content := render.AlignString(
		ctx, d.TextContent(), inner, d.Alignment(),
	)
	style := d.Style()
	content = style.Styled(content)
	ss := uv.NewStyledString(content)
	ws := d.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}
