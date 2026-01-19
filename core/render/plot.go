package render

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Plot calculates the outer bounding box and positioning coordinates for the
// Document Object Model (DOM) rooted at the supplied element, recursing
// through the tree in a depth-first fashion.
func Plot(
	ctx context.Context,
	e types.Element,
) {
	e.Plot(ctx)
	for _, child := range e.Children() {
		Plot(ctx, child)
	}
}
