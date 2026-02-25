package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/core/render"
	"github.com/mikul1999-pixel/osrs-sh/internal/core/xp"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

// -- Skill definitions ----------

type skill struct {
	name   string
	abbrev string
}

var skills = []skill{
	{"Attack", "ATK"},
	{"Hitpoints", "HP"},
	{"Mining", "MIN"},
	{"Strength", "STR"},
	{"Agility", "AGI"},
	{"Smithing", "SMI"},
	{"Defence", "DEF"},
	{"Herblore", "HRB"},
	{"Fishing", "FSH"},
	{"Ranged", "RNG"},
	{"Thieving", "THV"},
	{"Cooking", "COK"},
	{"Prayer", "PRA"},
	{"Crafting", "CRF"},
	{"Firemaking", "FMK"},
	{"Magic", "MAG"},
	{"Fletching", "FLT"},
	{"Woodcutting", "WC"},
	{"Runecraft", "RC"},
	{"Slayer", "SLY"},
	{"Farming", "FRM"},
	{"Construction", "CON"},
	{"Hunter", "HNT"},
	{"Sailing", "SAI"},
}

const gridCols = 3

// -- Input mode ----------

type inputMode int

const (
	modeCurrent inputMode = iota
	modeTarget
)

// -- Image message ----------

type imageLoadedMsg struct {
	skill string
	ansi  string
	err   error
}

func loadSkillImage(skillName string) tea.Cmd {
	return func() tea.Msg {
		if ansi, ok := render.GetSkillIcon(skillName); ok {
			return imageLoadedMsg{skill: skillName, ansi: ansi}
		}
		url := render.SkillIconURL(skillName)
		ansi, err := render.ImageToANSI(url, "20x12")
		return imageLoadedMsg{skill: skillName, ansi: ansi, err: err}
	}
}

// -- Model ----------

type XPModel struct {
	width    int
	height   int
	selected int

	levels  [24]int
	targets [24]int

	mode inputMode

	input    components.Input
	inputErr string

	currentImage string
	imageLoading bool
	imageErr     string
}

func NewXPModel() XPModel {
	input := components.NewInput(components.InputOptions{
		CharLimit:        12,
		Placeholder:      "type level or xp...",
		AccentFocused:    lipgloss.Color(ColorBlue),
		AccentUnfocused:  lipgloss.Color(ColorBorder),
		Background:       lipgloss.Color(ColorBgInput),
		TextStyle:        InputPrompt,
		PlaceholderStyle: InputPlaceholder,
		CursorStyle:      InputCursor,
		ShowBottomRow:    true,
		PaddingTop:       0,
		PaddingMiddle:    0,
		PaddingBottom:    0,
	})
	input.Focus()

	var levels [24]int
	var targets [24]int
	for i := range levels {
		levels[i] = 1
		targets[i] = 99
	}

	m := XPModel{input: input, levels: levels, targets: targets}
	m.syncInputToMode()
	return m
}

// syncInputToMode updates placeholder and pre-fills input value from stored state.
func (m *XPModel) syncInputToMode() {
	switch m.mode {
	case modeCurrent:
		stored := m.levels[m.selected]
		if stored > 1 {
			m.input.SetValue(fmt.Sprintf("%d", stored))
		} else {
			m.input.SetValue("")
		}
	case modeTarget:
		stored := m.targets[m.selected]
		if stored < 99 {
			m.input.SetValue(fmt.Sprintf("%d", stored))
		} else {
			m.input.SetValue("")
		}
	}
}

func (m *XPModel) SetSize(w, h int) {
	m.width = w
	m.height = h
}

func (m *XPModel) SetQuery(q string) {
	m.input.SetValue(q)
	m.input.Focus()
}

func (m XPModel) Init() tea.Cmd {
	return tea.Batch(m.input.Init(), loadSkillImage(skills[m.selected].name))
}

// -- Update ----------

func (m XPModel) Update(msg tea.Msg) (XPModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case imageLoadedMsg:
		if msg.err != nil {
			m.imageErr = msg.err.Error()
		} else {
			m.currentImage = msg.ansi
			m.imageErr = ""
		}
		m.imageLoading = false
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			m.selected -= gridCols
			if m.selected < 0 {
				col := (m.selected + gridCols) % gridCols
				prevCol := (col - 1 + gridCols) % gridCols
				lastRow := (len(skills) - 1) / gridCols
				m.selected = lastRow*gridCols + prevCol
				for m.selected >= len(skills) {
					m.selected -= gridCols
				}
			}
			m.syncInputToMode()
			m.imageLoading = true
			return m, loadSkillImage(skills[m.selected].name)

		case "down", "j":
			m.selected += gridCols
			if m.selected >= len(skills) {
				col := (m.selected - gridCols) % gridCols
				nextCol := (col + 1) % gridCols
				m.selected = nextCol
			}
			m.syncInputToMode()
			m.imageLoading = true
			return m, loadSkillImage(skills[m.selected].name)

		case "tab":
			if m.input.Focused() {
				// Save current input then toggle mode
				m.saveCurrentInput()
				if m.mode == modeCurrent {
					m.mode = modeTarget
				} else {
					m.mode = modeCurrent
				}
				m.syncInputToMode()
			} else {
				m.input.Focus()
			}

		case "enter":
			m.saveCurrentInput()
			m.inputErr = ""
			return m, nil
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *XPModel) saveCurrentInput() {
	_, level := parseXPInput(strings.TrimSpace(m.input.Value()))
	switch m.mode {
	case modeCurrent:
		if level >= 1 {
			m.levels[m.selected] = level
		}
	case modeTarget:
		if level >= 1 {
			m.targets[m.selected] = level
		}
	}
}

