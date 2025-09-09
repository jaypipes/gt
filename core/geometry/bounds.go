package geometry

import (
	"image"

	uv "github.com/charmbracelet/ultraviolet"
)

// Overlapping returns the overlapping viewport bounding box given two
// supplied bounding boxes.
func Overlapping(a, b uv.Rectangle) uv.Rectangle {
	return uv.Rectangle{
		Min: image.Point{
			X: max(a.Min.X, b.Min.X),
			Y: max(a.Min.Y, b.Min.Y),
		},
		Max: image.Point{
			X: min(a.Max.X, b.Max.X),
			Y: min(a.Max.Y, b.Max.Y),
		},
	}
}
