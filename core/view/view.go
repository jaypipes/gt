package view

import (
	"context"
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"
	uvscreen "github.com/charmbracelet/ultraviolet/screen"
	"github.com/samber/lo"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element/vdiv"
	"github.com/jaypipes/gt/types"
)

// New returns a new View instance.
func New(ctx context.Context, id string) *View {
	d := vdiv.New(ctx, "")
	d.SetID(id)
	v := &View{
		VDiv: *d,
	}
	return v
}

// View is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type View struct {
	vdiv.VDiv
	// restoreBuf is a pointer to a [uv.Buffer] that we save pre-rendered
	// content from the View to. This is restored when View.Restore is called.
	restoreBuf *uv.Buffer
	// currentViewKeyPress is the key combination that should trigger setting
	// this View as the current View in the Application.
	currentViewKeyPress string
	// keyPressMap stores the View's map of key press combinations to
	// callbacks.
	keyPressMap types.KeyPressMap
}

// String returns a short string representation of the View.
func (v *View) String() string {
	return fmt.Sprintf("<view id=%s>", v.ID())
}

// WithBounds sets the View's outer bounding box and returns the View.
func (v *View) WithBounds(bounds types.Rectangle) *View {
	v.SetBounds(bounds)
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

// KeyPressMap returns the View's map of key press combination strings to
// callbacks that will execute when that key press combination is entered.
func (v *View) KeyPressMap() types.KeyPressMap {
	ctx := context.TODO()
	res := types.KeyPressMap{}

	// copy in our view-scoped key press callbacks
	for k, cb := range v.keyPressMap {
		res[k] = cb
	}
	viewKPs := lo.Keys(v.keyPressMap)

	// now add all child key press maps
	children := v.Children()
	for _, child := range children {
		kp, ok := child.(types.HasKeyPressMap)
		if ok {
			kpMap := kp.KeyPressMap()
			for k, cb := range kpMap {
				if lo.Contains(viewKPs, k) {
					gtlog.Warn(
						ctx,
						"key press combination %q for view child %q "+
							"shadows view-level key press combination",
						k, child.NodeID(),
					)
				}
				_, exists := res[k]
				if exists {
					gtlog.Warn(
						ctx,
						"key press combination %q for view child %q "+
							"shadows prior key press combination",
						k, child.NodeID(),
					)
				}
				res[k] = cb
			}
		}
	}
	return res
}

// OnKeyPress registers an View-level callback to execute upon a key press
// combination.
func (v *View) OnKeyPress(key string, cb types.EventCallback) {
	v.keyPressMap[key] = cb
}

// SetContent sets the thing that will be rendered in the View.
func (v *View) SetContent(content types.Node) {
	v.RemoveAllChildren()
	v.AppendChild(content)
}

// WithContent sets the thing that will be rendered in the View and returns the
// View.
func (v *View) WithContent(content types.Node) *View {
	v.SetContent(content)
	return v
}

// AppendContent adds a child Element to the View's content and returns the
// View.
func (v *View) AppendContent(content types.Node) *View {
	v.AppendChild(content)
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
