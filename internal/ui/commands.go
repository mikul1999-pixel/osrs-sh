package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

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

// -- Helpers ----------

// parseCommand parses the input and returns a NavigateMsg or an error string.
func parseCommand(raw string) (*NavigateMsg, string) {
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
