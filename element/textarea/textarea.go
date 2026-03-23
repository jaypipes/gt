package textarea

import (
	"context"
	"strings"

	"github.com/jaypipes/gt"
	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/border"
	"github.com/jaypipes/gt/core/key"
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

var (
	DefaultEscapeKey = key.New("escape")
	DefaultClearKey  = key.New("alt+r")
)

// WithPlaceholder sets the TextArea's placeholder text.
func WithPlaceholder(placeholder string) types.ElementWithOption {
	return func(e types.Element) {
		ta, ok := e.(*TextArea)
		if ok {
			ta.SetPlaceholder(placeholder)
		}
	}
}

// WithEscapeKey sets the TextArea's escape key. The supplied argument can be a
// string, a [types.Key], a [types.KeyCode] or a [tcell.Key].
func WithEscapeKey(subject any) types.ElementWithOption {
	return func(e types.Element) {
		ta, ok := e.(*TextArea)
		if ok {
			ta.SetEscapeKey(subject)
		}
	}
}

// WithClearKey sets the TextArea's clear key. The supplied argument can be a
// string, a [types.Key], a [types.KeyCode] or a [tcell.Key].
func WithClearKey(subject any) types.ElementWithOption {
	return func(e types.Element) {
		ta, ok := e.(*TextArea)
		if ok {
			ta.SetClearKey(subject)
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
	t.SetEscapeKey(DefaultEscapeKey)
	t.SetClearKey(DefaultClearKey)
	for _, opt := range opts {
		opt(t)
	}
	t.OnFocus(
		func(ctx context.Context, ev types.FocusEvent) {
			focused := ev.Enabled()
			s := ev.Source()
			if focused {
				if s != nil {
					kpi, ok := s.(types.KeyPressEventInterceptor)
					if ok {
						kpi.InterceptKeyPressEvents(ctx, t.escapeKey, t)
					}
				}
			} else {
				if s != nil {
					kpi, ok := s.(types.KeyPressEventInterceptor)
					if ok {
						kpi.StopInterceptKeyPressEvents(ctx)
					}
				}
			}
		},
	)
	t.OnKeyPress(
		func(ctx context.Context, ev gt.KeyPressEvent) bool {
			k := ev.Key()
			if k.Equal(t.escapeKey) {
				// This should never be true, since the Application's main
				// event loop should have been intercepting this key press
				// combination...
				gtlog.Warn(ctx, "TextArea[%s] escape key received!", t.ID())
				return false
			}

			input := t.input
			code := k.Code()
			mods := k.Modifiers()
			if k.Equal(t.clearKey) {
				input.Reset()
			} else {
				// Handle some special keys.
				if mods.None() {
					switch {
					case code == gt.KeyCodeBackspace:
						removeLastRune(input)
					case code == gt.KeyCodeEnter:
						input.WriteRune('\n')
					case code == gt.KeyCodeTab:
						input.WriteString(strings.Repeat(" ", t.tabSize))
					case k.Printable():
						input.WriteRune(rune(code))
					}
				}
			}
			t.SetTextContent(input.String())
			return true
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
	// escapeKey is the key press combination that causes the focus on the
	// TextArea to be lost, resulting in the stoppage of the TextArea
	// processing key strokes.
	escapeKey types.Key
	// clearKey is the key press combination that clears the TextArea's text.
	clearKey types.Key
	// tabSize is the number of spaces a TAB character should consume in the
	// TextArea's text content.
	tabSize int
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

// SetTabSize sets the number of spaces to replace a TAB character in TextArea.
func (t *TextArea) SetTabSize(tabSize int) {
	t.tabSize = tabSize
}

// SetTabSize sets the number of spaces to replace a TAB character in TextArea
// and returns the TextArea.
func (t *TextArea) WithTabSize(tabSize int) *TextArea {
	t.SetTabSize(tabSize)
	return t
}

// TabSize returns the number of spaces to replace TAB character for the
// TextArea.
func (t *TextArea) TabSize() int {
	return t.tabSize
}

// SetEscapeKey sets the TextArea's escape key.
func (t *TextArea) SetEscapeKey(subject any) {
	t.escapeKey = key.New(subject)
}

// WithEscapeKey sets the TextArea's escape key and returns the TextArea.
func (t *TextArea) WithEscapeKey(subject any) *TextArea {
	t.SetEscapeKey(subject)
	return t
}

// EscapeKey returns the escape key for the TextArea.
func (t *TextArea) EscapeKey() types.Key {
	return t.escapeKey
}

// SetClearKey sets the TextArea's clear key.
func (t *TextArea) SetClearKey(subject any) {
	t.clearKey = key.New(subject)
}

// WithClearKey sets the TextArea's clear key and returns the TextArea.
func (t *TextArea) WithClearKey(subject any) *TextArea {
	t.SetClearKey(subject)
	return t
}

// ClearKey returns the clear key for the TextArea.
func (t *TextArea) ClearKey() types.Key {
	return t.clearKey
}

// Render implements the types.Renderable interface
func (t *TextArea) Render(ctx context.Context, h types.ScreenHandler) {
	bounds := t.Bounds()
	gtlog.Debug(ctx, "TextArea.Render[%s]: bounds=%s", t.Tag(), bounds)
	t.Box.Render(ctx, h)

	screen := h.Screen()
	cursor := h.Cursor()

	content := t.TextContent()
	focused := t.HasFocus()
	if len(content) == 0 {
		if !focused {
			content = t.placeholder
		}
	}
	s := t.Style()
	inner := t.InnerBounds()
	lines := strings.Split(content, "\n")
	startX := inner.Min.X
	startY := inner.Min.Y
	for y, line := range lines {
		screen.PutStrStyled(startX, startY+y, line, style.TCell(s))
	}
	// If we have the focus, show the cursor at the end of the TextArea's text
	// content.
	if focused {
		x := inner.Min.X
		y := inner.Min.Y
		y += len(lines) - 1
		x += len(lines[len(lines)-1])
		cursor.SetPosition(types.Point{X: x, Y: y})
	} else {
		cursor.Hide()
	}
}

func removeLastRune(input *strings.Builder) {
	if input.Len() == 0 {
		return
	}
	runes := []rune(input.String())
	lastIndex := len(runes) - 1
	runes = append(runes[:lastIndex], runes[lastIndex+1:]...)
	input.Reset()
	input.WriteString(string(runes))
}
