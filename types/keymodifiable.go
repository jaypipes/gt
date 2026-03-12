package types

// KeyModifiable describes something modified by a modifier key (e.g. "Ctrl" or
// "Shift")
type KeyModifiable interface {
	// KeyModifiers returns the key modifier bitmask.
	KeyModifiers() KeyModifiers
	// SetKeyModifiers sets the KeyModifiable's key modifier bitmask.
	SetKeyModifiers(KeyModifiers)
	// Ctrl returns true if the Ctrl modifier key was held.
	Ctrl() bool
	// Shift returns true if the Shift modifier key was held.
	Shift() bool
	// Alt returns true if the Alt modifier key was held.
	Alt() bool
}
