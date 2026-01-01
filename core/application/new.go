package application

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"
)

type ApplicationModifier func(a *Application)

// New returns a new Application.
func New(
	ctx context.Context,
	mods ...ApplicationModifier,
) *Application {
	a := &Application{}
	for _, mod := range mods {
		mod(a)
	}
	return a
}

// WithName sets the application's modional name, which by default sets the
// terminal screen's title.
func WithName(name string) ApplicationModifier {
	return func(a *Application) {
		a.name = name
	}
}

// WithRoot sets the default root renderable element for the Application.
func WithRoot(renderable uv.Drawable) ApplicationModifier {
	return func(a *Application) {
		a.root = renderable
	}
}
