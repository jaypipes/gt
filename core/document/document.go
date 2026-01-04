package document

import (
	"context"

	"github.com/jaypipes/gt/core/element"
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
	return &Document{Element: *e}
}

// Document is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type Document struct {
	element.Element
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
	// Propogate the outer bounds even if the user called SetBounds() before
	// SetRoot()...
	bounds := d.Bounds()
	if !bounds.Empty() {
		d.SetBounds(bounds)
	}
}

// Render ensures that any bounds placed on the Document are applied to all the
// Document's element tree.
func (d *Document) Render(
	ctx context.Context,
	screen types.Screen,
) {
	// If the document has had no bounds set, adopt the screen's max width and
	// height.
	bounds := d.Bounds()
	if bounds.Empty() {
		bounds = screen.Bounds()
		d.SetBounds(bounds)
	}
	d.Element.Render(ctx, screen)
}

var _ types.Document = (*Document)(nil)
var _ types.Renderable = (*Document)(nil)