// -- View ----------

const (
	sidebarW = 38
	statsW   = 32
)

var (
	sidebarPanel = components.New(sidebarW).
			Title(PanelTitle.Render("Skills")).
			BottomTitle(BodyDim.Render("jk ↑↓")).
			BottomTitleAlign(2).
			BgColor(ColorBg).
			Decorator(components.DecoratorDash).
			ActiveBorderColor(ColorBorder).
			InactiveBorderColor(ColorBorder).
			ActiveTitleColor(ColorText).
			InactiveTitleColor(ColorText).
			Padding(0, 1)

	iconPanel = components.New(sidebarW).
			TitleAlign(1). // set title in renderIcon()
			BgColor(ColorBg).
			Decorator(components.DecoratorDash).
			ActiveBorderColor(ColorBorder).
			InactiveBorderColor(ColorBorder).
			Padding(0, 1)

	statsPanel = components.New(statsW).
			Title(PanelTitle.Render("XP Info")).
			BgColor(ColorBg).
			Decorator(components.DecoratorDash).
			ActiveBorderColor(ColorBorder).
			InactiveBorderColor(ColorBorder).
			ActiveTitleColor(ColorText).
			InactiveTitleColor(ColorText).
			Padding(0, 1)

	presetsPanel = components.New(statsW).
			Title(PanelTitle.Render("Presets")).
			BgColor(ColorBg).
			Decorator(components.DecoratorDash).
			ActiveBorderColor(ColorBorder).
			InactiveBorderColor(ColorBorder).
			ActiveTitleColor(ColorText).
			InactiveTitleColor(ColorText).
			Padding(0, 1)
)

func (m XPModel) View() string {
	if m.width == 0 {
		return ""
	}

	sidebar := m.renderSidebar(sidebarW)
	icon := m.renderIcon(sidebarW)
	stats := m.renderStats(statsW)
	presets := m.renderPresets()

	colOne := lipgloss.JoinVertical(lipgloss.Left, sidebar, icon)
	colTwo := lipgloss.JoinVertical(lipgloss.Left, stats, presets)
	row := lipgloss.JoinHorizontal(lipgloss.Top, colOne, colTwo)

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Background(lipgloss.Color(ColorBg)).
		Align(lipgloss.Center, lipgloss.Center).
		Render(row)
}

// -- Sidebar ----------

func (m XPModel) renderSidebar(w int) string {
	var sb strings.Builder
	sb.WriteString("\n")

	cellW := (w - 4) / gridCols

	for i, s := range skills {
		col := i % gridCols
		level := m.levels[i]
		if level < 1 {
			level = 1
		}

		levelStr := fmt.Sprintf("%d/99", level)
		cell := fmt.Sprintf("%-3s %5s", s.abbrev, levelStr)

		var rendered string
		if i == m.selected {
			rendered = SidebarItemSelected.Width(cellW).Render(cell)
		} else {
			rendered = SidebarItem.Width(cellW).Render(cell)
		}

		sb.WriteString(rendered)
		if col == gridCols-1 {
			sb.WriteString("\n")
		}
	}

	sb.WriteString("\n  ")
	return sidebarPanel.Render(sb.String(), false)
}

// -- Icon ----------

func (m XPModel) renderIcon(w int) string {
	s := skills[m.selected]

	var imageContent string
	switch {
	case m.imageLoading:
		imageContent = lipgloss.NewStyle().
			Width(w).Height(10).
			Align(lipgloss.Center, lipgloss.Center).
			Render("loading...")
	case m.imageErr != "":
		imageContent = lipgloss.NewStyle().
			Width(w).Height(10).
			Align(lipgloss.Center, lipgloss.Center).
			Render("x " + m.imageErr)
	case m.currentImage != "":
		imageContent = lipgloss.NewStyle().
			Width(w - 4).
			Align(lipgloss.Center).
			Render(m.currentImage)
	default:
		imageContent = ImagePlaceholder.
			Width(w).Height(10).
			Render("scroll for skill icon")
	}

	imageBox := lipgloss.NewStyle().
		Padding(1, 1).
		Render(imageContent)
	imageTitle := PanelTitleAccent.Render(strings.ToUpper(s.name))

	return iconPanel.Title(imageTitle).Render(imageBox, false)
}

// -- Stats panel ----------

