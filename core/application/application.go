package application

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/core/box"
	kpevent "github.com/jaypipes/gt/core/event/keypress"
	mevent "github.com/jaypipes/gt/core/event/mouse"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/view"
	"github.com/jaypipes/gt/types"
)

const (
	defaultEventQueueSize  = 50
	defaultDrawMinInterval = 50 * time.Millisecond
)

// New returns a new Application.
func New(
	ctx context.Context,
) *Application {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	err = s.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	c := newController(s)
	return &Application{
		screen:      s,
		controller:  c,
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
	sync.RWMutex
	box.Box
	screen     tcell.Screen
	controller types.Controller

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
	// pasteEnabled is true if we support bracketed pasting of contents in the
	// terminal.
	pasteEnabled bool
	// focusEnabled is true if we support focus events in the terminal.
	focusEnabled bool

	// lastMouseEvent stores the event from the last mouse action.
	lastMouseEvent types.MouseEvent
	// mouseDownEvent stores the event when the user pressed a mouse button.
	mouseDownEvent types.MouseEvent
}

// Title returns the Application's optional title.
func (a *Application) Title() string {
	return a.title
}

// SetTitle sets the Application's optional title, which by default also sets the
// terminal's screen title.
func (a *Application) SetTitle(title string) {
	a.title = title
}

// EnableMouse enables mouse event handling for the Application.
func (a *Application) EnableMouse() {
	a.mouseEnabled = true
}

// EnablePaste enables bracketed paste for the Application.
func (a *Application) EnablePaste() {
	a.pasteEnabled = true
}

// EnableFocus enables focus events for the Application.
func (a *Application) EnableFocus() {
	a.focusEnabled = true
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
func (a *Application) SetBounds(bounds types.Rectangle) {
	a.Box.SetBounds(bounds)
}

// WithBounds sets the Application's outer bounding box and returns the
// Application.
func (a *Application) WithBounds(bounds types.Rectangle) *Application {
	a.SetBounds(bounds)
	return a
}

// SetBorder sets the Application's border.
func (a *Application) SetBorder(border types.Border) {
	a.Box.SetBorder(border)
}

// WithBorder sets the Application's border and returns the Application.
func (a *Application) WithBorder(border types.Border) *Application {
	a.SetBorder(border)
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
	s := a.screen
	if s == nil {
		return fmt.Errorf("cannot start Application will nil Screen.")
	}
	c := a.controller
	if s == nil {
		return fmt.Errorf("cannot start Application will nil Controller.")
	}

	if a.title != "" {
		s.SetTitle(a.title)
	}
	keyMap := a.buildKeyPressMap()
	a.controller.SetKeyPressMap(keyMap)

	if a.mouseEnabled {
		s.EnableMouse()
	}
	if a.focusEnabled {
		s.EnableFocus()
	}
	if a.pasteEnabled {
		s.EnablePaste()
	}
	// If the user has not overridden the bounds for the Application, we
	// default to the Screen area.
	appBounds := a.Box.Bounds()
	if appBounds.Empty() {
		w, h := s.Size()
		sb := types.Rect(0, 0, w, h)
		gtlog.Debug(
			ctx,
			"Application.Start: no bounds set. defaulting to screen bounds %s",
			sb,
		)
		a.SetBounds(sb)
	}

	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		if s != nil {
			s.Fini()
		}
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
	for {
		ev := <-s.EventQ()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				break loop
			}
			kev := kpevent.New(kpevent.WithTCell(ev))
			if c.HandleKeyPress(ctx, kev) {
				a.draw(ctx)
				// rebuild the key map since we may have changed views.
				keyMap = a.buildKeyPressMap()
				c.SetKeyPressMap(keyMap)
			}
		case *tcell.EventMouse:
			mev := mevent.New(mevent.WithTCell(ev)) //, a.lastMouseEvent, a.mouseDownEvent)
			pos := mev.Position()
			s.ShowCursor(pos.X, pos.Y)
			v := a.CurrentView()
			node := v.AtPoint(pos)
			if node != nil {
				el, ok := node.(types.Element)
				if ok {
					el.Click(ctx, mev)
					a.draw(ctx)
				}
			}
			a.lastMouseEvent = mev
			if mev.Action().MouseDown() {
				a.mouseDownEvent = mev
				gtlog.Debug(
					ctx,
					"mouse down event @%s", mev.Position(),
				)
			}

		case *tcell.EventError:
			return ev
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

	s.Clear()

	a.Box.Render(ctx, s)
	v.SetBounds(a.InnerBounds())
	v.Draw(ctx, s)
	s.Show()
}
