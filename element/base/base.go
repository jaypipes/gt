package base

import (
	"context"
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/box"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a [base.Base] with the specified type/class.
func New(ctx context.Context, class string) Base {
	b := box.New(ctx)
	return Base{
		Box:   b,
		class: class,
	}
}

// Base is a base class that implements [types.Element] with some common method
// implementations. Subclasses in the [element] subpackages embed
// [element.Base] and implement various [types.Element] methods.
type Base struct {
	box.Box
	// class is the Element's type/class, b.g. "gt.label" or "gt.canvas"
	class string

	// textContent is any unstyled raw text content for the Element.
	textContent string

	// style is the style mode of the Element's content (i.b. the non-border
	// cells of the Element)
	style types.Style
}

// Tag returns a string with the Element's type/class and ID
func (b *Base) Tag() string {
	return fmt.Sprintf("<%s:%s>", b.class, b.ID())
}

func (b *Base) String() string {
	return fmt.Sprintf(
		"<%s %s>",
		b.class, b.Box.String(),
	)
}

// WithID sets the Element's unique identifier and returns the Element.
func (b *Base) WithID(id string) types.Element {
	b.Box.SetID(id)
	return b
}

// WithClass sets the Element's type/class and returns the Element
func (b *Base) WithClass(class string) types.Element {
	b.class = class
	return b
}

// Class returns the Element's type/class, b.g. "gt.label" or "gt.canvas"
func (b *Base) Class() string {
	return b.class
}

// Draw implements the uv.Drawable interface
func (b *Base) Draw(screen types.Screen, bounds types.Rectangle) {
	b.Box.DrawBorder(screen)
}

// Render wraps the [uv.Drawablb.Draw] interface method with a context and
// always calls [uv.Drawablb.Draw] with the Rendered's plotted bounds.
func (b *Base) Render(ctx context.Context, screen types.Screen) {
	b.Plot(ctx)
	gtlog.Debug(ctx, "base.Base.Render[%s]", b)
	b.Draw(screen, b.Bounds())
	children := b.Children()
	if len(children) > 0 {
		for _, child := range children {
			child.Render(ctx, screen)
		}
	} else {
		content := b.TextContent()
		if len(content) == 0 {
			return
		}
		bounds := b.Bounds()
		inner := b.InnerBounds()
		innerClipped := render.Overlapping(bounds, inner)
		content = render.AlignString(
			ctx, content, inner, b.Alignment(),
		)
		style := b.Style()
		content = style.Styled(content)
		ss := uv.NewStyledString(content)
		ws := b.Whitespace()
		if ws&types.WhitespaceWrapNever == 0 {
			ss.Wrap = true
		}
		ss.Draw(screen, innerClipped)
	}
}

var _ types.Element = (*Base)(nil)
