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

	xp      [24]int
	targets [24]int

	mode inputMode

	input    components.Input
	inputErr string

	currentImage string
	imageLoading bool
	imageErr     string

	spinner *components.Spinner
}

func NewXPModel() XPModel {
	input := components.NewInput(components.InputOptions{
		CharLimit:        12,
		Placeholder:      "type level or xp...",
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

	var startXP [24]int
	var targets [24]int
	for i := range startXP {
		startXP[i] = xp.LevelToXP(1)
		targets[i] = 99
	}
	m := XPModel{input: input, xp: startXP, targets: targets}
	m.syncInputToMode()
	return m
}

// syncInputToMode updates placeholder and pre-fills input value from stored state.
func (m *XPModel) syncInputToMode() {
	switch m.mode {
	case modeCurrent:
		storedXP := m.xp[m.selected]
		storedLevel := xp.XPToLevel(storedXP)
		if storedLevel > 1 {
			m.input.SetValue(fmt.Sprintf("%d", storedXP))
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

// SetPlayerXP populates skill xp from hiscores
func (m *XPModel) SetPlayerXP(rawXP [24]int) {
	m.xp = rawXP
	m.syncInputToMode()
}

func (m XPModel) Init() (XPModel, tea.Cmd) {
	m.imageLoading = true
	m.spinner = components.NewSpinner().SetFrames(components.SpinnerBraille)
	return m, tea.Batch(loadSkillImage(skills[m.selected].name), m.spinner.Start())
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
		if m.spinner != nil {
			m.spinner.Stop()
		}
		return m, nil

	case components.SpinnerTickMsg:
		if m.spinner != nil && msg.ID == m.spinner.ID() {
			m.spinner.Tick()
			return m, m.spinner.TickCmd()
		}
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
			m.spinner = components.NewSpinner().SetFrames(components.SpinnerBraille)
			return m, tea.Batch(loadSkillImage(skills[m.selected].name), m.spinner.Start())

		case "down", "j":
			m.selected += gridCols
			if m.selected >= len(skills) {
				col := (m.selected - gridCols) % gridCols
				nextCol := (col + 1) % gridCols
				m.selected = nextCol
			}
			m.syncInputToMode()
			m.imageLoading = true
			m.spinner = components.NewSpinner().SetFrames(components.SpinnerBraille)
			return m, tea.Batch(loadSkillImage(skills[m.selected].name), m.spinner.Start())

		case "q", "e", "r":
			for _, p := range render.GetPresets() {
				if msg.String() == p.Hotkey {
					m.applyPreset(p)
					return m, func() tea.Msg { return PresetAppliedMsg{Name: p.Name} }
				}
			}

		case "esc":
			m.resetTargets()
			return m, func() tea.Msg { return PresetClearedMsg{} }

		case "tab":
			if m.input.Focused() {
				// Save current input then toggle mode
				cmd := m.saveCurrentInput()
				if m.mode == modeCurrent {
					m.mode = modeTarget
				} else {
					m.mode = modeCurrent
				}
				m.syncInputToMode()
				return m, cmd
			}

		case "enter":
			cmd := m.saveCurrentInput()
			m.inputErr = ""
			return m, cmd
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *XPModel) saveCurrentInput() tea.Cmd {
	_, level := parseXPInput(strings.TrimSpace(m.input.Value()))
	switch m.mode {
	case modeCurrent:
		if level >= 1 {
			m.xp[m.selected] = xp.LevelToXP(level)
			return func() tea.Msg {
				return LevelSetMsg{
					Message: strings.ToLower(skills[m.selected].abbrev) + " set to level " + StatusBlockMode1.Render(fmt.Sprintf(" %d ", level)),
					Sub:     "current",
					Style:   components.ToastInfo,
				}
			}
		}
	case modeTarget:
		if level >= 1 {
			m.targets[m.selected] = level
			return func() tea.Msg {
				return LevelSetMsg{
					Message: strings.ToLower(skills[m.selected].abbrev) + " set to level " + StatusBlockMode2.Render(fmt.Sprintf(" %d ", level)),
					Sub:     "target",
					Style:   components.ToastInfo,
				}
			}
		}
	}
	return nil
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
			BottomTitle(BodyDim.Render("esc ↻")).
			BottomTitleAlign(2).
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
		level := xp.XPToLevel(m.xp[i])
		if level < 1 {
			level = 1
		}

		var levelStr string
		if m.targets[i] < 99 { // only show when target is set
			levelStr = fmt.Sprintf("%d/%d", level, m.targets[i])
		} else {
			levelStr = fmt.Sprintf("%d/99", level)
		}
		cell := fmt.Sprintf("%-3s %5s", s.abbrev, levelStr)

		var dot string
		if m.targets[i] < 99 {
			currentLevel := xp.XPToLevel(m.xp[i])
			if currentLevel >= m.targets[i] {
				dot = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorGreen)).Render("•")
			} else {
				dot = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorRed)).Render("•")
			}
		} else {
			dot = " " // alignment when no dot
		}

		// gold text for 99
		currentLevel := xp.XPToLevel(m.xp[i])
		var cellRendered string
		if currentLevel >= 99 {
			if m.xp[i] >= 200000000 {
				cellRendered = SidebarItem200M.Width(cellW).Render(cell)
			} else {
				cellRendered = SidebarItemMaxed.Width(cellW).Render(cell)
			}
		} else {
			cellRendered = SidebarItem.Width(cellW).Render(cell)
		}

		var rendered string
		if i == m.selected {
			rendered = dot + SidebarItemSelected.Width(cellW).Render(cell)
		} else {
			rendered = dot + cellRendered
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
		frame := ""
		if m.spinner != nil {
			frame = m.spinner.View()
		}
		return StatusLine1.
			Width(w).Height(10).
			Align(lipgloss.Center, lipgloss.Center).
			Render(frame + " loading...")
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
	currentXP := m.xp[m.selected]
	currentLevel := xp.XPToLevel(currentXP)
	targetLevel := m.targets[m.selected]
	targetXP := xp.LevelToXP(targetLevel)

	// Level vs xp based on input
	rawInput := strings.TrimSpace(m.input.Value())
	if rawInput != "" {
		totalXP, parsedLevel := parseXPInput(rawInput)

		switch m.mode {
		case modeCurrent:
			currentLevel = parsedLevel
			currentXP = totalXP

		case modeTarget:
			targetLevel = parsedLevel
			targetXP = totalXP
		}
	}

	var sb strings.Builder
	sb.WriteString("\n")

	sb.WriteString(statRow("Skill", s.name) + "\n")
	sb.WriteString(statRow("Level", fmt.Sprintf("%d", currentLevel)) + "\n")
	sb.WriteString(statRow("Total XP", formatXP(currentXP)) + "\n\n")

	// -- Mode indicator ----------
	var modeBar string
	var modeColor string
	switch m.mode {
	case modeCurrent:
		modeBar = "Current"
		modeColor = ColorSecondary
		m.input.SetAccentFocused(lipgloss.Color(modeColor))

		sb.WriteString(StatHeader.Render("MILESTONES") + "\n")
		var maxText string
		if currentXP >= 200000000 {
			maxText = "200m! nerd"
		} else {
			maxText = "max level!"
		}

		if currentLevel < 99 {
			sb.WriteString(statRowMode("To next lvl", formatXP(xp.XPToNextLevel(currentXP)), modeColor) + "\n")
			sb.WriteString(statRowMode("To level 99", formatXP(xp.XPToLevel99(currentXP)), modeColor) + "\n\n")
		} else {
			sb.WriteString(statRowDim("To next lvl", Bg.Render(maxText)) + "\n")
			sb.WriteString(statRowDim("To level 99", Bg.Render(maxText)) + "\n\n")
		}
		for _, milestone := range []int{50, 70, 80, 90, 99} {
			needed := xp.XPBetween(currentXP, xp.LevelToXP(milestone))
			label := fmt.Sprintf("→ Lvl %d", milestone)
			if currentLevel >= milestone {
				sb.WriteString(statRowDim(label, Bg.Render("reached!")) + "\n")
			} else {
				sb.WriteString(statRowMode(label, formatXP(needed), modeColor) + "\n")
			}
		}

	case modeTarget:
		modeBar = "Target"
		modeColor = ColorPositive
		m.input.SetAccentFocused(lipgloss.Color(modeColor))

		sb.WriteString(StatHeader.Render("GOALS") + "\n")
		sb.WriteString(statRowMode("Target lvl", fmt.Sprintf("%d", targetLevel), modeColor) + "\n")
		sb.WriteString(statRowMode("Target XP", formatXP(targetXP), modeColor) + "\n")
		sb.WriteString("\n")
		if targetLevel > currentLevel {
			needed := xp.XPBetween(currentXP, targetXP)
			levelsLeft := targetLevel - currentLevel
			sb.WriteString(statRowMode("Levels left", fmt.Sprintf("%d", levelsLeft), modeColor) + "\n")
			sb.WriteString(statRowMode("XP needed", formatXP(needed), modeColor) + "\n")
		} else if targetLevel == currentLevel {
			sb.WriteString(statRowDim("Levels left", Bg.Render("= current")) + "\n")
			sb.WriteString(statRowDim("XP needed", Bg.Render("= current")) + "\n")
		} else {
			sb.WriteString(statRowDim("Levels left", Bg.Render("< current")) + "\n")
			sb.WriteString(statRowDim("XP needed", Bg.Render("< current")) + "\n")
		}
		sb.WriteString("\n\n\n")
	}

	sb.WriteString("\n\n")

	// -- Input box ----------
	m.input.SetWidth(w - 4)
	m.input.SetBottomLeft(BgInput.Foreground(lipgloss.Color(modeColor)).Faint(true).Render(modeBar))
	sb.WriteString(m.input.View() + "\n")

	// -- Instructions ----------
	if m.input.Focused() {
		sb.WriteString(HelpStyle.Render(" tab ") + HelpStyleMuted.Render("mode") + Space(2))
		sb.WriteString(HelpStyle.Render("enter ") + HelpStyleMuted.Render("set"))
	} else {
		sb.WriteString(HelpStyle.Render("tab ") + HelpStyleMuted.Render("focus") + Space(2))
	}

	return statsPanel.Render(sb.String(), false)
}

// -- Presets panel ----------

func (m XPModel) renderPresets() string {
	lines := make([]string, len(render.GetPresets()))
	for i, p := range render.GetPresets() {
		lines[i] = HomeKeybindStyle.Render(p.Hotkey) + StatValueDim.Render("  "+p.Name)
	}
	return presetsPanel.Render(strings.Join(lines, "\n"), false)
}

// -- Helpers ----------

func statRow(label, value string) string {
	return StatLabel.Render(label+":") + StatValue.Render(value)
}

func statRowMode(label, value string, color string) string {
	lipglossColor := color
	if color == "" {
		lipglossColor = ColorGold
	}
	return StatLabelMode.Render(label+":") +
		StatValueMode.Foreground(lipgloss.Color(lipglossColor)).Render(value)
}

func statRowDim(label, value string) string {
	return StatLabelMode.Render(label+":") + StatValueDim.Render(value)
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

// applyPreset sets all target levels from a preset
func (m *XPModel) applyPreset(p render.Preset) {
	for i, target := range p.Targets {
		if target > 0 {
			m.targets[i] = target
		}
	}
	m.syncInputToMode()
}

// resetTargets sets all targets back to 99
func (m *XPModel) resetTargets() {
	for i := range m.targets {
		m.targets[i] = 99
	}
	m.syncInputToMode()
}
