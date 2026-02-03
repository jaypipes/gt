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

	// TabGroup has built-in mouse click handlers for when you click on a tab
	// in the tab bar to switch the current tab. Enable mouse handling on the
	// gt.Application to take advantage of these handlers.
	app.EnableMouse()

	// gt.View is used to group displayable things that represent a
	// logically-related view of something.
	//
	// gt.Application.View returns a View with the supplied ID. If the
	// Application does not have a View with that ID, a new View is added to
	// the Application and then returned.
	v := app.View(ctx, "main")
	// You can set an outer border on your View.
	v.SetBorder(gt.RoundedBorder())

	// gt.TabGroup contains zero or more Tabs in a bar that, by default,
	// displays the tabs above the tab content.
	tg := tabgroup.New(ctx, "tg")

	// The bar on the TabGroup can be styled separately.
	bar := tg.Bar()
	bar.SetBorderForegroundColor(yellow)

	tab1 := tg.Tab(ctx, "tab-1")
	tab1.SetTitle("tab one")

	// Tabs default to receiving 100% of the width and height of the container,
	// which is the TabGroup. The TabGroup likewise defaults to 100% of the
	// width and height of its container, which is the View, which defaults to
	// 100% of the width and height of its container, which is the Screen.
	//
	// You can style and decorate the Tab like any other element.
	tab1.SetAlignment(gt.AlignmentMiddleCenter)
	tab1.SetForegroundColor(black)
	tab1.SetBackgroundColor(yellow)

	// You can set a key press combination to trigger the tab to become the
	// active tab.
	tab1.SetCurrentTabKeyPress("1")

	// Give the Tab some content by creating a new gt.Element or gt.Component
	// and using Tab.SetContent or Tab.AppendContent.
	div1 := gtdiv.New(
		ctx,
		gt.WithID("div-1"),
		gt.WithTextContent("tab 1 content"),
		gt.WithHeight(gt.Percent(100)),
	)
	tab1.SetContent(div1)

	tab2 := tg.Tab(ctx, "tab-2")
	tab2.SetTitle("tab two")
	tab2.SetAlignment(gt.AlignmentMiddleCenter)
	tab2.SetForegroundColor(black)
	tab2.SetBackgroundColor(pink)
	tab2.SetCurrentTabKeyPress("2")

	div2 := gtdiv.New(
		ctx,
		gt.WithID("div-2"),
		gt.WithTextContent("tab 2 content"),
		gt.WithHeight(gt.Percent(100)),
	)
	tab2.SetContent(div2)

	// Add the TabGroup to our View.
	v.AppendContent(tg)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
