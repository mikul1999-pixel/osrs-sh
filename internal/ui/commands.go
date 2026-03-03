package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

// -- Helper File for Commands ----------

// CommandNav describes a slash command that navigates tabs
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
	{slug: "/wiki", args: "query", desc: "Open OSRS wiki", targetTab: 0, action: ActionLookup},

	// Bookmarks
	{slug: "/rsn-add", args: "username", desc: "Bookmark a username", targetTab: 0, action: ActionBookmark},
	{slug: "/rsn-rm", args: "username", desc: "Remove a bookmark", targetTab: 0, action: ActionBookmark},
	{slug: "/rsn-list", args: "", desc: "Show saved usernames", targetTab: 0, action: ActionBookmark},

	// Navigate (no-arg tab switchers)
	{slug: "/xp", args: "", desc: "Switch to xp tab", targetTab: TabXP, action: ActionNavigate},
	{slug: "/npc", args: "", desc: "Switch to npc tab", targetTab: TabMonster, action: ActionNavigate},
	{slug: "/item", args: "", desc: "Switch to item tab", targetTab: TabItem, action: ActionNavigate},
	{slug: "/rsn", args: "", desc: "Switch to rsn tab", targetTab: TabPlayer, action: ActionNavigate},

	// Session
	{slug: "/history", args: "", desc: "Show recent commands", targetTab: 0, action: ActionSession},
	{slug: "/clear", args: "", desc: "Clear current output", targetTab: 0, action: ActionSession},
	{slug: "/new", args: "", desc: "New session", targetTab: 0, action: ActionSession},
	{slug: "/theme", args: "", desc: "Cycle themes", targetTab: 0, action: ActionSession},
	{slug: "/exit", args: "", desc: "Exit the app", targetTab: 0, action: ActionSession},
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

// HelpLine is the hint at the bottom of the text input
type HelpLine struct {
	Hint      string
	BeforeCmd string
	Command   string
	AfterCmd  string
}

// -- Helpers ----------

// parseCommand parses the input and returns a CommandMsg or an error string.
func parseCommand(raw string) (*CommandMsg, string) {
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
		return &CommandMsg{Tab: c.targetTab, Slug: slug, Query: query, Action: c.action}, ""
	}

	return nil, fmt.Sprintf("unknown command %q — try /xp, /npc, /item, /rsn", slug)
}

func insertCommand(cmd components.InputCommand) string {
	return "/" + cmd.Key + " "
}

// parseExecuteCommand handles immediate execution for no arg commands by parsing input + executing
func parseExecuteCommand(cmd components.InputCommand) tea.Cmd {
	exec, _ := parseCommandFromInput("/" + cmd.Key)
	if exec == nil {
		return nil
	}
	if exec.Action == ActionSession { // ActionSession have no args, but are not navigation cmds
		return executeCommand(*exec)
	}
	return func() tea.Msg { return *exec }
}

// executeCommand handles commands in priority order
func executeCommand(cmd CommandMsg) tea.Cmd {
	switch cmd.Action {
	case ActionLookup:
		switch cmd.Tab {
		case TabPlayer:
			if cmd.Query != "" {
				return func() tea.Msg { return LoadPlayerMsg{RSN: cmd.Query} }
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
		switch cmd.Slug {
		case "/history":
			// add later
		case "/clear":
			// add later
		case "/exit":
			return tea.Quit
		case "/theme":
			return func() tea.Msg { return OpenThemeMsg{} }
		default:
			return nil
		}
	}
	// ActionNavigate or unhandled lookup
	return func() tea.Msg { return cmd }
}

// parseCommandFromInput helps executeCommand call parseCommand
func parseCommandFromInput(raw string) (*CommandMsg, string) {
	raw = strings.TrimSpace(raw)
	parts := strings.Fields(raw)
	if len(parts) == 0 {
		return nil, ""
	}
	slug := strings.ToLower(parts[0])
	for _, c := range commands {
		if slug == c.slug && c.args == "" {
			return &CommandMsg{Tab: c.targetTab, Slug: slug, Query: "", Action: c.action}, ""
		}
	}
	return nil, fmt.Sprintf("unknown command %q", slug)
}

func GetTabCmds(activeTab int) []StatusKeybind {
	var activeCommands []StatusKeybind
	switch activeTab {
	case TabHome:
		// add later
	case TabXP:
		// activeCommands = []StatusKeybind{
		// 	{Key: "w", Label: "wiki"},
		// }

	case TabItem:
		// add later
	case TabMonster:
		// add later
	case TabPlayer:
		// add later
	}
	return activeCommands
}

// CommandHelp live generates command argument help for input
func CommandHelp(value string, bg lipgloss.Style, defaultHelpLine HelpLine) string {
	defaultHelp := renderCommandHelp(bg, defaultHelpLine)

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
			return renderCommandHelp(bg, HelpLine{
				Hint:     "",
				Command:  c.slug,
				AfterCmd: c.args + " ",
			})
		}
	}

	return defaultHelp
}

func renderCommandHelp(bg lipgloss.Style, h HelpLine) string {
	hintStyle := bg.Foreground(lipgloss.Color(ActiveTheme.Primary)).Faint(true)
	cmdStyle := bg.Foreground(lipgloss.Color(ActiveTheme.TextLight)).Faint(true)
	textStyle := bg.Foreground(lipgloss.Color(ActiveTheme.Muted))

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
	space := bg.Render(" ")

	return strings.Join(parts, space)
}
