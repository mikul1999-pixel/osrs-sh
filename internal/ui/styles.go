package ui

import "github.com/charmbracelet/lipgloss"

// -- Palette ----------

const (
	ColorBlack   = "#000000"
	ColorGold    = "#c8a951"
	ColorGoldDim = "#7a6535"
	ColorGreen   = "#98c379"
	ColorBlue    = "#61afef"
	ColorRed     = "#e06c75"
	ColorOrange  = "#d19a66"

	ColorBg      = "#13141a" // default background
	ColorBgInput = "#2a2b3d" // text boxes
	ColorBgPanel = "#1e2030" // panel fill
	ColorBorder  = "#2a2b3d" // panel borders
	// ColorBorder = ColorGreen // panel borders

	ColorText       = "#abb2bf"    // main text
	ColorTextLight  = "#90969f"    // main text light
	ColorTextDim    = "#3e4451"    // main text dimmed
	ColorTextDark   = "#151516"    // main text dimmed
	ColorMuted      = "#4b5263"    // help / info
	ColorPrimary    = ColorGold    // primary accent
	ColorPrimaryDim = ColorGoldDim // primary accent dimmed
	ColorSecondary  = ColorBlue    // secondary accent
	ColorPositive   = ColorGreen   // positive values
	ColorWarning    = ColorRed     // errors / warnings
	ColorHighlight  = ColorOrange  // keybind highlights
)

// -- Base Styles ----------

var (
	Base = lipgloss.NewStyle().
		Background(lipgloss.Color(ColorBg)).
		Foreground(lipgloss.Color(ColorText))

	Bold    = lipgloss.NewStyle().Bold(true)
	Body    = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorText))
	BodyDim = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorTextDim))
	Bg      = lipgloss.NewStyle().Background(lipgloss.Color(ColorBg))
	BgInput = lipgloss.NewStyle().Background(lipgloss.Color(ColorBgInput))
)

// -- Panel Styles ----------

var (
	PanelTitleAccent = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorPrimary)).
				Background(lipgloss.Color(ColorBg)).
				Bold(true).
				PaddingLeft(1).
				PaddingRight(1)

	PanelTitle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorTextLight)).
			Background(lipgloss.Color(ColorBg))
)

// -- Home Screen ----------

var (
	LogoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPrimary)).
			Bold(true)

	LogoDimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPrimaryDim)).
			Bold(true)

	HomeVersionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorMuted)).
				Background(lipgloss.Color(ColorBg)).
				Italic(true)

	HomeCmdStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorSecondary)).
			Background(lipgloss.Color(ColorBg)).
			Bold(true)

	HomeDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorText)).
			Background(lipgloss.Color(ColorBg))

	HomeKeybindStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorHighlight)).
				Background(lipgloss.Color(ColorBg))

	InputPrompt = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorText)).
			Background(lipgloss.Color(ColorBgInput))

	InputPlaceholder = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorTextDim)).
				Background(lipgloss.Color(ColorBgInput))

	InputCursor = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorText)).
			Foreground(lipgloss.Color(ColorTextDark))
)

// -- Sidebar ----------

var (
	SidebarItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorText)).
			Background(lipgloss.Color(ColorBg)).
			PaddingLeft(2)

	SidebarItemSelected = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorBg)).
				Background(lipgloss.Color(ColorPrimary)).
				Bold(true).
				PaddingLeft(1).
				PaddingRight(1)

	SidebarHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted)).
			Bold(true).
			PaddingLeft(1).
			MarginBottom(1)
)

// -- XP Stats Panel ----------

var (
	StatLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorSecondary)).
			Background(lipgloss.Color(ColorBg)).
			Width(18)

	StatValue = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPositive)).
			Background(lipgloss.Color(ColorBg)).
			Bold(true)

	StatValueDim = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted))

	XPBarFilled = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPrimary)).
			Background(lipgloss.Color(ColorPrimary))

	XPBarEmpty = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorBorder)).
			Background(lipgloss.Color(ColorBorder))
)

// -- Image ----------

var (
	ImagePlaceholder = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorMuted)).
		Align(lipgloss.Center, lipgloss.Center)
)

// -- Bottom Status Bar ----------

var (
	StatusBar = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorBgPanel)).
			Foreground(lipgloss.Color(ColorMuted))

	StatusKey = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorHighlight)).
			Background(lipgloss.Color(ColorBgPanel)).
			Bold(true)

	StatusVal = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorText)).
			Background(lipgloss.Color(ColorBgPanel))

	StatusTab = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted)).
			Background(lipgloss.Color(ColorBgPanel)).
			Padding(0, 1)

	StatusTabActive = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorBg)).
			Background(lipgloss.Color(ColorPrimary)).
			Bold(true).
			Padding(0, 1)

	// Error text inside panels
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorWarning))

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorText)).
			Background(lipgloss.Color(ColorBg)).
			Italic(true)

	HelpStyleDim = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorTextDim)).
			Background(lipgloss.Color(ColorBg)).
			Italic(true)

	HelpStyleMuted = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted)).
			Background(lipgloss.Color(ColorBg)).
			Italic(true)
)
