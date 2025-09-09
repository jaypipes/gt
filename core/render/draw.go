package render

import (
	uv "github.com/charmbracelet/ultraviolet"
)

// Clear clears any rendered cell contents for the supplied bounding box.
func Clear(buf uv.Screen, area uv.Rectangle) {
	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			buf.SetCell(x, y, nil)
		}
	}
}
