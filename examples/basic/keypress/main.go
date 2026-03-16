package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtdiv "github.com/jaypipes/gt/element/div"
)

const (
	textFormat           = "press some keys and see what happens\n\nlast event:\n\n%s"
	onKeyPressTextFormat = "key pressed (modifiers: %s rune: %c)"
)

var (
	lastEventText = ""
)

type myApp struct {
	*gt.Application
}

func content(e gt.Element) string {
	return fmt.Sprintf(
		textFormat,
		lastEventText,
	)
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

	// You can take some action when a key is pressed. Use the OnKeyPress
	// method to add a callback that will execute when any key is pressed. This
	// callback receives a gt.KeyPressEvent object that you can use to examine
	// the key combination that was pressed. The function should return true if
	// the element "consumed" or handled the event, false otherwise.
	d.OnKeyPress(
		func(ctx context.Context, ev gt.KeyPressEvent) bool {
			k := ev.Key()
			lastEventText = fmt.Sprintf(
				onKeyPressTextFormat,
				k.Modifiers(),
				k.Code(),
			)
			d.SetTextContent(content(d))
			return true
		},
	)

	v.AppendContent(d)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
