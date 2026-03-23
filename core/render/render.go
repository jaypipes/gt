package render

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Render calls Render on the supplied Renderable.
func Render(
	ctx context.Context,
	n types.Node,
	h types.ScreenHandler,
) {
	r, ok := n.(types.Renderable)
	if !ok {
		return
	}
	r.Render(ctx, h)
	for _, child := range n.Children() {
		Render(ctx, child, h)
	}
}
