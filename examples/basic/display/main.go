package main

import (
	"log"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtdiv "github.com/jaypipes/gt/element/div"
	gtpre "github.com/jaypipes/gt/element/pre"
	gtspan "github.com/jaypipes/gt/element/span"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	shortText        = "Short text"
	longText         = `Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas. Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.`
	preformattedText = `
	This text's		whitespace
				should be preserved			
instead of			being wrapped.

`
	paragraphText = `The quick brown fox
jumps over the lazy dog`
)

type myApp struct {
	*gt.Application
}

func main() {
	red, _ := colorful.Hex("#ff0000")
	yellow, _ := colorful.Hex("#ffff00")
	lightblue, _ := colorful.Hex("#add8e6")
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}
	// You can set an outer border on your Application.
	app.SetBorder(gt.ThickBorder())

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

	// gt.Div is similar to an HTML <div> element. It will display any content
	// within a bounding box that by default will begin its content on a new
	// line and consume the width (in cells on the screen) of its parent
	// container.
	short := gtdiv.New(
		ctx,
		gt.WithID("short"),
		gt.WithTextContent(shortText),
		// An Element's border can be controlled with the `gt.Element.SetBorder`
		// method, which accepts as a parameter a `gt.Border` struct. A `gt.Border`
		// struct is returned from helper functions like `gt.RoundedBorder`.
		gt.WithBorder(gt.RoundedBorder()),
		// You can give your border a foreground color with the
		// `gt.Element.SetBorderForegroundColor` method. Note that a border's
		// foreground color is the color of the character symbols that comprise the
		// border itself.
		gt.WithBorderForegroundColor(yellow),
		// An Element's padding can be controlled with the `gt.Element.SetPadding`
		// method. Here, we pad the left and right of the fixed div by two cells
		gt.WithPadding(gt.PadLR(2, 2)),
		// An Element's horizontal and vertical alignment is controlled with
		// the `gt.WithAlignment` modifier. Here, we place the text "Short
		// text" vertically in the middle and horizontally in the center of the
		// container.
		gt.WithAlignment(gt.AlignmentMiddleCenter),
		// If a Div's width and height are set to a fixed size (using the
		// `gt.WithSize` modifier), the Div's width and height will no longer
		// adjust dynamically to the containing element's bounding box.
		//
		// We know that the default width of short would normally be the width
		// of the parent container's inner bounding box (which is its outer
		// bounding box minus any border and padding), and that the "natural"
		// height of short would be the number of screen lines it would take to
		// output its text contents ("Short text").
		//
		// By calling gt.WithSize to give short a fixed width of 30 cells and a
		// height of 5 lines, we override this dynamic sizing behaviour.
		gt.WithSize(gt.FixedArea(30, 5)),
	)

	// Add short to our View.
	v.AppendContent(short)

	// We will *not* give long a fixed width and height, instead relying on the
	// default sizing of `gt.Div` elements.
	long := gtdiv.New(
		ctx,
		gt.WithTextContent(longText),
		gt.WithID("long"),
		gt.WithBorder(gt.RoundedBorder()),
		gt.WithBorderForegroundColor(red),
		// You can give your border a background color with the
		// `gt.Element.WithBorderBackgroundColor` method. A border's background
		// color is the background color of the cells that comprise the border.
		gt.WithBorderBackgroundColor(lightblue),
		// Give gt a top/bottom padding of 2 lines and a left/right padding of 4
		// cells.
		gt.WithPadding(gt.PadTBLR(2, 2, 4, 4)),
		// gt Elements all have a whitepace mode that controls how text is wrapped
		// and whether sequences of whitespace characters are collapsed.
		//
		// A gt.Div's default whitespace mode is "WhitespaceNormal", which means
		// that sequences of whitespace characters are collapsed and text will wrap
		// when necessary and when line breaks (i.e. \n or \r\n) are found.
		//
		// We can call `gt.WithWhitespace()` to change this whitespace mode.
		//
		// Uncomment the below line to set gt's whitespace mode to
		// WhitespaceWrapNever, which will force the text in gt to be clipped at
		// the container's right margin.
		// gt.WithWhitespace(gt.WhitespaceWrapNever),
	)

	// Add long to our View as a sibling of short.
	v.AppendContent(long)

	// gt.Pre is similar to an HTML <pre> element. It will display any content
	// within a bounding box that by default will begin its content on a new
	// line and consume the width (in cells on the screen) of its parent
	// container. However, unlike a gt.Div, a gt.Pre defaults to preserving
	// whitespace within its text content (as opposed to wrapping by default).
	preformatted := gtpre.New(
		ctx,
		gt.WithTextContent(preformattedText),
		gt.WithID("preformatted"),
		gt.WithBorder(gt.RoundedBorder()),
		// Preformatted text preserves whitespace even when aligned.
		gt.WithAlignment(gt.AlignmentMiddleRight),
	)

	// Add preformatted to our View as a sibling of short and long.
	v.AppendContent(preformatted)

	// gt.Span is similar to an HTML <span> element. It will display any
	// content within a bounding box that by default will begin its content to
	// the right of a previous sibling (display: inline), receive a width in
	// cells equal to the "natural" width of its content and receive a height
	// in lines equal to the "natural" height of its content.
	//
	// The default whitespace mode for gt.Span is "auto", which means that
	// longer text content in the gt.Span will be wrapped at the container's
	// bounding box at word boundaries. You can make a gt.Span behave like a
	// gt.Pre by setting the whitespace mode to gt.WhitespaceModePreserve using
	// the gt.WithWhitespace modifier.
	paragraph := gtspan.New(
		ctx,
		gt.WithTextContent(paragraphText),
		gt.WithID("paragraph"),
		gt.WithBorder(gt.ThickBorder()),
		gt.WithPadding(gt.PadHorizontal(2)),
	)

	// Add paragraph to our View as a sibling of short, long and preformatted.
	v.AppendContent(paragraph)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
