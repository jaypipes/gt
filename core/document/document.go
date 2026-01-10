package document

import (
	"context"

	"github.com/jaypipes/gt/core/element"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
)

const (
	ElementClass = "gt.document"
	ElementID    = "#document"
)

// New returns a new Document instance.
func New(ctx context.Context) *Document {
	e := element.New(ctx, ElementClass)
	e.SetID(ElementID)
	return &Document{Element: e}
}

// Document is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type Document struct {
	*element.Element
}

// SetSize sets the width and height of the Document. This has the effect of
// overriding the Document's Bounds.
func (d *Document) SetSize(width, height int) {
	d.Element.SetSize(width, height)
	bounds := types.Rectangle{}
	bounds.Min = d.TL()
	bounds.Max.X = bounds.Min.X + width
	bounds.Max.Y = bounds.Min.Y + height
	d.SetBounds(bounds)
}

// Anchor sets the top-left anchor point for the Document. This has the effect
// of providing an offset for the Document's Bounds if the Document has a fixed
// width and height (i.e. the Document.SetSize method has been called)
func (d *Document) Anchor(pt types.Point) {
	size := d.Size()
	d.Element.Anchor(pt)
	bounds := d.Bounds()
	if !size.Empty() {
		bounds.Max.X = bounds.Min.X + size.W
		bounds.Max.Y = bounds.Min.Y + size.H
	}
	d.SetBounds(bounds)
}

// ElementByID returns the Element with the specified ID or nil if the
// Document contains no Element with that identifier.
func (d *Document) ElementByID(id string) types.Element {
	return nil
}

// ElementsByClass returns all Elements having the specified type/class.
func (d *Document) ElementsByClass(class string) []types.Element {
	res := []types.Element{}
	return res
}

// SetRoot sets the Document's top-level renderable element. This is just
// syntactic sugar over the underlying Node.PushChild() method.
func (d *Document) SetRoot(el types.Element) {
	d.RemoveAllChildren()
	d.PushChild(el)
}

// Render ensures that any bounds placed on the Document are applied to all the
// Document's element tree.
func (d *Document) Render(
	ctx context.Context,
	screen types.Screen,
) {
	// If the document has had no bounds set, adopt the screen's max width and
	// height.
	outer := d.Bounds()
	if outer.Empty() {
		screenBounds := screen.Bounds()
		gtlog.Debug(
			ctx,
			"Document.Render: outer=%s screen=%s",
			outer, screenBounds,
		)
		outer = screenBounds
		d.SetBounds(outer)
	}
	// calculate the position and sizing for each element in the DOM.
	render.Plot(ctx, d)

	// clear the outer bounds before rendering the DOM.
	render.Clear(screen, outer)

	// draw each element in the DOM
	render.Draw(ctx, d, screen)
}

var _ types.Document = (*Document)(nil)
