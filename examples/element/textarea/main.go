package main

import (
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	"github.com/jaypipes/gt/core/border"
	gtstyle "github.com/jaypipes/gt/core/style"
	gttextarea "github.com/jaypipes/gt/element/textarea"
	"github.com/lucasb-eyer/go-colorful"
)

var (
	placeholder = "<placeholder text>"
)

type myApp struct {
	*gt.Application
}

func main() {
	yellow, _ := colorful.Hex("#ffff00")
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}
	app.EnableMouse()
	app.SetBorder(border.Normal())

	normalStyle := gtstyle.New()
	hoverStyle := gtstyle.New(gtstyle.WithForegroundColor(yellow))
	focusBorder := gt.DoubleBorder()
	hoverBorder := gt.NormalBorder()
	hoverBorder.SetForegroundColor(yellow)

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

	var userInput string

	// gt.TextArea is an Element that renders a multi-line input box for
	// user-entered text similar to the HTML <textarea> element
	ta := gttextarea.New(
		ctx,
		gt.WithID("textarea"),
		gt.WithTextContent(userInput),
		// You can style your gt.TextArea like any other gt.Element.
		gt.WithStyle(normalStyle),
		gt.WithHoverStyle(hoverStyle),
		gt.WithFocusBorder(focusBorder),
		gt.WithHoverBorder(hoverBorder),
		// Placeholder text is displayed in the absence of user-provided text
		// input and is hidden when focus is placed on the TextArea.
		gttextarea.WithPlaceholder(placeholder),
	)
	v.AppendContent(ta)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
