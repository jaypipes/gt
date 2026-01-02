package main

import (
	"log"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtbox "github.com/jaypipes/gt/core/box"
	gtctx "github.com/jaypipes/gt/core/context"
)

const (
	myAppName = "root with bounds demo"
)

type myApp struct {
	*gt.Application
}

func main() {
	// create a new context.Context from environs variables
	ctx := gtctx.FromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{
		gtapp.New(ctx, gtapp.WithName(myAppName)),
	}

	// gt.Box is a simple component that draws a box on the screen.
	box := gtbox.New(
		// A rectangle, anchored at cell (0,0) that is 100 cells wide and 20
		// cells high.
		gtbox.WithBounds(0, 0, 100, 20),
		gtbox.WithBorder(uv.RoundedBorder()),
	)

	// A rectangle, anchored at cell (10,10) that is 40 cells wide and 20 cells
	// high.
	rootBounds := uv.Rect(10, 10, 40, 20)

	// By specifying a a bounding box (bounds) when setting the Application
	// root element, we trigger gt to draw the box component in a viewport that
	// represents the maximum overlapping bounding rectangles for the root
	// bounds and box bounds.
	//
	// In this case, even though we specified a box to be 100 cells wide and 20
	// cells high, anchored at cell (0,0), the root bounds rectangle causes the
	// actual rendererd box to be 40 cells wide and 10 cells high, anchored at
	// cell (10,10).
	//
	// Play around with the anchoring cell values and widths/heights and see
	// the effect on the rendered box.
	app.SetRootWithBounds(box, rootBounds)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
