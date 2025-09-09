package application

import (
	"context"
	"fmt"
	"log"
	"os"

	uv "github.com/charmbracelet/ultraviolet"
)

// Application wraps the terminal screen and contains the main event-processing
// loop. It is intended to be wrapped in a struct that houses your own
// Application state, like so:
//
//	type MyApplication struct {
//	    *gt.Application
//	    myappstate string
//	}
//
//	gtapp := gt.NewApplication()
//	myapp := MyApplication{gtapp}
//
//	if err := myapp.Start(); err != nil {
//	    log.Fatal(err)
//	}
type Application struct {
	term *uv.Terminal

	// optional name of the application, used as a title for the outer
	// containing box for the TUI program.
	appName string

	// root is the top-level renderable element in the Application. This is a
	// Box that consumes the entire screen real-estate. Use WithRoot to
	// override this default renderable element.
	root uv.Drawable
	// rootBounds is the optional bounding box to draw the root renderable
	// element on the screen. If nil, the terminal's bounding box is used.
	rootBounds *uv.Rectangle
}

// SetName sets the Application's optional
func (a *Application) SetName(name string) {
	a.appName = name
}

// SetRootWithBounds sets the Application's top-level renderable element with a
// bounding box.
func (a *Application) SetRootWithBounds(
	root uv.Drawable,
	bounds uv.Rectangle,
) {
	a.root = root
	a.rootBounds = &bounds
}

// SetRoot sets the Application's top-level renderable element.
func (a *Application) SetRoot(root uv.Drawable) {
	a.root = root
}

// draw renders the Application to the Terminal screen.
func (a *Application) draw() {
	if a.term == nil {
		panic("called Application.draw() with nil terminal.")
	}
	bounds := a.term.Bounds()
	if a.rootBounds != nil {
		bounds = *a.rootBounds
	}
	a.root.Draw(a.term, bounds)
	if err := a.term.Display(); err != nil {
		log.Fatal(err)
	}
}

// Start starts up the Application and its event loop, blocking until the event
// loop is closed.
func (a *Application) Start() error {
	if a == nil {
		return fmt.Errorf("cannot start nil Application.")
	}
	t := uv.NewTerminal(os.Stdin, os.Stdout, os.Environ())
	if a.appName != "" {
		t.SetTitle(a.appName)
	}

start:
	// By entering alt screen we take control of the output of the terminal
	// which means when we exit the application, the terminal screen will be
	// returned to its original state.
	t.EnterAltScreen()

	if err := t.Start(); err != nil {
		return fmt.Errorf("failed to start terminal program: %w", err)
	}

	a.term = t

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	evch := make(chan uv.Event)
	go func() {
		defer close(evch)
		_ = t.StreamEvents(ctx, evch)
	}()

	for ev := range evch {
		switch ev := ev.(type) {
		case uv.WindowSizeEvent:
			t.Resize(ev.Width, ev.Height)
			t.Erase()
		case uv.KeyPressEvent:
			if ev.MatchStrings("q", "ctrl+c") {
				cancel() // This will stop the loop
			} else if ev.MatchString("ctrl+z") {
				t.Erase()
				if err := t.Display(); err != nil {
					log.Fatal(err)
				}
				if t.Shutdown(ctx) != nil {
					log.Fatal("failed to shutdown terminal")
				}

				cancel()

				uv.Suspend()

				goto start
			}
		}

		a.draw()
	}

	if err := t.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	return nil
}
