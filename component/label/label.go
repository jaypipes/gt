package label

import (
	"context"
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/element"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
)

const (
	ElementClass = "gt.label"
)

// New returns a new Label instance.
func New[T []byte | string | *uv.StyledString](
	ctx context.Context,
	content T,
) *Label {
	e := element.New(ctx, ElementClass)
	c := &Label{Element: *e}
	c.SetContent(content)
	return c
}

// Label is a [uv.Drawable] that renders some text to the screen.
type Label struct {
	element.Element
	// ss is the string content of the Label.
	ss *uv.StyledString
}

// SetContent sets the Label's content to the supplied thing. The supplied
// thing can be []byte, string, or *uv.StyledString
func (c *Label) SetContent(content any) {
	if c.ss == nil {
		c.ss = uv.NewStyledString("")
	}
	switch content := content.(type) {
	case string:
		c.ss.Text = content
	case []byte:
		c.ss.Text = string(content)
	case *uv.StyledString:
		c.ss = content
	default:
		msg := fmt.Sprintf(
			"must pass []byte, string or *uv.StyledString to SetContent(). "+
				"You passed a %T",
			content,
		)
		panic(msg)
	}
}

// SetWrap sets the Label's wrapping behaviour.
func (c *Label) SetWrap(enabled bool) {
	if c.ss == nil {
		c.ss = uv.NewStyledString("")
	}
	c.ss.Wrap = enabled
}

// Draw renders the Label to the given buffer at the specified area.
func (c *Label) Draw(buf uv.Screen, bounds types.Rectangle) {
	outer := c.Bounds()
	outerClipped := render.Overlapping(bounds, outer)
	c.Element.Draw(buf, outerClipped)
	if c.ss != nil {
		inner := c.InnerBounds()
		innerClipped := render.Overlapping(bounds, inner)
		c.ss.Draw(buf, innerClipped)
	}
}
