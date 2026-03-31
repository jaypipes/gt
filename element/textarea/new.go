package textarea

import (
	"context"
	"strings"

	"github.com/jaypipes/gt"
	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/border"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
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

// WithTabSize sets the TextArea's tab size (number of spaces to replace TAB
// characters).
func WithTabSize(size int) types.ElementWithOption {
	return func(e types.Element) {
		ta, ok := e.(*TextArea)
		if ok {
			ta.SetTabSize(size)
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
	t.SetTabSize(DefaultTabSize)
	// TextArea is an input element so should be able to receive the focus.
	t.SetFocusable(true)
	for _, opt := range opts {
		opt(t)
	}
	t.OnFocus(
		func(ctx context.Context, ev types.FocusEvent) {
			focused := ev.Focused()
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
			if !t.HasFocus() {
				return false
			}
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
