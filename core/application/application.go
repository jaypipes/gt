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
	"github.com/charmbracelet/x/termios"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/core/box"
	"github.com/jaypipes/gt/core/eventloop"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/view"
	"github.com/jaypipes/gt/types"
)

// New returns a new Application.
func New(
	ctx context.Context,
) *Application {
	return &Application{
		views:       map[string]*view.View{},
		keyPressMap: types.KeyPressMap{},
		events:      eventloop.New(ctx),
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

	// keyPressMap contains key press combination callbacks registered for the
	// Application itself -- i.e. global key press callbacks.
	keyPressMap types.KeyPressMap

	// events is the event loop for Application events (separate from the event
	// loop that the Application's Terminal runs)
	events *eventloop.EventLoop

	// mouseEnabled is true if we're trapping mouse events in the terminal.
	mouseEnabled bool
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

// EnableMouse enables mouse event handling for the Application.
func (a *Application) EnableMouse() {
	a.mouseEnabled = true
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

// KeyPressMap returns the Application's *global* map of key press
// combination strings to callbacks that will execute when that key press
// combination is entered.
func (a *Application) KeyPressMap() types.KeyPressMap {
	return a.keyPressMap
}

// OnKeyPress registers an Application-level (global)  callback to execute
// upon a key press combination.
func (a *Application) OnKeyPress(key string, cb types.EventCallback) {
	a.keyPressMap[key] = cb
}

// Start starts up the Application and its event loop, blocking until the event
// loop is closed.
func (a *Application) Start(ctx context.Context) error {
	if a == nil {
		return fmt.Errorf("cannot start nil Application.")
	}

	// Start our Application's internal event loop.
	a.events.Start()

	t := uv.DefaultTerminal()

	scr := t.Screen()
	scr.ShowCursor()

	// By entering alt screen we take control of the output of the terminal
	// which means when we exit the application, the terminal screen will be
	// returned to its original state.
	scr.EnterAltScreen()
	defer func() {
		if r := recover(); r != nil {
			_ = t.Stop()
			fmt.Fprintf(os.Stderr, "recovered from panic: %v\n", r)
			debug.PrintStack()
			if gtlog.Level() < slog.LevelInfo {
				fmt.Fprintf(os.Stderr, "%s", gtlog.Records())
			}
		}
	}()

	if err := t.Start(); err != nil {
		return fmt.Errorf("failed to start terminal program: %w", err)
	}

	modes := a.terminalModes()
	if len(modes) > 0 {
		scr.WriteString(ansi.SetMode(modes...))
	}

	a.term = t
	if a.title != "" {
		scr.WriteString(ansi.SetWindowTitle(a.title))
	}
	keyMap := a.buildKeyPressMap()

	a.draw(ctx)
loop:
	for ev := range t.Events() {
		switch ev := ev.(type) {
		case uv.WindowSizeEvent:
			scr.Resize(ev.Width, ev.Height)
			a.draw(ctx)
		case uv.KeyPressEvent:
			switch {
			case ev.MatchString("q", "ctrl+c"):
				break loop
			case ev.MatchString("ctrl+z"):
				if err := scr.Flush(); err != nil {
					log.Fatal(err)
				}

				uv.Suspend()

				if err := scr.Restore(); err != nil {
					log.Fatal(err)
				}
			}
			for kp, cb := range keyMap {
				if ev.MatchString(kp) {
					cb(ctx)
					a.draw(ctx)
					// rebuild the key map since we may have changed views.
					keyMap = a.buildKeyPressMap()
					break
				}
			}
		case uv.MouseClickEvent:
			m := ev.Mouse()
			scr.SetCursorPosition(m.X, m.Y)
			cur := scr.CellAt(m.X, m.Y)
			if cur == nil {
				// outside the screen bounds...
				break
			}
			pos := types.Point{X: m.X, Y: m.Y}
			v := a.CurrentView()
			node := v.AtPoint(pos)
			if node != nil {
				el, ok := node.(types.Element)
				if ok {
					el.Click(ctx, ev)
					a.draw(ctx)
				}
			}
		}
	}

	if len(modes) > 0 {
		scr.WriteString(ansi.ResetMode(modes...))
	}

	a.events.Stop()

	if err := t.Stop(); err != nil {
		log.Fatal(err)
	}
	if gtlog.Level() < slog.LevelInfo {
		fmt.Fprintf(os.Stderr, "%s", gtlog.Records())
	}
	return nil
}

// terminalModes returns the ANSI mode flags to set up the terminal.
func (a *Application) terminalModes() []ansi.Mode {
	if a.mouseEnabled {
		return []ansi.Mode{
			ansi.ButtonEventMouseMode,
			ansi.SgrExtMouseMode,
			ansi.FocusEventMode,
		}
	}
	return nil
}

// buildKeyPressMap builds the Application's outermost map of key press
// combinations to callback functions to execute when those key press
// combinations are entered.
//
// The outermost map will always be the "current view" key press combinations
// that the Application's registered Views have along with any key press
// combinations registered with the Application itself and any key press
// combinations that the *current* View contains.
func (a *Application) buildKeyPressMap() types.KeyPressMap {
	ctx := context.TODO()
	res := types.KeyPressMap{}

	// copy in our global key press callbacks
	for k, cb := range a.keyPressMap {
		res[k] = cb
	}
	globalKPs := lo.Keys(a.keyPressMap)

	// now add our "current view" key press callbacks
	for viewID, view := range a.views {
		currentViewKP := view.CurrentViewKeyPress()
		if currentViewKP != "" {
			if lo.Contains(globalKPs, currentViewKP) {
				gtlog.Warn(
					ctx,
					"current view key press combination %q for view %q "+
						"shadows global key press combination",
					currentViewKP, viewID,
				)
			}
			res[currentViewKP] = func(_ context.Context) {
				a.SetCurrentView(viewID)
			}
		}
	}

	// finally, add all the current View's key press callbacks
	curView := a.views[a.curView]
	curViewKPMap := curView.KeyPressMap()
	if len(curViewKPMap) > 0 {
		appKPs := lo.Keys(res)
		for k, cb := range curViewKPMap {
			if lo.Contains(appKPs, k) {
				gtlog.Warn(
					ctx,
					"view key press combination %q for view %q "+
						"shadows application key press combination",
					k, curView.ID(),
				)
			}
			res[k] = cb
		}
	}

	gtlog.Debug(
		ctx,
		"Application.buildKeyPressMap: built map for combinations %v",
		lo.Keys(res),
	)
	return res
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

	scr := a.term.Screen()

	// If the Application has had no bounds set, adopt the screen's max width and
	// height.
	bounds := a.Bounds()
	if bounds.Empty() {
		winSize, err := termios.GetWinsize(int(os.Stdin.Fd()))
		if err != nil {
			panic(err.Error())
		}
		screenBounds := types.Rectangle{Min: types.Point{X: 0, Y: 0}, Max: types.Point{X: int(winSize.Col), Y: int(winSize.Row)}}
		bounds = screenBounds
		a.SetBounds(bounds)
		gtlog.Debug(
			ctx,
			"Application.draw: setting application bounds to screen bounds %s",
			screenBounds,
		)
		bounds = screenBounds
		a.Box.SetBounds(bounds)
	}

	a.Box.Draw(scr, a.Bounds())
	v.SetBounds(a.InnerBounds())
	v.Render(ctx, scr)
	if err := scr.Render(); err != nil {
		log.Fatal(err)
	}
	if err := scr.Flush(); err != nil {
		log.Fatal(err)
	}
}
