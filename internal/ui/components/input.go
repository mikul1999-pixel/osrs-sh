package components

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

// -- Custom Input Box Component ----------

// InputOptions configures a new Input component
type InputOptions struct {
	// Placeholder text when buffer is empty
	Placeholder string

	// Prompt prefix rendered before the text, like "> "
	Prompt        string
	CharLimit     int
	ShowBottomRow bool

	// Dynamic content for the bottom row. Called every View()
	BottomLeft  string
	BottomRight string

	AccentFocused    lipgloss.Color
	AccentUnfocused  lipgloss.Color
	Background       lipgloss.Color
	PromptStyle      lipgloss.Style
	TextStyle        lipgloss.Style
	PlaceholderStyle lipgloss.Style
	CursorStyle      lipgloss.Style
	PaddingTop       int
	PaddingMiddle    int
	PaddingBottom    int
}

type CursorBlinkMsg struct{}

func CursorBlink() tea.Cmd {
	return tea.Tick(time.Millisecond*530, func(t time.Time) tea.Msg {
		return CursorBlinkMsg{}
	})
}

func defaultOptions() InputOptions {
	bg := lipgloss.Color("#1a1a1a")
	accent := lipgloss.Color("#f0c040")
	dim := lipgloss.Color("#555555")

	base := lipgloss.NewStyle().Background(bg)

	return InputOptions{
		Placeholder:      "type something...",
		Prompt:           "",
		CharLimit:        0,
		ShowBottomRow:    false,
		AccentFocused:    accent,
		AccentUnfocused:  dim,
		Background:       bg,
		PromptStyle:      base.Foreground(accent),
		TextStyle:        base.Foreground(lipgloss.Color("#ffffff")),
		PlaceholderStyle: base.Foreground(dim),
		CursorStyle:      lipgloss.NewStyle().Background(accent).Foreground(lipgloss.Color("#000000")),
		PaddingTop:       1,
		PaddingMiddle:    1,
		PaddingBottom:    1,
	}
}

// -- Model ----------

// Input is a single-line text input component for Bubble Tea
type Input struct {
	opts          InputOptions
	value         []rune
	cursor        int // byte index into value
	width         int
	focused       bool
	cursorVisible bool
}

// NewInput creates a new Input with the given options merged over defaults
func NewInput(opts InputOptions) Input {
	d := defaultOptions()

	if opts.Placeholder == "" {
		opts.Placeholder = d.Placeholder
	}
	if opts.Prompt == "" {
		opts.Prompt = d.Prompt
	}
	if opts.AccentFocused == "" {
		opts.AccentFocused = d.AccentFocused
	}
	if opts.AccentUnfocused == "" {
		opts.AccentUnfocused = d.AccentUnfocused
	}
	if opts.Background == "" {
		opts.Background = d.Background
	}

	return Input{opts: opts, cursorVisible: true}
}

// -- Public API ----------

func (m *Input) Focus() {
	m.focused = true
}

func (m *Input) Blur() {
	m.focused = false
}

func (m *Input) Focused() bool {
	return m.focused
}

// SetWidth sets the total render width of the component
func (m *Input) SetWidth(w int) {
	m.width = w
}

// Value returns the current input string
func (m *Input) Value() string {
	return string(m.value)
}

// SetValue replaces the buffer and moves the cursor to the end
func (m *Input) SetValue(s string) {
	m.value = []rune(s)
	m.cursor = len(m.value)
}

// Reset clears the buffer
func (m *Input) Reset() {
	m.value = nil
	m.cursor = 0
}

// SetPlaceholder updates the left bottom row content function
func (m *Input) SetPlaceholder(s string) {
	m.opts.Placeholder = s
}

// SetBottomLeft updates the left bottom row content function
func (m *Input) SetBottomLeft(s string) {
	m.opts.BottomLeft = s
}

// SetBottomRight updates the right bottom row content function
func (m *Input) SetBottomRight(s string) {
	m.opts.BottomRight = s
}

// -- Bubble Tea interface ----------

func (m Input) Init() tea.Cmd {
	return CursorBlink()
}

func (m Input) Update(msg tea.Msg) (Input, tea.Cmd) {
	switch msg := msg.(type) {
	case CursorBlinkMsg:
		if m.focused {
			m.cursorVisible = !m.cursorVisible
		}
		return m, CursorBlink()

	case tea.KeyPressMsg:
		if !m.focused {
			return m, nil
		}

		switch msg.String() {
		case "left":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right":
			if m.cursor < len(m.value) {
				m.cursor++
			}
		case "home", "ctrl+a":
			m.cursor = 0
		case "end", "ctrl+e":
			m.cursor = len(m.value)
		case "backspace", "ctrl+h":
			if m.cursor > 0 {
				m.value = append(m.value[:m.cursor-1], m.value[m.cursor:]...)
				m.cursor--
			}
		case "delete":
			if m.cursor < len(m.value) {
				m.value = append(m.value[:m.cursor], m.value[m.cursor+1:]...)
			}
		case "ctrl+w":
			if m.cursor > 0 {
				end := m.cursor
				for m.cursor > 0 && m.value[m.cursor-1] == ' ' {
					m.cursor--
				}
				for m.cursor > 0 && m.value[m.cursor-1] != ' ' {
					m.cursor--
				}
				m.value = append(m.value[:m.cursor], m.value[end:]...)
			}
		case "ctrl+u":
			m.value = m.value[m.cursor:]
			m.cursor = 0
		case "ctrl+k":
			m.value = m.value[:m.cursor]
		default:
			if key := msg.Key(); key.Text != "" {
				runes := []rune(key.Text)
				if m.opts.CharLimit > 0 && len(m.value) >= m.opts.CharLimit {
					break
				}
				newVal := make([]rune, len(m.value)+len(runes))
				copy(newVal, m.value[:m.cursor])
				copy(newVal[m.cursor:], runes)
				copy(newVal[m.cursor+len(runes):], m.value[m.cursor:])
				m.value = newVal
				m.cursor += len(runes)
			}
		}
	}
	return m, nil
}

