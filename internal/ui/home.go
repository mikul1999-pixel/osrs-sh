package ui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/core/render"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

// command describes a slash command shown on the home screen.
type CommandNav struct {
	cmd         string
	description string
	keybind     string
}

var CommandNavMenu = []CommandNav{
	{"/xp", "skill xp calculator", "2"},
	{"/npc", "monster stats & drops", "3"},
	{"/item", "item info & ge price", "4"},
	{"/rsn", "player lookup", "5"},
}

// ASCII logo
// Upper half bright, lower half dimmed
var logoLines = render.GetLogo()

// HomeModel is the home screen tab.
type HomeModel struct {
	width  int
	height int
	input  components.Input
	err    string
}

func NewHomeModel() HomeModel {
	input := components.NewInput(components.InputOptions{
		Placeholder: "type a command...",
		// Prompt:           "> ",
		// PromptStyle:      HomeInputPlaceholder,
		CharLimit:        80,
		AccentFocused:    ActiveTheme.Primary,
		AccentUnfocused:  ActiveTheme.Border,
		Background:       ActiveTheme.BgInput,
		TextStyle:        ActiveTheme.InputPrompt(),
		PlaceholderStyle: ActiveTheme.InputPlaceholder(),
		CursorStyle:      ActiveTheme.InputCursor(),
		ShowBottomRow:    true,
		PaddingTop:       1,
		PaddingMiddle:    0,
		PaddingBottom:    0,

		// Dropdown
		Commands:        buildInputCommands(),
		DropdownVisible: 8,
	})
	input.Focus()
	return HomeModel{input: input}
}

func (m *HomeModel) SetSize(w, h int) {
	m.width = w
	m.height = h
	m.input.SetWidth(w / 2)
}

func (m HomeModel) Init() tea.Cmd {
	return m.input.Init()
}

func (m HomeModel) Update(msg tea.Msg) (HomeModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.input.SetValue("")
			m.err = ""
			m.syncHelpText()
			return m, nil

		case "enter":
			// User selected command from dropdown
			if selected := m.input.CommitDropdownSelection(true); selected != nil {
				if selected.Args == "" {
					return m, executeCommand(*selected)
				}
				m.input.SetValue(insertCommand(*selected))
				m.syncHelpText()
				return m, nil
			}

			// User manually typed command
			nav, err := parseCommand(m.input.Value())
			m.err = err
			if nav != nil {
				m.input.SetValue("")
				m.syncHelpText()
				return m, executeNav(*nav)
			}
			return m, nil
		}
	}

	m.input, cmd = m.input.Update(msg)
	m.syncHelpText()
	return m, cmd
}

func (m HomeModel) View() string {
	if m.width == 0 {
		return ""
	}
	m.syncInputTheme()
	m.syncHelpText()

	var sections []string

	// -- Logo and Commands ----------
	header, _, _ := m.headerPosition()
	sections = append(sections, "\n", header)

	// Shared container for input elements
	inputColumn := ActiveTheme.Bg_().Width(m.width / 2).Align(lipgloss.Left)

	// -- Input Box ----------
	inputRow := ActiveTheme.Bg_().Width(m.width).Align(lipgloss.Center).
		Render(inputColumn.Render(m.input.View()))
	sections = append(sections, inputRow)

	// -- Input info ----------
	cwd := getCwdDisplay(CwdOptions{
		ShortenHome:   true,
		LastOnly:      true,
		RootLabel:     "/",
		FallbackValue: "~/?",
	})
	info := ActiveTheme.Help().Render("config ") + ActiveTheme.HelpMuted().Render(cwd)

	infoRow := ActiveTheme.Bg_().Width(m.width).Align(lipgloss.Center).PaddingLeft(2).
		Render(inputColumn.Render(info))
	sections = append(sections, "\n", infoRow)

	// -- Error ----------
	if m.err != "" {
		errRow := ActiveTheme.Bg_().Width(m.width).Align(lipgloss.Center).PaddingLeft(2).
			Render(inputColumn.Render(ActiveTheme.Error().Render("x " + m.err)))
		sections = append(sections, "\n", errRow)
	}

	// Combine header + input
	content := strings.Join(sections, "")

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Background(lipgloss.Color(ActiveTheme.Bg)).
		Render(content)
}

