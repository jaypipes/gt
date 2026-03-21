package application

import (
	"context"
	"time"

	mevent "github.com/jaypipes/gt/core/event/mouse"
	"github.com/jaypipes/gt/types"
)

// setHover sets the currently-hovered thing and calls MouseLoseHover() on the
// previously-hovered thing if the hovered thing has changed.
func (a *Application) setHover(
	ctx context.Context,
	m types.MouseEventHandler,
	ev types.MouseEvent,
) bool {
	changed := false
	if a.hovered != nil {
		if a.hovered != m {
			a.hovered.MouseLoseHover(ctx)
			changed = true
		}
	} else {
		// nothing previously being hovered.
		if m != nil {
			m.MouseHover(ctx, ev)
			changed = true
		}
	}
	a.hovered = m
	return changed
}

// handleMouseEvent determines what logical action the user took with the mouse
// and executes the appropriate mouse event handler for the target element.
func (a *Application) handleMouseEvent(
	ctx context.Context,
	ev types.MouseEvent,
) {
	// We determine if the mouse is over a thing that can handle to mouse
	// events. If the mouse is over something that can handle mouse events and
	// is not disabled, we will fire the OnMouseHover event if the thing does
	// not have the focus.
	var target types.MouseEventHandler

	pos := ev.Position()
	v := a.CurrentView()
	node := v.AtPoint(pos)
	redraw := false
	if node != nil {
		el, ok := node.(types.Element)
		if ok && !el.Disabled() {
			target = el
		}
	}

	// We determine what logical mouse action the user has taken by
	// examining our stored state of previous mouse events.
	x, y := pos.X, pos.Y
	button := ev.Button()
	buttonWasDown := a.mouseDownEvent != nil
	buttonNowDown := button.Pressable()

	downX, downY := x, y
	if buttonWasDown {
		downPos := a.mouseDownEvent.Position()
		downX, downY = downPos.X, downPos.Y
	}

	downMoved := x != downX || y != downY

	switch {
	case buttonWasDown && buttonNowDown && downMoved:
		// mouse has moved while a button was pressed -- i.e. a drag operation.
		if target != nil {
			de := mevent.NewDragEvent(ev, a.mouseDownEvent)
			target.MouseDragMove(ctx, de)
			redraw = true
		}
		a.mouseDragged = true
	case !buttonWasDown && buttonNowDown && !downMoved && !a.mouseDragged:
		// mouse was clicked or double-clicked.
		a.mouseDownEvent = ev
		if a.lastMouseClickTime.Add(types.DefaultMouseDoubleClickInterval).Before(time.Now()) {
			a.lastMouseClickTime = time.Now()
			if target != nil {
				ce := mevent.NewClickEvent(ev, false)
				target.MouseClick(ctx, ce)
				f, ok := target.(types.FocusEventHandler)
				if ok {
					a.setFocus(ctx, f)
				}
				redraw = true
			} else {
				// mouse was clicked on a part of the screen represented by no
				// element, so we remove the focus from whatever element had
				// the focus.
				redraw = a.setFocus(ctx, nil)
			}
		} else {
			a.lastMouseClickTime = time.Time{}
			if target != nil {
				ce := mevent.NewClickEvent(ev, true)
				target.MouseClick(ctx, ce)
				f, ok := target.(types.FocusEventHandler)
				if ok {
					a.setFocus(ctx, f)
				}
				redraw = true
			} else {
				// mouse was clicked on a part of the screen represented by no
				// element, so we remove the focus from whatever element had
				// the focus.
				redraw = a.setFocus(ctx, nil)
			}
		}
	case buttonWasDown && !buttonNowDown:
		if a.mouseDragged && target != nil {
			// mouse drag operation has stopped.
			de := mevent.NewDragEvent(ev, a.mouseDownEvent)
			target.MouseDragStop(ctx, de)
			redraw = true
		}
		a.mouseDownEvent = nil
		a.mouseDragged = false
	case !buttonWasDown && !buttonNowDown:
		// mouse move.
		if target != nil {
			f, ok := target.(types.FocusEventHandler)
			if ok {
				if !f.HasFocus() {
					redraw = a.setHover(ctx, target, ev)
				}
			}
		} else {
			redraw = a.setHover(ctx, nil, nil)
		}
	}

	if redraw {
		a.draw(ctx)
	}
}
