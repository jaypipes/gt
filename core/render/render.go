package render

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Render calls Render on the supplied Renderable.
func Render(
	ctx context.Context,
	n types.Node,
	screen types.Screen,
) {
	r, ok := n.(types.Renderable)
	if !ok {
		return
	}
	r.Render(ctx, screen)
	for _, child := range n.Children() {
		Render(ctx, child, screen)
	}
}
