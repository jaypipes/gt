package button

import (
	"context"
	"image/color"
	"strings"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt/core/border"
	fevent "github.com/jaypipes/gt/core/event/focus"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.button"
)

var (
	veryLightGrey, _       = colorful.Hex("#ececec")
	black, _               = colorful.Hex("#000000")
	DefaultBackgroundColor = color.Transparent
	DefaultForegroundColor = veryLightGrey
	DefaultStyle           = style.New(
		style.WithForegroundColor(
			DefaultForegroundColor,
		),
		style.WithBackgroundColor(
			DefaultBackgroundColor,
		),
	)
	DefaultBorder = border.Rounded().
			WithBackgroundColor(DefaultBackgroundColor).
			WithForegroundColor(DefaultForegroundColor)
	// Default hover is inverse of normal style.
	DefaultHoverBackgroundColor = veryLightGrey
	DefaultHoverForegroundColor = black
	DefaultHoverStyle           = style.New(
		style.WithForegroundColor(
			black,
		),
		style.WithBackgroundColor(
			DefaultHoverBackgroundColor,
		),
	)
	DefaultHoverBorder = border.InnerHalfBlock().
				WithBackgroundColor(color.Transparent).
				WithForegroundColor(DefaultHoverBackgroundColor)
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
	b.SetBorder(DefaultBorder)
	b.SetHoverBorder(DefaultHoverBorder)
	b.SetStyle(DefaultStyle)
	b.SetHoverStyle(DefaultHoverStyle)
	// Button is an input element so should be able to receive the focus.
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

	border := b.Border()
	b.Box.SetBorder(border)

	b.Box.Render(ctx, h)

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
