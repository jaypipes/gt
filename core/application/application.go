package application

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/charmbracelet/x/ansi"
	"github.com/gdamore/tcell/v3"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/core/box"
	"github.com/jaypipes/gt/core/keypress"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/mouse"
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
	screen     tcell.Screen
	controller *Controller

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
		v = view.New(ctx, a.controller, id)
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
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	err = s.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	c := newController(s)
	a.screen = s
	a.controller = c

	if a.title != "" {
		s.SetTitle(a.title)
	}
	keyMap := a.buildKeyPressMap()
	a.controller.SetKeyPressMap(keyMap)

	s.EnableMouse()
	s.EnableFocus()
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if gtlog.Level() < slog.LevelInfo {
			fmt.Fprintf(os.Stderr, "%s", gtlog.Records())
		}
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	a.draw(ctx)

loop:
	for ev := range s.EventQ() {
		fmt.Fprintf(os.Stderr, "%v", ev)
		switch ev := ev.(type) {
		case *tcell.EventResize:
			a.draw(ctx)
			s.Sync()
		case *tcell.EventKey:
			switch {
			case ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC:
				break loop
			}
			gev := keypress.EventFromTCell(ev)
			if a.controller.HandleKeyPress(ctx, gev) {
				a.draw(ctx)
				// rebuild the key map since we may have changed views.
				keyMap = a.buildKeyPressMap()
				a.controller.SetKeyPressMap(keyMap)
			}
		case *tcell.EventMouse:
			gev := mouse.EventFromTCell(ev)
			pos := gev.Position()
			s.ShowCursor(pos.X, pos.Y)
			v := a.CurrentView()
			node := v.AtPoint(pos)
			if node != nil {
				el, ok := node.(types.Element)
				if ok {
					el.Click(ctx, gev)
					a.draw(ctx)
				}
			}
		}
	}
	s.Fini()
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
	s := a.screen
	if s == nil {
		panic("called Application.draw() with nil screen.")
	}
	v := a.CurrentView()
	if v == nil {
		v = view.New(ctx, a.controller, "main")
		a.views["main"] = v
		a.curView = "main"
	}

	a.Box.Render(ctx, s)
	v.SetBounds(a.InnerBounds())
	v.Render(ctx, s)
	s.Show()
}
