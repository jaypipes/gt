package element

import (
	"context"
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/box"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a [element.Element] with the specified type/class.
func New(ctx context.Context, class string) Element {
	b := box.New(ctx)
	return Element{
		Box:   b,
		class: class,
	}
}

// Element is a base class that implements [types.Element] with some common
// method implementations. Subclasses in the [element] subpackages embed
// [element.Element] and implement various [types.Element] methods.
type Element struct {
	box.Box
	// class is the Element's type/class, e.g. "gt.label" or "gt.canvas"
	class string

	// textContent is any unstyled raw text content for the Element.
	textContent string

	// style is the style mode of the Element's content (i.e. the non-border
	// cells of the Element)
	style types.Style
}

// Tag returns a string with the Element's type/class and ID
func (e *Element) Tag() string {
	return fmt.Sprintf("<%s:%s>", e.class, e.ID())
}

func (e *Element) String() string {
	return fmt.Sprintf(
		"<%s %s>",
		e.class, e.Box.String(),
	)
}

// WithID sets the Element's unique identifier and returns the Element.
func (e *Element) WithID(id string) types.Element {
	e.Box.SetID(id)
	return e
}

// WithClass sets the Element's type/class and returns the Element
func (e *Element) WithClass(class string) types.Element {
	e.class = class
	return e
}

// Class returns the Element's type/class, e.g. "gt.label" or "gt.canvas"
func (e *Element) Class() string {
	return e.class
}

// Draw implements the uv.Drawable interface
func (e *Element) Draw(screen types.Screen, bounds types.Rectangle) {
	ctx := context.TODO()
	gtlog.Debug(ctx, "Element.Draw[%s]: bounds=%s", e.Tag(), bounds)
	e.Box.Draw(screen, bounds)
	content := e.TextContent()
	if len(content) == 0 {
		return
	}
	inner := e.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
	content = render.AlignString(
		ctx, content, inner, e.Alignment(),
	)
	style := e.Style()
	content = style.Styled(content)
	ss := uv.NewStyledString(content)
	ws := e.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}

var _ types.Element = (*Element)(nil)
