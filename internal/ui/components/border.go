package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// -- Custom Border Component ----------

// -- Alignment ----------

// Align controls where a title sits along the top or bottom border
type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

// -- Decorator ----------

// Decorator controls the visual style wrapping each title
type Decorator int

const (
	DecoratorNone Decorator = iota
	DecoratorDash
	DecoratorBrackets
	DecoratorAngle
)

func (d Decorator) wrap() (string, string) {
	switch d {
	case DecoratorDash:
		return "─ ", " "
	case DecoratorBrackets:
		return "[ ", " ]"
	case DecoratorAngle:
		return "< ", " >"
	default:
		return " ", " "
	}
}

// -- State styles ----------

// StateStyle holds the colors that change when the panel toggles focus
type StateStyle struct {
	BorderColor lipgloss.Color
	TitleColor  lipgloss.Color
	TitleBg     lipgloss.Color
	BadgeColor  lipgloss.Color
}

// -- Panel ----------

// Panel is the main component
type Panel struct {
	width  int
	height int // 0 = auto

	borderStyle lipgloss.Border

	// Top title
	title       string
	titleAlign  Align
	titleOffset int
	decorator   Decorator

	// Top badge
	badge       string
	badgeOffset int

	// Bottom title
	bottomTitle      string
	bottomTitleAlign Align

	// Bottom badge
	bottomBadge string

	// Focus states
	active   StateStyle
	inactive StateStyle

	// Background
	bgColor lipgloss.Color

	// Inner padding
	paddingTop    int
	paddingRight  int
	paddingBottom int
	paddingLeft   int
}

// New returns a Panel with defaults and the given inner width
func New(width int) *Panel {
	return &Panel{
		width:       width,
		borderStyle: lipgloss.RoundedBorder(),
		titleOffset: 1,
		badgeOffset: 1,
		decorator:   DecoratorDash,
		titleAlign:  AlignLeft,
		active: StateStyle{
			BorderColor: lipgloss.Color("#7aa2f7"),
			TitleColor:  lipgloss.Color("#c0caf5"),
		},
		inactive: StateStyle{
			BorderColor: lipgloss.Color("#3b4261"),
			TitleColor:  lipgloss.Color("#565f89"),
		},
		paddingRight: 1,
		paddingLeft:  1,
	}
}

// -- Chainable setters ----------

func (p *Panel) Width(w int) *Panel                   { p.width = w; return p }
func (p *Panel) Height(h int) *Panel                  { p.height = h; return p }
func (p *Panel) BorderStyle(b lipgloss.Border) *Panel { p.borderStyle = b; return p }

func (p *Panel) Title(t string) *Panel     { p.title = t; return p }
func (p *Panel) TitleAlign(a Align) *Panel { p.titleAlign = a; return p }
func (p *Panel) TitleOffset(n int) *Panel  { p.titleOffset = n; return p }

func (p *Panel) Decorator(d Decorator) *Panel { p.decorator = d; return p }

func (p *Panel) Badge(b string) *Panel           { p.badge = b; return p }
func (p *Panel) BadgeOffset(n int) *Panel        { p.badgeOffset = n; return p }
func (p *Panel) BottomTitle(t string) *Panel     { p.bottomTitle = t; return p }
func (p *Panel) BottomTitleAlign(a Align) *Panel { p.bottomTitleAlign = a; return p }
func (p *Panel) BottomBadge(b string) *Panel     { p.bottomBadge = b; return p }

func (p *Panel) ActiveBorderColor(c lipgloss.Color) *Panel   { p.active.BorderColor = c; return p }
func (p *Panel) InactiveBorderColor(c lipgloss.Color) *Panel { p.inactive.BorderColor = c; return p }

