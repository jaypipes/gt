package element

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/box"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/types"
	"github.com/samber/lo"
)

// Element is a base class that implements [types.Element] with some common
// method implementations. Subclasses in the [element] subpackages embed
// [element.Element] and implement various [types.Element] methods.
type Element struct {
	*sync.RWMutex
	core.Identifiable
	box.Box

	// controller can be used by the element can use to control the
	// screen's cursor.
	controller types.Controller

	// childIndex is the index of this Box in the parent's children.
	childIndex int
	// parent is the this Node's parent, if any.
	parent types.Node
	// children is the collection of Nodes that are the direct children of this
	// Node, if any.
	children []types.Node
	// class is the Element's type/class, e.g. "gt.div" or "gt.span"
	class string

	// style this is the style of the Element's content (i.e. the non-border
	// cells of the Element)
	style types.Style

	// textContent is any unstyle raw text content for the Element.
	textContent string

	// disabled indicates the Element is not focusable and does not accept
	// input.
	disabled bool
	// focused is true if the Element has the current focus.
	focused bool
	// hovered is true if the mouse is currently over the Element but the
	// Element does *not* have the current focus.
	hovered bool

	// onFocus contains the stack of callbacks that execute when the Element
	// receives focus.
	onFocus []types.FocusCallback
	// onLoseFocus contains the stack of callbacks that execute when the
	// Element loses focus.
	onLoseFocus []types.FocusCallback
	// onKeyPress contains the stack of callbacks that execute when a keypress
	// event occurs.
	onKeyPress []types.KeyPressEventCallback
	// onScroll contains the stack of callbacks that execute when a scroll
	// event occurs.
	onScroll []types.ScrollEventCallback
	// onMouseHover contains the stack of callbacks that execute when the
	// Element is hovered over by the mouse but the Element does *not* have the
	// focus.
	onMouseHover []types.MouseEventCallback
	// onMouseLoseHover contains the stack of callbacks that execute when the
	// mouse is no longer hovering over the Element (but was previously).
	onMouseLoseHover []types.EventCallback
	// onMouseClick contains the stack of callbacks that execute when the
	// Element is clicked on by the mouse.
	onMouseClick []types.MouseClickEventCallback
	// onMouseDoubleClick contains the stack of callbacks that execute when the
	// Element is double-clicked on by the mouse.
	onMouseDoubleClick []types.MouseClickEventCallback
	// onMouseScroll contains the stack of callbacks that execute when the
	// mouse wheel is scrolled while the Element has focus.
	onMouseScroll []types.MouseEventCallback
	// onMouseDragMove contains the stack of callbacks that execute when the
	// user performs a mouse drag action.
	onMouseDragMove []types.MouseDragEventCallback
	// onMouseDragStop contains the stack of callbacks that execute when the
	// user ends a mouse drag action.
	onMouseDragStop []types.MouseDragEventCallback
}

// Tag returns a string with the Element's type/class and ID
func (e *Element) Tag() string {
	return fmt.Sprintf("<%s:%s>", e.class, e.ID())
}

func (e *Element) String() string {
	parentStr := "nil"
	if e.parent != nil {
		parentEl, ok := e.parent.(types.Element)
		if ok {
			parentStr = parentEl.Tag()
		} else {
			parentID, ok := e.parent.(types.Identifiable)
			if ok {
				parentStr = parentID.ID()
			}
		}
	}
	return fmt.Sprintf(
		"<%s id=%s child_index=%d parent=%s children=%d %s",
		e.class, e.ID(),
		e.childIndex, parentStr, len(e.children),
		e.Box.String(),
	)
}

// WithID sets the Element's unique identifier and returns the Element.
func (e *Element) WithID(id string) types.Element {
	e.SetID(id)
	return e
}

// WithClass sets the Element's type/class and returns the Element
func (e *Element) WithClass(class string) types.Element {
	e.class = class
	return e
}

// Class returns the Element's type/class, e.g. "gt.span" or "gt.div"
func (e *Element) Class() string {
	return e.class
}

// Controller returns the controller.
func (e *Element) Controller() types.Controller {
	return e.controller
}

// SetController sets the controller.
func (e *Element) SetController(c types.Controller) {
	e.controller = c
}

// WithController sets the controller and returns the Element.
func (e *Element) WithController(c types.Controller) types.Element {
	e.SetController(c)
	return e
}

// SetDisabled sets whether the Element is disabled. Disabled Elements cannot
// receive the focus.
func (e *Element) SetDisabled(on bool) {
	e.disabled = on
}

// Disabled returns true if the Element cannot get the focus.
func (e *Element) Disabled() bool {
	return e.disabled
}

// WithDisabled sets whether the Element is disabled and returns the Element.
func (e *Element) WithDisabled(on bool) types.Element {
	e.SetDisabled(on)
	return e
}

// Render implements the types.Renderable interface
func (e *Element) Render(ctx context.Context, screen types.Screen) {
	s := e.Style()
	styleStr := "none"
	if s != nil {
		styleStr = s.String()
	}
	bounds := e.Bounds()
	gtlog.Debug(
		ctx, "Element.Render[%s]: bounds=%s style=%s",
		e.Tag(), bounds, styleStr,
	)
	e.Box.Render(ctx, screen)
	content := e.TextContent()
	if len(content) == 0 {
		return
	}
	inner := e.InnerBounds()
	// If there is no alignment set, inherit from the nearest parent with
	// non-auto alignment.
	align := e.Alignment()
	if align == types.AlignmentAuto {
		parentNode := e.Parent()
		parent, ok := parentNode.(types.Plottable)
		if ok {
			parentAlign := parent.Alignment()
			if parentAlign != types.AlignmentAuto {
				align = parentAlign
			}
		}
	}
	whitespace := e.Whitespace()
	if whitespace&types.WhitespacePreserve != 0 {
		// Preserve the whitespace by making the text content string we supply
		// to render.Align already pre-padded with spaces.
		sb := &strings.Builder{}
		content = strings.ReplaceAll(content, "\t", "    ")
		lines := strings.Split(content, "\n")
		maxWidth := lo.Max(lo.Map(lines, func(line string, _ int) int {
			return len(line)
		}))
		for x, line := range lines {
			diffFromMax := len(line) - maxWidth
			if diffFromMax > 0 {
				pad := strings.Repeat(" ", diffFromMax)
				sb.WriteString(pad)
			}
			sb.WriteString(line)
			if x < len(line)-1 {
				sb.WriteRune('\n')
			}
		}
		content = sb.String()
	}
	content = render.Align(
		ctx, content, inner, align, whitespace,
	)
	lines := strings.Split(content, "\n")
	startX := inner.Min.X
	startY := inner.Min.Y
	for y, line := range lines {
		for x := range line {
			screen.Put(startX+x, startY+y, string(line[x]), style.TCell(s))
		}
	}
}

var _ types.Element = (*Element)(nil)
