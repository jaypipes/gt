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

	doc := app.Document()

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
	spanC.SetForegroundColor(black)
	spanC.SetBackgroundColor(lightblue)
	spanC.SetAlignment(gt.AlignmentRight)

	divD := gtdiv.New(ctx, "D")
	divD.SetID("D")
	divD.SetForegroundColor(black)
	divD.SetBackgroundColor(lightgreen)
	divD.SetAlignment(gt.AlignmentCenter)

	doc.PushChild(spanA)
	doc.PushChild(spanB)
	doc.PushChild(spanC)
	doc.PushChild(divD)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
