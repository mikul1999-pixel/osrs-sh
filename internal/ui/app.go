package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikul1999-pixel/osrs-sh/internal/core/api"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui/components"
)

// Theme selection
var ActiveTheme Theme

// var ActiveTheme = loadThemeFromConfig("modern") // manual for now. change to cfg.Theme later

// Tab indices
const (
	TabHome    = 0
	TabXP      = 1
	TabMonster = 2
	TabItem    = 3
	TabPlayer  = 4
)

var tabNames = []string{"HOME", "XP", "MONSTER", "ITEM", "PLAYER"}

// NavigateMsg is sent by child tabs to request a tab switch
type NavigateMsg struct {
	Tab    int
	Query  string
	Action CommandAction
}

// LoadPlayerMsg is recieved from home.go. Indicates user looked up an RSN
type LoadPlayerMsg struct{ RSN string }

// PlayerLoadedMsg/PlayerErrMsg confirm Hiscores load status
type PlayerLoadedMsg struct {
	rsn string
	xp  [24]int
}
type PlayerErrMsg struct {
	err string
}

// PresetAppliedMsg/PresetClearedMsg confirm stat preset was applied to targets
type PresetAppliedMsg struct{ Name string }
type PresetClearedMsg struct{}

// LevelSet confirms a current/target level was stored
type LevelSetMsg struct {
	Message string
	Sub     string
	Style   components.ToastStyle
}

// PlayerState stores info on player's RSN
type PlayerState struct {
	RSN     string
	xp      [24]int
	Loaded  bool
	Loading bool
	Err     string
}

// StatusContext is the right text on the status bar
type StatusKeybind struct {
	Key   string
	Label string
}
type StatusContext struct {
	Label    string
	Keybinds []StatusKeybind
}
type SetStatusContextMsg struct{ Context StatusContext }

// Options to display dir in app
type CwdOptions struct {
	ShortenHome   bool   // replace $HOME with ~
	LastOnly      bool   // return only last folder
	RootLabel     string // what to display for root ("/" or "C:\")
	FallbackValue string // what to return if Getwd() fails
}

// AppModel is the root Bubble Tea model
type AppModel struct {
	activeTab     int
	width         int
	height        int
	player        PlayerState
	home          HomeModel
	xp            XPModel
	spinner       *components.Spinner
	toast         *components.Toast
	activePreset  string
	statusContext StatusContext
	palette       PaletteModel
	theme         Theme
}

func NewAppModel() AppModel {
	t := loadThemeFromConfig("modern") // manual for now. change to cfg.Theme later
	ActiveTheme = t
	return AppModel{
		activeTab: TabHome,
		home:      NewHomeModel(),
		xp:        NewXPModel(),
		palette:   NewPaletteModel(),
		theme:     t,
	}
}

// -- Init ----------

func (a AppModel) Init() tea.Cmd {
	return tea.Batch(
		a.home.Init(),
	)
}

