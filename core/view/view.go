package view

import (
	"context"
	"fmt"

	"github.com/samber/lo"

	"github.com/jaypipes/gt/core/key"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/element/vdiv"
	"github.com/jaypipes/gt/types"
)

// New returns a new View instance with the given ID.
func New(
	ctx context.Context,
	id string,
) *View {
	d := vdiv.New(ctx, element.WithID(id))
	v := &View{
		VDiv: *d,
	}
	return v
}

// View is the virtual representation of the tree of elements that will be
// rendered to a Screen.
type View struct {
	vdiv.VDiv
	// currentViewKey is the key combination that should trigger setting
	// this View as the current View in the Application.
	currentViewKey types.Key
	// keyMap stores the View's map of key press combinations to callbacks.
	keyMap types.KeyMap
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

// SetCurrentViewKey sets the key combination that should trigger setting this
// View as the current View in the Application.
//
// The keypress combination can be a string -- e.g. "Ctrl+C", "Esc" -- or a
// [tcell.Key] code -- e.g. tcell.KeyCtrlC, tcell.KeyEscape.
func (v *View) SetCurrentViewKey(subject any) {
	v.currentViewKey = key.New(subject)
}

// WithCurrentViewKey sets the key combination that should trigger setting this
// View as the current View in the Application and returns the View.
//
// The keypress combination can be a string -- e.g. "Ctrl+C", "Esc" -- or a
// [tcell.Key] code -- e.g. tcell.KeyCtrlC, KeyEscape.
func (v *View) WithCurrentViewKey(kp string) *View {
	v.SetCurrentViewKey(kp)
	return v
}

// CurrentViewKey returns the key combination that triggers setting this View
// as the current View in the Application
func (v *View) CurrentViewKey() types.Key {
	return v.currentViewKey
}

// KeyMap returns the View's map of key press combination strings to
// callbacks that will execute when that key press combination is entered.
func (v *View) KeyMap() types.KeyMap {
	ctx := context.TODO()
	res := types.KeyMap{}

	// copy in our view-scoped key press callbacks
	for k, cb := range v.keyMap {
		res[k] = cb
	}
	viewKPs := lo.Keys(v.keyMap)

	// now add all child key maps
	children := v.Children()
	for _, child := range children {
		kp, ok := child.(types.HasKeyMap)
		if ok {
			kpMap := kp.KeyMap()
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
//
// The keypress combinations can be strings -- e.g. "Ctrl+C", "Esc" -- or
// [tcell.Key] codes -- e.g. tcell.KeyCtrlC, KeyEscape.
func (v *View) OnKeyPress(subject any, cb types.EventCallback) {
	k := key.New(subject)
	v.keyMap[k] = cb
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

// Draw ensures that any bounds placed on the View are applied to all the
// View's element tree and draws all elements in the DOM to the supplied
// Screen.
func (v *View) Draw(
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

	// Then recursively plot all content in the View.
	render.Plot(ctx, v, inner)

	// And finally draw all the content to the Screen.
	render.Render(ctx, v, screen)
}
