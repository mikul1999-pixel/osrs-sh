package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

const (
	paletteWidth  = 64
	paletteHeight = 22
)

// PaletteModel is the command palette modal overlay
type PaletteModel struct {
	input   components.Input
	visible bool
	width   int
	height  int
}

func NewPaletteModel() PaletteModel {
	input := components.NewInput(components.InputOptions{
		Placeholder:      "begin typing /",
		CharLimit:        80,
		AccentFocused:    lipgloss.Color(ColorAccent),
		AccentUnfocused:  lipgloss.Color(ColorMuted),
		Background:       lipgloss.Color(ColorBgPanel),
		TextStyle:        InputPrompt,
		PlaceholderStyle: InputPlaceholder,
		CursorStyle:      InputCursor,
		ShowBottomRow:    false,
		PaddingTop:       0,
		PaddingMiddle:    0,
		PaddingBottom:    0,

		// Dropdown
		Commands:        buildInputCommands(),
		DropdownVisible: 18,
		ForceDropdown:   true,
	})
	input.Focus()
	return PaletteModel{input: input}
}

func (p PaletteModel) Init() tea.Cmd {
	return p.input.Init()
}

func (p PaletteModel) Open() PaletteModel {
	p.visible = true
	p.input.SetValue("/")
	p.input.Focus()
	p.input.RefreshDropdown()
	return p
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
	if !p.visible {
		return p, nil
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return p.Close(), nil
		case "enter":
			if selected := p.input.CommitDropdownSelection(false); selected != nil {
				if selected.Args == "" {
					p.visible = false
					return p, executeCommand(*selected)
				}
				p.input.SetValue(insertCommand(*selected))
				return p, nil
			}
			nav, _ := parseCommand(p.input.Value())
			if nav != nil {
				p.visible = false
				return p, executeNav(*nav)
			}
			return p, nil
		}
	}

	p.input, cmd = p.input.Update(msg)
	return p, cmd
}

func (p PaletteModel) View() string {
	if !p.visible {
		return ""
	}

	pad := lipgloss.NewStyle().
		Background(lipgloss.Color(ColorBgPanel)).
		Width(paletteWidth + 2).
		Height(1).
		Render("")

	panel := components.New(paletteWidth).
		Height(paletteHeight).
		PaddingFull(1, 1, 0, 1).
		Title("Commands").
		Badge(BodyDim.Render("esc")).
		BgColor(lipgloss.Color(ColorBgPanel)).
		ActiveBorderColor(lipgloss.Color(ColorBgPanel)).
		InactiveBorderColor(lipgloss.Color(ColorBgPanel))

	inputView := p.input.View()
	dropdownView := p.input.DropdownView()

	var content string
	if dropdownView != "" {
		content = inputView + "\n" + dropdownView
	} else {
		content = inputView
	}

	return lipgloss.JoinVertical(lipgloss.Left, pad, panel.Render(content, true))
}
