package components

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

// -- Custom Toast Component ----------

// ToastStyle controls the visual treatment of a notification
type ToastStyle int

const (
	ToastSuccess ToastStyle = iota // good
	ToastError                     // bad
	ToastInfo                      // info
)

// ToastDismissMsg is sent when the dismiss timer fires
type ToastDismissMsg struct{ ID int }

// Toast is a standalone, embeddable notification component
type Toast struct {
	id      int
	message string
	sub     string
	style   ToastStyle
	visible bool
	width   int
	dismiss time.Duration
}

var toastIDCounter int

const defaultToastWidth = 30
const defaultDismiss = 2 * time.Second

var (
	colorSuccess = lipgloss.Color("#57a854")
	colorError   = lipgloss.Color("#c94f4f")
	colorInfo    = lipgloss.Color("#c8a84b")
	colorMain    = lipgloss.Color("#cbd2dd")
	colorSub     = lipgloss.Color("#6b6b6b")
)

// NewToast creates a toast
func NewToast() *Toast {
	toastIDCounter++
	return &Toast{
		id:      toastIDCounter,
		style:   ToastInfo,
		width:   defaultToastWidth,
		dismiss: defaultDismiss,
		visible: false,
	}
}

// SetMessage sets the primary message line
func (t *Toast) SetMessage(msg string) *Toast {
	t.message = msg
	return t
}

// SetSub sets an optional secondary line
func (t *Toast) SetSub(sub string) *Toast {
	t.sub = sub
	return t
}

// SetStyle sets the visual style
func (t *Toast) SetStyle(s ToastStyle) *Toast {
	t.style = s
	return t
}

// SetWidth sets the width of the toast box
func (t *Toast) SetWidth(w int) *Toast {
	t.width = w
	return t
}

// SetDismiss overrides the dismiss duration
func (t *Toast) SetDismiss(d time.Duration) *Toast {
	t.dismiss = d
	return t
}

// ID returns the unique ID for this toast instance
func (t *Toast) ID() int {
	return t.id
}

// Show marks the toast as visible and returns the dismiss command
func (t *Toast) Show() tea.Cmd {
	t.visible = true
	id := t.id
	dismiss := t.dismiss
	return tea.Tick(dismiss, func(time.Time) tea.Msg {
		return ToastDismissMsg{ID: id}
	})
}

// Hide marks the toast as invisible
func (t *Toast) Hide() {
	t.visible = false
}

// Visible returns if toast is shown
func (t *Toast) Visible() bool {
	return t.visible
}

// View renders the toast as a string
func (t *Toast) View() string {
	if !t.visible {
		return ""
	}

	accentColor := t.accentColor()
	accentBar := lipgloss.NewStyle().
		Foreground(accentColor).
		Render("▎")
	inner := lipgloss.NewStyle().Foreground(colorMain).Faint(true).Width(t.width - 4) // 4 = bar + padding
	line1 := accentBar + inner.Render(t.message)
	content := line1

	if t.sub != "" {
		line2 := accentBar + inner.Foreground(colorSub).Render(t.sub)
		content = lipgloss.JoinVertical(lipgloss.Left, line1, line2)
	}

	box := lipgloss.NewStyle().
		Padding(0, 1).
		Width(t.width).
		Render(content)

	return box
}

// accentColor returns the status color
func (t *Toast) accentColor() lipgloss.Color {
	switch t.style {
	case ToastSuccess:
		return colorSuccess
	case ToastError:
		return colorError
	default: // ToastInfo
		return colorInfo
	}
}
