package view

import (
	"context"
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"
	uvscreen "github.com/charmbracelet/ultraviolet/screen"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element/div"
	"github.com/jaypipes/gt/types"
)

// New returns a new View instance.
func New(ctx context.Context, id string) *View {
	return &View{
		id: id,
	}
}

// View is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type View struct {
	core.Box
	// id is the unique identifier for the View.
	id string
	// bounds is the outer bounding box for the View.
	bounds types.Rectangle
	// root is the root Element for the View.
	root types.Element
	// restoreBuf is a pointer to a [uv.Buffer] that we save pre-rendered
	// content from the View to. This is restored when View.Restore is called.
	restoreBuf *uv.Buffer
	// currentViewKeyPress is the key combination that should trigger setting
	// this View as the current View in the Application.
	currentViewKeyPress string
}

// ID returns the View's unique identifier.
func (v *View) ID() string {
	return v.id
}

// String returns a short string representation of the View.
func (v *View) String() string {
	root := "none"
	if v.root != nil {
		root = v.root.Tag()
	}
	return fmt.Sprintf("<view id=%s root=%s>", v.id, root)
}

// SetRoot sets the View's top-level renderable Element.
func (v *View) SetRoot(root types.Element) *View {
	v.root = root
	return v
}

// Root returns the View's top-level renderable Element.
func (v *View) Root() types.Element {
	return v.root
}

// SetBounds sets the View's outer bounding box.
func (v *View) SetBounds(bounds types.Rectangle) *View {
	v.Box.SetBounds(bounds)
	return v
}

// SetCurrentViewKeyPress sets the key combination that should trigger setting this
// View as the current View in the Application.
func (v *View) SetCurrentViewKeyPress(key string) *View {
	v.currentViewKeyPress = key
	return v
}

// CurrentViewKeyPress returns the key combination that triggers setting this
// View as the current View in the Application
func (v *View) CurrentViewKeyPress() string {
	return v.currentViewKeyPress
}

// AppendElement adds a new Element to the supplied View's root Element.
func (v *View) AppendElement(el types.Element) *View {
	if v.root == nil {
		v.root = div.New(context.TODO(), "")
	}
	v.root.AppendChild(el)
	return v
}

// Restore redraws any previously-saved content from the View to the supplied
// [uv.Terminal].
func (v *View) Restore(
	ctx context.Context,
	screen uv.Screen,
) {
	if v.restoreBuf != nil && v.root != nil {
		bounds := v.root.Bounds()
		v.restoreBuf.Draw(screen, bounds)
	}
}

// Save saves any pre-rendered content from the View's root Element's bounds to
// a buffer that can be quickly restored when View.Restore is called.
func (v *View) Save(
	ctx context.Context,
	screen uv.Screen,
) {
	if v.root == nil {
		return
	}
	bounds := v.Bounds()
	v.restoreBuf = uvscreen.CloneArea(screen, bounds)
}

// Render ensures that any bounds placed on the View are applied to all the
// View's element tree and draws all elements in the DOM to the supplied
// Screen.
func (v *View) Render(
	ctx context.Context,
	screen types.Screen,
) {
	gtlog.Debug(ctx, "View.Render(%s)", v.id)
	root := v.root
	if root == nil {
		return
	}
	root.SetBounds(v.InnerBounds())

	// calculate the position and sizing for each element in the DOM rooted at
	// the root Element.
	for _, child := range root.Children() {
		child.Plot(ctx)
	}

	// clear the outer bounds before rendering the DOM rooted at the root
	// Element.
	render.Clear(screen, v.Bounds())

	v.DrawBorder(screen)

	// draw each element in the DOM rooted at the root Element.
	for _, child := range root.Children() {
		child.Render(ctx, screen)
	}
}
