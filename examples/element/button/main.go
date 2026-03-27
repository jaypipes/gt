package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	"github.com/jaypipes/gt/core/border"
	gtbutton "github.com/jaypipes/gt/element/button"
	gtdiv "github.com/jaypipes/gt/element/div"
)

var (
	buttonText = "click me!"
)

const (
	textFormat = "button clicked?\n\n%t"
)

type myApp struct {
	*gt.Application
}

var (
	buttonClicked = false
)

func content() string {
	return fmt.Sprintf(
		textFormat,
		buttonClicked,
	)
}

func main() {
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}
	app.EnableMouse()
	app.SetBorder(border.Normal())

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

	d := gtdiv.New(
		ctx,
		gt.WithBorder(gt.NormalBorder()),
		gt.WithPadding(gt.PadHorizontal(2)),
		gt.WithWidth(gt.Fixed(60)),
		gt.WithHeight(gt.Fixed(30)),
		gt.WithAlignment(gt.AlignmentMiddleCenter),
		gt.WithWhitespace(gt.WhitespaceNormal),
	)
	d.SetTextContent(content())

	v.AppendContent(d)

	// gt.Button is an Element that renders a clickable button to the terminal
	// screen.
	ta := gtbutton.New(
		ctx,
		gt.WithID("button"),
		gt.WithTextContent(buttonText),
	)

	v.AppendContent(ta)

	ta.OnMouseClick(
		func(ctx context.Context, ev gt.MouseClickEvent) {
			buttonClicked = true
			d.SetTextContent(content())
		},
	)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
