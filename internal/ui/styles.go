package ui

import "github.com/charmbracelet/lipgloss"

// -- Raw Colors ----------

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
	ColorRedDark    = "#4d2929"
	ColorOrange     = "#d19a66"
)

// -- Theme ----------

// Theme holds all semantic colors for a given visual theme
type Theme struct {
	// Core
	Bg          lipgloss.Color // terminal background
	BgDim       lipgloss.Color // overlay/modal dim background
	BgInput     lipgloss.Color // text input backgrounds
	BgInputDim  lipgloss.Color
	BgPanel     lipgloss.Color // panel backgrounds
	BgPanelDim  lipgloss.Color
	BgModal     lipgloss.Color // modal backgrounds
	BgModalList lipgloss.Color // modal list backgrounds
	Border      lipgloss.Color // panel borders
	BorderDim   lipgloss.Color

	// Text
	Text          lipgloss.Color // default text color
	TextLight     lipgloss.Color
	TextDim       lipgloss.Color
	TextDark      lipgloss.Color // text on light backgrounds
	Muted         lipgloss.Color // help/placeholder text
	MutedModal    lipgloss.Color
	ModalText     lipgloss.Color // default text color for modals
	ModalTextDark lipgloss.Color // darker text color for modals
	Cursor        lipgloss.Color
	CursorText    lipgloss.Color

	// Semantic
	Primary      lipgloss.Color // main accent
	PrimaryModal lipgloss.Color // main accent
	PrimaryDim   lipgloss.Color
	PrimaryDark  lipgloss.Color
	Secondary    lipgloss.Color // second accent
	SecondaryDim lipgloss.Color
	Tertiary     lipgloss.Color // third accent
	TertiaryDim  lipgloss.Color
	Warning      lipgloss.Color
	Rare         lipgloss.Color // rare highlight
	Highlight    lipgloss.Color // interactive element highlight
	Green        lipgloss.Color // fixed green color
	GreenDim     lipgloss.Color
	Red          lipgloss.Color // fixed red color
	RedDim       lipgloss.Color

	// Status blocks
	BlockDefault     lipgloss.Color // fg for default status block
	BlockDefaultBg   lipgloss.Color
	BlockInfo        lipgloss.Color
	BlockInfoBg      lipgloss.Color
	BlockSecondary   lipgloss.Color
	BlockSecondaryBg lipgloss.Color
	BlockTertiary    lipgloss.Color
	BlockTertiaryBg  lipgloss.Color
	BlockMuted       lipgloss.Color
	BlockMutedBg     lipgloss.Color
}

// -- Theme Definitions ----------

// ThemeModern = app default palette
var ThemeModern = Theme{
	Bg:          lipgloss.Color("#13141a"),
	BgDim:       lipgloss.Color("#0d0d12"),
	BgInput:     lipgloss.Color("#2a2b3d"),
	BgInputDim:  lipgloss.Color("#1a1b2a"),
	BgPanel:     lipgloss.Color("#10111a"),
	BgPanelDim:  lipgloss.Color("#0a0b10"),
	BgModal:     lipgloss.Color("#10111a"),
	BgModalList: lipgloss.Color("#1c1e25"),
	Border:      lipgloss.Color("#2a2b3d"),
	BorderDim:   lipgloss.Color("#1a1b26"),

	Text:          lipgloss.Color("#abb2bf"),
	TextLight:     lipgloss.Color("#90969f"),
	TextDim:       lipgloss.Color("#3e4451"),
	TextDark:      lipgloss.Color("#151516"),
	Muted:         lipgloss.Color("#4b5263"),
	MutedModal:    lipgloss.Color("#4b5263"),
	ModalText:     lipgloss.Color("#b5bdcb"),
	ModalTextDark: lipgloss.Color("#4b5263"),
	Cursor:        lipgloss.Color("#abb2bf"),
	CursorText:    lipgloss.Color("#151516"),

	Primary:      lipgloss.Color(ColorGold),
	PrimaryModal: lipgloss.Color(ColorGold),
	PrimaryDim:   lipgloss.Color(ColorGoldDim),
	PrimaryDark:  lipgloss.Color(ColorGoldDark),
	Secondary:    lipgloss.Color(ColorBlue),
	SecondaryDim: lipgloss.Color(ColorBlueDark),
	Tertiary:     lipgloss.Color(ColorGreenLight),
	TertiaryDim:  lipgloss.Color(ColorGreenDark),
	Warning:      lipgloss.Color(ColorRedLight),
	Rare:         lipgloss.Color(ColorPink),
	Highlight:    lipgloss.Color(ColorOrange),
	Green:        lipgloss.Color(ColorGreen),
	GreenDim:     lipgloss.Color(ColorGreenDark),
	Red:          lipgloss.Color(ColorRed),
	RedDim:       lipgloss.Color(ColorRedDark),

	BlockDefault:     lipgloss.Color(ColorPurple),
	BlockDefaultBg:   lipgloss.Color(ColorPurpleDark),
	BlockInfo:        lipgloss.Color(ColorGold),
	BlockInfoBg:      lipgloss.Color(ColorGoldDark),
	BlockSecondary:   lipgloss.Color(ColorBlue),
	BlockSecondaryBg: lipgloss.Color(ColorBlueDark),
	BlockTertiary:    lipgloss.Color(ColorGreenLight),
	BlockTertiaryBg:  lipgloss.Color(ColorGreenDark),
	BlockMuted:       lipgloss.Color(ColorGrey),
	BlockMutedBg:     lipgloss.Color(ColorGreyDark),
}

