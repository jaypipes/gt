package core

import (
	"github.com/gdamore/tcell/v3"
	"github.com/jaypipes/gt/types"
)

// KeyModifiable stores key modifiers and implements [types.KeyModifiable]
type KeyModifiable struct {
	// modifiers are the modifier keys that were held down
	modifiers types.KeyModifiers
}

// KeyModifiers returns the key modifier bitmask.
func (m *KeyModifiable) KeyModifiers() types.KeyModifiers {
	return m.modifiers
}

// SetKeyModifiers sets the key modifier bitmask.
func (m *KeyModifiable) SetKeyModifiers(mods types.KeyModifiers) {
	m.modifiers = mods
}

// Shift returns true if the Shift modifier key was held.
func (m *KeyModifiable) Shift() bool {
	return tcell.ModMask(m.modifiers)&tcell.ModShift != 0
}

// Ctrl returns true if the Ctrl modifier key was held.
func (m *KeyModifiable) Ctrl() bool {
	return tcell.ModMask(m.modifiers)&tcell.ModCtrl != 0
}

// Alt returns true if the Alt modifier key was held.
func (m *KeyModifiable) Alt() bool {
	return tcell.ModMask(m.modifiers)&tcell.ModAlt != 0
}

var _ types.KeyModifiable = (*KeyModifiable)(nil)
