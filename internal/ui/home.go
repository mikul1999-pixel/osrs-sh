package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/core/render"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

// command describes a slash command shown on the home screen.
type command struct {
	cmd         string
	description string
	keybind     string
	targetTab   int
}

var commands = []command{
	{"/xp", "skill xp calculator", "2", TabXP},
	{"/npc", "monster stats & drops", "3", TabMonster},
	{"/item", "item info & ge price", "4", TabItem},
	{"/rsn", "player lookup", "5", TabPlayer},
}

// ASCII logo
// Upper half bright, lower half dimmed
var logoLines = render.GetLogo()

// HelpLine is the hint at the bottom of the text input
type HelpLine struct {
	Hint      string
	BeforeCmd string
	Command   string
	AfterCmd  string
}

// HomeModel is the home screen tab.
type HomeModel struct {
	width  int
	height int
	input  components.Input
	err    string
}

func NewHomeModel() HomeModel {
	helpText := RenderHelpLine(BgInput, HelpLine{
		Hint:      "● hint",
		BeforeCmd: "run",
		Command:   "/rsn <name>",
		AfterCmd:  " ",
	})

	input := components.NewInput(components.InputOptions{
		Placeholder: "type a command...",
		// Prompt:           "> ",
		// PromptStyle:      HomeInputPlaceholder,
		CharLimit:        80,
		AccentFocused:    lipgloss.Color(ColorGold),
		AccentUnfocused:  lipgloss.Color(ColorBorder),
		Background:       lipgloss.Color(ColorBgInput),
		TextStyle:        InputPrompt,
		PlaceholderStyle: InputPlaceholder,
		CursorStyle:      InputCursor,
		ShowBottomRow:    true,
		BottomRight:      helpText,
		PaddingTop:       1,
		PaddingMiddle:    0,
		PaddingBottom:    0,
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
		case "enter":
			nav, err := m.parseCommand(m.input.Value())
			m.err = err
			if nav != nil {
				m.input.SetValue("")
				// lookup /rsn
				if nav.Tab == TabPlayer && nav.Query != "" {
					helpText := RenderHelpLine(BgInput, HelpLine{
						Hint:      "● hint",
						BeforeCmd: "begin typing",
						Command:   "/",
						AfterCmd:  " ",
					})
					m.input.SetBottomRight(helpText)
					return m, func() tea.Msg {
						return LoadPlayerMsg{RSN: nav.Query}
					}
				}
				return m, func() tea.Msg { return *nav }
			}
			return m, nil
		case "esc":
			m.input.SetValue("")
			m.err = ""
			return m, nil
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// parseCommand parses the input and returns a NavigateMsg or an error string.
func (m HomeModel) parseCommand(raw string) (*NavigateMsg, string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, ""
	}

	parts := strings.Fields(raw)
	slug := strings.ToLower(parts[0])

	var query string
	if len(parts) > 1 {
		query = strings.Join(parts[1:], " ")
	}

	for _, c := range commands {
		if slug == c.cmd {
			return &NavigateMsg{Tab: c.targetTab, Query: query}, ""
		}
	}
	return nil, fmt.Sprintf("unknown command %q — try /xp, /npc, /item, /rsn", slug)
}

func (m HomeModel) View() string {
	if m.width == 0 {
		return ""
	}

	var sections []string

	// -- Logo ----------
	mid := len(logoLines) / 2
	var logoBlock strings.Builder
	for i, line := range logoLines {
		if i < mid {
			logoBlock.WriteString(LogoStyle.Render(line))
		} else {
			logoBlock.WriteString(LogoDimStyle.Render(line))
		}
		if i < len(logoLines)-1 {
			logoBlock.WriteString("\n")
		}
	}
	logo := Bg.Width(m.width).Align(lipgloss.Center).Render(logoBlock.String())

	subTitlePadding := Space(2)
	subTitleLabel := GetCwdDisplay(CwdOptions{
		ShortenHome:   true,
		LastOnly:      true,
		RootLabel:     "/",
		FallbackValue: "~/?",
	})
	subTitle := Bg.Width(m.width/2 + 18).Align(lipgloss.Right).
		Render(HomeSubTitleStyle.Render(subTitleLabel + subTitlePadding))

	sections = append(sections, "\n\n", logo, "\n", subTitle, "\n\n\n")

	// -- Command List ----------
	// Columns: command | description | keybind | tab
	colCmdW := 12
	colDescW := 28
	var cmdLines strings.Builder
	for _, c := range commands {
		cmdCol := HomeCmdStyle.Width(colCmdW).Render(c.cmd)
		descCol := HomeDescStyle.Width(colDescW).Render(c.description)
		keyCol := HomeKeybindStyle.Render("alt+" + c.keybind)
		tabCol := HomeKeybindStyle.Render(c.keybind)
		cmdLines.WriteString(cmdCol + Space(2) + descCol + Space(2) + keyCol + Space(2) + tabCol + "\n")
	}
	cmdBlock := Bg.Width(m.width).Align(lipgloss.Center).Render(cmdLines.String())
	sections = append(sections, cmdBlock, "\n\n")

	// Shared container for input elements
	inputColumn := Bg.Width(m.width / 2).Align(lipgloss.Left)

	// -- Input Box ----------
	inputRow := Bg.Width(m.width).Align(lipgloss.Center).
		Render(inputColumn.Render(m.input.View()))
	sections = append(sections, inputRow)

	// -- Input info ----------
	info := HelpStyle.Render("enter") + HelpStyleMuted.Render(" run")

	infoRow := Bg.Width(m.width).Align(lipgloss.Center).PaddingLeft(2).
		Render(inputColumn.Render(info))
	sections = append(sections, "\n", infoRow)

	// -- Error ----------
	if m.err != "" {
		errRow := Bg.Width(m.width).Align(lipgloss.Center).PaddingLeft(2).
			Render(inputColumn.Render(ErrorStyle.Render("x " + m.err)))
		sections = append(sections, "\n", errRow)
	}

	// Vertically center everything
	content := strings.Join(sections, "")
	contentH := lipgloss.Height(content)
	topPad := (m.height - contentH) / 2
	if topPad > 0 {
		content = strings.Repeat("\n", topPad) + content
	}

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Background(lipgloss.Color(ColorBg)).
		Render(content)
}

// -- Helpers ----------

func RenderHelpLine(bg lipgloss.Style, h HelpLine) string {
	hintStyle := bg.Foreground(lipgloss.Color(ColorPrimary)).Faint(true)
	cmdStyle := bg.Foreground(lipgloss.Color(ColorTextLight)).Faint(true)
	textStyle := bg.Foreground(lipgloss.Color(ColorMuted))

	var parts []string

	if h.Hint != "" {
		parts = append(parts, hintStyle.Render(h.Hint))
	}
	if h.BeforeCmd != "" {
		parts = append(parts, textStyle.Render(h.BeforeCmd))
	}
	if h.Command != "" {
		parts = append(parts, cmdStyle.Render(h.Command))
	}
	if h.AfterCmd != "" {
		parts = append(parts, textStyle.Render(h.AfterCmd))
	}

	return strings.Join(parts, SpaceInput(1))
}
