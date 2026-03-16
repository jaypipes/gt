package core

import (
	"github.com/jaypipes/gt/types"
)

// KeyModifiable stores key modifiers and implements [types.KeyModifiable]
type KeyModifiable struct {
	// modifiers are the modifier keys that were held down
	modifiers types.KeyModifiers
}

// Modifiers returns the key modifier bitmask.
func (m *KeyModifiable) Modifiers() types.KeyModifiers {
	return m.modifiers
}

// SetModifiers sets the key modifier bitmask.
func (m *KeyModifiable) SetModifiers(mods types.KeyModifiers) {
	m.modifiers = mods
}

// Shift returns true if the Shift modifier key was held.
func (m *KeyModifiable) Shift() bool {
	return m.modifiers&types.KeyModifierShift != 0
}

// Ctrl returns true if the Ctrl modifier key was held.
func (m *KeyModifiable) Ctrl() bool {
	return m.modifiers&types.KeyModifierCtrl != 0
}

// Alt returns true if the Alt modifier key was held.
func (m *KeyModifiable) Alt() bool {
	return m.modifiers&types.KeyModifierAlt != 0
}

var _ types.KeyModifiable = (*KeyModifiable)(nil)
