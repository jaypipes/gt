package types

// Alignment is an alignment mode
type Alignment int

const (
	// AlignmentAuto is the default alignment, which means the parent's alignment
	// will be used or "top-left" if there is no parent.
	AlignmentAuto Alignment = 0
	// AlignmentTop indicates the element's content will be positioned at the
	// top edge of the element's inner bounding box.
	AlignmentTop = 1
	// AlignmentBottom indicates the element's content will be positioned at the
	// bottom edge of the element's inner bounding box.
	AlignmentBottom = 1 << 1
	// AlignmentLeft indicates the element's content will be positioned at the
	// left edge of the element's inner bounding box.
	AlignmentLeft = 1 << 2
	// AlignmentRight indicates the element's content will be positioned at the
	// right edge of the element's inner bounding box.
	AlignmentRight = 1 << 3
	// AlignmentCenter indicates the element's content will be positioned at the
	// horizontal center of the element's inner bounding box.
	AlignmentCenter = 1 << 4
	// AlignmentMiddle indicates the element's content will be positioned at the
	// vertical middle of the element's inner bounding box.
	AlignmentMiddle = 1 << 5
	// AlignmentTopLeft indicates the element's content will be positioned at the
	// left and top edge of the element's inner bounding box.
	AlignmentTopLeft = AlignmentTop | AlignmentLeft
	// AlignmentTopRight indicates the element's content will be positioned at the
	// right and top edge of the element's inner bounding box.
	AlignmentTopRight = AlignmentTop | AlignmentRight
	// AlignmentTopCenter indicates the element's content will be positioned at the
	// horizontal center and top edge of the element's inner bounding box.
	AlignmentTopCenter = AlignmentTop | AlignmentCenter
	// AlignmentBottomLeft indicates the element's content will be positioned at the
	// left and bottom edge of the element's inner bounding box.
	AlignmentBottomLeft = AlignmentBottom | AlignmentLeft
	// AlignmentBottomRight indicates the element's content will be positioned at the
	// right and bottom edge of the element's inner bounding box.
	AlignmentBottomRight = AlignmentBottom | AlignmentRight
	// AlignmentBottomCenter indicates the element's content will be positioned at
	// the horizontal center and bottom edge of the element's inner bounding
	// box.
	AlignmentBottomCenter = AlignmentBottom | AlignmentCenter
	// AlignmentMiddleLeft indicates the element's content will be positioned at
	// the left and vertical middle of the element's inner bounding box.
	AlignmentMiddleLeft = AlignmentMiddle | AlignmentLeft
	// AlignmentMiddleRight indicates the element's content will be positioned at
	// the right and vertical middle of the element's inner bounding box.
	AlignmentMiddleRight = AlignmentMiddle | AlignmentRight
	// AlignmentMiddleCenter indicates the element's content will be positioned at
	// the horizontal center and vertical middle of the element's inner
	// bounding box.
	AlignmentMiddleCenter = AlignmentMiddle | AlignmentCenter
)

var (
	alignmentStrings = map[Alignment]string{
		AlignmentAuto:         "auto",
		AlignmentTop:          "top",
		AlignmentBottom:       "bottom",
		AlignmentLeft:         "left",
		AlignmentRight:        "right",
		AlignmentCenter:       "center",
		AlignmentMiddle:       "middle",
		AlignmentTopLeft:      "top-left",
		AlignmentTopRight:     "top-right",
		AlignmentTopCenter:    "top-center",
		AlignmentBottomLeft:   "bottom-left",
		AlignmentBottomRight:  "bottom-right",
		AlignmentBottomCenter: "bottom-center",
		AlignmentMiddleLeft:   "middle-left",
		AlignmentMiddleRight:  "middle-right",
		AlignmentMiddleCenter: "middle-center",
	}
)

func (a Alignment) String() string {
	return alignmentStrings[a]
}
