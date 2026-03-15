package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt"
	gtapp "github.com/jaypipes/gt/core/application"
	gtdiv "github.com/jaypipes/gt/element/div"
)

const (
	textFormat           = "perform mouse actions on this box and see what happens\n\nhas focus? %t\n\nlast mouse event:\n\n%s"
	onHoverTextFormat    = "hovering (pos: %s)"
	onClickTextFormat    = "click (pos: %s double-click? %t button: %s)"
	onDragMoveTextFormat = "drag move (start pos: %s current pos: %s)"
	onDragStopTextFormat = "drag stop (start pos: %s end pos: %s)"
)

var (
	lastEventText = ""
)

type myApp struct {
	*gt.Application
}

func content(e gt.Element) string {
	return fmt.Sprintf(
		textFormat,
		e.HasFocus(),
		lastEventText,
	)
}

func main() {
	white, _ := colorful.Hex("#ffffff")
	red, _ := colorful.Hex("#ff0000")
	yellow, _ := colorful.Hex("#ffff00")
	// create a new context.Context from environs variables
	ctx := gt.ContextFromEnv()
	// create a new myApp that wraps the gt.Application
	app := myApp{gtapp.New(ctx)}
	app.EnableMouse()

	// gt.View is used to group displayable things that represent a
	// logically-related view of something.
	v := app.View(ctx, "main")

	d := gtdiv.New(
		ctx,
		gt.WithBorder(gt.NormalBorder()),
		gt.WithBorderForegroundColor(white),
		gt.WithPadding(gt.PadHorizontal(2)),
		gt.WithWidth(gt.Fixed(60)),
		gt.WithHeight(gt.Fixed(30)),
		gt.WithAlignment(gt.AlignmentMiddleCenter),
		gt.WithWhitespace(gt.WhitespaceNormal),
	)
	d.SetTextContent(content(d))

	// The focus is given to an element when the element is clicked on and the
	// element is not disabled. You can take some action when the element gets
	// the focus by adding a callback to the element with the
	// gt.Element.OnFocus() method.
	d.OnFocus(
		func(ctx context.Context) {
			d.SetBorderForegroundColor(red)
			d.SetTextContent(content(d))
		},
	)
	d.OnLoseFocus(
		func(ctx context.Context) {
			d.SetBorderForegroundColor(white)
			d.SetTextContent(content(d))
		},
	)

	// You can take some action when the mouse hovers over some element. Mouse
	// hover fires when the mouse is over an element but the element does *not*
	// have the focus.
	d.OnMouseHover(
		func(ctx context.Context, ev gt.MouseEvent) {
			d.SetBorderForegroundColor(yellow)
			lastEventText = fmt.Sprintf(
				onHoverTextFormat, ev.Position(),
			)
			d.SetTextContent(content(d))
		},
	)

	// You can take some action when the mouse clicks or double-clicks on some
	// element.
	d.OnMouseClick(
		func(ctx context.Context, ev gt.MouseClickEvent) {
			lastEventText = fmt.Sprintf(
				onClickTextFormat,
				ev.Position(), ev.DoubleClicked(), ev.Button().String(),
			)
			d.SetTextContent(content(d))
		},
	)

	// You can take some action when a mouse drag operation is in progress and
	// has completed.
	d.OnMouseDragMove(
		func(ctx context.Context, ev gt.MouseDragEvent) {
			lastEventText = fmt.Sprintf(
				// The MouseDragEvent.Start() returns the MouseEvent when the
				// user originally pressed a mouse button and began to drag the
				// mouse.
				//
				// The MouseDragEvent.Position() returns the current position
				// of the mouse.
				onDragMoveTextFormat, ev.Start().Position(), ev.Position(),
			)
			d.SetTextContent(content(d))
		},
	)
	d.OnMouseDragStop(
		func(ctx context.Context, ev gt.MouseDragEvent) {
			lastEventText = fmt.Sprintf(
				// The MouseDragEvent.Start() returns the MouseEvent when the
				// user originally pressed a mouse button and began to drag the
				// mouse.
				//
				// The MouseDragEvent.Position() returns the position of the
				// mouse when the user released the mouse button.
				onDragStopTextFormat, ev.Start().Position(), ev.Position(),
			)
			d.SetTextContent(content(d))
		},
	)

	v.AppendContent(d)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
