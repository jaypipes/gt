package types

import (
	"context"
)

// Buildable allows components to generate their content dynamically.
type Buildable interface {
	Node
	// Build generates any dynamic content in the Buildable before the content
	// is plotted.
	Build(context.Context)
}