// ThemeClassic = original RuneScape palette
var ThemeClassic = Theme{
	Bg:          lipgloss.Color("#1b140b"),
	BgDim:       lipgloss.Color("#120c05"),
	BgInput:     lipgloss.Color("#2a1d0f"),
	BgInputDim:  lipgloss.Color("#1c1208"),
	BgPanel:     lipgloss.Color("#171006"),
	BgPanelDim:  lipgloss.Color("#100a04"),
	BgModal:     lipgloss.Color("#171006"),
	BgModalList: lipgloss.Color("#22170c"),
	Border:      lipgloss.Color("#6b4e2a"),
	BorderDim:   lipgloss.Color("#3a2915"),

	Text:          lipgloss.Color("#e2c98a"),
	TextLight:     lipgloss.Color("#c2a86d"),
	TextDim:       lipgloss.Color("#5c4324"),
	TextDark:      lipgloss.Color("#0c0804"),
	Muted:         lipgloss.Color("#7a5c35"),
	MutedModal:    lipgloss.Color("#7a5c35"),
	ModalText:     lipgloss.Color("#f0d79c"),
	ModalTextDark: lipgloss.Color("#5c4324"),
	Cursor:        lipgloss.Color("#f0d79c"),
	CursorText:    lipgloss.Color("#0c0804"),

	Primary:      lipgloss.Color("#d4af37"),
	PrimaryModal: lipgloss.Color("#d4af37"),
	PrimaryDim:   lipgloss.Color("#8a6d2e"),
	PrimaryDark:  lipgloss.Color("#4a3718"),
	Secondary:    lipgloss.Color("#6faa6f"),
	SecondaryDim: lipgloss.Color("#2d3d20"),
	Tertiary:     lipgloss.Color("#8fd18f"),
	TertiaryDim:  lipgloss.Color("#2d3d20"),
	Warning:      lipgloss.Color("#b94a3c"),
	Rare:         lipgloss.Color("#c77dff"),
	Highlight:    lipgloss.Color("#e2c98a"),
	Green:        lipgloss.Color("#6faa6f"),
	GreenDim:     lipgloss.Color("#2d3d20"),
	Red:          lipgloss.Color("#b94a3c"),
	RedDim:       lipgloss.Color("#4a1e18"),

	BlockDefault:     lipgloss.Color("#e2c98a"),
	BlockDefaultBg:   lipgloss.Color("#4a3718"),
	BlockInfo:        lipgloss.Color("#d4af37"),
	BlockInfoBg:      lipgloss.Color("#4a3718"),
	BlockSecondary:   lipgloss.Color("#6faa6f"),
	BlockSecondaryBg: lipgloss.Color("#2d3d20"),
	BlockTertiary:    lipgloss.Color("#8fd18f"),
	BlockTertiaryBg:  lipgloss.Color("#2d3d20"),
	BlockMuted:       lipgloss.Color("#a08963"),
	BlockMutedBg:     lipgloss.Color("#2a1d0f"),
}