func (p *Panel) ActiveTitleColor(c lipgloss.Color) *Panel   { p.active.TitleColor = c; return p }
func (p *Panel) InactiveTitleColor(c lipgloss.Color) *Panel { p.inactive.TitleColor = c; return p }
func (p *Panel) ActiveTitleBg(c lipgloss.Color) *Panel      { p.active.TitleBg = c; return p }
func (p *Panel) InactiveTitleBg(c lipgloss.Color) *Panel    { p.inactive.TitleBg = c; return p }

func (p *Panel) ActiveBadgeColor(c lipgloss.Color) *Panel   { p.active.BadgeColor = c; return p }
func (p *Panel) InactiveBadgeColor(c lipgloss.Color) *Panel { p.inactive.BadgeColor = c; return p }

func (p *Panel) ActiveState(s StateStyle) *Panel   { p.active = s; return p }
func (p *Panel) InactiveState(s StateStyle) *Panel { p.inactive = s; return p }

func (p *Panel) BgColor(c lipgloss.Color) *Panel { p.bgColor = c; return p }

// Padding sets uniform inner padding
func (p *Panel) Padding(vertical, horizontal int) *Panel {
	p.paddingTop = vertical
	p.paddingBottom = vertical
	p.paddingLeft = horizontal
	p.paddingRight = horizontal
	return p
}

// PaddingFull sets per side inner padding
func (p *Panel) PaddingFull(top, right, bottom, left int) *Panel {
	p.paddingTop = top
	p.paddingRight = right
	p.paddingBottom = bottom
	p.paddingLeft = left
	return p
}

// -- Render ----------

// Render produces the final string
func (p *Panel) Render(content string, active bool) string {
	state := p.inactive
	if active {
		state = p.active
	}

	borderColor := state.BorderColor

	// Dash character from the current border style's top segment
	dash := p.borderStyle.Top
	if dash == "" {
		dash = "─"
	}
	// lipgloss border uses a single rune for the top. grab that
	if len([]rune(dash)) > 1 {
		dash = string([]rune(dash)[:1])
	}

	topLine := p.buildTopBorder(state, borderColor, dash)
	bottomLine := p.buildBottomBorder(state, borderColor, dash)

	// Style the three remaining sides
	sideStyle := lipgloss.NewStyle().
		Border(p.borderStyle).
		BorderTop(false).
		BorderBottom(false).
		BorderForeground(borderColor).
		BorderBackground(p.bgColor).
		Background(p.bgColor).
		PaddingTop(p.paddingTop).
		PaddingRight(p.paddingRight).
		PaddingBottom(p.paddingBottom).
		PaddingLeft(p.paddingLeft).
		Width(p.width)

	if p.height > 0 {
		// Subtract top & bottom border lines from inner height
		innerHeight := p.height - 2
		if innerHeight < 0 {
			innerHeight = 0
		}
		sideStyle = sideStyle.Height(innerHeight)
	}

	middle := sideStyle.Render(content)

	return lipgloss.JoinVertical(lipgloss.Left,
		topLine,
		middle,
		bottomLine,
	)
}

// -- Border line builders ----------

func (p *Panel) buildTopBorder(state StateStyle, borderColor lipgloss.Color, dash string) string {
	tl := p.borderStyle.TopLeft
	tr := p.borderStyle.TopRight
	if tl == "" {
		tl = "╭"
	}
	if tr == "" {
		tr = "╮"
	}

	// Total inner width available for dashes + titles.
	innerWidth := p.width

	titleStr := p.buildTitleSegment(p.title, state)
	badgeStr := ""
	if p.badge != "" {
		badgeStr = p.buildBadgeSegment(p.badge, state, dash)
	}

	return p.assembleBorderLine(tl, tr, dash, innerWidth, titleStr, badgeStr, p.titleAlign, borderColor)
}

