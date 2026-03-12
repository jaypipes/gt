package mouse

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// NOTE(jaypipes): Some of this code adapted from:
// https://github.com/rivo/tview/blob/f39b95c73dbb30877f4b5145b835333002afb2a8/application.go

// Event exposes an easy-to-use interface for handling mouse-related events.
// Implements [types.MouseEvent].
type Event struct {
	event.Event
	core.KeyModifiable
	// pos contains the coordinates of the mouse when the event fired.
	pos types.Point
	// button is the mouse button that was pressed, if any.
	button types.MouseButton
	// action is the semantic action that was taken by the user.
	action types.MouseAction
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	action := e.action.String()
	button := e.button.String()
	pos := e.pos.String()
	mods := e.KeyModifiers().String()
	return fmt.Sprintf(
		"mouse:%s:%s%s@%s",
		action, button, mods, pos,
	)
}

// Action returns the semantic action taken by the user.
func (e *Event) Action() types.MouseAction {
	return e.action
}

// SetAction sets the semantic action taken by the user.
func (e *Event) SetAction(action types.MouseAction) {
	e.action = action
}

// Position returns the coordinates of the mouse when the event fired.
func (e *Event) Position() types.Point {
	return e.pos
}

// SetPosition sets the coordinates of the mouse when the event fired.
func (e *Event) SetPosition(pos types.Point) {
	e.pos = pos
}

// Button returns the mouse button that was depressed when the event fired.
func (e *Event) Button() types.MouseButton {
	return e.button
}

// SetButton sets the mouse button that was depressed when the event fired.
func (e *Event) SetButton(button types.MouseButton) {
	e.button = button
}

// mouseButtonFromTCell translates a tcell ButtonMask to a single MouseButton.
func mouseButtonFromTCell(bm tcell.ButtonMask) types.MouseButton {
	switch {
	case bm&tcell.ButtonPrimary != 0:
		return types.MouseButtonPrimary
	case bm&tcell.ButtonSecondary != 0:
		return types.MouseButtonSecondary
	case bm&tcell.ButtonMiddle != 0:
		return types.MouseButtonMiddle
	case bm&tcell.WheelUp != 0:
		return types.MouseWheelUp
	case bm&tcell.WheelDown != 0:
		return types.MouseWheelDown
	}
	return types.MouseButtonNone
}

// mouseActionFromEvents calculates the MouseAction given a current mouse event
// and a mouse event representing the moment the user pressed a mouse button.
func mouseActionFromEvents(
	curEvent types.MouseEvent,
	lastMouseEvent types.MouseEvent,
	mouseDownEvent types.MouseEvent,
) types.MouseAction {
	curPos := curEvent.Position()
	curX, curY := curPos.X, curPos.Y
	curButton := curEvent.Button()
	lastButton := curButton
	lastX, lastY := curX, curY
	if lastMouseEvent != nil {
		lastPos := lastMouseEvent.Position()
		lastX, lastY = lastPos.X, lastPos.Y
		lastButton = lastMouseEvent.Button()
	}

	if curX != lastX || curY != lastY {
		return types.MouseActionMove
	}

	downX, downY := -1, -1
	if mouseDownEvent != nil {
		downPos := mouseDownEvent.Position()
		downX, downY = downPos.X, downPos.Y
	}

	clickMoved := curX != downX || curY != downY
	buttonChanged := curButton != lastButton

	if buttonChanged {
		lastClickTime := time.Time{}
		if lastMouseEvent != nil {
			lastClickTime = lastMouseEvent.When()
		}
		for _, buttonEvent := range []struct {
			button                  types.MouseButton
			down, up, click, dclick types.MouseAction
		}{
			{types.MouseButtonPrimary, types.MouseActionLeftDown, types.MouseActionLeftUp, types.MouseActionLeftClick, types.MouseActionLeftDoubleClick},
			{types.MouseButtonMiddle, types.MouseActionMiddleDown, types.MouseActionMiddleUp, types.MouseActionMiddleClick, types.MouseActionMiddleDoubleClick},
			{types.MouseButtonSecondary, types.MouseActionRightDown, types.MouseActionRightUp, types.MouseActionRightClick, types.MouseActionRightDoubleClick},
		} {
			if curButton == buttonEvent.button {
				return buttonEvent.down
			} else {
				if !clickMoved {
					if lastClickTime.Add(types.DefaultMouseDoubleClickInterval).Before(time.Now()) {
						return buttonEvent.click
					} else {
						return buttonEvent.dclick
					}
				}
				return buttonEvent.up
			}
		}
	}

	return types.MouseActionNone
}

var _ tcell.Event = (*Event)(nil)
var _ types.MouseEvent = (*Event)(nil)
