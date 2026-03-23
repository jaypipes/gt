package types

import "github.com/gdamore/tcell/v3"

type Screen = tcell.Screen

// ScreenHandler provides an easy-to-use interface for authors of Elements and
// Components to draw to the Screen and control the Cursor and Events.
type ScreenHandler interface {
	// Screen returns the tcell.Screen object.
	Screen() Screen
	// Cursor returns the Cursor managed by ScreenHandler.
	Cursor() Cursor
	// SetCursor sets the Cursor managed by ScreenHandler.
	SetCursor(Cursor)
}
