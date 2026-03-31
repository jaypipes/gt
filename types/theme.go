package types

// ThemeClass represents a class of thing that a Theme will style.
//
// There are four built-in ThemeClasses, but since ThemeClass is a string,
// users can create their own ThemeClasses and add greater extensibility to
// their styled Elements and Components.
type ThemeClass string

const (
	// ThemeClassNone is the zero value for ThemeClass
	ThemeClassNone ThemeClass = ""
	// ThemeClassInput is a class of Elements that comprise input and form
	// components. By default, the following Elements have a ThemeClass of
	// ThemeClassInput:
	// * TextArea
	// * RadioButton
	// * Checkbox
	ThemeClassInput = "gt.input"
	// ThemeClassNavigation is for Elements that provide navigation
	// functionality for the user.
	ThemeClassNavigation = "gt.navigation"
	// ThemeClassPrimary is for Elements that make up the primary non-input,
	// non-navigation components of the application.
	ThemeClassPrimary = "gt.primary"
	// ThemeClassSecondary is for Elements that make up the non-primary,
	// non-input, non-navigation components of the application.
	ThemeClassSecondary = "gt.secondary"
)

// Theme describes a set of Motifs, Style and Border properties that can apply
// to different sets of Elements.
type Theme interface {
	// Motif returns the Motif used for styling and borders of things having
	// the supplied ThemeClass.
	Motif(ThemeClass) Motif
	// SetMotif sets the Motif used for styling and borders of things having
	// the supplied ThemeClass.
	SetMotif(ThemeClass, Motif)
	// Style returns the Style used for styling things having the supplied
	// ThemeClass.
	Style(ThemeClass) Style
	// SetStyle sets the Style used for things having the supplied ThemeClass.
	SetStyle(ThemeClass, Style)
	// Border returns the Border used for things having the supplied
	// ThemeClass.
	Border(ThemeClass) Border
	// SetBorder sets the Border used for things having the supplied
	// ThemeClass.
	SetBorder(ThemeClass, Border)
}

// ThemeWithOption describes an optional varg parameter to [theme.New] that
// modifies the returned Theme.
type ThemeWithOption func(Theme)
