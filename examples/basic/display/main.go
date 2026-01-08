package main

import (
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtdiv "github.com/jaypipes/gt/element/div"
)

const (
	myAppName = "basic demo of display, border and padding properties"
	shortText = "Short text"
	longText  = `
Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque
faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium
tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar
vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl
malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class
aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos
himenaeos.
`
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

	// gt.Document represents the Document Object Model (DOM) for your
	// Application's display views.
	doc := app.Document()
	// You can set an outer border on your Application by setting a border on
	// the Document.
	doc.SetBorder(gt.ThickBorder())

	// gt.Div is similar to an HTML <div> element. It will display any content
	// within a bounding box that by default will begin its content on a new
	// line and consume the width (in cells on the screen) of its parent
	// container.
	divA := gtdiv.New(ctx, shortText)
	divA.SetID("A")
	// Give the div a rounded border.
	divA.SetBorder(gt.RoundedBorder())
	// An Element's padding can be controlled with the `gt.Element.SetPadding`
	// method. Here, we pad the left and right of the fixed div by two cells
	divA.SetPadding(gt.PadLR(2, 2))
	// If a Div's width and height are set to a fixed size (using
	// `gt.Element.SetSize`), the Div's width and height will no longer adjust
	// dynamically to the containing element's bounding box.
	//
	// We know that the default width of divA would normally be the width of
	// the parent container's inner bounding box (which is its outer bounding
	// box minus any border and padding), and that the "natural" height of divA
	// would be the number of screen lines it would take to output its text
	// contents ("Short text").
	//
	// By calling SetSize to give divA a fixed width of 30 cells and a height
	// of 5 lines, we override this dynamic sizing behaviour.
	divA.SetSize(30, 5)
	// Add divA to our Application's Document.
	doc.PushChild(divA)

	// We will *not* give divB a fixed width and height, instead relying on the
	// default sizing of `gt.Div` elements.
	divB := gtdiv.New(ctx, longText)
	divB.SetID("B")
	// Give divB a double-lined border to make it distinguishable from divA on
	// the screen.
	divB.SetBorder(gt.DoubleBorder())
	// Add divB to our Application's Document as a sibling of divA.
	doc.PushChild(divB)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
