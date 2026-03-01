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

// CommandAction describes what a command does beyond tab navigation
type CommandAction int

const (
	ActionNavigate CommandAction = iota // switch tab
	ActionLookup                        // query an API or tab func
	ActionBookmark                      // config bookmarks
	ActionSession                       // app session
)

// command is the single source of truth for all slash commands
type command struct {
	slug      string
	args      string
	desc      string
	targetTab int
	action    CommandAction
}

var commands = []command{
	// Lookups
	{slug: "/rsn", args: "username", desc: "Lookup player stats", targetTab: TabPlayer, action: ActionLookup},
	{slug: "/xp", args: "lv1 lvl2", desc: "XP between two levels", targetTab: TabXP, action: ActionLookup},
	{slug: "/item", args: "name", desc: "Item price & info", targetTab: TabItem, action: ActionLookup},
	{slug: "/npc", args: "name", desc: "Monster stats & drops", targetTab: TabMonster, action: ActionLookup},
	{slug: "/wiki", args: "query", desc: "Open OSRS wiki", targetTab: -1, action: ActionLookup},

	// Bookmarks
	{slug: "/rsn-add", args: "username", desc: "Bookmark a username", targetTab: -1, action: ActionBookmark},
	{slug: "/rsn-rm", args: "username", desc: "Remove a bookmark", targetTab: -1, action: ActionBookmark},
	{slug: "/rsn-list", args: "", desc: "Show saved usernames", targetTab: -1, action: ActionBookmark},

	// Navigate (no-arg tab switchers)
	{slug: "/xp", args: "", desc: "Switch to xp tab", targetTab: TabXP, action: ActionNavigate},
	{slug: "/npc", args: "", desc: "Switch to npc tab", targetTab: TabMonster, action: ActionNavigate},
	{slug: "/item", args: "", desc: "Switch to item tab", targetTab: TabItem, action: ActionNavigate},
	{slug: "/rsn", args: "", desc: "Switch to rsn tab", targetTab: TabPlayer, action: ActionNavigate},

	// Session
	{slug: "/history", args: "", desc: "Show recent commands", targetTab: -1, action: ActionSession},
	{slug: "/clear", args: "", desc: "Clear current output", targetTab: -1, action: ActionSession},
	{slug: "/new", args: "", desc: "New session", targetTab: -1, action: ActionSession},
	{slug: "/theme", args: "", desc: "Cycle themes", targetTab: -1, action: ActionSession},
	{slug: "/exit", args: "", desc: "Exit the app", targetTab: -1, action: ActionSession},
}

func buildInputCommands() []components.InputCommand {
	out := make([]components.InputCommand, len(commands))
	for i, c := range commands {
		out[i] = components.InputCommand{
			Key:  strings.TrimPrefix(c.slug, "/"),
			Args: c.args,
			Desc: c.desc,
		}
	}
	return out
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
			if selected := m.input.CommitDropdownSelection(); selected != nil {
				if selected.Args == "" {
					return m, executeCommand(*selected)
				}
				m.input.SetValue(insertCommand(*selected))
				m.syncHelpText()
				return m, nil
			}

			// User manually typed command
			nav, err := m.parseCommand(m.input.Value())
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

	var sections []string

	// -- Logo and Commands ----------
	header, _, _ := m.HeaderPosition()
	sections = append(sections, "\n", header)

	// Shared container for input elements
	inputColumn := Bg.Width(m.width / 2).Align(lipgloss.Left)

	// -- Input Box ----------
	inputRow := Bg.Width(m.width).Align(lipgloss.Center).
		Render(inputColumn.Render(m.input.View()))
	sections = append(sections, inputRow)

	// -- Input info ----------
	cwd := GetCwdDisplay(CwdOptions{
		ShortenHome:   true,
		LastOnly:      true,
		RootLabel:     "/",
		FallbackValue: "~/?",
	})
	info := HelpStyle.Render("config ") + HelpStyleMuted.Render(cwd)

	infoRow := Bg.Width(m.width).Align(lipgloss.Center).PaddingLeft(2).
		Render(inputColumn.Render(info))
	sections = append(sections, "\n", infoRow)

	// -- Error ----------
	if m.err != "" {
		errRow := Bg.Width(m.width).Align(lipgloss.Center).PaddingLeft(2).
			Render(inputColumn.Render(ErrorStyle.Render("x " + m.err)))
		sections = append(sections, "\n", errRow)
	}

	// Combine header + input
	content := strings.Join(sections, "")

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Background(lipgloss.Color(ColorBg)).
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
			logoBlock.WriteString(LogoStyle.Render(line))
		} else {
			logoBlock.WriteString(LogoDimStyle.Render(line))
		}
		if i < len(logoLines)-1 {
			logoBlock.WriteString("\n")
		}
	}
	logo := Bg.Width(m.width).Align(lipgloss.Center).Render(logoBlock.String())
	sections = append(sections, "\n\n", logo, "\n\n")

	// -- Command List ----------
	// Columns: command | description | keybind | tab
	colCmdW := 12
	colDescW := 28
	var cmdLines strings.Builder
	for _, c := range CommandNavMenu {
		cmdCol := HomeCmdStyle.Width(colCmdW).Render(c.cmd)
		descCol := HomeDescStyle.Width(colDescW).Render(c.description)
		keyCol := HomeKeybindStyle.Render("alt+" + c.keybind)
		tabCol := HomeKeybindStyle.Render(c.keybind)
		cmdLines.WriteString(cmdCol + Space(2) + descCol + Space(2) + keyCol + Space(2) + tabCol + "\n")
	}
	cmdBlock := Bg.Width(m.width).Align(lipgloss.Center).Render(cmdLines.String())
	sections = append(sections, cmdBlock, "\n\n")

	return sections
}

