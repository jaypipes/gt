package application

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"runtime/debug"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/core/box"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/view"
	"github.com/jaypipes/gt/types"
)

// New returns a new Application.
func New(
	ctx context.Context,
) *Application {
	b := box.New(ctx)
	b.SetID("gt.app")
	return &Application{
		Box:   b,
		views: map[string]*view.View{},
	}
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
	box.Box
	term *uv.Terminal

	// title is an optional title for the application, used as a title for the
	// terminal when set.
	title string

	// views is a map, keyed by the View ID, of Views that the Application is
	// managing.
	views map[string]*view.View
	// curView is the ID of the currently active (displayed) View.
	curView string
}

// SetTitle sets the Application's optional title, which by default also sets the
// terminal's screen title.
func (a *Application) SetTitle(title string) {
	a.title = title
}

// Title returns the Application's optional title.
func (a *Application) Title() string {
	return a.title
}

// View returns the View with the supplied ID. If no such View exists, a new
// empty View with that ID is returned.
func (a *Application) View(ctx context.Context, id string) *view.View {
	v, ok := a.views[id]
	if !ok {
		v = view.New(ctx, id)
		a.views[id] = v
		a.curView = id
	}
	return v
}

// Views returns the collection of the Application's Views.
func (a *Application) Views() []*view.View {
	return lo.Values(a.views)
}

// CurrentView returns the currently active (displaying) View.
func (a *Application) CurrentView() *view.View {
	return a.views[a.curView]
}

// SetCurrentView sets the currently active (displaying) View.
func (a *Application) SetCurrentView(id string) *Application {
	a.curView = id
	return a
}

// SetBounds sets the View's outer bounding box.
func (a *Application) SetBounds(bounds types.Rectangle) *Application {
	a.Box.SetBounds(bounds)
	return a
}

// SetRect sets the Element's bounding rectangle
func (a *Application) SetBorder(border types.Border) *Application {
	a.Box.SetBorder(border)
	return a
}

// SetBorderForegroundColor sets the Application's border foreground color (i.e
// the color of the border cell's underlying grapheme).
func (a *Application) SetBorderForegroundColor(c types.Color) *Application {
	a.Box.SetBorderForegroundColor(c)
	return a
}

// SetBorderBackgroundColor sets the Application's border background color (i.e
// the background color of the border's cells.
func (a *Application) SetBorderBackgroundColor(c types.Color) *Application {
	a.Box.SetBorderBackgroundColor(c)
	return a
}

// draw renders the Application's active View to the Terminal screen.
func (a *Application) draw(ctx context.Context) {
	if a.term == nil {
		panic("called Application.draw() with nil terminal.")
	}
	v := a.CurrentView()
	if v == nil {
		v = view.New(context.TODO(), "main")
		a.views["main"] = v
		a.curView = "main"
	}

	// If the Application has had no bounds set, adopt the screen's max width and
	// height.
	bounds := a.Bounds()
	if bounds.Empty() {
		screenBounds := a.term.Bounds()
		gtlog.Debug(
			ctx,
			"Application.draw: setting application bounds to screen bounds %s",
			screenBounds,
		)
		bounds = screenBounds
		a.Box.SetBounds(bounds)
	}

	a.Box.Draw(a.term, a.Bounds())
	v.SetBounds(a.InnerBounds())
	v.Render(ctx, a.term)
	if err := a.term.Display(); err != nil {
		log.Fatal(err)
	}
}

func (a *Application) currentViewKeyMap() map[string]string {
	res := map[string]string{}
	for viewID, view := range a.views {
		currentKP := view.CurrentViewKeyPress()
		if currentKP != "" {
			res[currentKP] = viewID
		}
	}
	return res
}

// Start starts up the Application and its event loop, blocking until the event
// loop is closed.
func (a *Application) Start(ctx context.Context) error {
	if a == nil {
		return fmt.Errorf("cannot start nil Application.")
	}
	t := uv.NewTerminal(os.Stdin, os.Stdout, os.Environ())

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
	if a.title != "" {
		t.WriteString(ansi.SetWindowTitle(a.title))
	}
	currentViewKeyMap := a.currentViewKeyMap()

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
				if t.Pause() != nil {
					log.Fatal("failed to pause terminal")
				}

				uv.Suspend()

				if err := t.Resume(); err != nil {
					log.Fatal("failed to resume terminal")
				}
			}
			for viewKP, viewID := range currentViewKeyMap {
				if ev.MatchString(viewKP) && a.curView != viewID {
					a.SetCurrentView(viewID)
				}
			}
		}

		a.draw(ctx)
	}

	if err := t.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	if gtlog.Level() < slog.LevelInfo {
		fmt.Fprintf(os.Stderr, "%s", gtlog.Records())
	}
	return nil
}
