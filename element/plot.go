package element

import (
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// AtPoint returns the child element at the supplied position, or nil if no the
// position is out of the element's bounding box. We perform a depth-first
// search of the child nodes since we do not allow overlapping boxes therefore
// the first matched leaf node is our match.
func (e *Element) AtPoint(pos types.Point) types.Node {
	return render.AtPoint(e, pos)
}
