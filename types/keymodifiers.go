package types

import (
	"strings"
)

// KeyModifiers is a bitmask for modifier keys
type KeyModifiers uint8

const (
	KeyModifierNone  KeyModifiers = 0
	KeyModifierShift              = 1 << iota
	KeyModifierCtrl
	KeyModifierAlt
)

func (m KeyModifiers) String() string {
	if m == KeyModifierNone {
		return ""
	}
	mods := []string{}
	if m&KeyModifierShift != 0 {
		mods = append(mods, "shift")
	}
	if m&KeyModifierCtrl != 0 {
		mods = append(mods, "ctrl")
	}
	if m&KeyModifierAlt != 0 {
		mods = append(mods, "alt")
	}
	return ":" + strings.Join(mods, "+")
}

var (
	KeyModifierStrings = []string{
		"shift",
		"ctrl",
		"alt",
	}
)
