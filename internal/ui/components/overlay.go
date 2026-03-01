package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// -- Helper for Custom Components ----------

// allows overlay component to float on top
func PlaceOverlay(x, y int, overlay, base string, width int) string {
	baseLines := strings.Split(base, "\n")
	overlayLines := strings.Split(overlay, "\n")

	for i, ol := range overlayLines {
		row := y + i
		if row < 0 || row >= len(baseLines) {
			continue
		}
		bl := baseLines[row]
		// Pad base line to width if needed
		blW := lipgloss.Width(bl)
		if blW < width {
			bl += strings.Repeat(" ", width-blW)
		}
		// Replace chars at column x with overlay line
		baseLines[row] = blLeft(bl, x) + ol + blRight(bl, x+lipgloss.Width(ol))
	}
	return strings.Join(baseLines, "\n")
}

// baseLineLeft returns the visible characters of s up to column col
func blLeft(s string, col int) string {
	return lipgloss.NewStyle().MaxWidth(col).Render(s)
}

// baseLineRight returns the visible characters of s starting at column col
func blRight(s string, col int) string {
	// Strip everything left of col
	leftPart := lipgloss.NewStyle().MaxWidth(col).Render(s)
	leftW := lipgloss.Width(leftPart)

	// Walk runes counting visible width
	inEscape := false
	pos := 0
	for i, r := range s {
		if r == '\x1b' {
			inEscape = true
		}
		if inEscape {
			if r == 'm' {
				inEscape = false
			}
			continue
		}
		if pos >= leftW {
			return s[i:]
		}
		pos += lipgloss.Width(string(r))
	}
	return ""
}
