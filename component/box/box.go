package box

import (
	"context"

	"github.com/jaypipes/gt/core/element"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
)

const (
	ElementClass = "gt.box"
)

// New returns a new instance of a Box.
func New(ctx context.Context, id string) *Box {
	e := element.New(ctx, ElementClass)
	e.SetID(id)
	return &Box{Element: *e}
}

// Box is the simplest component. It's just a box.
type Box struct {
	element.Element
}

// Draw implements the uv.Renderable interface
func (b *Box) Draw(screen types.Screen, area types.Rectangle) {
	ctx := context.TODO()
	b.Debug(ctx, "Box(%s).Draw: bounding box %s\n", b.ID(), area)
	// determine the overlapping bounding element and clear its cells before
	// rendering the element.
	bb := render.Overlapping(area, b.Bounds())
	render.Clear(screen, bb)

	// If we have a border, draw it around the outer bounding box.
	border := b.Border()
	if border != nil {
		b.Debug(ctx, "Box(%s).Draw: drawing border around %s\n", b.ID(), area)
		border.Draw(screen, bb)
	}
	b.Element.Draw(screen, area)
}

var _ types.Renderable = (*Box)(nil)
