package ui

import (
	"sort"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

const (
	themeWidth  = 52
	themeHeight = 18
)

// ThemeModel is the theme picker modal overlay
type ThemeModel struct {
	input   components.Input
	help    string
	visible bool
	width   int
	height  int
}

func buildInputThemes() []components.InputCommand {
	keys := make([]string, 0, len(Themes))
	for key := range Themes {
		keys = append(keys, key)
	}

	// Sort alphabetically
	sort.Strings(keys)
	out := make([]components.InputCommand, 0, len(keys))
	for _, key := range keys {
		out = append(out, components.InputCommand{
			Key: key,
		})
	}

	return out
}

func NewThemeModel() ThemeModel {
	// do not update: prevent from dimmed() impact
	input := components.NewInput(components.InputOptions{
		Placeholder:      "Search",
		CharLimit:        80,
		AccentFocused:    lipgloss.Color(ActiveTheme.Primary),
		AccentUnfocused:  lipgloss.Color(ActiveTheme.Muted),
		Background:       lipgloss.Color(ActiveTheme.BgModal),
		TextStyle:        ActiveTheme.InputPrompt(),
		PlaceholderStyle: ActiveTheme.InputPlaceholder(),
		CursorStyle:      ActiveTheme.InputCursor(),
		ShowBottomRow:    false,
		PaddingTop:       0,
		PaddingMiddle:    0,
		PaddingBottom:    0,

		// Dropdown
		CommandPrefix:   "",
		Commands:        buildInputThemes(),
		DropdownVisible: 10,
		DropdownAccent:  lipgloss.Color(ActiveTheme.BgModalList),
		ForceDropdown:   true,
		FilterDropdown:  true,
	})
	input.Focus()
	return ThemeModel{input: input}
}

func (t ThemeModel) Init() tea.Cmd {
	return nil
}

func (t ThemeModel) Open() (ThemeModel, tea.Cmd) {
	t.visible = true
	t.input.SetValue("")
	t.help = ""
	t.input.Focus()
	t.input.RefreshDropdown()
	return t, t.input.Init()
}

func (t ThemeModel) Close() ThemeModel {
	t.visible = false
	t.input.Blur()
	return t
}

func (t ThemeModel) SetSize(w, h int) ThemeModel {
	t.width = w
	t.height = h
	t.input.SetWidth(themeWidth - 4) // account for panel padding
	return t
}

func (t ThemeModel) Update(msg tea.Msg) (ThemeModel, tea.Cmd) {
	t.syncInputTheme()
	var cmd tea.Cmd

	// Force cursor blink by always forwarding messages to child
	t.input, cmd = t.input.Update(msg)

	if !t.visible {
		return t, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return t.Close(), nil
		// case "enter":
		case "enter", "up", "down":
			if selected := t.input.CommitDropdownSelection(false); selected != nil {
				t.input.SetValue("/" + selected.Key)
				cmd = setTheme(selected.Key)
				t.input.Reset()
				return t, cmd
			}
			cleanValue := strings.ReplaceAll(t.input.Value(), "/", "")
			cmd = setTheme(cleanValue)
			t.input.Reset()
			return t, cmd
		}
	}

	return t, cmd
}

func (t ThemeModel) View() string {
	if !t.visible {
		return ""
	}
	t.syncInputTheme()

	pad := lipgloss.NewStyle().
		Background(lipgloss.Color(ActiveTheme.BgModal)).
		Width(themeWidth + 2).
		Height(1).
		Render("")

	panel := components.New(themeWidth).
		Height(themeHeight).
		PaddingFull(1, 1, 0, 1).
		Title(ActiveTheme.PanelTitleMenu().Render("Themes")).
		Badge(ActiveTheme.PanelBadgeMenu().Render("esc")).
		BgColor(lipgloss.Color(ActiveTheme.BgModal)).
		ActiveBorderColor(lipgloss.Color(ActiveTheme.BgModal)).
		InactiveBorderColor(lipgloss.Color(ActiveTheme.BgModal))

	inputView := t.input.View()
	dropdownView := t.input.DropdownView()

	var content string
	if dropdownView != "" {
		content = inputView + "\n" + "\n" + dropdownView
	} else {
		content = inputView + "\n"
	}

	return lipgloss.JoinVertical(lipgloss.Left, pad, panel.Render(content, true))
}

// -- Helpers ----------

func setTheme(themeVal string) tea.Cmd {
	return func() tea.Msg { return ChangeThemeMsg{theme: loadThemeFromConfig(themeVal), name: themeVal} }
}

// syncInputTheme ensures input box styling can render in View and update
func (t *ThemeModel) syncInputTheme() {
	t.input.SetAccentFocused(lipgloss.Color(ActiveTheme.PrimaryModal))
	t.input.SetAccentUnfocused(lipgloss.Color(ActiveTheme.PrimaryModal))
	t.input.SetBackground(lipgloss.Color(ActiveTheme.BgModal))
	t.input.SetTextStyle(ActiveTheme.InputPromptModal())
	t.input.SetPlaceholderStyle(ActiveTheme.InputPlaceholderModal())
	t.input.SetCursorStyle(ActiveTheme.InputCursorModal())
	t.input.SetDropdownAccent(lipgloss.Color(ActiveTheme.BgModalList))
}
