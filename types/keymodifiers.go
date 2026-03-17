package types

import (
	"strings"
)

// KeyModifiers is a bitmask for modifier keys
type KeyModifiers uint8

const (
	KeyModifierShift KeyModifiers = 1 << iota
	KeyModifierCtrl
	KeyModifierAlt
	KeyModifierNone = 0
)

// None returns true if none of the KeyModifiers bits are on.
func (m KeyModifiers) None() bool {
	return m == KeyModifierNone
}

// Shift returns true if the KeyModifiers Shift bit is on.
func (m KeyModifiers) Shift() bool {
	return m&KeyModifierShift != 0
}

// Shift returns true if the KeyModifiers Ctrl bit is on.
func (m KeyModifiers) Ctrl() bool {
	return m&KeyModifierCtrl != 0
}

// Shift returns true if the KeyModifiers Alt bit is on.
func (m KeyModifiers) Alt() bool {
	return m&KeyModifierAlt != 0
}

// String returns a string representation of all enabled bits in the
// KeyModifiers bitmask.
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
	return strings.Join(mods, "+")
}
