package box

import (
	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/render"
)

type Option func(*Box)

// WithBounds sets the rectangle boundaries of the Box.
func WithBounds(x, y, width, height int) Option {
	return func(b *Box) {
		b.rect = uv.Rect(x, y, width, height)
	}
}

// WithBorder sets a Border style for the Box.
func WithBorder(border uv.Border) Option {
	return func(b *Box) {
		b.border = &border
	}
}

// New returns a new Box instance.
func New(opts ...Option) *Box {
	b := &Box{}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// Box is a [uv.Drawable] that renders a styled box on the screen. All
// renderable elements and components will inherit from Box.
type Box struct {
	rect   uv.Rectangle
	border *uv.Border
}

// Height returns the height of the Box
func (b *Box) Height() int {
	return b.rect.Max.Y - b.rect.Min.Y
}

// Width returns the width of the Box
func (b *Box) Width() int {
	return b.rect.Max.X - b.rect.Min.X
}

// Bounds returns the bounding box for the Box
func (b *Box) Bounds() uv.Rectangle {
	return b.rect
}

// Draw renders the Box to the given buffer at the specified area.
func (b *Box) Draw(buf uv.Screen, area uv.Rectangle) {
	// determine the overlapping bounding box and clear its cells before
	// rendering the box.
	bb := render.Overlapping(area, b.rect)
	render.Clear(buf, bb)

	// If we have a border, draw it around the Box's bounding box.
	if b.border != nil {
		b.border.Draw(buf, bb)
	}
}

var _ uv.Drawable = (*Box)(nil)
