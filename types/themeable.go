package types

// Themeable represents something that can have a Theme
type Themeable interface {
	// Theme returns the Themeable's ThemeClass, if any
	ThemeClass() ThemeClass
	// SetThemeClass sets the Element's theme class.
	SetThemeClass(ThemeClass)
	// Theme returns the Themeable's Theme, if any
	Theme() Theme
	// SetTheme sets the Themeable's Theme.
	SetTheme(Theme)
}
