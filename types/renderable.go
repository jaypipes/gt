package types

import "context"

// Renderable things can render their contents to a [Screen]
type Renderable interface {
	// Render draws the contents of the Renderable to the supplied
	// ScreenHandler.
	Render(context.Context, ScreenHandler)
}
