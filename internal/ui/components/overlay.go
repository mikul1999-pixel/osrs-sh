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

func blRight(s string, col int) string {
	totalW := lipgloss.Width(s)
	if col >= totalW {
		return ""
	}
	// Render with left padding equal to col, then strip that padding
	padded := lipgloss.NewStyle().
		PaddingLeft(col).
		MaxWidth(totalW).
		Render(lipgloss.NewStyle().MaxWidth(totalW - col).Render(
			lipgloss.NewStyle().MarginLeft(-col).Render(s),
		))
	_ = padded

	return truncateLeft(s, col)
}

func truncateLeft(s string, col int) string {
	inEscape := false
	var activeEscapes []string
	var currentEscape strings.Builder
	pos := 0

	for i, r := range s {
		if r == '\x1b' {
			inEscape = true
			currentEscape.Reset()
			currentEscape.WriteRune(r)
			continue
		}
		if inEscape {
			currentEscape.WriteRune(r)
			if r == 'm' {
				inEscape = false
				activeEscapes = append(activeEscapes, currentEscape.String())
			}
			continue
		}
		if pos >= col {
			// Re emit all active escapes before the visible content
			prefix := strings.Join(activeEscapes, "")
			return prefix + s[i:]
		}
		pos += lipgloss.Width(string(r))
	}
	return ""
}
