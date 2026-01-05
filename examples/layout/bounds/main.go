package main

import (
	"log"

	"github.com/jaypipes/gt"
)

const (
	myAppName = "bounding box demo"
)

type myApp struct {
	*gt.Application
}

func main() {
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	gt.Debug(ctx, "before NewApplication")
	// create a new myApp that wraps the gt.Application
	app := myApp{gt.NewApplication(ctx)}
	app.SetName(myAppName)

	// gt.Box is a simple component that draws a box on the screen.
	box := gt.NewBox(ctx, "mybox")
	// We will constrain the box to a bounding box (viewport) that is a
	// rectangle, anchored at cell (0,0) that is 100 cells wide and 20 cells
	// high.
	box.SetBounds(gt.Rect(0, 0, 100, 20))
	// Make a rounded border around the box
	box.SetBorder(gt.RoundedBorder())
	// By default, the width and height of the box will be width and height of
	// its bounding box, if set. We can change this using the SetSize() method,
	// however note that the width and height specified in SetSize() will never
	// exceed any bounding box's width and height.
	//
	// Play around with these values and see the effect on the rendered box.
	box.SetSize(80, 20)

	// We make the bounding box of the Document (the outermost renderable of the
	// Application) a rectangle, anchored at cell (10,10) that is 40 cells wide
	// and 20 cells high.
	docBounds := gt.Rect(10, 10, 40, 20)

	// By specifying a bounding box (bounds) on the Application's Document, we
	// trigger gt to draw the box component in a viewport that represents the
	// maximum overlapping bounding rectangles for the root bounds and box
	// bounds.
	//
	// In this case, even though we specified a box to be displayed within a
	// bounding box 100 cells wide and 20 cells high, anchored at cell (0,0),
	// the Document's bounding box causes the actual rendered box to be "clipped"
	// in a viewport 40 cells wide and 10 cells high, anchored at cell (10,10).
	//
	// Play around with the anchoring cell values and widths/heights and see
	// the effect on the rendered box.
	doc := app.Document()
	doc.SetBounds(docBounds)
	doc.SetRoot(box)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
