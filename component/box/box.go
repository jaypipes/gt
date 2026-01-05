package box

import (
	"context"

	"github.com/jaypipes/gt/core/element"
	gtlog "github.com/jaypipes/gt/core/log"
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
func (b *Box) Prerender(ctx context.Context, screen types.Screen, area types.Rectangle) {
	gtlog.Debug(ctx, "Box(%s).Prerender: bounding box %s\n", b.ID(), area)
	// Draw the border, if any, and clear the inner bounding box of this
	// Element.
	b.Element.Draw(screen, area)
}

var _ types.Renderable = (*Box)(nil)
