package main

import (
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtspan "github.com/jaypipes/gt/element/span"
	"github.com/lucasb-eyer/go-colorful"
)

type myApp struct {
	*gt.Application
}

func main() {
	yellow, _ := colorful.Hex("#ffff00")
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}
	app.SetTitle("bounds demo")
	// You can set an outer border on your Application.
	app.SetBorder(gt.ThickBorder())

	// gt.View is used to group displayable things that represent a
	// logically-related view of something.
	//
	// Here, we use the gt.Application.View method to return a View with the ID
	// "main". Note that if no such View exists in the Application that that
	// ID, a new empty View with that ID is created, added to the Application,
	// and returned.
	//
	// Below, we add a set of gt.Elements to this View. gt.Elements are
	// displayable primitives that function very much like an HTML element. The
	// View can be seen as the root of a sort of Document Object Model (DOM) of
	// gt.Elements.
	v := app.View(ctx, "main")

	// gt.Span is a simple element that behaves like an HTML <span> element,
	// meaning that by default, the element will be rendered inline with either
	// the previous sibling element, or if no previous sibling, the parent
	// container. By default, Spans have a dynamic width and height, where the
	// width of the Span defaults to the width of the Span's text contents (in
	// cells on the screen) and the height of the Span is the natural number of
	// lines that the text content will wrap within the Span's parent
	// container.
	span := gtspan.New(ctx, "some text")
	// To demonstrate how bounds work, we will give this Span a fixed width and
	// height instead of using the default width and height of a Span. The
	// `gt.Element.SetSize()` method allows you to set the fixed width and
	// height of an element.
	span.SetSize(gt.FixedArea(60, 10))
	// Make a rounded yellow border around the span and place the text in the
	// centered middle of the span.
	span.SetBorder(gt.RoundedBorder())
	span.SetBorderForegroundColor(yellow)
	span.SetAlignment(gt.AlignmentMiddleCenter)
	v.AppendElement(span)

	// We make the bounding box of the Application's View (the outermost
	// renderable of the Application) a rectangle with the top-left coordinates
	// at (10,10) and the bottom-right coordinates at (40,20). In other words,
	// a rectangle anchored at (10,10) that is 30 cells wide and 10 cells high.
	bounds := gt.Rect(10, 10, 40, 20)

	// By specifying a bounding box (bounds) on the Application, we trigger gt
	// to draw the box component in a viewport that represents the maximum
	// overlapping bounding rectangles for the root bounds and box bounds.
	//
	// In this case, even though we specified a box to be displayed within a
	// bounding box 100 cells wide and 20 cells high, anchored at cell (0,0),
	// the View's bounding box causes the actual rendered box to be "clipped"
	// in a viewport 40 cells wide and 10 cells high, anchored at cell (10,10).
	//
	// Play around with the anchoring cell values and widths/heights and see
	// the effect on the rendered box.
	app.SetBounds(bounds)

	// Note that instead of specifying the bounding box using SetBounds(), we
	// could have specified the anchor point of (10,10) using:
	//
	//doc.SetAbsolutePosition(gt.Pt(10, 10))
	//
	// and set the view's width and height:
	//
	//doc.SetSize(gt.FixedArea(30, 10))

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
