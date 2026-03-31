package theme

import "github.com/jaypipes/gt/types"

// Theme describes a set of Motifs, Style and Border properties that can apply
// to different sets of Elements.
type Theme struct {
	// motifs is the set of Motifs to use for different ThemeClasses.
	motifs map[types.ThemeClass]types.Motif
	// styles is the set of Styles to use for different ThemeClasses.
	styles map[types.ThemeClass]types.Style
	// borders is the set of Borders to use for different ThemeClasses.
	borders map[types.ThemeClass]types.Border
}

// Motif returns the Motif used for styling and borders of things having the
// supplied ThemeClass.
func (t *Theme) Motif(class types.ThemeClass) types.Motif {
	return t.motifs[class]
}

// SetMotif sets the Motif used for styling and borders of things having the
// supplied ThemeClass.
func (t *Theme) SetMotif(class types.ThemeClass, m types.Motif) {
	t.motifs[class] = m
}

// Style returns the Style used for styling of things having the supplied
// ThemeClass.
func (t *Theme) Style(class types.ThemeClass) types.Style {
	return t.styles[class]
}

// SetStyle sets the Style used for styling of things having the supplied
// ThemeClass.
func (t *Theme) SetStyle(class types.ThemeClass, m types.Style) {
	t.styles[class] = m
}

// Border returns the Border used for borders of things having the supplied
// ThemeClass.
func (t *Theme) Border(class types.ThemeClass) types.Border {
	return t.borders[class]
}

// SetBorder sets the Border used for borders of things having the supplied
// ThemeClass.
func (t *Theme) SetBorder(class types.ThemeClass, m types.Border) {
	t.borders[class] = m
}

var _ types.Theme = (*Theme)(nil)