func (m XPModel) renderStats(w int) string {
	s := skills[m.selected]

	// Source of truth: stored values
	currentLevel := m.levels[m.selected]
	currentXP := xp.LevelToXP(currentLevel)
	targetLevel := m.targets[m.selected]
	targetXP := xp.LevelToXP(targetLevel)

	// Level vs xp based on input
	rawInput := strings.TrimSpace(m.input.Value())
	var inputVal int
	if rawInput != "" {
		fmt.Sscan(rawInput, &inputVal)
		_, parsed := parseXPInput(rawInput)
		switch m.mode {
		case modeCurrent:
			currentLevel = parsed
			currentXP = xp.LevelToXP(currentLevel)
		case modeTarget:
			targetLevel = parsed
			targetXP = xp.LevelToXP(targetLevel)
		}
	}

	var sb strings.Builder
	sb.WriteString("\n")

	sb.WriteString(statRow("Skill", s.name) + "\n")
	sb.WriteString(statRow("Level", fmt.Sprintf("%d", currentLevel)) + "\n")
	sb.WriteString(statRow("Total XP", formatXP(currentXP)) + "\n\n")

	switch m.mode {
	case modeCurrent:
		if currentLevel < 99 {
			sb.WriteString(statRow("To next lvl", formatXP(xp.XPToNextLevel(currentXP))) + "\n")
			sb.WriteString(statRow("To level 99", formatXP(xp.XPToLevel99(currentXP))) + "\n\n")
		} else {
			sb.WriteString(StatValue.Render("  MAX LEVEL") + "\n\n")
		}
		sb.WriteString(SidebarHeader.Render("MILESTONES") + "\n")
		for _, milestone := range []int{50, 70, 80, 90, 99} {
			needed := xp.XPBetween(currentXP, xp.LevelToXP(milestone))
			label := fmt.Sprintf("→ Lvl %d", milestone)
			if currentLevel >= milestone {
				sb.WriteString(statRowDim(label, Bg.Render("reached!")) + "\n")
			} else {
				sb.WriteString(statRow(label, formatXP(needed)) + "\n")
			}
		}

	case modeTarget:
		sb.WriteString(statRow("Target lvl", fmt.Sprintf("%d", targetLevel)) + "\n")
		if targetLevel > currentLevel {
			needed := xp.XPBetween(currentXP, targetXP)
			levelsLeft := targetLevel - currentLevel
			sb.WriteString(statRow("XP needed", formatXP(needed)) + "\n")
			sb.WriteString(statRow("Levels left", fmt.Sprintf("%d", levelsLeft)) + "\n")
		} else if targetLevel == currentLevel {
			sb.WriteString(statRowDim("Target", Bg.Render("= current")) + "\n")
		} else {
			sb.WriteString(statRowDim("Target", Bg.Render("< current")) + "\n")
		}
		sb.WriteString("\n")
		sb.WriteString(statRow("Target XP", formatXP(targetXP)) + "\n")
	}

	sb.WriteString("\n\n")

	// -- Mode indicator ----------
	var modeBar string
	if m.mode == modeCurrent {
		modeBar = "Current"

	} else {
		modeBar = "Target"

	}

	// -- Input box ----------
	m.input.SetWidth(w - 4)
	m.input.SetBottomLeft(BgInput.Foreground(lipgloss.Color(ColorSecondary)).Render(modeBar))
	sb.WriteString(m.input.View() + "\n")

	// -- Instructions ----------
	if m.input.Focused() {
		sb.WriteString(HelpStyle.Render(" tab ") + HelpStyleMuted.Render("mode") + Space(2))
		if m.mode == modeCurrent {
			sb.WriteString(HelpStyle.Render("enter ") + HelpStyleMuted.Render("set"))
		} else {
			sb.WriteString("")
		}
	} else {
		sb.WriteString(HelpStyle.Render("tab ") + HelpStyleMuted.Render("focus") + Space(2))
	}

	return statsPanel.Render(sb.String(), false)
}

// -- Presets panel ----------

func (m XPModel) renderPresets() string {
	lines := make([]string, len(render.GetPresets()))
	for i, p := range render.GetPresets() {
		lines[i] = HelpStyle.Render(p.Hotkey) + HelpStyleMuted.Render("  "+p.Name)
	}
	return presetsPanel.Render(strings.Join(lines, "\n"), false)
}

// -- Helpers ----------

func statRow(label, value string) string {
	return StatLabel.Render(label+":") + StatValue.Render(value)
}

func statRowDim(label, value string) string {
	return StatLabel.Render(label+":") + StatValueDim.Render(value)
}

func parseXPInput(raw string) (totalXP, level int) {
	if raw == "" {
		return 0, 1
	}
	var val int
	_, err := fmt.Sscan(raw, &val)
	if err != nil || val <= 0 {
		return 0, 1
	}
	if val <= 99 {
		level = val
		totalXP = xp.LevelToXP(level)
	} else {
		totalXP = val
		level = xp.XPToLevel(totalXP)
	}
	return
}

func formatXP(n int) string {
	if n == 0 {
		return "0"
	}
	s := fmt.Sprintf("%d", n)
	var b strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteRune(',')
		}
		b.WriteRune(c)
	}
	return b.String()
}
