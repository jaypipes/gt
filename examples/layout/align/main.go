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

	a := gtspan.New(
		ctx,
		gt.WithTextContent("A"),
		gt.WithID("A"),
		gt.WithWidth(gt.Fixed(20)),
		gt.WithForegroundColor(black),
		gt.WithBackgroundColor(yellow),
		gt.WithAlignment(gt.AlignmentLeft),
	)

	b := gtspan.New(
		ctx,
		gt.WithTextContent("B"),
		gt.WithID("B"),
		gt.WithWidth(gt.Fixed(20)),
		gt.WithForegroundColor(black),
		gt.WithBackgroundColor(pink),
		gt.WithAlignment(gt.AlignmentCenter),
	)

	c := gtspan.New(
		ctx,
		gt.WithTextContent("C"),
		gt.WithID("C"),
		gt.WithWidth(gt.Fixed(20)),
		// We set a fixed height of 2 for this span to demonstrate how inline-block
		// Elements will lay out on the Screen. This will have the effect of making
		// this span twice the height of spans "A" and "B", since the natural
		// (content) height of those spans is 1.
		gt.WithHeight(gt.Fixed(2)),
		gt.WithForegroundColor(black),
		gt.WithBackgroundColor(lightblue),
		gt.WithAlignment(gt.AlignmentRight),
	)

	d := gtdiv.New(
		ctx,
		gt.WithTextContent("D"),
		gt.WithID("D"),
		gt.WithForegroundColor(black),
		gt.WithBackgroundColor(lightgreen),
		gt.WithAlignment(gt.AlignmentCenter),
	)

	// Calling gt.View.AppendContent pushes the element into the View. Elements
	// pushed into the View are displayed top-down on the Screen in the order
	// they are pushed/appended.
	v.AppendContent(a)
	v.AppendContent(b)
	v.AppendContent(c)
	v.AppendContent(d)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
