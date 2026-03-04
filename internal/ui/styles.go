package ui

import (
	"log"
	"reflect"

	"github.com/charmbracelet/lipgloss"
	th "github.com/mikul1999-pixel/osrs-sh/internal/themes"
)

// -- Theme on Startup ----------
const DefaultTheme = "modern"

// -- Raw Colors ----------

var ColorStore = map[string]string{
	"ColorBlack":      "#000000",
	"ColorGold":       "#c8a951",
	"ColorGoldDim":    "#7a6535",
	"ColorGoldDark":   "#3d2c1f",
	"ColorGreen":      "#57a854",
	"ColorGreenLight": "#98c379",
	"ColorGreenDark":  "#2e3a4d",
	"ColorBlue":       "#3b98d4",
	"ColorBlueDark":   "#2e3a4d",
	"ColorPink":       "#ff66b2",
	"ColorPurple":     "#b49dd8",
	"ColorPurpleDark": "#373349",
	"ColorGrey":       "#8a93b2",
	"ColorGreyDark":   "#2a2b3d",
	"ColorRed":        "#c94f4f",
	"ColorRedLight":   "#e06c75",
	"ColorRedDark":    "#4d2929",
	"ColorOrange":     "#d19a66",
}

// -- Theme Struct ----------

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

// -- Load Preset Themes ----------

// Registry for runtime lookup
var knownThemes = []string{"modern", "classic", "desert", "runelite", "guthix", "saradomin", "zamorak"}
var themeCache = map[string]Theme{}

// Helpers for loading jsons
func resolveColor(val string) lipgloss.Color {
	if hex, ok := ColorStore[val]; ok {
		return lipgloss.Color(hex)
	}
	return lipgloss.Color(val)
}
func themeFromMap(m map[string]string) Theme {
	var t Theme
	v := reflect.ValueOf(&t).Elem()

	for key, val := range m {
		// Replace constant names with hex values
		field := v.FieldByName(key)
		if !field.IsValid() || !field.CanSet() {
			log.Printf("warning: unknown theme key %q - skipping", key)
			continue
		}
		field.Set(reflect.ValueOf(resolveColor(val)))
	}

	return t
}

// loadTheme gets a theme from its json
func loadTheme(name string) Theme {
	if t, ok := themeCache[name]; ok {
		return t
	}
	m, err := th.LoadThemeJson(name)
	if err != nil {
		log.Printf("warning: could not load theme %q: %v — falling back to %q", name, err, DefaultTheme)
		if name == DefaultTheme {
			panic("failed to load default theme")
		}
		return loadTheme(DefaultTheme)
	}
	t := themeFromMap(m)
	themeCache[name] = t // cache the theme
	return t
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