// ThemeRuneLite = dark RuneLite palette
var ThemeRuneLite = Theme{
	Bg:          lipgloss.Color("#0f1115"),
	BgDim:       lipgloss.Color("#0a0c10"),
	BgInput:     lipgloss.Color("#1b1f26"),
	BgInputDim:  lipgloss.Color("#141820"),
	BgPanel:     lipgloss.Color("#12151b"),
	BgPanelDim:  lipgloss.Color("#0d1015"),
	BgModal:     lipgloss.Color("#12151b"),
	BgModalList: lipgloss.Color("#1b1f26"),
	Border:      lipgloss.Color("#2b313d"),
	BorderDim:   lipgloss.Color("#1b1f26"),

	Text:          lipgloss.Color("#cfd6e6"),
	TextLight:     lipgloss.Color("#a9b1c6"),
	TextDim:       lipgloss.Color("#3a4150"),
	TextDark:      lipgloss.Color("#0f1115"),
	Muted:         lipgloss.Color("#4f5666"),
	MutedModal:    lipgloss.Color("#4f5666"),
	ModalText:     lipgloss.Color("#dbe2f3"),
	ModalTextDark: lipgloss.Color("#4f5666"),
	Cursor:        lipgloss.Color("#00d4ff"),
	CursorText:    lipgloss.Color("#0f1115"),

	Primary:      lipgloss.Color("#00d4ff"),
	PrimaryModal: lipgloss.Color("#00d4ff"),
	PrimaryDim:   lipgloss.Color("#007a99"),
	PrimaryDark:  lipgloss.Color("#003844"),
	Secondary:    lipgloss.Color("#ffb454"),
	SecondaryDim: lipgloss.Color("#6a4a1b"),
	Tertiary:     lipgloss.Color("#6adf91"),
	TertiaryDim:  lipgloss.Color("#1f3d2e"),
	Warning:      lipgloss.Color("#ff5f56"),
	Rare:         lipgloss.Color("#c792ea"),
	Highlight:    lipgloss.Color("#ffb454"),
	Green:        lipgloss.Color("#6adf91"),
	GreenDim:     lipgloss.Color("#1f3d2e"),
	Red:          lipgloss.Color("#ff5f56"),
	RedDim:       lipgloss.Color("#4a1e1c"),

	BlockDefault:     lipgloss.Color("#00d4ff"),
	BlockDefaultBg:   lipgloss.Color("#003844"),
	BlockInfo:        lipgloss.Color("#ffb454"),
	BlockInfoBg:      lipgloss.Color("#6a4a1b"),
	BlockSecondary:   lipgloss.Color("#ffb454"),
	BlockSecondaryBg: lipgloss.Color("#6a4a1b"),
	BlockTertiary:    lipgloss.Color("#6adf91"),
	BlockTertiaryBg:  lipgloss.Color("#1f3d2e"),
	BlockMuted:       lipgloss.Color("#8b92a6"),
	BlockMutedBg:     lipgloss.Color("#1b1f26"),
}

