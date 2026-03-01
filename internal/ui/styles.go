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
	ModalText     lipgloss.Color // default text color for modals
	ModalTextDark lipgloss.Color // darker text color for modals
	Cursor        lipgloss.Color
	CursorText    lipgloss.Color

	// Semantic
	Primary      lipgloss.Color // main accent
	PrimaryDim   lipgloss.Color
	PrimaryDark  lipgloss.Color
	Secondary    lipgloss.Color // supporting accent
	SecondaryDim lipgloss.Color
	Positive     lipgloss.Color
	PositiveDim  lipgloss.Color
	Warning      lipgloss.Color
	Rare         lipgloss.Color // rare highlight
	Highlight    lipgloss.Color // interactive element highlight
	Green        lipgloss.Color // fixed green color
	GreenDim     lipgloss.Color
	Red          lipgloss.Color // fixed red color
	RedDim       lipgloss.Color

	// Status blocks
	BlockDefault   lipgloss.Color // fg for default status block
	BlockDefaultBg lipgloss.Color
	BlockInfo      lipgloss.Color
	BlockInfoBg    lipgloss.Color
	BlockMode1     lipgloss.Color
	BlockMode1Bg   lipgloss.Color
	BlockMode2     lipgloss.Color
	BlockMode2Bg   lipgloss.Color
	BlockMuted     lipgloss.Color
	BlockMutedBg   lipgloss.Color
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
	BgModalList: lipgloss.Color("#10111a"),
	Border:      lipgloss.Color("#2a2b3d"),
	BorderDim:   lipgloss.Color("#1a1b26"),

	Text:          lipgloss.Color("#abb2bf"),
	TextLight:     lipgloss.Color("#90969f"),
	TextDim:       lipgloss.Color("#3e4451"),
	TextDark:      lipgloss.Color("#151516"),
	Muted:         lipgloss.Color("#4b5263"),
	ModalText:     lipgloss.Color("#b5bdcb"),
	ModalTextDark: lipgloss.Color("#4b5263"),
	Cursor:        lipgloss.Color("#abb2bf"),
	CursorText:    lipgloss.Color("#151516"),

	Primary:      lipgloss.Color(ColorGold),
	PrimaryDim:   lipgloss.Color(ColorGoldDim),
	PrimaryDark:  lipgloss.Color(ColorGoldDark),
	Secondary:    lipgloss.Color(ColorBlue),
	SecondaryDim: lipgloss.Color(ColorBlueDark),
	Positive:     lipgloss.Color(ColorGreenLight),
	PositiveDim:  lipgloss.Color(ColorGreenDark),
	Warning:      lipgloss.Color(ColorRedLight),
	Rare:         lipgloss.Color(ColorPink),
	Highlight:    lipgloss.Color(ColorOrange),
	Green:        lipgloss.Color(ColorGreen),
	GreenDim:     lipgloss.Color(ColorGreenDark),
	Red:          lipgloss.Color(ColorRed),
	RedDim:       lipgloss.Color(ColorRedDark),

	BlockDefault:   lipgloss.Color(ColorPurple),
	BlockDefaultBg: lipgloss.Color(ColorPurpleDark),
	BlockInfo:      lipgloss.Color(ColorGold),
	BlockInfoBg:    lipgloss.Color(ColorGoldDark),
	BlockMode1:     lipgloss.Color(ColorBlue),
	BlockMode1Bg:   lipgloss.Color(ColorBlueDark),
	BlockMode2:     lipgloss.Color(ColorGreenLight),
	BlockMode2Bg:   lipgloss.Color(ColorGreenDark),
	BlockMuted:     lipgloss.Color(ColorGrey),
	BlockMutedBg:   lipgloss.Color(ColorGreyDark),
}

