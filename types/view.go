package types

import (
	"context"
	"fmt"
)

// View describes an overlay on the terminal screen. An Application can have
// many Views but only one active View. The active View is the one that is
// displayed when the Application draws to the terminal screen.
type View interface {
	fmt.Stringer
	Identifiable
	KeyPressEventHandler
	Plottable

	// WithID sets the View's unique identifier and returns the View.
	WithID(string) View

	// ActiveKey returns the key combination that triggers setting this View as
	// the active View in the Application
	ActiveKey() Key
	// SetActiveKey sets the key combination that should trigger setting this
	// View as the active (displayed) View.
	SetActiveKey(any)
	// WithActiveKey sets the key combination that should trigger setting this
	// View as the active View in the Application and returns the View.
	//
	// The keypress combination can be a string ("Tab"), a [tcell.Key] code
	// (tcell.KeyTab), a [types.KeyCode] value (types.KeyCodeTab) or a
	// [types.Key] object (core.key.KeyTab)
	WithActiveKey(any) View

	// NextFocusable returns the next focusable thing, or nil if there is no
	// next focusable thing. The View's children will first be inspected and
	// then the next sibling View.
	NextFocusable(context.Context) FocusEventHandler

	// AtPoint returns the child element at the supplied position, or nil if no the
	// position is out of the element's bounding box. We perform a depth-first
	// search of the child nodes since we do not allow overlapping boxes therefore
	// the first matched leaf node is our match.
	AtPoint(Point) Node
	// SetBounds sets the View's outer bounding box.
	SetBounds(Rectangle)
	// WithBounds sets the View's outer bounding box and returns the View.
	WithBounds(Rectangle) View

	// Border returns the View's border.
	Border() Border
	// SetBorder sets the View's border.
	SetBorder(Border)
	// WithBorder sets the View's border and returns the View.
	WithBorder(Border) View

	// SetContent sets the thing that will be rendered in the View.
	SetContent(Node)
	// WithContent sets the thing that will be rendered in the View and returns
	// the View.
	WithContent(Node) View
	// AppendContent adds a child Element to the View's content and returns the
	// View.
	AppendContent(Node) View

	// Draw ensures that any bounds placed on the View are applied to all the
	// View's element tree and draws all elements in the DOM to the supplied
	// Screen.
	Draw(context.Context, ScreenHandler)
}

// ViewWithOption describes an optional varg parameter to [core.view.New] that
// modifies the returned View.
type ViewWithOption func(View)
