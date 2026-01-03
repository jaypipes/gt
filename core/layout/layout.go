package layout

import (
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/element"
	"github.com/jaypipes/gt/core/types"
)

type LayoutModifier func(*Layout)

// WithBounds sets the outer rectangle boundaries of the Layout. By default, a
// Layout will consume the entire width and height of the available Screen.
func WithBounds(x, y, width, height int) LayoutModifier {
	return func(e *Layout) {
		e.SetBounds(uv.Rect(x, y, width, height))
	}
}

// WithBorder sets a Border style for the Layout. This border is the outermost
// border of the Layout's Element. This is NOT the border style of the Layout's
// Panes.
func WithBorder(border uv.Border) LayoutModifier {
	return func(e *Layout) {
		e.SetBorder(border)
	}
}

// New returns a new Layout instance.
func New(opts ...LayoutModifier) *Layout {
	g := &Layout{}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

// Layout represents an abstraction of multiple Elements arranged in a set of
// rows and columns.
type Layout struct {
	element.Element
}

type SplitDirection int

const (
	Horizontal SplitDirection = iota
	Vertical
)

// Split is a section of a Layout. It contains *instructions* for how the
// Layout should create the section.
type Split struct {
	r         types.Renderable
	direction SplitDirection
	con       uv.Constraint
}

// SplitHorizontal splits the Layout's last child element, adding a new
// horizontal pane containing the supplied Renderable. If the Layout was empty,
// the supplied Renderable is simply added as the child element.
func (e *Layout) SplitHorizontal(r types.Renderable, con uv.Constraint) {
	lastChild := e.PopChild()
	if lastChild == nil {
		e.PushChild(r)
		return
	}
	left, right := uv.SplitHorizontal(lastChild.Bounds(), con)
	fmt.Println("left", left, "right", right)
	lastChild.SetBounds(left)
	e.PushChild(lastChild)
	r.SetBounds(right)
	e.PushChild(r)
}

// SplitVertical splits the Layout's last child element, adding a new
// vertical pane containing the supplied Renderable. If the Layout was empty,
// the supplied Renderable is simply added as the child element.
func (e *Layout) SplitVertical(r types.Renderable, con uv.Constraint) {
	lastChild := e.PopChild()
	if lastChild == nil {
		e.PushChild(r)
		return
	}
	top, bottom := uv.SplitVertical(lastChild.Bounds(), con)
	fmt.Println("top", top, "bottom", bottom)
	lastChild.SetBounds(top)
	e.PushChild(lastChild)
	r.SetBounds(bottom)
	e.PushChild(r)
}
