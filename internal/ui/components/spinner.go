package components

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

// -- Custom Spinner Component ----------

// SpinnerFrames defines the animation frames for spinner
type SpinnerFrames []string

// Built in frame set
var (
	SpinnerBraille SpinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	SpinnerDots    SpinnerFrames = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	SpinnerASCII   SpinnerFrames = []string{"|", "/", "-", "\\"}
	SpinnerBlock   SpinnerFrames = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"}
)

// SpinnerTickMsg is sent on each animation tick
type SpinnerTickMsg struct{ ID int }

// Spinner is a standalone, embeddable spinner component
type Spinner struct {
	id       int
	frames   SpinnerFrames
	current  int
	interval time.Duration
	style    lipgloss.Style
	active   bool
}

var spinnerIDCounter int

// NewSpinner creates a spinner
func NewSpinner() *Spinner {
	spinnerIDCounter++
	return &Spinner{
		id:       spinnerIDCounter,
		frames:   SpinnerBraille,
		interval: 80 * time.Millisecond,
		style:    lipgloss.NewStyle(),
		active:   false,
	}
}

// SetFrames sets the animation frame set
func (s *Spinner) SetFrames(frames SpinnerFrames) *Spinner {
	s.frames = frames
	s.current = 0
	return s
}

// SetInterval sets how fast the spinner animates
func (s *Spinner) SetInterval(d time.Duration) *Spinner {
	s.interval = d
	return s
}

// SetStyle sets the lipgloss style applied to each frame
func (s *Spinner) SetStyle(style lipgloss.Style) *Spinner {
	s.style = style
	return s
}

// Start marks the spinner as active
func (s *Spinner) Start() tea.Cmd {
	s.active = true
	s.current = 0
	return s.TickCmd()
}

// Stop marks the spinner as inactive
func (s *Spinner) Stop() {
	s.active = false
}

func (s *Spinner) Active() bool {
	return s.active
}

// ID returns the unique ID for this spinner instance
func (s *Spinner) ID() int {
	return s.id
}

// Tick advances the spinner by one frame
func (s *Spinner) Tick() {
	if len(s.frames) == 0 {
		return
	}
	s.current = (s.current + 1) % len(s.frames)
}

// TickCmd returns the tea.Cmd that fires the next SpinnerTickMsg
func (s *Spinner) TickCmd() tea.Cmd {
	id := s.id
	interval := s.interval
	return tea.Tick(interval, func(time.Time) tea.Msg {
		return SpinnerTickMsg{ID: id}
	})
}

// View returns the current frame as a string
func (s *Spinner) View() string {
	if !s.active || len(s.frames) == 0 {
		return ""
	}
	return s.style.Render(s.frames[s.current])
}
