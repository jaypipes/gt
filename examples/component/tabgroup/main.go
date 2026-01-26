package main

import (
	"log"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt"
	"github.com/jaypipes/gt/component/tabgroup"
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

	// gt.View is used to group displayable things that represent a
	// logically-related view of something.
	//
	// gt.Application.View returns a View with the supplied ID. If the
	// Application does not have a View with that ID, a new View is added to
	// the Application and then returned.
	v := app.View(ctx, "main")

	tg := tabgroup.New(ctx, "tg")
	tab1 := tg.Tab(ctx, "tab-1")
	tab1.SetTitle("tab one")

	div1 := gtdiv.New(ctx, "tab 1 content")
	div1.SetID("div-1")
	div1.SetHeight(gt.Percent(100))
	div1.SetAlignment(gt.AlignmentMiddleCenter)
	div1.SetForegroundColor(black)
	div1.SetBackgroundColor(yellow)
	tab1.SetContent(div1)

	tab2 := tg.Tab(ctx, "tab-2")
	tab2.SetTitle("tab two")

	div2 := gtdiv.New(ctx, "content 2")
	div2.SetID("div-2")
	div2.SetHeight(gt.Percent(100))
	div2.SetAlignment(gt.AlignmentMiddleCenter)
	div2.SetForegroundColor(black)
	div2.SetBackgroundColor(pink)
	tab2.SetContent(div2)

	v.SetContent(tg)

	app.SetCurrentView(v.ID())

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
