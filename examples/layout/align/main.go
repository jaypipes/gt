package main

import (
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtdiv "github.com/jaypipes/gt/element/div"
	gtspan "github.com/jaypipes/gt/element/span"
	"github.com/lucasb-eyer/go-colorful"
)

type myApp struct {
	*gt.Application
}

func main() {
	black, _ := colorful.Hex("#000000")
	yellow, _ := colorful.Hex("#ffff00")
	pink, _ := colorful.Hex("#ffcccc")
	lightblue, _ := colorful.Hex("#add8e6")
	lightgreen, _ := colorful.Hex("#d1ffbd")

	ctx := gt.ContextFromEnv()
	app := myApp{gtapp.New(ctx)}

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

	spanA := gtspan.New(ctx, "A")
	spanA.SetID("A")
	spanA.SetWidth(gt.Fixed(20))
	spanA.SetForegroundColor(black)
	spanA.SetBackgroundColor(yellow)
	spanA.SetAlignment(gt.AlignmentLeft)

	spanB := gtspan.New(ctx, "B")
	spanB.SetID("B")
	spanB.SetWidth(gt.Fixed(20))
	spanB.SetForegroundColor(black)
	spanB.SetBackgroundColor(pink)
	spanB.SetAlignment(gt.AlignmentCenter)

	spanC := gtspan.New(ctx, "C")
	spanC.SetID("C")
	spanC.SetWidth(gt.Fixed(20))
	// We set a fixed height of 2 for this span to demonstrate how inline-block
	// Elements will lay out on the Screen. This will have the effect of making
	// this span twice the height of spans "A" and "B", since the natural
	// (content) height of those spans is 1.
	spanC.SetHeight(gt.Fixed(2))
	spanC.SetForegroundColor(black)
	spanC.SetBackgroundColor(lightblue)
	spanC.SetAlignment(gt.AlignmentRight)

	divD := gtdiv.New(ctx, "D")
	divD.SetID("D")
	divD.SetForegroundColor(black)
	divD.SetBackgroundColor(lightgreen)
	divD.SetAlignment(gt.AlignmentCenter)

	// Calling gt.View.AppendContent pushes the element into the View. Elements
	// pushed into the View are displayed top-down on the Screen in the order
	// they are pushed/appended.
	v.AppendContent(spanA)
	v.AppendContent(spanB)
	v.AppendContent(spanC)
	v.AppendContent(divD)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
