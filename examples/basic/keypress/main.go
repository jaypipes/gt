package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtdiv "github.com/jaypipes/gt/element/div"
)

const (
	tabSize    = 4
	textFormat = `
press some keys and see what happens

== key pressed ==
 %s

== input ==

 %s

=========================================
Ctrl-C: exit the app
Alt-R:  clear the collected input string
`
	onKeyPressTextFormat = `
modifiers: %s
code: %#02x (%d)
string: %s
`
)

var (
	lastEventText = ""
	input         strings.Builder
)

type myApp struct {
	*gt.Application
}

func content(e gt.Element) string {
	return fmt.Sprintf(
		textFormat,
		lastEventText,
		input.String(),
	)
}

func removeLastRune() {
	if input.Len() == 0 {
		return
	}
	runes := []rune(input.String())
	lastIndex := len(runes) - 1
	runes = append(runes[:lastIndex], runes[lastIndex+1:]...)
	input.Reset()
	input.WriteString(string(runes))
}

func main() {
	white, _ := colorful.Hex("#ffffff")
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}

	// gt.View is used to group displayable things that represent a
	// logically-related view of something.
	v := app.View(ctx, "main")

	d := gtdiv.New(
		ctx,
		gt.WithBorder(gt.NormalBorder()),
		gt.WithBorderForegroundColor(white),
		gt.WithPadding(gt.PadHorizontal(2)),
		gt.WithWidth(gt.Fixed(60)),
		gt.WithHeight(gt.Fixed(30)),
		gt.WithAlignment(gt.AlignmentMiddleCenter),
		gt.WithWhitespace(gt.WhitespaceNormal),
	)
	d.SetTextContent(content(d))

	// Use the gt.NewKey function to return a gt.Key object that you can use to
	// compare to the Key you receive from the gt.KeyPressEvent
	altR := gt.NewKey("alt+r")

	// You can take some action when a key is pressed. Use the OnKeyPress
	// method to add a callback that will execute when any key is pressed. This
	// callback receives a gt.KeyPressEvent object that you can use to examine
	// the key combination that was pressed. The function should return true if
	// the element "consumed" or handled the event, false otherwise.
	d.OnKeyPress(
		func(ctx context.Context, ev gt.KeyPressEvent) bool {
			k := ev.Key()
			// gt.Key.Code() returns the gt.KeyCode, which is the Unicode code
			// point for the trapped key. This is the same as Go's rune type
			// and so can be directly typecast to a rune() below.
			code := k.Code()
			// gt.Key.String() returns a string representation of the gt.Key,
			// e.g. "ctrl+a" or "backspace" or "alt+pgup"
			str := k.String()
			// gt.Key.Modifiers() returns the gt.KeyModifiers for the Key.
			// gt.KeyModifiers has a String() method returning a string
			// representation of all enabled bits on the KeyModifiers bitmask.
			mods := k.Modifiers()
			lastEventText = fmt.Sprintf(
				onKeyPressTextFormat,
				mods.String(),
				code, code,
				str,
			)
			if k.Equal(altR) {
				input.Reset()
			} else {
				// Handle some special keys.
				if mods.None() {
					switch {
					case code == gt.KeyCodeBackspace:
						removeLastRune()
					case code == gt.KeyCodeEnter:
						input.WriteRune('\n')
					case code == gt.KeyCodeTab:
						input.WriteString(strings.Repeat(" ", tabSize))
					case k.Printable():
						input.WriteRune(rune(code))
					}
				}
			}
			d.SetTextContent(content(d))
			return true
		},
	)

	v.AppendContent(d)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
