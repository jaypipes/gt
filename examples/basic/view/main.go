package main

import (
	"log"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtdiv "github.com/jaypipes/gt/element/div"
)

type myApp struct {
	*gt.Application
}

func main() {
	black, _ := colorful.Hex("#000000")
	yellow, _ := colorful.Hex("#ffff00")
	pink, _ := colorful.Hex("#ffcccc")

	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}
	// Application has an optional border and padding.
	app.SetBorder(gt.ThickBorder())
	app.SetPadding(gt.PadHorizontal(1))

	// gt.View is used to group displayable things that represent a
	// logically-related view of something.
	//
	// gt.Application.View returns a View with the supplied ID. If the
	// Application does not have a View with that ID, a new View is added to
	// the Application and then returned.
	v1 := app.View(ctx, "1")

	// Add a keyboard shortcut that will set the Application's active
	// (displayed) View.
	v1.SetCurrentViewKeyPress("1")

	// Views can have borders and padding, too!
	v1.SetBorder(gt.RoundedBorder())
	v1.SetPadding(gt.Pad(1))

	div1 := gtdiv.New(ctx, "content 1")
	div1.SetID("div-1")
	div1.SetHeight(gt.Percent(100))
	div1.SetAlignment(gt.AlignmentMiddleCenter)
	div1.SetForegroundColor(black)
	div1.SetBackgroundColor(yellow)
	v1.SetContent(div1)

	v2 := app.View(ctx, "2")
	v2.SetCurrentViewKeyPress("2")

	v2.SetBorder(gt.RoundedBorder())
	v2.SetPadding(gt.Pad(1))

	div2 := gtdiv.New(ctx, "content 2")
	div2.SetID("div-2")
	div2.SetHeight(gt.Percent(100))
	div2.SetAlignment(gt.AlignmentMiddleCenter)
	div2.SetForegroundColor(black)
	div2.SetBackgroundColor(pink)
	v2.SetContent(div2)

	// gt.Application.SetCurrentView can be used to programmatically switch the
	// active (displayed) View for an Application.
	app.SetCurrentView(v1.ID())

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
