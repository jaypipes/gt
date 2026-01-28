package view

import (
	"context"
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"
	uvscreen "github.com/charmbracelet/ultraviolet/screen"

	"github.com/jaypipes/gt/core/box"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// New returns a new View instance.
func New(ctx context.Context, id string) *View {
	b := box.New(ctx)
	b.SetID(id)
	return &View{Box: b}
}

// View is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type View struct {
	box.Box
	// bounds is the outer bounding box for the View.
	bounds types.Rectangle
	// restoreBuf is a pointer to a [uv.Buffer] that we save pre-rendered
	// content from the View to. This is restored when View.Restore is called.
	restoreBuf *uv.Buffer
	// currentViewKeyPress is the key combination that should trigger setting
	// this View as the current View in the Application.
	currentViewKeyPress string
}

// String returns a short string representation of the View.
func (v *View) String() string {
	return fmt.Sprintf("<view id=%s>", v.ID())
}

// WithBounds sets the View's outer bounding box and returns the View.
func (v *View) WithBounds(bounds types.Rectangle) *View {
	v.Box.SetBounds(bounds)
	return v
}

// SetCurrentViewKeyPress sets the key combination that should trigger setting this
// View as the current View in the Application.
func (v *View) SetCurrentViewKeyPress(key string) {
	v.currentViewKeyPress = key
}

// SetCurrentViewKeyPress sets the key combination that should trigger setting this
// View as the current View in the Application and returns the View.
func (v *View) WithCurrentViewKeyPress(key string) *View {
	v.SetCurrentViewKeyPress(key)
	return v
}

// CurrentViewKeyPress returns the key combination that triggers setting this
// View as the current View in the Application
func (v *View) CurrentViewKeyPress() string {
	return v.currentViewKeyPress
}

// SetContent sets the thing that will be rendered in the View.
func (v *View) SetContent(content types.Plottable) {
	v.Box.RemoveAllChildren()
	v.Box.AppendChild(content)
}

// WithContent sets the thing that will be rendered in the View and returns the
// View.
func (v *View) WithContent(content types.Plottable) *View {
	v.SetContent(content)
	return v
}

// AppendContent adds a child Element to the View's content and returns the
// View.
func (v *View) AppendContent(content types.Plottable) *View {
	v.Box.AppendChild(content)
	return v
}

// Restore redraws any previously-saved content from the View to the supplied
// [uv.Terminal].
func (v *View) Restore(
	ctx context.Context,
	screen uv.Screen,
) {
	if v.restoreBuf != nil {
		bounds := v.Bounds()
		v.restoreBuf.Draw(screen, bounds)
	}
}

// Save saves any pre-rendered content from the View's root Element's bounds to
// a buffer that can be quickly restored when View.Restore is called.
func (v *View) Save(
	ctx context.Context,
	screen uv.Screen,
) {
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
	bounds := v.Bounds()
	inner := v.InnerBounds()
	gtlog.Debug(
		ctx, "View.Render[%s]: bounds=%s inner_bounds=%s",
		v.ID(), bounds, inner,
	)

	// Allow any components to dynamically create renderable content.
	render.Build(ctx, v)

	// clear the outer bounds before rendering the DOM rooted at the root
	// Element.
	render.Clear(screen, bounds)

	// Then recursively plot all content in the View.
	render.Plot(ctx, v, inner)

	// And finally draw all the content to the Screen.
	render.Draw(ctx, v, screen)
}
