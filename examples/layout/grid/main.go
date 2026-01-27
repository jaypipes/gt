package main

import (
	"log"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	"github.com/jaypipes/gt/element/div"
	"github.com/jaypipes/gt/element/span"
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
	top := div.New(ctx, "Top")
	top.SetID("top")
	top.SetForegroundColor(black)
	top.SetBackgroundColor(yellow)
	top.SetHeight(gt.Percent(25))
	top.SetAlignment(gt.AlignmentMiddleCenter)
	v.AppendContent(top)

	mid := div.New(ctx, "Mid")
	mid.SetID("mid")
	mid.SetHeight(gt.Percent(50))
	mid.SetAlignment(gt.AlignmentMiddleCenter)
	v.AppendContent(mid)

	// gt.Span is similar to an HTML <span> element. It will display any
	// content within a bounding box that by default will begin its content to
	// the right of a previous sibling (display: inline), receive a width in
	// cells equal to the "natural" width of its content and receive a height
	// in lines equal to the "natural" height of its content. By setting these
	// gt.Span's width to either a fixed or percent value, the display mode is
	// automatically switched to "inline-block" which will cause the width to
	// be either a fixed value or a percent of the parent container's width.
	midA := span.New(ctx, "Mid A")
	midA.SetID("mid-a")
	midA.SetWidth(gt.Fixed(10))
	// Setting a height of 100% here forces the height to be equal to the
	// parent container's height.
	midA.SetHeight(gt.Percent(100))
	midA.SetAlignment(gt.AlignmentMiddleCenter)
	midA.SetBorder(gt.RoundedBorder())
	midA.SetForegroundColor(black)
	midA.SetBackgroundColor(pink)
	mid.AppendChild(midA)

	midB := span.New(ctx, "Mid B")
	midB.SetID("mid-b")
	midB.SetWidth(gt.Percent(20))
	midB.SetHeight(gt.Percent(100))
	midB.SetAlignment(gt.AlignmentMiddleCenter)
	midB.SetBorder(gt.RoundedBorder())
	midB.SetForegroundColor(black)
	midB.SetBackgroundColor(pink)
	mid.AppendChild(midB)

	midB1 := div.New(ctx, "Mid B-1")
	midB1.SetID("mid-b1")
	midB1.SetHeight(gt.Percent(33))
	midB1.SetAlignment(gt.AlignmentMiddleCenter)
	midB1.SetForegroundColor(black)
	midB1.SetBackgroundColor(lightblue)
	midB.AppendChild(midB1)

	midB2 := div.New(ctx, "Mid B-2")
	midB2.SetID("mid-b2")
	midB2.SetHeight(gt.Percent(33))
	midB2.SetAlignment(gt.AlignmentMiddleCenter)
	midB2.SetForegroundColor(black)
	midB2.SetBackgroundColor(lightblue)
	midB.AppendChild(midB2)

	midB3 := div.New(ctx, "Mid B-3")
	midB3.SetID("mid-b3")
	midB3.SetHeight(gt.Percent(34))
	midB3.SetAlignment(gt.AlignmentMiddleCenter)
	midB3.SetForegroundColor(black)
	midB3.SetBackgroundColor(lightblue)
	midB.AppendChild(midB3)

	midC := span.New(ctx, "Mid C")
	midC.SetID("mid-c")
	midC.SetWidth(gt.Percent(80))
	midC.SetHeight(gt.Percent(100))
	midC.SetAlignment(gt.AlignmentMiddleCenter)
	midC.SetBorder(gt.RoundedBorder())
	midC.SetForegroundColor(black)
	midC.SetBackgroundColor(pink)
	mid.AppendChild(midC)

	bottom := div.New(ctx, "Bottom")
	bottom.SetID("bottom")
	bottom.SetForegroundColor(black)
	bottom.SetBackgroundColor(lightgreen)
	bottom.SetHeight(gt.Percent(25))
	bottom.SetAlignment(gt.AlignmentMiddleCenter)
	v.AppendContent(bottom)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