// -- Update ----------

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Let modal process messages
	if a.palette.visible {
		a.palette, cmd = a.palette.Update(msg)
		return a, cmd
	}

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		// Propagate size to all tabs
		a.home.SetSize(msg.Width, a.contentHeight())
		a.xp.SetSize(msg.Width, a.contentHeight())
		a.palette = a.palette.SetSize(msg.Width, msg.Height)
		return a, nil

	case NavigateMsg:
		a.activeTab = msg.Tab
		if msg.Tab == TabXP && msg.Query != "" {
			a.xp.SetQuery(msg.Query)
		}
		return a, nil

	case LoadPlayerMsg:
		a.player.Loading = true
		a.player.RSN = msg.RSN
		a.spinner = components.NewSpinner().
			SetFrames(components.SpinnerBraille).
			SetStyle(ActiveTheme.StatusLine())
		return a, tea.Batch(loadPlayerCmd(msg.RSN), a.spinner.Start())

	case PlayerLoadedMsg:
		a.player = PlayerState{
			RSN:    msg.rsn,
			xp:     msg.xp,
			Loaded: true,
		}
		a.xp.SetPlayerXP(msg.xp)
		if a.spinner != nil {
			a.spinner.Stop()
		}
		// Show success toast
		a.toast = components.NewToast().
			SetMessage(strings.ToLower(msg.rsn) + " loaded").
			SetStyle(components.ToastSuccess)
		return a, a.toast.Show()

	case PlayerErrMsg:
		a.player.Loading = false
		a.player.Err = msg.err
		if a.spinner != nil {
			a.spinner.Stop()
		}
		a.toast = components.NewToast().
			SetMessage("lookup failed").
			SetSub(msg.err).
			SetStyle(components.ToastError)
		return a, a.toast.Show()

	case PresetAppliedMsg:
		a.activePreset = msg.Name
		a.statusContext = StatusContext{Label: a.activePreset, Keybinds: GetTabCmds(TabXP)}
		a.toast = components.NewToast().
			SetMessage(msg.Name + " applied").
			SetStyle(components.ToastSuccess)
		return a, a.toast.Show()

	case PresetClearedMsg:
		a.activePreset = ""
		a.statusContext = StatusContext{Label: a.activePreset, Keybinds: GetTabCmds(TabXP)}
		a.toast = components.NewToast().
			SetMessage("preset cleared").
			SetStyle(components.ToastInfo)
		return a, a.toast.Show()

	case LevelSetMsg:
		a.toast = components.NewToast().
			SetMessage(msg.Message).
			SetSub(msg.Sub).
			SetStyle(msg.Style)
		return a, a.toast.Show()

	case SetStatusContextMsg:
		a.statusContext = msg.Context
		return a, nil

	case components.SpinnerTickMsg:
		if a.spinner != nil && msg.ID == a.spinner.ID() {
			a.spinner.Tick()
			return a, a.spinner.TickCmd()
		}
		return a, nil

	case components.ToastDismissMsg:
		if a.toast != nil && msg.ID == a.toast.ID() {
			a.toast.Hide()
		}
		return a, nil

	case tea.KeyMsg: // tea.Keymsg handles both press and release, tea.KeyPressMsg for press only
		switch msg.String() {
		case "ctrl+p":
			a.palette, _ = a.palette.Open()
			return a, nil

		// Global tab switching
		case "alt+1":
			a.activeTab = TabHome
			a.statusContext = StatusContext{}
			return a, nil
		case "alt+2":
			a.activeTab = TabXP
			a.statusContext = StatusContext{Label: a.activePreset, Keybinds: GetTabCmds(a.activeTab)}
			return a, nil
		case "alt+3":
			a.activeTab = TabMonster
			a.statusContext = StatusContext{}
			return a, nil
		case "alt+4":
			a.activeTab = TabItem
			a.statusContext = StatusContext{}
			return a, nil
		case "alt+5":
			a.activeTab = TabPlayer
			a.statusContext = StatusContext{}
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
	// Set background colors to dim style if overlay
	if a.palette.visible {
		ActiveTheme = a.theme.Dimmed()
		defer func() { ActiveTheme = a.theme }()
	}

	content := a.activeTabView()
	statusBar := a.renderStatusBar()
	full := lipgloss.JoinVertical(lipgloss.Left, content, statusBar)

	// Force full width & height
	full = lipgloss.Place(
		a.width, a.height,
		lipgloss.Left, lipgloss.Top,
		full,
		lipgloss.WithWhitespaceBackground(ActiveTheme.Bg),
	)

	// Overlay command dropdown if active
	if a.activeTab == TabHome {
		if panel, x, y := a.home.DropdownOverlay(); panel != "" {
			full = components.PlaceOverlay(x, y, panel, full, a.width)
		}
	}

	// Overlay toast in top right if visible
	if a.toast != nil && a.toast.Visible() {
		toastStr := a.toast.View()
		toastW := lipgloss.Width(toastStr)
		x := a.width - toastW - 1
		full = components.PlaceOverlay(x, 1, toastStr, full, a.width)
	}

	// Overlay command palette if visible
	if a.palette.visible {
		modal := a.palette.View()

		// center everything
		full = components.PlaceOverlay(
			(a.width-lipgloss.Width(modal))/2,
			(a.height-lipgloss.Height(modal))/2,
			modal,
			full,
			a.width,
		)
	}

	place := lipgloss.Place(
		a.width, a.height,
		lipgloss.Left, lipgloss.Top,
		full,
		lipgloss.WithWhitespaceBackground(lipgloss.Color(ActiveTheme.Bg)),
	)
	v := tea.NewView(place)
	v.AltScreen = true
	v.BackgroundColor = lipgloss.Color(ActiveTheme.Bg)
	v.WindowTitle = "osrs-sh"
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
	msg := fmt.Sprintf("\n\n  %s: TODO\n", name)
	return lipgloss.NewStyle().
		Width(a.width).
		Height(a.contentHeight()).
		Foreground(lipgloss.Color(ActiveTheme.Muted)).
		Render(msg)
}

// contentHeight is the usable height minus the status bar
func (a AppModel) contentHeight() int {
	h := a.height - 2
	if h < 0 {
		return 0
	}
	return h
}

// -- Commands ----------

func loadPlayerCmd(rsn string) tea.Cmd {
	client := api.NewClient(api.Options{})
	hs := api.New(client)

	return func() tea.Msg {
		result, err := hs.Lookup(rsn)
		if err != nil {
			return PlayerErrMsg{err: err.Error()}
		}
		return PlayerLoadedMsg{rsn: rsn, xp: result.XP}
	}
}

// -- Helpers ----------

// -- General ----------

func Space(rpt int) string {
	space := strings.Repeat(" ", rpt)
	return ActiveTheme.Bg_().Render(space)
}

func SpaceInput(rpt int) string {
	space := strings.Repeat(" ", rpt)
	return ActiveTheme.BgInput_().Render(space)
}

func getCwdDisplay(opts CwdOptions) string {
	cwd, err := os.Getwd()
	if err != nil {
		if opts.FallbackValue != "" {
			return opts.FallbackValue
		}
		return "unknown"
	}

	cwd = filepath.Clean(cwd)

	// Handle root
	if isRoot(cwd) {
		if opts.RootLabel != "" {
			return opts.RootLabel
		}
		return cwd
	}

	// Replace HOME with "~"
	if opts.ShortenHome {
		if home, err := os.UserHomeDir(); err == nil {
			home = filepath.Clean(home)
			if strings.HasPrefix(cwd, home) {
				cwd = "~" + strings.TrimPrefix(cwd, home)
			}
		}
	}

	if opts.LastOnly {
		return "~/" + filepath.Base(cwd)
	}

	return cwd
}

func isRoot(path string) bool {
	parent := filepath.Dir(path)
	return parent == path
}
