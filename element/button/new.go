package button

import (
	"context"

	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

// New returns a new Button instance with the given options.
func New(
	ctx context.Context,
	opts ...types.ElementWithOption,
) *Button {
	e := element.New(ctx, ElementClass)
	b := &Button{
		Element: e,
	}
	// Button defaults to top-left alignment and preserving the user's exact
	// input whitespacing and a square, thin-line border.
	b.SetDisplay(types.DisplayInlineBlock)
	b.SetAlignment(types.AlignmentTopLeft)
	b.SetWhitespace(types.WhitespacePreserve)
	b.SetMotif(DefaultMotif)
	// Button is an input element so should be able to receive the focus. Note
	// that the MouseClick action for a Button releases the focus immediately
	// after the mouse clicks on the button. This is done so that the hover
	// effect (which does not fire when the element has the focus) is restored
	// when the mouse clicks on the button.
	b.SetFocusable(true)
	for _, opt := range opts {
		opt(b)
	}
	return b
}
