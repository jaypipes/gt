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
	sevent "github.com/jaypipes/gt/core/event/scroll"
	"github.com/jaypipes/gt/core/key"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/view"
	"github.com/jaypipes/gt/types"
)

const (
	defaultEventQueueSize  = 50
	defaultDrawMinInterval = 50 * time.Millisecond
)

var (
	defaultExitKey = key.New("ctrl+c")
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
	return &Application{
		screen:   s,
		exitKeys: []types.Key{defaultExitKey},
		views:    map[string]*view.View{},
		keyMap:   types.KeyMap{},
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
	screen tcell.Screen

	// title is an optional title for the application, used as a title for the
	// terminal when set.
	title string

	// views is a map, keyed by the View ID, of Views that the Application is
	// managing.
	views map[string]*view.View
	// curView is the ID of the currently active (displayed) View.
	curView string

	// exitKeys contains the supplied exit key combinations. These
	// key combinations are always evaluated first when a key event
	// is received by the Application.
	//
	// If no exit key combinations are set for the Application, it
	// defaults to "Ctrl+C".
	exitKeys []types.Key
	// keyMap contains key press combination callbacks registered for the
	// Application itself -- i.e. global key press callbacks.
	keyMap types.KeyMap
	// keyInterceptor points to a types.KeyPressEventHandler that receives all
	// keyboard input after InterceptKey has been called.
	keyInterceptor types.KeyPressEventHandler
	// keyInterceptEscape is the key combination that will trigger the
	// interceptor to be removed.
	keyInterceptEscape types.Key

	// mouseEnabled is true if we're trapping mouse events in the terminal.
	mouseEnabled bool
	// pasteEnabled is true if we support bracketed pasting of contents in the
	// terminal.
	pasteEnabled bool
	// focusEnabled is true if we support focus events in the terminal.
	focusEnabled bool
	// focused contains the thing that currently has the focus.
	focused types.FocusEventHandler
	// hovered contains the thing that the mouse is currently over.
	hovered types.MouseEventHandler

	// mouseDragged is true if a mouse button was pressed and the mouse moved.
	mouseDragged bool
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

	if a.title != "" {
		s.SetTitle(a.title)
	}

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
			kev := kpevent.New(kpevent.WithTCell(ev))
			if a.exitKeyPressed(kev) {
				break loop
			}
			a.handleKeyPressEvent(ctx, kev)
		case *tcell.EventMouse:
			sev := sevent.New(sevent.WithTCell(ev))
			if sev.Direction() != types.ScrollDirectionNone {
				a.handleScrollEvent(ctx, sev)
			} else {
				mev := mevent.New(mevent.WithTCell(ev)) //, a.lastMouseEvent, a.mouseDownEvent)
				a.handleMouseEvent(ctx, mev)
			}
		case *tcell.EventError:
			return ev
		}
	}

	return nil
}

// draw renders the Application's active View to the Terminal screen.
func (a *Application) draw(ctx context.Context) {
	s := a.screen
	if s == nil {
		panic("called Application.draw() with nil screen.")
	}
	v := a.CurrentView()
	if v == nil {
		v = view.New(ctx, "main")
		a.views["main"] = v
		a.curView = "main"
	}

	s.Clear()

	a.Box.Render(ctx, s)
	v.SetBounds(a.InnerBounds())
	v.Draw(ctx, s)
	s.Show()
}
