package view

import (
	"context"
	"fmt"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element/vdiv"
	"github.com/jaypipes/gt/types"
)

// View is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type View struct {
	vdiv.VDiv

	// activeKey is the key combination that should trigger setting this View
	// as the active View in the Application.
	activeKey types.Key

	// keyShortcuts stores the View's set of key shortcuts.
	keyShortcuts []types.KeyShortcut
}

// String returns a short string representation of the View.
func (v *View) String() string {
	return fmt.Sprintf("<view id=%s>", v.ID())
}

// WithID sets the View's identifier and returns the View.
func (v *View) WithID(id string) types.View {
	v.VDiv.SetID(id)
	return v
}

// WithBounds sets the View's outer bounding box and returns the View.
func (v *View) WithBounds(bounds types.Rectangle) types.View {
	v.SetBounds(bounds)
	return v
}

// SetContent sets the thing that will be rendered in the View.
func (v *View) SetContent(content types.Node) {
	v.RemoveAllChildren()
	v.AppendChild(content)
}

// WithContent sets the thing that will be rendered in the View and returns the
// View.
func (v *View) WithContent(content types.Node) types.View {
	v.SetContent(content)
	return v
}

// AppendContent adds a child Element to the View's content and returns the
// View.
func (v *View) AppendContent(content types.Node) types.View {
	v.AppendChild(content)
	return v
}

// AtPoint returns the child element at the supplied position, or nil if no the
// position is out of the element's bounding box. We perform a depth-first
// search of the child nodes since we do not allow overlapping boxes therefore
// the first matched leaf node is our match.
func (v *View) AtPoint(pos types.Point) types.Node {
	return v.VDiv.AtPoint(pos)
}

// Draw ensures that any bounds placed on the View are applied to all the
// View's element tree and draws all elements in the DOM to the supplied
// Screen.
func (v *View) Draw(
	ctx context.Context,
	h types.ScreenHandler,
) {
	bounds := v.Bounds()
	inner := v.InnerBounds()
	gtlog.Debug(
		ctx, "View.Render[%s]: bounds=%s inner_bounds=%s",
		v.ID(), bounds, inner,
	)

	// Allow any components to dynamically create renderable content.
	render.Build(ctx, v)

	// Then recursively plot all content in the View.
	render.Plot(ctx, v, inner)

	// And finally draw all the content to the Screen.
	render.Render(ctx, v, h)
}
