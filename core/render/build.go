package render

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Build calls Build on the supplied Buildable.
func Build(
	ctx context.Context,
	n types.Node,
) {
	for _, child := range n.Children() {
		Build(ctx, child)
	}
	b, ok := n.(types.Buildable)
	if ok {
		b.Build(ctx)
	}
}
