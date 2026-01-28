package types

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"
)

// Drawable wraps the [uv.Drawable] interface allowing context to be passed.
type Drawable interface {
	uv.Drawable
	Node

	// DrawWithContext wraps the [uv.Drawable.Draw] interface method with a
	// context and always calls [uv.Drawable.Draw] with the Drawables's
	// pre-plotted bounds.
	DrawWithContext(context.Context, Screen)
}
