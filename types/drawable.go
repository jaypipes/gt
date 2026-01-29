package types

import (
	uv "github.com/charmbracelet/ultraviolet"
)

// Drawable wraps the [uv.Drawable] interface allowing context to be passed.
type Drawable interface {
	uv.Drawable
	Node
	Bounded
}