// HeaderPosition centers the header vertically and returns terminal postion info
func (m HomeModel) HeaderPosition() (string, int, int) {
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

	_, headerH, pad := m.HeaderPosition()
	inputBoxH := lipgloss.Height(m.input.View())
	y = pad + headerH + inputBoxH

	return panel, x, y
}

// -- Helpers ----------

// parseCommand parses the input and returns a NavigateMsg or an error string.
func (m HomeModel) parseCommand(raw string) (*NavigateMsg, string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, ""
	}

	parts := strings.Fields(raw)
	slug := strings.ToLower(parts[0])
	query := ""
	if len(parts) > 1 {
		query = strings.Join(parts[1:], " ")
	}

	hasArgs := query != ""

	for _, c := range commands {
		if slug != c.slug {
			continue
		}
		// match no arg commands
		if hasArgs && c.args == "" {
			continue
		}
		if !hasArgs && c.args != "" {
			continue
		}
		return &NavigateMsg{Tab: c.targetTab, Query: query, Action: c.action}, ""
	}

	return nil, fmt.Sprintf("unknown command %q — try /xp, /npc, /item, /rsn", slug)
}

func insertCommand(cmd components.InputCommand) string {
	return "/" + cmd.Key + " "
}

// executeCommand handles immediate execution for no arg commands
func executeCommand(cmd components.InputCommand) tea.Cmd {
	nav, _ := parseCommandFromInput("/" + cmd.Key)
	if nav == nil {
		return nil
	}
	return func() tea.Msg { return *nav }
}

// executeNav intercepts and handles command in priority order
func executeNav(nav NavigateMsg) tea.Cmd {
	switch nav.Action {
	case ActionLookup:
		switch nav.Tab {
		case TabPlayer:
			if nav.Query != "" {
				return func() tea.Msg { return LoadPlayerMsg{RSN: nav.Query} }
			}
		case TabXP:
			// add later
		case TabItem:
			// add later
		case TabMonster:
			// add later
		}
	case ActionBookmark:
		// add later
	case ActionSession:
		// /history, /clear, /theme, /exit ...add later
	}
	// ActionNavigate or unhandled lookup
	return func() tea.Msg { return nav }
}

// parseCommandFromInput helps executeCommand call parseCommand
func parseCommandFromInput(raw string) (*NavigateMsg, string) {
	raw = strings.TrimSpace(raw)
	parts := strings.Fields(raw)
	if len(parts) == 0 {
		return nil, ""
	}
	slug := strings.ToLower(parts[0])
	for _, c := range commands {
		if slug == c.slug && c.args == "" {
			return &NavigateMsg{Tab: c.targetTab, Query: "", Action: c.action}, ""
		}
	}
	return nil, fmt.Sprintf("unknown command %q", slug)
}

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

// deriveHelpText live generates help text for input based on what user types
func deriveHelpText(value string) string {
	defaultHelp := RenderHelpLine(BgInput, HelpLine{
		Hint:     "● hint",
		Command:  "begin typing /",
		AfterCmd: " ",
	})

	value = strings.TrimSpace(value)
	if value == "" || !strings.HasPrefix(value, "/") {
		return defaultHelp
	}

	parts := strings.Fields(value)
	slug := strings.ToLower(parts[0])

	// Still typing the command name
	if len(parts) == 1 && !strings.HasSuffix(value, " ") {
		return defaultHelp
	}

	// Find match command with args
	for _, c := range commands {
		if c.slug == slug && c.args != "" {
			return RenderHelpLine(BgInput, HelpLine{
				Hint:     "",
				Command:  c.slug,
				AfterCmd: c.args + " ",
			})
		}
	}

	return defaultHelp
}

func (m *HomeModel) syncHelpText() {
	m.input.SetBottomRight(deriveHelpText(m.input.Value()))
}
