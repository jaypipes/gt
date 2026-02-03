package element

import (
	"context"
	"fmt"
	"sync"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/box"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// Element is a base class that implements [types.Element] with some common
// method implementations. Subclasses in the [element] subpackages embed
// [element.Element] and implement various [types.Element] methods.
type Element struct {
	*sync.RWMutex
	core.Identifiable
	box.Box

	// childIndex is the index of this Box in the parent's children.
	childIndex int
	// parent is the this Node's parent, if any.
	parent types.Node
	// children is the collection of Nodes that are the direct children of this
	// Node, if any.
	children []types.Node
	// class is the Element's type/class, e.g. "gt.label" or "gt.canvas"
	class string

	// textContent is any unstyled raw text content for the Element.
	textContent string

	// style is the style mode of the Element's content (i.e. the non-border
	// cells of the Element)
	style types.Style
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

// Draw implements the uv.Drawable interface
func (e *Element) Draw(screen types.Screen, bounds types.Rectangle) {
	ctx := context.TODO()
	gtlog.Debug(ctx, "Element.Draw[%s]: bounds=%s", e.Tag(), bounds)
	e.Box.Draw(screen, bounds)
	content := e.TextContent()
	if len(content) == 0 {
		return
	}
	inner := e.InnerBounds()
	innerClipped := render.Overlapping(bounds, inner)
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
	content = render.AlignString(
		ctx, content, inner, align,
	)
	style := e.Style()
	content = style.Styled(content)
	ss := uv.NewStyledString(content)
	ws := e.Whitespace()
	if ws&types.WhitespaceWrapNever == 0 {
		ss.Wrap = true
	}
	ss.Draw(screen, innerClipped)
}

var _ types.Element = (*Element)(nil)
