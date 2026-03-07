package mouse

import (
	"fmt"

	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling mouse-related events.
// Implements [types.MouseEvent].
type Event struct {
	event.Event
	// modifiers are the modifier keys that were held down
	modifiers types.KeyModifiers
	// pos contains the coordinates of the mouse when the event fired.
	pos types.Point
	// button is the mouse button that was pressed, if any.
	button types.MouseButton
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	button := e.button.String()
	pos := e.pos.String()
	mods := e.modifiers.String()
	return fmt.Sprintf(
		"mouse:%s%s@%s",
		button, mods, pos,
	)
}

// Shift returns true if the Shift modifier key was held.
func (e *Event) Shift() bool {
	return tcell.ModMask(e.modifiers)&tcell.ModShift != 0
}

// Ctrl returns true if the Ctrl modifier key was held.
func (e *Event) Ctrl() bool {
	return tcell.ModMask(e.modifiers)&tcell.ModCtrl != 0
}

// Alt returns true if the Alt modifier key was held.
func (e *Event) Alt() bool {
	return tcell.ModMask(e.modifiers)&tcell.ModAlt != 0
}

// Position returns the coordinates of the mouse when the event fired.
func (e *Event) Position() types.Point {
	return e.pos
}

// Button returns the mouse button that was depressed when the event fired.
func (e *Event) Button() types.MouseButton {
	return e.button
}

// EventFromTCell returns an Event from a [tcell.EventMouse]
func EventFromTCell(
	te *tcell.EventMouse,
) *Event {
	x, y := te.Position()
	mods := te.Modifiers()
	e := &Event{
		Event:     event.New(),
		modifiers: types.KeyModifiers(mods),
		pos: types.Point{
			X: x, Y: y,
		},
		button: mouseButtonFromTCell(te.Buttons()),
	}
	e.SetWhen(te.When())

	return e
}

// mouseButtonFromTCell translates a tcell ButtonMask to a single MouseButton.
func mouseButtonFromTCell(bm tcell.ButtonMask) types.MouseButton {
	switch {
	case bm&tcell.ButtonPrimary != 0:
		return types.MouseLeft
	case bm&tcell.ButtonSecondary != 0:
		return types.MouseRight
	case bm&tcell.ButtonMiddle != 0:
		return types.MouseMiddle
	case bm&tcell.WheelUp != 0:
		return types.MouseWheelUp
	case bm&tcell.WheelDown != 0:
		return types.MouseWheelDown
	}
	return types.MouseNone
}

var _ tcell.Event = (*Event)(nil)
var _ types.MouseEvent = (*Event)(nil)
