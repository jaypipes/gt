package label

import (
	"fmt"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/jaypipes/gt/core/element"
)

type LabelModifier func(*Label)

// WithBounds sets the rectangle boundaries of the Label.
func WithBounds(x, y, width, height int) LabelModifier {
	return func(e *Label) {
		e.SetBounds(uv.Rect(x, y, width, height))
	}
}

// WithBorder sets a Border style for the Label.
func WithBorder(border uv.Border) LabelModifier {
	return func(e *Label) {
		e.SetBorder(border)
	}
}

// WithContent sets the Label's content to the supplied thing. The supplied
// thing can be []byte, string, or *uv.StyledString
func WithContent[T []byte | string | *uv.StyledString](s T) LabelModifier {
	return func(e *Label) {
		e.SetContent(s)
	}
}

// WithWrap enables or disables text wrapping for the Label.
func WithWrap(enabled bool) LabelModifier {
	return func(e *Label) {
		e.SetWrap(enabled)
	}
}

// New returns a new Label instance.
func New(mods ...LabelModifier) *Label {
	e := &Label{}
	for _, mod := range mods {
		mod(e)
	}
	return e
}

// Label is a [uv.Drawable] that renders some text to the screen.
type Label struct {
	element.Element
	// ss is the string content of the Label.
	ss *uv.StyledString
}

// SetContent sets the Label's content to the supplied thing. The supplied
// thing can be []byte, string, or *uv.StyledString
func (e *Label) SetContent(content any) {
	if e.ss == nil {
		e.ss = uv.NewStyledString("")
	}
	switch content := content.(type) {
	case string:
		e.ss.Text = content
	case []byte:
		e.ss.Text = string(content)
	case *uv.StyledString:
		e.ss = content
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
func (e *Label) SetWrap(enabled bool) {
	if e.ss == nil {
		e.ss = uv.NewStyledString("")
	}
	e.ss.Wrap = enabled
}

// Draw renders the Label to the given buffer at the specified area.
func (e *Label) Draw(buf uv.Screen, area uv.Rectangle) {
	e.Element.Draw(buf, area)
	if e.ss != nil {
		bb := e.Element.InnerBounds()
		e.ss.Draw(buf, bb)
	}
}
