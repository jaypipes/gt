package border

import (
	"github.com/jaypipes/gt/core/cell"
	"github.com/jaypipes/gt/core/graphic"
	"github.com/jaypipes/gt/types"
)

// None returns an empty border.
func None() types.Border {
	return &Border{}
}

// Normal returns a standard-type border with a normal weight and 90
// degree corners.
func Normal() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.BoxDrawingLightHorizontal)),
		b:  cell.New(cell.WithContent(graphic.BoxDrawingLightHorizontal)),
		l:  cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		r:  cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		tl: cell.New(cell.WithContent(graphic.BoxDrawingLightDownAndRight)),
		tr: cell.New(cell.WithContent(graphic.BoxDrawingLightDownAndLeft)),
		bl: cell.New(cell.WithContent(graphic.BoxDrawingLightUpAndRight)),
		br: cell.New(cell.WithContent(graphic.BoxDrawingLightUpAndLeft)),
	}
}

// Rounded returns a border with rounded corners.
func Rounded() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.BoxDrawingLightHorizontal)),
		b:  cell.New(cell.WithContent(graphic.BoxDrawingLightHorizontal)),
		l:  cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		r:  cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		tl: cell.New(cell.WithContent(graphic.BoxDrawingLightArcDownAndRight)),
		tr: cell.New(cell.WithContent(graphic.BoxDrawingLightArcDownAndLeft)),
		bl: cell.New(cell.WithContent(graphic.BoxDrawingLightArcUpAndRight)),
		br: cell.New(cell.WithContent(graphic.BoxDrawingLightArcUpAndLeft)),
	}
}

// Block returns a border that takes the whole block.
func Block() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.BlockFullBlock)),
		b:  cell.New(cell.WithContent(graphic.BlockFullBlock)),
		l:  cell.New(cell.WithContent(graphic.BlockFullBlock)),
		r:  cell.New(cell.WithContent(graphic.BlockFullBlock)),
		tl: cell.New(cell.WithContent(graphic.BlockFullBlock)),
		tr: cell.New(cell.WithContent(graphic.BlockFullBlock)),
		bl: cell.New(cell.WithContent(graphic.BlockFullBlock)),
		br: cell.New(cell.WithContent(graphic.BlockFullBlock)),
	}
}

// OuterHalfBlock returns a half-block border that sits outside the frame.
func OuterHalfBlock() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.BlockUpperHalfBlock)),
		b:  cell.New(cell.WithContent(graphic.BlockLowerHalfBlock)),
		l:  cell.New(cell.WithContent(graphic.BlockLeftHalfBlock)),
		r:  cell.New(cell.WithContent(graphic.BlockRightHalfBlock)),
		tl: cell.New(cell.WithContent(graphic.BlockQuadrantUpperLeftAndUpperRightAndLowerLeft)),
		tr: cell.New(cell.WithContent(graphic.BlockQuadrantUpperLeftAndUpperRightAndLowerRight)),
		bl: cell.New(cell.WithContent(graphic.BlockQuadrantUpperLeftAndLowerLeftAndLowerRight)),
		br: cell.New(cell.WithContent(graphic.BlockQuadrantUpperRightAndLowerLeftAndLowerRight)),
	}
}

// InnerHalfBlock returns a half-block border that sits inside the frame.
func InnerHalfBlock() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.BlockLowerHalfBlock)),
		b:  cell.New(cell.WithContent(graphic.BlockUpperHalfBlock)),
		l:  cell.New(cell.WithContent(graphic.BlockRightHalfBlock)),
		r:  cell.New(cell.WithContent(graphic.BlockLeftHalfBlock)),
		tl: cell.New(cell.WithContent(graphic.BlockQuadrantLowerRight)),
		tr: cell.New(cell.WithContent(graphic.BlockQuadrantLowerLeft)),
		bl: cell.New(cell.WithContent(graphic.BlockQuadrantUpperRight)),
		br: cell.New(cell.WithContent(graphic.BlockQuadrantUpperLeft)),
	}
}

// Thick returns a border that's thicker than the one returned by
// Normal.
func Thick() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.BoxDrawingHeavyHorizontal)),
		b:  cell.New(cell.WithContent(graphic.BoxDrawingHeavyHorizontal)),
		l:  cell.New(cell.WithContent(graphic.BoxDrawingHeavyVertical)),
		r:  cell.New(cell.WithContent(graphic.BoxDrawingHeavyVertical)),
		tl: cell.New(cell.WithContent(graphic.BoxDrawingHeavyDownAndRight)),
		tr: cell.New(cell.WithContent(graphic.BoxDrawingHeavyDownAndLeft)),
		bl: cell.New(cell.WithContent(graphic.BoxDrawingHeavyUpAndRight)),
		br: cell.New(cell.WithContent(graphic.BoxDrawingHeavyUpAndLeft)),
	}
}

// Double returns a border comprised of two thin strokes.
func Double() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.BoxDrawingDoubleHorizontal)),
		b:  cell.New(cell.WithContent(graphic.BoxDrawingDoubleHorizontal)),
		l:  cell.New(cell.WithContent(graphic.BoxDrawingDoubleVertical)),
		r:  cell.New(cell.WithContent(graphic.BoxDrawingDoubleVertical)),
		tl: cell.New(cell.WithContent(graphic.BoxDrawingDoubleDownAndRight)),
		tr: cell.New(cell.WithContent(graphic.BoxDrawingDoubleDownAndLeft)),
		bl: cell.New(cell.WithContent(graphic.BoxDrawingDoubleUpAndRight)),
		br: cell.New(cell.WithContent(graphic.BoxDrawingDoubleUpAndLeft)),
	}
}

// Hidden returns a border that renders as a series of single-cell spaces. It's
// useful for cases when you want to remove a standard border but maintain
// layout positioning. This said, you can still apply a background color to a
// hidden border.
func Hidden() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.Space)),
		b:  cell.New(cell.WithContent(graphic.Space)),
		l:  cell.New(cell.WithContent(graphic.Space)),
		r:  cell.New(cell.WithContent(graphic.Space)),
		tl: cell.New(cell.WithContent(graphic.Space)),
		tr: cell.New(cell.WithContent(graphic.Space)),
		bl: cell.New(cell.WithContent(graphic.Space)),
		br: cell.New(cell.WithContent(graphic.Space)),
	}
}

// Markdown return a table border in markdown style.
func Markdown() types.Border {
	return &Border{
		l:  cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		r:  cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		tl: cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		tr: cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		bl: cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
		br: cell.New(cell.WithContent(graphic.BoxDrawingLightVertical)),
	}
}

// ASCII returns a table border with ASCII characters.
func ASCII() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(graphic.Hyphen)),
		b:  cell.New(cell.WithContent(graphic.Hyphen)),
		l:  cell.New(cell.WithContent(graphic.Bar)),
		r:  cell.New(cell.WithContent(graphic.Bar)),
		tl: cell.New(cell.WithContent(graphic.Plus)),
		tr: cell.New(cell.WithContent(graphic.Plus)),
		bl: cell.New(cell.WithContent(graphic.Plus)),
		br: cell.New(cell.WithContent(graphic.Plus)),
	}
}
