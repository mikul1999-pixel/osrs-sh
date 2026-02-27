package ui

import "github.com/charmbracelet/lipgloss"

// -- Palette ----------

const (
	ColorBlack      = "#000000"
	ColorGold       = "#c8a951"
	ColorGoldDim    = "#7a6535"
	ColorGoldDark   = "#3d2c1f"
	ColorGreen      = "#57a854"
	ColorGreenLight = "#98c379"
	ColorGreenDark  = "#2e3a4d"
	ColorBlue       = "#3b98d4"
	ColorBlueDark   = "#2e3a4d"
	ColorPink       = "#ff66b2"
	ColorPurple     = "#b49dd8"
	ColorPurpleDark = "#373349"
	ColorGrey       = "#8a93b2"
	ColorGreyDark   = "#2a2b3d"
	ColorRed        = "#c94f4f"
	ColorRedLight   = "#e06c75"
	ColorOrange     = "#d19a66"

	ColorBg      = "#13141a" // default background
	ColorBgInput = "#2a2b3d" // text boxes

	ColorBorder = "#2a2b3d" // panel borders

	ColorText         = "#abb2bf"       // main text
	ColorTextLight    = "#90969f"       // main text light
	ColorTextDim      = "#3e4451"       // main text dimmed
	ColorTextDark     = "#151516"       // main text dimmed
	ColorMuted        = "#4b5263"       // help / info
	ColorPrimary      = ColorGold       // primary accent
	ColorPrimaryDim   = ColorGoldDim    // primary accent dimmed
	ColorPrimaryDark  = ColorGoldDark   // primary accent dark
	ColorSecondary    = ColorBlue       // secondary accent
	ColorSecondaryDim = ColorBlueDark   // secondary accent dimmed
	ColorRare         = ColorPink       // rare accent
	ColorPositive     = ColorGreenLight // positive values
	ColorPositiveDim  = ColorGreenDark  // positive values dimmed
	ColorWarning      = ColorRedLight   // errors / warnings
	ColorHighlight    = ColorOrange     // keybind highlights
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

	HomeSubTitleStyle = lipgloss.NewStyle().
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
			Background(lipgloss.Color(ColorBg))

	SidebarItemSelected = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorBg)).
				Background(lipgloss.Color(ColorPrimary)).
				Bold(true).
				PaddingRight(1)

	SidebarItemMaxed = lipgloss.NewStyle().
				Foreground(lipgloss.Color(ColorPrimary)).
				Background(lipgloss.Color(ColorBg))

	SidebarItem200M = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorRare)).
			Background(lipgloss.Color(ColorBg)).
			Faint(true)
)

// -- XP Stats Panel ----------

var (
	StatHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted)).
			Bold(true).
			MarginBottom(1)

	StatLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPrimary)).
			Background(lipgloss.Color(ColorBg)).
			Bold(true).
			Italic(true).
			Faint(true).
			Width(18)

	StatLabelMode = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorText)).
			Background(lipgloss.Color(ColorBg)).
			Width(18)

	StatValue = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorPrimary)).
			Background(lipgloss.Color(ColorBg)).
			Bold(true).
			Italic(true).
			Faint(true)

	StatValueMode = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorBg))

	StatValueDim = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted)).
			Background(lipgloss.Color(ColorBg))
)

// -- Image ----------

var (
	ImagePlaceholder = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorMuted)).
		Align(lipgloss.Center, lipgloss.Center)
)

// -- Bottom Status Bar ----------

var (
	StatusKey = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorTextLight)).
			Background(lipgloss.Color(ColorBg))

	StatusVal = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ColorMuted)).
			Background(lipgloss.Color(ColorBg))

	StatusLine1 = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorBg)).
			Foreground(lipgloss.Color(ColorMuted))

	StatusLine2 = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorBg)).
			Foreground(lipgloss.Color(ColorMuted))

	StatusBlock = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorPurpleDark)).
			Foreground(lipgloss.Color(ColorPurple))

	StatusBlockInfo = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorPrimaryDark)).
			Foreground(lipgloss.Color(ColorPrimary))

	StatusBlockMode1 = lipgloss.NewStyle().
				Background(lipgloss.Color(ColorSecondaryDim)).
				Foreground(lipgloss.Color(ColorSecondary))

	StatusBlockMode2 = lipgloss.NewStyle().
				Background(lipgloss.Color(ColorPositiveDim)).
				Foreground(lipgloss.Color(ColorPositive))

	StatusBlockMuted = lipgloss.NewStyle().
				Background(lipgloss.Color(ColorGreyDark)).
				Foreground(lipgloss.Color(ColorGrey))

	StatusError = lipgloss.NewStyle().
			Background(lipgloss.Color(ColorBg)).
			Foreground(lipgloss.Color(ColorWarning))

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
