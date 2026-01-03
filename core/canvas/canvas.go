package canvas

import (
	"context"

	"github.com/jaypipes/gt/core/box"
	"github.com/jaypipes/gt/core/types"
)

// New returns a new Canvas instance.
func New() *Canvas {
	return &Canvas{}
}

// Canvas is used to manage displaying to a Screen.
type Canvas struct {
	box.Box
	// root is the top-level renderable element in the Canvas. This is a
	// Box that consumes the entire screen real-estate. Use WithRoot to
	// override this default renderable element.
	root types.Renderable
}

// SetRoot sets the Canvas's top-level renderable element.
func (c *Canvas) SetRoot(root types.Renderable) {
	c.root = root
}

// Draw draws the Canvas to the supplied screen.
func (c *Canvas) Draw(
	screen types.Screen,
	bounds types.Rectangle,
) {
	c.Box.Draw(screen, bounds)
	if c.root != nil {
		c.root.Draw(screen, bounds)
	}
}

// Render positions Canvas's elements in the supplied Screen and then draws to
// the Screen.
func (c *Canvas) Render(
	ctx context.Context,
	screen types.Screen,
) {
	bounds := c.Bounds()
	if bounds.Empty() {
		bounds = screen.Bounds()
		c.SetBounds(bounds)
	}

	// Position the root element within the inner bounding box. The root
	// element is responsible for propogating this positioning change to any
	// child elements.
	inner := c.InnerBounds()
	root := c.root
	if root != nil {
		if root.Relative() {
			root.SetRelativePosition(inner.Min.X, inner.Min.Y)
		}
	}
	c.Draw(screen, bounds)
}

var _ types.Renderable = (*Canvas)(nil)
