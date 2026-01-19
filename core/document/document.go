package document

import (
	"context"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element/base"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.document"
	ElementID    = "#document"
)

// New returns a new Document instance.
func New(ctx context.Context) *Document {
	b := base.New(ctx, ElementClass)
	d := &Document{Base: b}
	d.SetID(ElementID)
	return d
}

// Document is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type Document struct {
	base.Base
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
func (d *Document) SetRoot(node types.Element) {
	d.RemoveAllChildren()
	d.PushChild(node)
}

// Render ensures that any bounds placed on the Document are applied to all the
// Document's element tree and draws all elements in the DOM to the supplied
// Screen.
func (d *Document) Render(
	ctx context.Context,
	screen types.Screen,
) {
	// If the document has had no bounds set, adopt the screen's max width and
	// height.
	bounds := d.Bounds()
	if bounds.Empty() {
		screenBounds := screen.Bounds()
		gtlog.Debug(
			ctx,
			"Document.Render: setting document bounds to screen bounds %s",
			screenBounds,
		)
		bounds = screenBounds
		d.SetBounds(bounds)
	}

	// calculate the position and sizing for each element in the DOM.
	for _, child := range d.Children() {
		child.Plot(ctx)
	}

	// clear the outer bounds before rendering the DOM.
	render.Clear(screen, bounds)

	// draw each element in the DOM
	for _, child := range d.Children() {
		child.Render(ctx, screen)
	}
}