// ThemeZamorak = red Zammy palette
var ThemeZamorak = Theme{
	Bg:          lipgloss.Color("#120c0c"),
	BgDim:       lipgloss.Color("#0c0808"),
	BgPanel:     lipgloss.Color("#160f0f"),
	BgPanelDim:  lipgloss.Color("#0e0909"),
	BgInput:     lipgloss.Color("#1c1414"),
	BgInputDim:  lipgloss.Color("#140e0e"),
	BgModal:     lipgloss.Color("#160f0f"),
	BgModalList: lipgloss.Color("#1c1414"),
	Border:      lipgloss.Color("#5a1e1e"),
	BorderDim:   lipgloss.Color("#2b0f0f"),

	Text:          lipgloss.Color("#e6d6d6"),
	TextLight:     lipgloss.Color("#caa"),
	TextDim:       lipgloss.Color("#402020"),
	TextDark:      lipgloss.Color("#120c0c"),
	Muted:         lipgloss.Color("#6a3a3a"),
	MutedModal:    lipgloss.Color("#6a3a3a"),
	ModalText:     lipgloss.Color("#f2e6e6"),
	ModalTextDark: lipgloss.Color("#6a3a3a"),
	Cursor:        lipgloss.Color("#ff4d4d"),
	CursorText:    lipgloss.Color("#120c0c"),

	Primary:      lipgloss.Color("#ff4d4d"),
	PrimaryModal: lipgloss.Color("#ff4d4d"),
	PrimaryDim:   lipgloss.Color("#8a1f1f"),
	PrimaryDark:  lipgloss.Color("#4a0f0f"),
	Secondary:    lipgloss.Color("#d4af37"),
	SecondaryDim: lipgloss.Color("#4a3718"),
	Tertiary:     lipgloss.Color("#6adf91"),
	TertiaryDim:  lipgloss.Color("#1f3d2e"),
	Warning:      lipgloss.Color("#ffb454"),
	Rare:         lipgloss.Color("#c792ea"),
	Highlight:    lipgloss.Color("#ff4d4d"),
	Green:        lipgloss.Color("#6adf91"),
	GreenDim:     lipgloss.Color("#1f3d2e"),
	Red:          lipgloss.Color("#ff4d4d"),
	RedDim:       lipgloss.Color("#4a1e1c"),

	BlockDefault:     lipgloss.Color("#ff4d4d"),
	BlockDefaultBg:   lipgloss.Color("#4a1e1c"),
	BlockInfo:        lipgloss.Color("#ff4d4d"),
	BlockInfoBg:      lipgloss.Color("#4a1e1c"),
	BlockSecondary:   lipgloss.Color("#d4af37"),
	BlockSecondaryBg: lipgloss.Color("#4a3718"),
	BlockTertiary:    lipgloss.Color("#6adf91"),
	BlockTertiaryBg:  lipgloss.Color("#1f3d2e"),
	BlockMuted:       lipgloss.Color("#5a1e1e"),
	BlockMutedBg:     lipgloss.Color("#261b1b"),
}

// Themes is the registry for config file / runtime lookup
var Themes = map[string]Theme{
	"modern":   ThemeModern,
	"classic":  ThemeClassic,
	"runelite": ThemeRuneLite,
	"zamorak":  ThemeZamorak,
}

// -- Config theme ----------

// DefaultTheme is used when no config value is set
const DefaultTheme = "modern"

// loadThemeFromConfig loads a theme from a key
func loadThemeFromConfig(name string) Theme {
	if t, ok := Themes[name]; ok {
		return t
	}
	return Themes[DefaultTheme]
}

// -- Dimmed ----------

// Dimmed returns a copy of the theme with all colors shifted to darker variants
func (t Theme) Dimmed() Theme {
	return Theme{
		Bg:          t.BgDim,
		BgDim:       t.BgDim,
		BgInput:     t.BgInputDim,
		BgPanel:     t.BgPanelDim,
		BgModal:     t.BgModal,     // modal color, leave undimmed
		BgModalList: t.BgModalList, // modal color, leave undimmed
		Border:      t.BorderDim,

		Text:          t.TextDim,
		TextLight:     t.Muted,
		TextDim:       t.TextDim,
		TextDark:      t.TextDark,
		Muted:         t.TextDim,
		MutedModal:    t.Muted,
		ModalText:     t.ModalText,     // modal color, leave undimmed
		ModalTextDark: t.ModalTextDark, // modal color, leave undimmed
		Cursor:        t.BgInputDim,    // hide cursor,
		CursorText:    t.TextDim,       // hide cursor,

		Primary:      t.PrimaryDim,
		PrimaryModal: t.PrimaryModal,
		PrimaryDim:   t.PrimaryDim,
		PrimaryDark:  t.PrimaryDark,
		Secondary:    t.SecondaryDim,
		SecondaryDim: t.SecondaryDim,
		Tertiary:     t.TertiaryDim,
		TertiaryDim:  t.TertiaryDim,
		Warning:      t.Muted,
		Rare:         t.Muted,
		Highlight:    t.Muted,
		Green:        t.GreenDim,
		Red:          t.RedDim,

		BlockDefault:     t.Muted,
		BlockDefaultBg:   t.BgDim,
		BlockInfo:        t.Muted,
		BlockInfoBg:      t.BgDim,
		BlockSecondary:   t.Muted,
		BlockSecondaryBg: t.BgDim,
		BlockTertiary:    t.Muted,
		BlockTertiaryBg:  t.BgDim,
		BlockMuted:       t.Muted,
		BlockMutedBg:     t.BgDim,
	}
}

// -- Style Methods ----------
// Each method returns a lipgloss.Style derived from the theme

// -- Base ----------

