package types

import (
	"strings"

	"github.com/gdamore/tcell/v3"
)

// KeyModifiers is a bitmask for modifier keys
type KeyModifiers tcell.ModMask

func (mods KeyModifiers) String() string {
	m := tcell.ModMask(mods)
	if m == tcell.ModNone {
		return ""
	}
	res := []string{}
	if m&tcell.ModCtrl != 0 {
		res = append(res, "ctrl")
	}
	if m&tcell.ModShift != 0 {
		res = append(res, "shift")
	}
	if m&tcell.ModAlt != 0 {
		res = append(res, "alt")
	}
	return ":" + strings.Join(res, "+")
}

var (
	KeyModifierStrings = []string{
		"shift",
		"ctrl",
		"alt",
	}
)

// KeyModifiable describes something modified by a modifier key (e.g. "Ctrl" or
// "Shift")
type KeyModifiable interface {
	// Ctrl returns true if the Ctrl modifier key was held.
	Ctrl() bool
	// Shift returns true if the Shift modifier key was held.
	Shift() bool
	// Alt returns true if the Alt modifier key was held.
	Alt() bool
}
