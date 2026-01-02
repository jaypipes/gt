package main

import (
	"log"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/jaypipes/gt"
	gtlabel "github.com/jaypipes/gt/component/label"
	gtapp "github.com/jaypipes/gt/core/application"
)

const (
	myAppName = "label demo"
	history   = "The Romans learned from the Greeks that quinces slowly cooked with honey would “set” when cool. The Apicius gives a recipe for preserving whole quinces, stems and leaves attached, in a bath of honey diluted with defrutum: Roman marmalade. Preserves of quince and lemon appear (along with rose, apple, plum and pear) in the Book of ceremonies of the Byzantine Emperor Constantine VII Porphyrogennetos."
)

type myApp struct {
	*gt.Application
}

func main() {
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{
		gtapp.New(ctx, gtapp.WithName(myAppName)),
	}

	// gt.Label is a simple component that writes lines of text to a box
	// that is drawn on the screen.
	label := gtlabel.New(
		// A rectangle, anchored at cell (0,0) that is 100 cells wide and 20
		// cells high.
		gtlabel.WithBounds(0, 0, 100, 20),
		// Give the label box a rounded border.
		gtlabel.WithBorder(uv.RoundedBorder()),
		// Set the source of data that will be displayed.
		gtlabel.WithContent(history),
		// Enable wrapping on the text
		gtlabel.WithWrap(true),
	)

	app.SetRoot(label)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
