package types

import "image/color"

// ScreenController represents something that controls screen and cursor
// display.
type ScreenController interface {
	// Returns the Cursor for the Screen.
	// Cursor() Cursor
	// CursorStyle returns the Screen cursor's shape and blink.
	CursorStyle() (CursorShape, bool)
	// CursorVisible returns whether the Screen's cursor is visible.
	CursorVisible() bool
	// CursorColor returns the color of the Screen's cursor.
	CursorColor() color.Color
	// SetCursorPosition sets the position of the Screen's cursor.
	SetCursorPosition(int, int) error
	// SetCursorStyle sets the shape and blink of the Screen's cursor.
	SetCursorStyle(CursorShape, bool) error
	// SetCursorColor sets the color of the Screen's cursor.
	SetCursorColor(color.Color) error
	// ShowCursor makes the Screen's cursor visible.
	ShowCursor() error
	// HideCursor hides the Screen's cursor.
	HideCursor() error
}

// ScreenControllable has a ScreenController
type ScreenControllable interface {
	// ScreenController returns the screen controller.
	ScreenController() ScreenController
	// SetScreenController sets the screen controller.
	SetScreenController(ScreenController)
}
