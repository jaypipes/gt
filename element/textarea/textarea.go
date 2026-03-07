package textarea

import (
	"context"
	"fmt"
	"strings"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/border"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/style"
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
	t := &TextArea{
		Element: e,
		input:   &strings.Builder{},
	}
	// TextArea defaults to top-left alignment and preserving the user's exact
	// input whitespacing and a square, thin-line border.
	t.SetDisplay(types.DisplayInlineBlock)
	t.SetAlignment(types.AlignmentTopLeft)
	t.SetWhitespace(types.WhitespacePreserve)
	t.SetBorder(border.Normal())
	// TextArea defaults to a width of 20 cells and a height of 2 lines, the
	// same as the HTML element of the same name.
	t.SetWidth(core.Fixed(DefaultWidth))
	t.SetHeight(core.Fixed(DefaultHeight))
	for _, opt := range opts {
		opt(t)
	}
	t.OnFocus(
		func(ctx context.Context) {
			c := t.Controller()
			if c != nil {
				c.InterceptKeyPress("tab", t.input)
			}
		},
	)
	t.OnLoseFocus(
		func(ctx context.Context) {
			c := t.Controller()
			if c != nil {
				c.RestoreKeyPress()
				c.HideCursor()
			}
		},
	)
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
	// input allows us to receive key press content
	input *strings.Builder
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

// Render implements the types.Renderable interface
func (t *TextArea) Render(ctx context.Context, screen types.Screen) {
	bounds := t.Bounds()
	gtlog.Debug(ctx, "TextArea.Render[%s]: bounds=%s", t.Tag(), bounds)
	t.Box.Render(ctx, screen)
	content := t.TextContent()
	focused := t.HasFocus()
	if focused {
		input := t.input
		// If we've got some input text, update the stored text content
		if input.Len() > 0 {
			content = fmt.Sprintf("%s%s", content, input.String())
			t.SetTextContent(content)
			input.Reset()
		}
	}
	if len(content) == 0 {
		if !focused {
			content = t.placeholder
		}
	}
	defStyle := t.Style()
	inner := t.InnerBounds()
	lines := strings.Split(content, "\n")
	startX := inner.Min.X
	startY := inner.Min.Y
	for y, line := range lines {
		for x := range line {
			screen.Put(startX+x, startY+y, string(line[x]), style.TCell(defStyle))
		}
	}
	// If we have the focus, show the cursor at the end of the user-input text
	// to indicate this is an editable thing.
	if focused {
		c := t.Controller()
		if c != nil {
			x := inner.Max.X
			y := inner.Max.Y
			c.ShowCursor(x, y)
			c.SetCursorStyle(types.CursorStyleBar)
		}
	}
}
