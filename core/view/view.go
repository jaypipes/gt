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
	// By default, a View's content is a div that consumes the entire area of
	// the View.
	container := div.New(ctx, "")
	container.SetID(fmt.Sprintf("view-%s-container", id))
	container.SetHeight(core.Percent(100))
	container.SetWidth(core.Percent(100))
	return &View{
		id:      id,
		content: container,
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
	// content is the thing that will be rendered in the View.
	content types.Element
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
	return fmt.Sprintf("<view id=%s>", v.id)
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

// SetContent sets the thing that will be rendered in the View.
func (v *View) SetContent(content types.Element) *View {
	v.content = content
	return v
}

// AppendContent adds a child Element to the View's content.
func (v *View) AppendContent(content types.Element) *View {
	if v.content == nil {
		return v.SetContent(content)
	}
	v.content.AppendChild(content)
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
	content := v.content
	if content == nil {
		return
	}
	gtlog.Debug(ctx, "View.Render(%s)", v.id)
	content.SetBounds(v.InnerBounds())

	// clear the outer bounds before rendering the DOM rooted at the root
	// Element.
	render.Clear(screen, v.Bounds())

	v.DrawBorder(screen)

	content.Render(ctx, screen)
}