// -- View Components ----------

// HeaderView returns the logo and command panel (above input box)
func (m HomeModel) HeaderView() (sections []string) {

	// -- Logo ----------
	mid := len(logoLines) / 2
	var logoBlock strings.Builder
	for i, line := range logoLines {
		if i < mid {
			logoBlock.WriteString(ActiveTheme.Logo().Render(line))
		} else {
			logoBlock.WriteString(ActiveTheme.LogoDim().Render(line))
		}
		if i < len(logoLines)-1 {
			logoBlock.WriteString("\n")
		}
	}
	logo := ActiveTheme.Bg_().Width(m.width).Align(lipgloss.Center).Render(logoBlock.String())
	sections = append(sections, "\n\n", logo, "\n\n")

	// -- Command List ----------
	// Columns: command | description | keybind | tab
	colCmdW := 12
	colDescW := 28
	var cmdLines strings.Builder
	for _, c := range CommandNavMenu {
		cmdCol := ActiveTheme.HomeCmd().Width(colCmdW).Render(c.cmd)
		descCol := ActiveTheme.HomeDesc().Width(colDescW).Render(c.description)
		keyCol := ActiveTheme.HomeKeybind().Render("alt+" + c.keybind)
		tabCol := ActiveTheme.HomeKeybind().Render(c.keybind)
		cmdLines.WriteString(cmdCol + Space(2) + descCol + Space(2) + keyCol + Space(2) + tabCol + "\n")
	}
	cmdBlock := ActiveTheme.Bg_().Width(m.width).Align(lipgloss.Center).Render(cmdLines.String())
	sections = append(sections, cmdBlock, "\n\n")

	return sections
}

// HeaderPosition centers the header vertically and returns terminal postion info
func (m HomeModel) headerPosition() (string, int, int) {
	// Height of the header
	logoH := len(logoLines)
	cmdH := len(CommandNavMenu)
	headerH := 2 + logoH + 2 + cmdH + 2

	// Compute top padding
	pad := (m.height - headerH) / 2
	if pad < 0 {
		pad = 0
	}

	// Render header
	header := strings.Join(m.HeaderView(), "")
	centered := strings.Repeat("\n", pad) + header

	return centered, headerH, pad
}

// DropdownOverlay returns the rendered dropdown panel and its position within view
func (m HomeModel) DropdownOverlay() (panel string, x, y int) {
	panel = m.input.DropdownView()
	if panel == "" {
		return "", 0, 0
	}

	x = (m.width - m.width/2) / 2

	_, headerH, pad := m.headerPosition()
	inputBoxH := lipgloss.Height(m.input.View())
	y = pad + headerH + inputBoxH

	return panel, x, y
}

// -- Helpers ----------

func (m *HomeModel) syncHelpText() {
	bg := ActiveTheme.BgInput_()
	defaultHelp := HelpLine{
		Hint:     "● hint",
		Command:  "begin typing /",
		AfterCmd: " ",
	}
	m.input.SetBottomRight(CommandHelp(m.input.Value(), bg, defaultHelp))
}

// syncInputTheme ensures input box styling can render in View and update
func (m *HomeModel) syncInputTheme() {
	m.input.SetAccentFocused(lipgloss.Color(ActiveTheme.Primary))
	m.input.SetAccentUnfocused(lipgloss.Color(ActiveTheme.Border))
	m.input.SetBackground(lipgloss.Color(ActiveTheme.BgInput))
	m.input.SetTextStyle(ActiveTheme.InputPrompt())
	m.input.SetPlaceholderStyle(ActiveTheme.InputPlaceholder())
	m.input.SetCursorStyle(ActiveTheme.InputCursor())
}
