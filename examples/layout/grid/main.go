package main

import (
	"log"

	"github.com/jaypipes/gt"
)

const (
	myAppName = "layout demo"
)

type myApp struct {
	*gt.Application
}

func main() {
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gt.NewApplication(ctx)}
	app.SetName(myAppName)

	// gt.Canvas can be used to lay out a screen into different panes
	canvas := app.Canvas()

	// We create a layout with three horizontal panes consuming 25%, 50% and 25%
	// of the screen's height, respectively. The middle horizontal pane will be
	// divided into three vertical panes consuming 5 cells (fixed), 20% of the
	// remaining and the rest of the remaining width of the screen:
	//
	// +---------------------------------------------------------------------+
	// |                                                                     |
	// |                             Top                                     |
	// |                                                                     |
	// +---------------------------------------------------------------------+
	// |        |             |                                              |
	// |  Mid 1 |    Mid 2    |                   Mid 3                      |
	// |        |             |                                              |
	// +---------------------------------------------------------------------+
	// |                                                                     |
	// |                           Bottom                                    |
	// |                                                                     |
	// +---------------------------------------------------------------------+

	top := gt.NewLabel("Top")
	canvas.SplitVertical(top, gt.Percent(25))

	mid := gt.NewBox()
	mid.SetBorder(gt.RoundedBorder())

	mid1 := gt.NewLabel("Mid 1")
	mid2 := gt.NewLabel("Mid 2")
	mid3 := gt.NewLabel("Mid 3")
	mid.SplitHorizontal(mid1, gt.Fixed(5))
	mid.SplitHorizontal(mid2, gt.Percent(20))
	mid.SplitHorizontal(mid3, gt.Percent(80))
	canvas.SplitVertical(mid, gt.Percent(50))

	bottom := gt.NewLabel("Bottom")
	canvas.SplitVertical(bottom, gt.Percent(25))

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