func (p *Panel) buildBottomBorder(state StateStyle, borderColor lipgloss.Color, dash string) string {
	bl := p.borderStyle.BottomLeft
	br := p.borderStyle.BottomRight
	if bl == "" {
		bl = "╰"
	}
	if br == "" {
		br = "╯"
	}

	innerWidth := p.width

	titleStr := p.buildTitleSegment(p.bottomTitle, state)
	badgeStr := ""
	if p.bottomBadge != "" {
		badgeStr = p.buildBadgeSegment(p.bottomBadge, state, dash)
	}

	return p.assembleBorderLine(bl, br, dash, innerWidth, titleStr, badgeStr, p.bottomTitleAlign, borderColor)
}

// assembleBorderLine composes a full border line string
func (p *Panel) assembleBorderLine(
	left, right string,
	dash string,
	innerWidth int,
	titleStr, badgeStr string,
	align Align,
	borderColor lipgloss.Color,
) string {

	titleWidth := lipgloss.Width(titleStr)
	badgeWidth := lipgloss.Width(badgeStr)

	// Available dash space after title + badge
	dashSpace := innerWidth - titleWidth - badgeWidth
	if dashSpace < 0 {
		dashSpace = 0
	}

	colorBorder := func(s string) string {
		return lipgloss.NewStyle().Foreground(borderColor).Background(p.bgColor).Render(s)
	}

	var line string

	switch align {
	case AlignCenter:
		half := dashSpace / 2
		extra := dashSpace % 2
		line = colorBorder(left) +
			colorBorder(strings.Repeat(dash, half)) +
			titleStr +
			colorBorder(strings.Repeat(dash, half+extra)) +
			badgeStr +
			colorBorder(right)

	case AlignRight:
		line = colorBorder(left) +
			colorBorder(strings.Repeat(dash, dashSpace)) +
			titleStr +
			badgeStr +
			colorBorder(right)

	default: // AlignLeft
		leadingDashes := p.titleOffset
		if leadingDashes > dashSpace {
			leadingDashes = dashSpace
		}
		remaining := dashSpace - leadingDashes
		// If badge, keep trailing space for it
		line = colorBorder(left) +
			colorBorder(strings.Repeat(dash, leadingDashes)) +
			titleStr +
			colorBorder(strings.Repeat(dash, remaining)) +
			badgeStr +
			colorBorder(right)
	}

	return line
}

// buildTitleSegment returns a styled title segment
func (p *Panel) buildTitleSegment(title string, state StateStyle) string {
	if title == "" {
		return ""
	}

	prefix, suffix := p.decorator.wrap()

	// Replace spaces inside the decorator with dashes when using dash decorator
	if p.decorator == DecoratorDash {
		prefix = strings.ReplaceAll(prefix, " ", "")
		suffix = strings.ReplaceAll(suffix, " ", "")
		prefix = "─ "
		suffix = " "
	}

	bg := state.TitleBg
	if bg == "" {
		bg = p.bgColor
	}
	ts := lipgloss.NewStyle().Foreground(state.TitleColor).Background(bg)
	bc := lipgloss.NewStyle().Foreground(state.BorderColor).Background(bg)

	return bc.Render(prefix) + ts.Render(title) + bc.Render(suffix)
}

// buildBadgeSegment returns a styled badge segment
func (p *Panel) buildBadgeSegment(badge string, state StateStyle, dash string) string {
	if badge == "" {
		return ""
	}

	prefix, suffix := p.decorator.wrap()
	if p.decorator == DecoratorDash {
		prefix = "─ "
		suffix = " "
	}

	badgeColor := state.BadgeColor
	if badgeColor == "" {
		badgeColor = state.TitleColor
	}

	bg := state.TitleBg
	if bg == "" {
		bg = p.bgColor
	}
	bs := lipgloss.NewStyle().Foreground(badgeColor).Background(bg)
	bc := lipgloss.NewStyle().Foreground(state.BorderColor).Background(bg)

	leadingDashes := strings.Repeat(dash, p.badgeOffset)

	return bc.Render(leadingDashes) + bc.Render(prefix) + bs.Render(badge) + bc.Render(suffix)
}