// -- View ----------

func (m Input) View() string {
	if m.width == 0 {
		return ""
	}

	accentColor := m.opts.AccentUnfocused
	if m.focused {
		accentColor = m.opts.AccentFocused
	}

	// accent bar on the left
	accentBar := lipgloss.NewStyle().
		Foreground(accentColor).
		Render("▎")

	// Inner width = total width minus accent bar and gap
	innerWidth := m.width - 2
	if innerWidth < 1 {
		innerWidth = 1
	}

	bgStyle := lipgloss.NewStyle().Background(m.opts.Background)

	topRow := m.renderInputRow(innerWidth, bgStyle)
	accentGap := bgStyle.Width(1).Render(" ")

	var rows []string

	// Build the slice with padding
	pad := accentBar + accentGap + bgStyle.Render(strings.Repeat(" ", innerWidth))

	for i := 0; i < m.opts.PaddingTop; i++ {
		rows = append(rows, pad)
	}
	rows = append(rows, accentBar+accentGap+topRow)
	if m.opts.ShowBottomRow {
		for i := 0; i < m.opts.PaddingMiddle; i++ {
			rows = append(rows, pad)
		}
		bottomRow := m.renderBottomRow(innerWidth, bgStyle)
		rows = append(rows, accentBar+accentGap+bottomRow)
	}
	for i := 0; i < m.opts.PaddingBottom; i++ {
		rows = append(rows, pad)
	}

	return strings.Join(rows, "\n")
}

// renderInputRow builds the top line: prompt + text + cursor + padding
func (m Input) renderInputRow(innerWidth int, bgStyle lipgloss.Style) string {
	prompt := m.opts.PromptStyle.Background(m.opts.Background).Render(m.opts.Prompt)
	promptW := lipgloss.Width(prompt)

	textAreaWidth := innerWidth - promptW
	if textAreaWidth < 1 {
		textAreaWidth = 1
	}

	var textPart string

	if len(m.value) == 0 && !m.focused {
		// No focus, no text: show placeholder in text area
		ph := []rune(m.opts.Placeholder)
		visible := ph
		if len(visible) > textAreaWidth {
			visible = visible[:textAreaWidth]
		}
		padded := string(visible) + strings.Repeat(" ", textAreaWidth-len(visible))
		textPart = m.opts.PlaceholderStyle.Background(m.opts.Background).Render(padded)

	} else if len(m.value) == 0 {
		// Focused, empty: cursor on first char of placeholder
		ph := []rune(m.opts.Placeholder)

		// cursor covers first placeholder char
		cursorChar := " "
		if len(ph) > 0 {
			cursorChar = string(ph[0])
		}
		cursorRendered := m.renderCursor(cursorChar)

		// rest of placeholder after cursor
		var rest string
		if len(ph) > 1 {
			remaining := textAreaWidth - 1
			tail := ph[1:]
			if len(tail) > remaining {
				tail = tail[:remaining]
			}
			rest = m.opts.PlaceholderStyle.Background(m.opts.Background).
				Render(string(tail) + strings.Repeat(" ", remaining-len(tail)))
		} else {
			pad := strings.Repeat(" ", textAreaWidth-1)
			rest = bgStyle.Render(pad)
		}

		textPart = cursorRendered + rest

	} else {
		// Has text: render with cursor inline
		runes := m.value
		start := 0
		end := len(runes)

		if end-start > textAreaWidth-1 {
			// keep cursor in view
			if m.cursor >= textAreaWidth-1 {
				start = m.cursor - (textAreaWidth - 2)
			}
			end = start + textAreaWidth - 1
			if end > len(runes) {
				end = len(runes)
			}
		}

		var sb strings.Builder

		for i := start; i < end; i++ {
			ch := string(runes[i])
			if i == m.cursor {
				sb.WriteString(m.renderCursor(ch))
			} else {
				sb.WriteString(m.opts.TextStyle.Background(m.opts.Background).Render(ch))
			}
		}

		// cursor at end of text
		if m.cursor == len(runes) && m.cursor >= start {
			sb.WriteString(m.renderCursor(" "))
		}

		rendered := sb.String()
		renderedW := lipgloss.Width(rendered)

		// fill remaining space
		remaining := textAreaWidth - renderedW
		if remaining > 0 {
			sb.WriteString(bgStyle.Width(remaining).Render(""))
		}

		textPart = sb.String()
	}

	return prompt + textPart
}

// renderCursor builds the blinking cursor
func (m Input) renderCursor(ch string) string {
	if !m.cursorVisible {
		return m.opts.PlaceholderStyle.Background(m.opts.Background).Render(ch)
	}
	return m.opts.CursorStyle.Render(ch)
}

// renderBottomRow builds the second line with left/right content
func (m Input) renderBottomRow(innerWidth int, bgStyle lipgloss.Style) string {
	left := ""
	right := ""

	if m.opts.BottomLeft != "" {
		left = m.opts.BottomLeft
	}
	if m.opts.BottomRight != "" {
		right = m.opts.BottomRight
	}

	leftRendered := left
	rightRendered := right

	leftW := lipgloss.Width(leftRendered)
	rightW := lipgloss.Width(rightRendered)

	gap := innerWidth - leftW - rightW
	if gap < 0 {
		gap = 0
	}

	middle := bgStyle.Width(gap).Render("")

	return leftRendered + middle + rightRendered
}
