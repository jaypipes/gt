package textarea

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass  = "gt.textarea"
	DefaultWidth  = 20
	DefaultHeight = 2
)

// WithPlaceholder modifies the Element with Placeholder text.
func WithPlaceholder(placeholder string) types.ElementWithOption {
	return func(e types.Element) {
		ta, ok := e.(*TextArea)
		if ok {
			ta.SetPlaceholder(placeholder)
		}
	}
}

// New returns a new TextArea instance with the given options.
func New(
	ctx context.Context,
	opts ...types.ElementWithOption,
) *TextArea {
	e := element.New(ctx, ElementClass)
	t := &TextArea{Element: e}
	// TextArea defaults to top-left alignment and preserving the user's exact
	// input whitespacing and a square, thin-line border.
	t.SetDisplay(types.DisplayInlineBlock)
	t.SetAlignment(types.AlignmentTopLeft)
	t.SetWhitespace(types.WhitespacePreserve)
	t.SetBorder(uv.NormalBorder())
	// TextArea defaults to a width of 20 cells and a height of 2 lines, the
	// same as the HTML element of the same name.
	t.SetWidth(core.Fixed(DefaultWidth))
	t.SetHeight(core.Fixed(DefaultHeight))
	for _, opt := range opts {
		opt(t)
	}
	return t
}

// TextArea is an Element that renders a multi-line input box for user-entered
// text.
//
// TextArea defaults to a width of 20 cells and a height of 2 lines, the same
// as the HTML element of the same name.
type TextArea struct {
	element.Element
	// placeholder contains the text content that will be displayed in the
	// absence of user-provided text content.
	placeholder string
}

// SetPlaceholder sets the TextArea's placeholder text. Placeholder text is
// displayed in the absence of user-provided text content.
func (t *TextArea) SetPlaceholder(placeholder string) {
	t.placeholder = placeholder
}

// WithPlaceholder sets the TextArea's placeholder text and returns the
// TextArea.
func (t *TextArea) WithPlaceholder(placeholder string) *TextArea {
	t.SetPlaceholder(placeholder)
	return t
}

// Placeholder returns the placeholder text for the TextArea. Placeholder text
// is displayed in the absence of user-provided text content.
func (t *TextArea) Placeholder() string {
	return t.placeholder
}

// Draw draws the TextArea to the supplied Screen.
func (t *TextArea) Draw(screen types.Screen, bounds types.Rectangle) {
	t.Box.Draw(screen, bounds)
	focused := t.HasFocus()
	content := t.TextContent()
	if len(content) == 0 {
		if !focused {
			content = t.placeholder
		}
	}
	ss := uv.NewStyledString(content)
	ss.Draw(screen, t.InnerBounds())
	// If we have the focus, show the cursor at the end of the user-input text
	// to indicate this is an editable thing.
	if focused {
		sc := t.ScreenController()
		if sc != nil {
			sc.ShowCursor()
			sc.SetCursorStyle(types.CursorBar, true)
			x := ss.Bounds().Max.X + 1
			y := ss.Bounds().Max.Y
			sc.SetCursorPosition(x, y)
		}
	}
}

// defaultOnFocus is executed when a TextArea receives the focus. If no
// user-supplied text content has been set on the TextArea, when receiving
// focus, the placeholder text is replaced with an empty cursor indicating the
// user can input text in the TextArea.
func (t *TextArea) defaultOnFocus(ctx context.Context) {
	content := t.TextContent()
	if len(content) == 0 {

	}
}
