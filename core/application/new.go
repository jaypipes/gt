package application

import uv "github.com/charmbracelet/ultraviolet"

type Option func(a *Application)

// New returns a new Application.
func New(
	opts ...Option,
) *Application {
	a := &Application{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// WithRoot sets the default root renderable element for the Application.
func WithRoot(renderable uv.Drawable) Option {
	return func(a *Application) {
		a.root = renderable
	}
}
