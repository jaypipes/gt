package types

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"
)

// Renderable things can be drawn to a screen and expose information about
// their position on the screen.
type Renderable interface {
	uv.Drawable
	Plotted
	// Render wraps the uv.Drawable.Draw interface and passes in the context
	// that can be used by the tree of Renderables being drawn to the screen.
	Render(context.Context, uv.Screen)
}
