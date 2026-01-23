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
	lightgreen, _ := colorful.Hex("#d1ffbd")
	lightblue, _ := colorful.Hex("#add8e6")

	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}

	// gt.Document represents the Document Object Model (DOM) for your
	// Application's display views.
	doc := app.Document()

	// We create a layout with three horizontal panes consuming 25%, 50% and
	// 25% of the screen's height, respectively. The middle horizontal pane
	// will be divided into three vertical panes consuming 10 cells (fixed),
	// 20% of the remaining and the rest of the remaining width of the screen
	// and the center pane of the middle pane will be further vertically
	// divided into three equally-sized panes:
	//
	// +---------------------------------------------------------------------+
	// |                                                                     |
	// |                                                                     |
	// |                             Top                                     |
	// |                                                                     |
	// |                                                                     |
	// +---------------------------------------------------------------------+
	// |            |             |                                          |
	// |            |             |                                          |
	// |            |   Mid B-1   |                                          |
	// |            |             |                                          |
	// |            |             |                                          |
	// |            +-------------+                                          |
	// |            |             |                                          |
	// |            |             |                                          |
	// |   Mid A    |   Mid B-2   |                 Mid C                    |
	// |            |             |                                          |
	// |            |             |                                          |
	// |            +-------------+                                          |
	// |            |             |                                          |
	// |            |             |                                          |
	// |            |   Mid B-3   |                                          |
	// |            |             |                                          |
	// |            |             |                                          |
	// +---------------------------------------------------------------------+
	// |                                                                     |
	// |                                                                     |
	// |                           Bottom                                    |
	// |                                                                     |
	// |                                                                     |
	// +---------------------------------------------------------------------+

	// gt.Div is similar to an HTML <div> element. It will display any content
	// within a bounding box that by default will begin its content on a new
	// line and consume the width (in cells on the screen) of its parent
	// container and consume the natural height (in lines on the screen) of its
	// content.
	divTop := gtdiv.New(ctx, "Top")
	divTop.SetID("top")
	divTop.SetForegroundColor(black)
	divTop.SetBackgroundColor(yellow)
	divTop.SetHeight(gt.Percent(25))
	divTop.SetAlignment(gt.AlignmentMiddleCenter)
	doc.PushChild(divTop)

	divMid := gtdiv.New(ctx, "Mid")
	divMid.SetID("mid")
	divMid.SetHeight(gt.Percent(50))
	divMid.SetAlignment(gt.AlignmentMiddleCenter)
	doc.PushChild(divMid)

	// gt.Span is similar to an HTML <span> element. It will display any
	// content within a bounding box that by default will begin its content to
	// the right of a previous sibling (display: inline), receive a width in
	// cells equal to the "natural" width of its content and receive a height
	// in lines equal to the "natural" height of its content. By setting these
	// gt.Span's width to either a fixed or percent value, the display mode is
	// automatically switched to "inline-block" which will cause the height to
	// default to the height of the parent container, which is 50% of the
	// screen's number of lines.
	spanMidA := gtspan.New(ctx, "Mid A")
	spanMidA.SetID("mid-a")
	spanMidA.SetWidth(gt.Fixed(10))
	spanMidA.SetAlignment(gt.AlignmentMiddleCenter)
	spanMidA.SetBorder(gt.RoundedBorder())
	spanMidA.SetForegroundColor(black)
	spanMidA.SetBackgroundColor(pink)
	divMid.PushChild(spanMidA)

	spanMidB := gtspan.New(ctx, "Mid B")
	spanMidB.SetID("mid-b")
	spanMidB.SetWidth(gt.Percent(20))
	spanMidB.SetAlignment(gt.AlignmentMiddleCenter)
	spanMidB.SetBorder(gt.RoundedBorder())
	spanMidB.SetForegroundColor(black)
	spanMidB.SetBackgroundColor(pink)
	divMid.PushChild(spanMidB)

	divMidB1 := gtdiv.New(ctx, "Mid B-1")
	divMidB1.SetID("mid-b1")
	divMidB1.SetHeight(gt.Percent(33))
	divMidB1.SetAlignment(gt.AlignmentMiddleCenter)
	divMidB1.SetForegroundColor(black)
	divMidB1.SetBackgroundColor(lightblue)
	spanMidB.PushChild(divMidB1)

	divMidB2 := gtdiv.New(ctx, "Mid B-2")
	divMidB2.SetID("mid-b2")
	divMidB2.SetHeight(gt.Percent(33))
	divMidB2.SetAlignment(gt.AlignmentMiddleCenter)
	divMidB2.SetForegroundColor(black)
	divMidB2.SetBackgroundColor(lightblue)
	spanMidB.PushChild(divMidB2)

	divMidB3 := gtdiv.New(ctx, "Mid B-3")
	divMidB3.SetID("mid-b3")
	divMidB3.SetHeight(gt.Percent(33))
	divMidB3.SetAlignment(gt.AlignmentMiddleCenter)
	divMidB3.SetForegroundColor(black)
	divMidB3.SetBackgroundColor(lightblue)
	spanMidB.PushChild(divMidB3)

	spanMidC := gtspan.New(ctx, "Mid C")
	spanMidC.SetID("mid-c")
	spanMidC.SetWidth(gt.Percent(80))
	spanMidC.SetAlignment(gt.AlignmentMiddleCenter)
	spanMidC.SetBorder(gt.RoundedBorder())
	spanMidC.SetForegroundColor(black)
	spanMidC.SetBackgroundColor(pink)
	divMid.PushChild(spanMidC)

	divBottom := gtdiv.New(ctx, "Bottom")
	divBottom.SetID("bottom")
	divBottom.SetForegroundColor(black)
	divBottom.SetBackgroundColor(lightgreen)
	divBottom.SetHeight(gt.Percent(25))
	divBottom.SetAlignment(gt.AlignmentMiddleCenter)
	doc.PushChild(divBottom)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
