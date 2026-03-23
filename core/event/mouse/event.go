package mouse

import (
	"fmt"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling mouse-related events.
// Implements [types.MouseEvent].
type Event struct {
	*event.Event
	core.KeyModifiable
	// pos contains the coordinates of the mouse when the event fired.
	pos types.Point
	// button is the mouse button that was pressed, if any.
	button types.MouseButton
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	button := e.button.String()
	pos := e.pos.String()
	mods := e.Modifiers().String()
	return fmt.Sprintf(
		"mouse:%s%s@%s",
		button, mods, pos,
	)
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
	case bm&tcell.Button4 != 0:
		return types.MouseButtonForward
	case bm&tcell.Button5 != 0:
		return types.MouseButtonBackward
	}
	return types.MouseButtonNone
}

var _ tcell.Event = (*Event)(nil)
var _ types.MouseEvent = (*Event)(nil)
