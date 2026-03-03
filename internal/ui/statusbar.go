package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	coreXP "github.com/mikul1999-pixel/osrs-sh/internal/core/xp"
)

// renderStatusBar renders two lines of the status bar
func (a AppModel) renderStatusBar() string {
	line1 := a.renderStatusLine1()
	line2 := a.renderStatusLine2()
	return lipgloss.JoinVertical(lipgloss.Left, line1, line2)
}

// renderStatusLine1 renders the top line
func (a AppModel) renderStatusLine1() string {
	right := a.renderStatusBlock()
	spacer := ActiveTheme.StatusLine().Render(strings.Repeat(" ", max(0, a.width-lipgloss.Width(right))))
	return ActiveTheme.StatusLine().Width(a.width).Render(spacer + right)
}

// renderStatusLine1 renders the context line
func (a AppModel) renderStatusLine2() string {
	left := a.renderPlayerBlock()
	right := a.renderStatusActions() + a.renderGlobalActions()

	gap := a.width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 0 {
		gap = 0
	}
	spacer := ActiveTheme.StatusLine().Render(strings.Repeat(" ", gap))

	return ActiveTheme.StatusLine().Width(a.width).Render(left + spacer + right)
}

// renderPlayerBlock renders the left segment of line 1
func (a AppModel) renderPlayerBlock() string {
	// Loading state
	if a.player.Loading {
		frame := ""
		if a.spinner != nil {
			frame = a.spinner.View()
		}
		return ActiveTheme.StatusLine().Render(
			" " + ActiveTheme.StatusLine().Render(frame) +
				ActiveTheme.StatusLine().Render(" loading ") +
				ActiveTheme.StatusLine().Render(strings.ToLower(a.player.RSN)+"..."),
		)
	}

	// Error state
	if a.player.Err != "" {
		return ActiveTheme.StatusLine().Render(
			ActiveTheme.StatusError().Render(" ! " + a.player.Err + " "),
		)
	}

	// Loaded state
	if a.player.Loaded {
		total := coreXP.XPToTotalLevel(a.player.xp)
		flexMsg := ""
		if total >= 2376 {
			flexMsg = "(btw)"
		}
		return ActiveTheme.StatusBlock().Render(" "+strings.ToLower(a.player.RSN)+" ") +
			ActiveTheme.StatusLine().Render(fmt.Sprintf(" %d", total)) +
			ActiveTheme.StatusLine().Render(" "+flexMsg)
	}

	// Default - no player loaded
	return ActiveTheme.StatusBlock().Render(" username ") +
		ActiveTheme.StatusLine().Render(" not set")
}

// renderStatusBlock renders the right segment of line 1
func (a AppModel) renderStatusBlock() string {
	if a.statusContext.Label == "" {
		return ""
	}
	return ActiveTheme.StatusBlockMuted().Render(" "+a.statusContext.Label+" ") + ActiveTheme.StatusLine().Render(" ")
}

// renderStatusActions renders the status keybind hints on line 1
func (a AppModel) renderStatusActions() string {
	if len(a.statusContext.Keybinds) == 0 {
		return ""
	}

	var parts []string
	for _, kb := range a.statusContext.Keybinds {
		parts = append(parts, ActiveTheme.StatusKey().Render(kb.Key)+ActiveTheme.StatusVal().Render(" "+kb.Label))
	}
	return " " + strings.Join(parts, "  ")
}

// renderGlobalActions renders the global keybind hints on line 2
func (a AppModel) renderGlobalActions() string {
	return Space(3) +
		ActiveTheme.StatusKey().Render("ctrl+p") + ActiveTheme.StatusVal().Render(" commands") + Space(3) +
		ActiveTheme.StatusKey().Render("ctrl+t") + ActiveTheme.StatusVal().Render(" themes") + Space(3) +
		ActiveTheme.StatusKey().Render("ctrl+c") + ActiveTheme.StatusVal().Render(" quit") + Space(1)
}
