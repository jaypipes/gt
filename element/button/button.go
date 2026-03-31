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
