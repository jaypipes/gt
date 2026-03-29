package button

import (
	"context"
	"strings"

	fevent "github.com/jaypipes/gt/core/event/focus"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.button"
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

// Button is an Element that renders a clickable Button to the terminal screen.
type Button struct {
	element.Element
}

// Button elements can get the focus but when clicked, they release the
// focus...
func (b *Button) MouseClick(ctx context.Context, ev types.MouseClickEvent) {
	b.Element.MouseClick(ctx, ev)
	b.Focus(ctx, fevent.New(
		fevent.WithFocused(false),
		fevent.WithSource(ev.Source()),
	))
}

// Render implements the types.Renderable interface
func (b *Button) Render(ctx context.Context, h types.ScreenHandler) {
	bounds := b.Bounds()
	gtlog.Debug(ctx, "Button.Render[%s]: bounds=%s", b.Tag(), bounds)

	b.RenderBox(ctx, h)

	screen := h.Screen()

	content := b.TextContent()
	s := b.Style()
	inner := b.InnerBounds()
	lines := strings.Split(content, "\n")
	startX := inner.Min.X
	startY := inner.Min.Y
	for y, line := range lines {
		screen.PutStrStyled(startX, startY+y, line, style.TCell(s))
	}
}