// ThemeClassic = original RuneScape palette
var ThemeClassic = Theme{
	Bg:          lipgloss.Color("#1a1208"),
	BgDim:       lipgloss.Color("#110d05"),
	BgInput:     lipgloss.Color("#2e2010"),
	BgInputDim:  lipgloss.Color("#1e1508"),
	BgPanel:     lipgloss.Color("#150f06"),
	BgPanelDim:  lipgloss.Color("#0e0904"),
	BgModal:     lipgloss.Color("#150f06"),
	BgModalList: lipgloss.Color("#2e2010"),
	Border:      lipgloss.Color("#4a3520"),
	BorderDim:   lipgloss.Color("#2e1f0f"),

	Text:          lipgloss.Color("#c7a96b"),
	TextLight:     lipgloss.Color("#a08050"),
	TextDim:       lipgloss.Color("#4a3520"),
	TextDark:      lipgloss.Color("#0e0a04"),
	Muted:         lipgloss.Color("#6b4e2a"),
	ModalText:     lipgloss.Color("#dabc81"),
	ModalTextDark: lipgloss.Color("#6b4e2a"),
	Cursor:        lipgloss.Color("#c7a96b"),
	CursorText:    lipgloss.Color("#0e0a04"),

	Primary:      lipgloss.Color("#c8a951"),
	PrimaryDim:   lipgloss.Color("#7a6535"),
	PrimaryDark:  lipgloss.Color("#3d2c1f"),
	Secondary:    lipgloss.Color("#7aaa6a"),
	SecondaryDim: lipgloss.Color("#2e3a1e"),
	Positive:     lipgloss.Color("#98c379"),
	PositiveDim:  lipgloss.Color("#2e3a1e"),
	Warning:      lipgloss.Color("#cc5533"),
	Rare:         lipgloss.Color("#cc77aa"),
	Highlight:    lipgloss.Color("#d4a84b"),
	Green:        lipgloss.Color(ColorGreen),
	GreenDim:     lipgloss.Color(ColorGreenDark),
	Red:          lipgloss.Color(ColorRed),
	RedDim:       lipgloss.Color(ColorRedDark),

	BlockDefault:   lipgloss.Color("#c7a96b"),
	BlockDefaultBg: lipgloss.Color("#3d2c1f"),
	BlockInfo:      lipgloss.Color("#c8a951"),
	BlockInfoBg:    lipgloss.Color("#3d2c1f"),
	BlockMode1:     lipgloss.Color("#7aaa6a"),
	BlockMode1Bg:   lipgloss.Color("#2e3a1e"),
	BlockMode2:     lipgloss.Color("#98c379"),
	BlockMode2Bg:   lipgloss.Color("#2e3a1e"),
	BlockMuted:     lipgloss.Color("#8a8070"),
	BlockMutedBg:   lipgloss.Color("#2e2010"),
}

// Themes is the registry for config file / runtime lookup
var Themes = map[string]Theme{
	"modern":  ThemeModern,
	"classic": ThemeClassic,
}

// DefaultTheme is used when no config value is set
const DefaultTheme = "modern"

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
		ModalText:     t.ModalText,     // modal color, leave undimmed
		ModalTextDark: t.ModalTextDark, // modal color, leave undimmed
		Cursor:        t.BgInputDim,    // hide cursor,
		CursorText:    t.TextDim,       // hide cursor,

		Primary:      t.PrimaryDim,
		PrimaryDim:   t.PrimaryDim,
		PrimaryDark:  t.PrimaryDark,
		Secondary:    t.SecondaryDim,
		SecondaryDim: t.SecondaryDim,
		Positive:     t.PositiveDim,
		PositiveDim:  t.PositiveDim,
		Warning:      t.Muted,
		Rare:         t.Muted,
		Highlight:    t.Muted,
		Green:        t.GreenDim,
		Red:          t.RedDim,

		BlockDefault:   t.Muted,
		BlockDefaultBg: t.BgDim,
		BlockInfo:      t.Muted,
		BlockInfoBg:    t.BgDim,
		BlockMode1:     t.Muted,
		BlockMode1Bg:   t.BgDim,
		BlockMode2:     t.Muted,
		BlockMode2Bg:   t.BgDim,
		BlockMuted:     t.Muted,
		BlockMutedBg:   t.BgDim,
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
		Faint(true).
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
		Italic(true).
		Faint(true)
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
		Background(t.BlockMode1Bg).
		Foreground(t.BlockMode1)
}

func (t Theme) StatusBlockMode2() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(t.BlockMode2Bg).
		Foreground(t.BlockMode2)
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

// -- Config theme ----------

// dummy function to load from config file. Will make work later
func loadThemeFromConfig(name string) Theme {
	if t, ok := Themes[name]; ok {
		return t
	}
	return ThemeModern
}
