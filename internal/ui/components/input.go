package components

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

// -- Custom Input Box Component ----------

// InputCommand is a command entry and description shown in the dropdown
type InputCommand struct {
	Key  string
	Args string
	Desc string
}

// dropdownState holds runtime state for the command dropdown
type dropdownState struct {
	filtered  []InputCommand
	selected  int
	scrollOff int
}

// InputOptions configures a new Input component
type InputOptions struct {
	Placeholder      string // Placeholder text when buffer is empty
	Prompt           string // Prompt prefix rendered before the text, like "> "
	CharLimit        int
	ShowBottomRow    bool
	BottomLeft       string // Dynamic content for the bottom row. Called every View()
	BottomRight      string // Dynamic content for the bottom row. Called every View()
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
	Commands         []InputCommand
	CommandPrefix    string // first character before each command
	DropdownTrigger  rune   // rune to open the dropdown
	DropdownVisible  int    // max visible rows
	DropdownAccent   lipgloss.Color
	ForceDropdown    bool
	FilterDropdown   bool
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
		CommandPrefix:    "/",
		DropdownTrigger:  '/',
		DropdownVisible:  8,
		ForceDropdown:    false,
		FilterDropdown:   true,
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
	dropdown      *dropdownState
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

// RefreshDropdown refreshes the dropdown panel
func (m *Input) RefreshDropdown() {
	m.syncDropdown()
}

// -- Update opts ----------

func (m *Input) SetPlaceholder(s string)              { m.opts.Placeholder = s }
func (m *Input) SetBottomLeft(s string)               { m.opts.BottomLeft = s }
func (m *Input) SetBottomRight(s string)              { m.opts.BottomRight = s }
func (m *Input) SetAccentFocused(c lipgloss.Color)    { m.opts.AccentFocused = c }
func (m *Input) SetAccentUnfocused(c lipgloss.Color)  { m.opts.AccentUnfocused = c }
func (m *Input) SetBackground(c lipgloss.Color)       { m.opts.Background = c }
func (m *Input) SetPromptStyle(s lipgloss.Style)      { m.opts.PromptStyle = s }
func (m *Input) SetTextStyle(s lipgloss.Style)        { m.opts.TextStyle = s }
func (m *Input) SetPlaceholderStyle(s lipgloss.Style) { m.opts.PlaceholderStyle = s }
func (m *Input) SetCursorStyle(s lipgloss.Style)      { m.opts.CursorStyle = s }
func (m *Input) SetDropdownAccent(c lipgloss.Color)   { m.opts.DropdownAccent = c }
func (m *Input) SetCommands(cmd []InputCommand)       { m.opts.Commands = cmd }

// CommitDropdownSelection should be called by the parent when user presses enter
func (m *Input) CommitDropdownSelection(close bool) *InputCommand {
	if m.dropdown == nil || len(m.dropdown.filtered) == 0 {
		return nil
	}
	cmd := m.dropdown.filtered[m.dropdown.selected]
	if close {
		m.dropdown = nil
		m.Reset()
	}
	return &cmd
}

// syncDropdown opens, filters, or closes the dropdown
func (m *Input) syncDropdown() {
	if len(m.opts.Commands) == 0 {
		m.dropdown = nil
		return
	}

	runes := m.value
	trigger := m.opts.DropdownTrigger

	// ForceDropdown: gnore trigger entirely
	if m.opts.ForceDropdown {
		if m.dropdown == nil {
			m.dropdown = &dropdownState{}
		}

		if m.opts.FilterDropdown {
			// Filter on the first word
			firstWord := strings.Fields(string(runes))
			query := ""
			if len(firstWord) > 0 {
				query = strings.ToLower(firstWord[0])
			}
			m.dropdown.filtered = filterCommands(m.opts.Commands, query)
		} else {
			m.dropdown.filtered = m.opts.Commands
		}

		// Close dropdown if nothing matches
		if len(m.dropdown.filtered) == 0 {
			m.dropdown = nil
			return
		}

		// Clamp selection
		if m.dropdown.selected >= len(m.dropdown.filtered) {
			m.dropdown.selected = len(m.dropdown.filtered) - 1
		}
		if m.dropdown.scrollOff >= len(m.dropdown.filtered) {
			m.dropdown.scrollOff = 0
		}

		return
	}

	// Only open when trigger is typed
	if trigger != 0 {
		if len(runes) == 0 || runes[0] != trigger {
			m.dropdown = nil
			return
		}

		if m.dropdown == nil {
			m.dropdown = &dropdownState{}
		}

		if m.opts.FilterDropdown {
			// Filter on text after trigger
			query := strings.ToLower(string(runes[1:]))
			m.dropdown.filtered = filterCommands(m.opts.Commands, query)
		} else {
			m.dropdown.filtered = m.opts.Commands
		}

		if len(m.dropdown.filtered) == 0 {
			m.dropdown = nil
			return
		}

		// Clamp selection
		if m.dropdown.selected >= len(m.dropdown.filtered) {
			m.dropdown.selected = len(m.dropdown.filtered) - 1
		}
		if m.dropdown.scrollOff >= len(m.dropdown.filtered) {
			m.dropdown.scrollOff = 0
		}

		return
	}

	// No trigger and not forced: dropdown stays closed
	m.dropdown = nil
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

		// Dropdown navigation
		if m.dropdown != nil {
			switch msg.String() {
			case "up":
				if m.dropdown.selected > 0 {
					m.dropdown.selected--
					maxVis := m.opts.DropdownVisible
					if maxVis < 1 {
						maxVis = 8
					}
					if m.dropdown.selected < m.dropdown.scrollOff {
						m.dropdown.scrollOff = m.dropdown.selected
					}
				}
				return m, nil
			case "down":
				if m.dropdown.selected < len(m.dropdown.filtered)-1 {
					m.dropdown.selected++
					maxVis := m.opts.DropdownVisible
					if maxVis < 1 {
						maxVis = 8
					}
					if m.dropdown.selected >= m.dropdown.scrollOff+maxVis {
						m.dropdown.scrollOff = m.dropdown.selected - maxVis + 1
					}
				}
				return m, nil
			case "esc":
				m.dropdown = nil
				return m, nil
			}
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

		m.syncDropdown()
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

// DropdownView returns the rendered dropdown panel in isolation. Allows parent to position the overlay
func (m Input) DropdownView() string {
	if m.dropdown == nil || len(m.dropdown.filtered) == 0 {
		return ""
	}

	const (
		keyColW = 12
		argColW = 14
	)

	accentColor := m.opts.AccentUnfocused
	if m.focused {
		accentColor = m.opts.AccentFocused
	}

	dropdownAccentColor := m.opts.DropdownAccent
	if dropdownAccentColor == "" {
		dropdownAccentColor = m.opts.Background
	}

	bgStyle := lipgloss.NewStyle().Background(m.opts.Background)
	accentBar := lipgloss.NewStyle().Foreground(dropdownAccentColor).Render("▎")
	accentGap := bgStyle.Width(1).Render(" ")

	innerWidth := m.width - 2
	if innerWidth < 1 {
		innerWidth = 1
	}

	// Styles - unselected
	keyStyle := m.opts.TextStyle.
		Background(m.opts.Background).
		Width(keyColW)
	argsStyle := m.opts.TextStyle.
		Faint(true).
		Background(m.opts.Background).
		Width(argColW)
	descStyle := m.opts.TextStyle.
		Faint(true).
		Background(m.opts.Background)

	// Styles - selected
	selBase := lipgloss.NewStyle().
		Background(accentColor).
		Foreground(m.opts.Background)
	selKeyStyle := selBase.Width(keyColW)
	selArgsStyle := selBase.Width(argColW)
	selDescStyle := selBase
	selGap := selBase.Width(1).Render(" ")

	descW := innerWidth - keyColW - argColW // 2 = accentBar + gap
	if descW < 0 {
		descW = 0
	}

	dd := m.dropdown
	maxVis := m.opts.DropdownVisible
	if maxVis < 1 {
		maxVis = 8
	}
	end := dd.scrollOff + maxVis
	if end > len(dd.filtered) {
		end = len(dd.filtered)
	}
	visible := dd.filtered[dd.scrollOff:end]

	rows := make([]string, 0, len(visible))
	for i, cmd := range visible {
		absIdx := dd.scrollOff + i
		key := m.opts.CommandPrefix + cmd.Key
		args := cmd.Args
		desc := cmd.Desc

		if absIdx == dd.selected {
			row := accentBar +
				selGap +
				selKeyStyle.Render(key) +
				selArgsStyle.Render(args) +
				selDescStyle.Width(descW).Render(desc)
			rows = append(rows, row)
		} else {
			row := accentBar +
				accentGap +
				keyStyle.Render(key) +
				argsStyle.Render(args) +
				descStyle.Width(descW).Render(desc)
			rows = append(rows, row)
		}
	}

	return strings.Join(rows, "\n")
}

// -- Helpers ----------

// filterCommands
func filterCommands(cmds []InputCommand, query string) []InputCommand {
	if query == "" {
		return cmds
	}

	q := strings.ToLower(query)
	out := make([]InputCommand, 0, len(cmds))

	for _, c := range cmds {
		// Match on Key prefix
		if strings.HasPrefix(strings.ToLower(c.Key), q) {
			out = append(out, c)
		}
	}

	return out
}
