package application

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime/debug"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/document"
	"github.com/jaypipes/gt/core/types"
)

// New returns a new Application.
func New(
	ctx context.Context,
) *Application {
	a := &Application{}
	return a
}

// Application wraps the terminal screen and contains the main event-processing
// loop. It is intended to be wrapped in a struct that houses your own
// Application state, like so:
//
//	type MyApplication struct {
//	 	*gt.Application
//	 	myappstate string
//	}
type Application struct {
	term *uv.Terminal

	// name is an optional name for the application, used as a title for the
	// outer containing box for the TUI program.
	name string
	// log is the application-level `log/slog.Logger`
	log *slog.Logger

	// document contains the tree of elements to render the Application to a
	// screen.
	document *document.Document
}

// SetName sets the Application's optional name, which by default also sets the
// terminal's screen title.
func (a *Application) SetName(name string) {
	a.name = name
}

// Document returns the Application's Document.
func (a *Application) Document() *document.Document {
	if a.document == nil {
		a.document = document.New()
	}
	return a.document
}

// SetRoot instructs the Application which Element to put at the root of the
// render tree (the Document).
func (a *Application) SetRoot(el types.Element) {
	d := a.Document()
	d.SetRoot(el)
}

// SetRootWithBounds instructs the Application which Element to put at the
// root of the render tree (the Document) and a bounding box to use for the
// Document.
func (a *Application) SetRootWithBounds(
	el types.Element,
	bounds types.Rectangle,
) {
	d := a.Document()
	d.SetRoot(el)
	d.SetBounds(bounds)
}

// draw renders the Application to the Terminal screen.
func (a *Application) draw(ctx context.Context) {
	if a.term == nil {
		panic("called Application.draw() with nil terminal.")
	}
	doc := a.Document()
	doc.Render(ctx, a.term)
	if err := a.term.Display(); err != nil {
		log.Fatal(err)
	}
}

// Start starts up the Application and its event loop, blocking until the event
// loop is closed.
func (a *Application) Start(ctx context.Context) error {
	if a == nil {
		return fmt.Errorf("cannot start nil Application.")
	}
	t := uv.NewTerminal(os.Stdin, os.Stdout, os.Environ())
	//if a.name != "" {
	//	t.SetTitle(a.name)
	//}

	// By entering alt screen we take control of the output of the terminal
	// which means when we exit the application, the terminal screen will be
	// returned to its original state.
	t.EnterAltScreen()
	defer func() {
		if r := recover(); r != nil {
			_ = t.Teardown()
			fmt.Fprintf(os.Stderr, "recovered from panic: %v", r)
			debug.PrintStack()
		}
	}()

	if err := t.Start(); err != nil {
		return fmt.Errorf("failed to start terminal program: %w", err)
	}

	a.term = t

loop:
	for ev := range t.Events() {
		switch ev := ev.(type) {
		case uv.WindowSizeEvent:
			t.Resize(ev.Width, ev.Height)
			t.Erase()
		case uv.KeyPressEvent:
			switch {
			case ev.MatchString("q", "ctrl+c"):
				break loop
			case ev.MatchString("ctrl+z"):
				t.Erase()
				if err := t.Display(); err != nil {
					log.Fatal(err)
				}
				if t.Shutdown(ctx) != nil {
					log.Fatal("failed to shutdown terminal")
				}

				uv.Suspend()

				goto loop
			}
		}

		a.draw(ctx)
	}

	if err := t.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	return nil
}
