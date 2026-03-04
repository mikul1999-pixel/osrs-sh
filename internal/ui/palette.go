package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

// -- Command Palette Modal ----------

const (
	paletteWidth  = 64
	paletteHeight = 22
)

// PaletteModel is the command palette modal overlay
type PaletteModel struct {
	input   components.Input
	help    string
	visible bool
	width   int
	height  int
}

func NewPaletteModel() PaletteModel {
	// do not update: prevent from dimmed() impact
	input := components.NewInput(components.InputOptions{
		Placeholder:      "begin typing /",
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
		Commands:        buildInputCommands(),
		CommandPrefix:   "/",
		DropdownVisible: 17,
		DropdownTrigger: '/',
		DropdownAccent:  lipgloss.Color(ActiveTheme.BgModalList),
		ForceDropdown:   false,
		FilterDropdown:  true,
	})
	input.Focus()
	return PaletteModel{input: input, help: ""}
}

func (p PaletteModel) Init() tea.Cmd {
	return nil
}

func (p PaletteModel) Open() (PaletteModel, tea.Cmd) {
	p.visible = true
	p.input.SetValue("/")
	p.help = ""
	p.input.Focus()
	p.input.RefreshDropdown()
	return p, p.input.Init()
}

func (p PaletteModel) Close() PaletteModel {
	p.visible = false
	p.input.Blur()
	return p
}

func (p PaletteModel) SetSize(w, h int) PaletteModel {
	p.width = w
	p.height = h
	p.input.SetWidth(paletteWidth - 4) // account for panel padding
	return p
}

func (p PaletteModel) Update(msg tea.Msg) (PaletteModel, tea.Cmd) {
	p.syncInputTheme()
	var cmd tea.Cmd

	// Force cursor blink by always forwarding messages to child
	p.input, cmd = p.input.Update(msg)

	if !p.visible {
		return p, cmd
	}

	// var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p.Close(), nil
		case "enter":
			if selected := p.input.CommitDropdownSelection(false); selected != nil {
				if selected.Args == "" {
					p.visible = false
					return p, parseExecuteCommand(*selected)
				}
				p.input.SetValue(insertCommand(*selected))
				p.help = p.inputHelpText()
				return p, nil
			}
			nav, _ := parseCommand(p.input.Value())
			if nav != nil {
				p.visible = false
				return p, executeCommand(*nav)
			}
			return p, nil
		}
	}

	p.help = p.inputHelpText()
	return p, cmd
}

func (p PaletteModel) View() string {
	if !p.visible {
		return ""
	}
	p.syncInputTheme()

	pad := lipgloss.NewStyle().
		Background(lipgloss.Color(ActiveTheme.BgModal)).
		Width(paletteWidth + 2).
		Height(1).
		Render("")

	panel := components.New(paletteWidth).
		Height(paletteHeight).
		PaddingFull(1, 1, 0, 1).
		Title(ActiveTheme.PanelTitleMenu().Render("Commands")).
		Badge(ActiveTheme.PanelBadgeMenu().Render("esc")).
		BgColor(lipgloss.Color(ActiveTheme.BgModal)).
		ActiveBorderColor(lipgloss.Color(ActiveTheme.BgModal)).
		InactiveBorderColor(lipgloss.Color(ActiveTheme.BgModal))

	inputView := p.input.View()
	dropdownView := p.input.DropdownView()
	helpText := p.help
	helpGap := lipgloss.NewStyle().Background(lipgloss.Color(ActiveTheme.BgModal)).Render("  ")

	var content string
	if dropdownView != "" {
		content = inputView + "\n" + helpGap + helpText + "\n" + dropdownView
	} else {
		content = inputView + "\n" + helpGap + helpText
	}

	return lipgloss.JoinVertical(lipgloss.Left, pad, panel.Render(content, true))
}

// -- Helpers ----------

func (p *PaletteModel) inputHelpText() string {
	bg := lipgloss.NewStyle().Background(lipgloss.Color(ActiveTheme.BgModal))
	defaultHelp := HelpLine{
		Hint:     "",
		Command:  "",
		AfterCmd: "",
	}
	return CommandHelp(p.input.Value(), bg, defaultHelp)
}

// syncInputTheme ensures input box styling can render in View and update
func (p *PaletteModel) syncInputTheme() {
	p.input.SetAccentFocused(lipgloss.Color(ActiveTheme.PrimaryModal))
	p.input.SetAccentUnfocused(lipgloss.Color(ActiveTheme.PrimaryModal))
	p.input.SetBackground(lipgloss.Color(ActiveTheme.BgModal))
	p.input.SetTextStyle(ActiveTheme.InputPromptModal())
	p.input.SetPlaceholderStyle(ActiveTheme.InputPlaceholderModal())
	p.input.SetCursorStyle(ActiveTheme.InputCursorModal())
	p.input.SetDropdownAccent(lipgloss.Color(ActiveTheme.BgModalList))
}
