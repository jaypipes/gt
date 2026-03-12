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
	// focused contains the thing that currently has the focus.
	focused types.Focusable

	// lastMouseEvent stores the event from the last mouse action.
	lastMouseEvent types.MouseEvent
	// lastMouseClickTime stores the time when the last mouse click happened.
	lastMouseClickTime time.Time
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

// setFocus sets the currently-focused thing and calls Focus(false) on the
// previously-focused thing, returning whether there was a change in focus.
func (a *Application) setFocus(ctx context.Context, f types.Focusable) bool {
	if a.focused != nil {
		if f == nil {
			a.focused.SetFocus(ctx, false)
			a.focused = nil
			return true
		}
		if f.HasFocus() {
			// already has the focus, no need to do anything...
			return false
		}
		a.focused.SetFocus(ctx, false)
	} else {
		if f == nil {
			return false
		}
	}
	if f != nil {
		f.SetFocus(ctx, true)
	}
	a.focused = f
	return true
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
			redraw := a.handleMouseEvent(ctx, mev)
			if redraw {
				a.draw(ctx)
			}
		case *tcell.EventError:
			return ev
		}
	}

	return nil
}

// NOTE(jaypipes): Some of this code adapted from:
// https://github.com/rivo/tview/blob/f39b95c73dbb30877f4b5145b835333002afb2a8/application.go

// handleMouseEvent determines what logical action the user took with the mouse
// and executes the appropriate mouse event handler for the target element.
// The method returns a bool indicating whether the screen should be redrawn.
func (a *Application) handleMouseEvent(
	ctx context.Context,
	ev types.MouseEvent,
) bool {
	// First, we determine if the mouse is over a thing that can handle to
	// mouse events. If the mouse is over something that can handle mouse
	// events and is not disabled, we set the focus on that thing.
	var target types.MouseEventHandler

	pos := ev.Position()
	v := a.CurrentView()
	node := v.AtPoint(pos)
	changedFocus := false
	if node != nil {
		el, ok := node.(types.Element)
		if ok && !el.Disabled() {
			target = el.(types.MouseEventHandler)
			changedFocus = a.setFocus(ctx, el.(types.Focusable))
		}
	} else {
		changedFocus = a.setFocus(ctx, nil)
	}

	// Next, we determine what logical mouse action the user has taken by
	// examining our stored state of previous mouse events.
	x, y := pos.X, pos.Y
	button := ev.Button()
	lastX, lastY := x, y
	if a.lastMouseEvent != nil {
		lastPos := a.lastMouseEvent.Position()
		lastX, lastY = lastPos.X, lastPos.Y
	}

	buttonWasDown := a.mouseDownEvent != nil
	buttonNowDown := button.Pressable()

	if x != lastX || y != lastY {
		// The mouse has moved. If a mouse button had previously been down and
		// is now *not* down, we fire off a MouseDragStop event. If the mouse
		// button had previously *not* been down and is now down, we save the
		// current mouse event as the "mouse down" event to use in later
		// constructing a MouseDragStopEvent.
		if buttonWasDown {
			if !buttonNowDown {
				if target != nil {
					de := mevent.NewDragStopEvent(ev, a.lastMouseEvent)
					target.MouseDragStop(ctx, de)
				}
				a.lastMouseEvent = ev
				a.mouseDownEvent = nil
				return true
			} else {
				if target != nil {
					de := mevent.NewDragMoveEvent(ev, a.lastMouseEvent)
					target.MouseDragMove(ctx, de)
				}
				a.lastMouseEvent = ev
				a.mouseDownEvent = ev
				return true
			}
		}
	}

	downX, downY := x, y
	if buttonWasDown {
		downPos := a.mouseDownEvent.Position()
		downX, downY = downPos.X, downPos.Y
	}

	clickMoved := x != downX || y != downY

	if !buttonWasDown {
		if buttonNowDown {
			a.mouseDownEvent = ev
			if !clickMoved {
				if a.lastMouseClickTime.Add(types.DefaultMouseDoubleClickInterval).Before(time.Now()) {
					if target != nil {
						ce := mevent.NewClickEvent(ev, false)
						target.MouseClick(ctx, ce)
					}
					a.lastMouseClickTime = time.Now()
				} else {
					if target != nil {
						ce := mevent.NewClickEvent(ev, true)
						target.MouseClick(ctx, ce)
						a.lastMouseClickTime = time.Time{}
					}
				}
			}
		}
	}

	a.lastMouseEvent = ev
	return changedFocus
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
