package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui"
)

func main() {
	p := tea.NewProgram(
		ui.NewAppModel(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
