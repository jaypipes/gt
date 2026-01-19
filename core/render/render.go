package render

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Render draws the Document Object Model (DOM) rooted at the supplied element
// to the supplied screen, recursing through the tree in a depth-first fashion.
func Render(
	ctx context.Context,
	e types.Element,
	screen types.Screen,
) {
	e.Render(ctx, screen)
	for _, child := range e.Children() {
		Render(ctx, child, screen)
	}
}
