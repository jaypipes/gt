package main

import (
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtmotif "github.com/jaypipes/gt/core/motif"
	gtdiv "github.com/jaypipes/gt/element/div"
)

const (
	divContent = `
this is a <div> styled
with the NordDarkPrimary motif.
`
	helpContent = `
Click on an element to see the element's focus style.

Hover over an element to see the element's hover style.

Exit with 'Ctrl+C'
`
)

func main() {
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := gtapp.New(ctx)
	app.EnableMouse()

	v := app.View(ctx, "main")
	v.SetBorder(gt.RoundedBorder())

	d := gtdiv.New(
		ctx,
		gt.WithBorder(gt.NormalBorder()),
		gt.WithPadding(gt.PadHorizontal(2)),
		gt.WithWidth(gt.Fixed(60)),
		gt.WithHeight(gt.Fixed(30)),
		gt.WithAlignment(gt.AlignmentMiddleCenter),
		gt.WithWhitespace(gt.WhitespaceNormal),
		gt.WithMotif(gtmotif.NordDarkPrimary),
	)
	d.SetTextContent(divContent)
	v.AppendContent(d)

	help := gtdiv.New(
		ctx,
		gt.WithID("help"),
		gt.WithTextContent(helpContent),
		gt.WithAlignment(gt.AlignmentMiddleCenter),
		gt.WithHeight(gt.Fixed(7)),
	)
	v.AppendContent(help)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
