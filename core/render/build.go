package render

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Build calls Build on the supplied Buildable.
func Build(
	ctx context.Context,
	b types.Buildable,
) {
	for _, child := range b.Children() {
		cb := child.(types.Buildable)
		Build(ctx, cb)
	}
	b.Build(ctx)
}
