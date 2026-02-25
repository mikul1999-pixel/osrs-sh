package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

// Tab indices
const (
	TabHome    = 0
	TabXP      = 1
	TabMonster = 2
	TabPlayer  = 3
	TabItem    = 4
)

var tabNames = []string{"HOME", "XP", "MONSTER", "PLAYER", "ITEM"}

// NavigateMsg is sent by child tabs to request a tab switch
type NavigateMsg struct {
	Tab   int
	Query string
}

// AppModel is the root Bubble Tea model
type AppModel struct {
	activeTab int
	width     int
	height    int

	home HomeModel
	xp   XPModel
	// monster, player, item ...
}

func NewAppModel() AppModel {
	return AppModel{
		activeTab: TabHome,
		home:      NewHomeModel(),
		xp:        NewXPModel(),
	}
}

// -- Init ----------

func (a AppModel) Init() tea.Cmd {
	return a.home.Init()
}

// -- Update ----------

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		// Propagate size to all tabs
		a.home.SetSize(msg.Width, a.contentHeight())
		a.xp.SetSize(msg.Width, a.contentHeight())
		return a, nil

	case NavigateMsg:
		a.activeTab = msg.Tab
		if msg.Tab == TabXP && msg.Query != "" {
			a.xp.SetQuery(msg.Query)
		}
		return a, nil

	case tea.KeyMsg:
		// Global tab switching
		switch msg.String() {
		case "alt+1":
			a.activeTab = TabHome
			return a, nil
		case "alt+2":
			a.activeTab = TabXP
			return a, nil
		case "alt+3":
			a.activeTab = TabMonster
			return a, nil
		case "alt+4":
			a.activeTab = TabPlayer
			return a, nil
		case "alt+5":
			a.activeTab = TabItem
			return a, nil
		case "ctrl+c":
			return a, tea.Quit
		}
	}

	// Delegate to active tab
	return a.updateActiveTab(msg)
}

func (a AppModel) updateActiveTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch a.activeTab {
	case TabHome:
		a.home, cmd = a.home.Update(msg)
	case TabXP:
		a.xp, cmd = a.xp.Update(msg)
	}
	return a, cmd
}

// -- View ----------

func (a AppModel) View() tea.View {
	content := a.activeTabView()
	statusBar := a.renderStatusBar()

	full := lipgloss.JoinVertical(lipgloss.Left, content, statusBar)
	place := lipgloss.Place(
		a.width, a.height,
		lipgloss.Left, lipgloss.Top,
		full,
		lipgloss.WithWhitespaceBackground(lipgloss.Color(ColorBg)),
	)
	v := tea.NewView(place)
	v.AltScreen = true // create alt full screen
	v.BackgroundColor = lipgloss.Color(ColorBg)
	return v
}

func (a AppModel) activeTabView() string {
	switch a.activeTab {
	case TabHome:
		return a.home.View()
	case TabXP:
		return a.xp.View()
	default:
		return a.placeholderView()
	}
}

func (a AppModel) placeholderView() string {
	name := tabNames[a.activeTab]
	msg := fmt.Sprintf("\n\n  %s — coming soon\n", name)
	return lipgloss.NewStyle().
		Width(a.width).
		Height(a.contentHeight()).
		Foreground(lipgloss.Color(ColorMuted)).
		Render(msg)
}

func (a AppModel) renderStatusBar() string {
	// Left side
	left := StatusKey.Render("ctrl+c") + StatusVal.Render(" quit")

	// Right side
	var tabParts []string
	for i, name := range tabNames {
		label := fmt.Sprintf("%d %s", i+1, name)
		if i == a.activeTab {
			tabParts = append(tabParts, StatusTabActive.Render(label))
		} else {
			tabParts = append(tabParts, StatusTab.Render(label))
		}
	}
	right := strings.Join(tabParts, " ")

	// Pad to full width
	gap := a.width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 0 {
		gap = 0
	}
	spacer := StatusBar.Render(strings.Repeat(" ", gap))

	return StatusBar.Width(a.width).Render(left + spacer + right)
}

// contentHeight is the usable height minus the status bar
func (a AppModel) contentHeight() int {
	h := a.height - 1
	if h < 0 {
		return 0
	}
	return h
}
