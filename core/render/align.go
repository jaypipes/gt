package render

import (
	"context"
	"strings"

	"github.com/charmbracelet/x/ansi"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// AlignString returns a string adjusted with padding to make the supplied string
// align horizontally and vertically to the supplied alignment mode.
func AlignString(
	ctx context.Context,
	content string,
	bounds types.Rectangle,
	align types.Alignment,
) string {
	gtlog.Debug(
		ctx, "render.AlignString: bounds=%s alignment=%s",
		bounds, align,
	)
	width := bounds.Dx()
	height := bounds.Dy()

	var b strings.Builder
	// We know that the returned string is going to be equal to the number of
	// cells in the bounding area.
	area := width * height
	if area > 0 {
		b.Grow(area)
	}

	// First we process the vertical alignment of the supplied string and then
	// we will pad each individual line of the vertically-aligned string in
	// order to horizontally-align it.
	numLines := strings.Count(content, "\n") + 1
	if height >= numLines {
		linesToPad := height - numLines
		if align&types.AlignmentBottom != 0 {
			// pad the string with new lines to the top of the bounding box
			b.WriteString(strings.Repeat("\n", linesToPad))
			b.WriteString(content)
		} else if align&types.AlignmentMiddle != 0 {
			// pad the string with new lines at the top and bottom of the bounding
			// box
			linesToPadTop := linesToPad / 2
			linesToPadBottom := linesToPad / 2
			// handle overflow/underflow from integer division rounding...
			totalLines := numLines + linesToPadTop + linesToPadBottom
			if totalLines > height {
				linesToPadTop--
			} else if totalLines < height {
				linesToPadBottom++
			}
			b.WriteString(strings.Repeat("\n", linesToPadTop))
			b.WriteString(content)
			b.WriteString(strings.Repeat("\n", linesToPadBottom))
		} else {
			// pad the string with new lines to the bottom of the bounding box
			b.WriteString(content)
			b.WriteString(strings.Repeat("\n", linesToPad))
		}
	}
	lines := strings.Split(b.String(), "\n")
	b.Reset()

	for x, line := range lines {
		numCells := ansi.StringWidth(line)

		cellsToPad := width - numCells

		if cellsToPad > 0 {
			if align&types.AlignmentRight != 0 {
				b.WriteString(strings.Repeat(" ", cellsToPad))
				b.WriteString(line)
			} else if align&types.AlignmentCenter != 0 {
				cellsToPadLeft := cellsToPad / 2
				cellsToPadRight := cellsToPad / 2
				// handle overflow/underflow from integer division rounding...
				totalCells := numCells + cellsToPadLeft + cellsToPadRight
				if totalCells > width {
					cellsToPadRight--
				} else if totalCells < width {
					cellsToPadRight++
				}
				b.WriteString(strings.Repeat(" ", cellsToPadLeft))
				b.WriteString(line)
				b.WriteString(strings.Repeat(" ", cellsToPadRight))
			} else {
				b.WriteString(line)
				b.WriteString(strings.Repeat(" ", cellsToPad))
			}
		} else {
			b.WriteString(line)
		}
		if x < len(lines)-1 {
			b.WriteRune('\n')
		}
	}

	return b.String()
}
