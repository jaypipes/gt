package main

import (
	"log"

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
	app := myApp{gtapp.New(ctx)}
	app.SetName(myAppName)

	// gt.Label is a simple component that writes lines of text to a box
	// that is drawn on the screen.
	label := gtlabel.New(history)
	// Constrain the label to 100 cells wide and 20 lines high
	label.SetSize(100, 20)
	// pad the left and right of the label by two cells
	label.SetPadding(gt.PadLR(2, 2))
	// Offset the label 1 cell below and 1 cell to the right of the anchor
	// point (0, 0) of the containing box. The containing box in this case is
	// the canvas's inner bounding rectangle.
	label.SetRelativePosition(1, 1)
	// Give the label box a rounded border.
	label.SetBorder(gt.RoundedBorder())
	// Enable wrapping on the text
	label.SetWrap(true)

	canvas := app.Canvas()
	canvas.SetBorder(gt.ThickBorder())
	canvas.SetRoot(label)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
