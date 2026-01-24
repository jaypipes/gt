package main

import (
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gthr "github.com/jaypipes/gt/element/hr"
	"github.com/lucasb-eyer/go-colorful"
)

type myApp struct {
	*gt.Application
}

func main() {
	yellow, _ := colorful.Hex("#ffff00")
	blue, _ := colorful.Hex("#0000ff")

	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
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

	// gt.HR is a horizontal rule similar to an HTML <hr> element.
	hrA := gthr.New(ctx)
	hrA.SetID("A")
	// The default width is the width of the parent container, which in this
	// case is the View which defaults to the available width and height of
	// the terminal screen.
	//
	// To override the width of the HR, use the `gt.Element.SetWidth` method.
	hrA.SetWidth(gt.Fixed(20))
	// Note that the default behaviour of an HR is to draw its line centered in
	// the parent container. To change this behaviour, use the
	// `gt.Element.SetAlignment` method.
	hrA.SetAlignment(gt.AlignmentLeft)
	hrA.SetForegroundColor(yellow)
	v.AppendElement(hrA)

	// We'll add another HR without a fixed width or alignment mode to
	// demonstrate the default behaviour of a width equal to the parent
	// container.
	hrB := gthr.New(ctx)
	hrB.SetID("B")
	hrB.SetForegroundColor(blue)
	v.AppendElement(hrB)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
