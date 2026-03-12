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