func (t Theme) Base() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.Bg).
		Foreground(t.Text)
}

func (t Theme) Bg_() lipgloss.Style {
	return lipgloss.NewStyle().Background(t.Bg)
}

func (t Theme) BgInput_() lipgloss.Style {
	return lipgloss.NewStyle().Background(t.BgInput)
}

func (t Theme) Body() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.Text)
}

func (t Theme) BodyDim() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.TextDim)
}

// -- Panel ----------

func (t Theme) PanelTitleActive() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Primary).
		Background(t.Bg).
		Bold(true).
		PaddingLeft(1).
		PaddingRight(1)
}

func (t Theme) PanelTitleInactive() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.TextLight).
		Background(t.Bg)
}

func (t Theme) PanelTitleMenu() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.ModalText)
}

func (t Theme) PanelBadgeMenu() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.ModalTextDark)
}

// -- Home ----------

func (t Theme) Logo() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Primary).
		Bold(true)
}

func (t Theme) LogoDim() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.PrimaryDim).
		Bold(true)
}

func (t Theme) HomeCmd() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Secondary).
		Background(t.Bg).
		Bold(true)
}

func (t Theme) HomeDesc() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Text).
		Background(t.Bg)
}

func (t Theme) HomeKeybind() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Highlight).
		Background(t.Bg)
}

func (t Theme) InputPrompt() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Text).
		Background(t.BgInput)
}

func (t Theme) InputPlaceholder() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.TextDim).
		Background(t.BgInput)
}

func (t Theme) InputCursor() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.Cursor).
		Foreground(t.CursorText)
}

func (t Theme) InputPromptModal() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.ModalText).
		Background(t.BgModal)
}

func (t Theme) InputPlaceholderModal() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.ModalTextDark).
		Background(t.BgModal)
}

func (t Theme) InputCursorModal() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.ModalText).
		Foreground(t.ModalTextDark)
}

// -- Sidebar ----------

func (t Theme) SidebarItem() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Text).
		Background(t.Bg)
}

func (t Theme) SidebarItemSelected() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Bg).
		Background(t.Primary).
		Bold(true).
		PaddingRight(1)
}

func (t Theme) SidebarItemMaxed() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Primary).
		Background(t.Bg)
}

func (t Theme) SidebarItem200M() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Rare).
		Background(t.Bg).
		Faint(true)
}

// -- XP Stats ----------

func (t Theme) StatHeader() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Muted).
		Bold(true).
		MarginBottom(1)
}

func (t Theme) StatLabel() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Primary).
		Background(t.Bg).
		Bold(true).
		Italic(true).
		Width(18)
}

func (t Theme) StatLabelActive() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Text).
		Background(t.Bg).
		Width(18)
}

func (t Theme) StatValue() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Primary).
		Background(t.Bg).
		Bold(true).
		Italic(true)
}

func (t Theme) StatValueActive() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.Bg)
}

func (t Theme) StatValueDim() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Muted).
		Background(t.Bg)
}

// -- Image ----------

func (t Theme) ImagePlaceholder() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Muted).
		Align(lipgloss.Center, lipgloss.Center)
}

// -- Status Bar ----------

func (t Theme) StatusKey() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.TextLight).
		Background(t.Bg)
}

func (t Theme) StatusVal() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Muted).
		Background(t.Bg)
}

func (t Theme) StatusLine() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.Bg).
		Foreground(t.Muted)
}

func (t Theme) StatusError() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.Bg).
		Foreground(t.Warning)
}

func (t Theme) StatusBlock() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.BlockDefaultBg).
		Foreground(t.BlockDefault)
}

func (t Theme) StatusBlockInfo() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.BlockInfoBg).
		Foreground(t.BlockInfo)
}

func (t Theme) StatusBlockMode1() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.BlockSecondaryBg).
		Foreground(t.BlockSecondary)
}

func (t Theme) StatusBlockMode2() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.BlockTertiaryBg).
		Foreground(t.BlockTertiary)
}

func (t Theme) StatusBlockMuted() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.BlockMutedBg).
		Foreground(t.BlockMuted)
}

// -- Help / Error ----------

func (t Theme) Help() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Text).
		Background(t.Bg).
		Italic(true)
}

func (t Theme) HelpMuted() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Muted).
		Background(t.Bg).
		Italic(true)
}

func (t Theme) Error() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(t.Warning)
}
