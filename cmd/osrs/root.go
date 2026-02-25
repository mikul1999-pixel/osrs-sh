package osrs

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mikul1999-pixel/osrs-sh/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "osrs",
	Short: "OSRS CLI — stats, XP, drops and more",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(
			ui.NewAppModel(),
		)
		if _, err := p.Run(); err != nil {
			return fmt.Errorf("TUI error: %w", err)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
