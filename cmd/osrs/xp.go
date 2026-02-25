package osrs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikul1999-pixel/osrs-sh/internal/core/xp"
	"github.com/spf13/cobra"
)

// Cmd returns the "osrs xp" command with subcommands
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xp",
		Short: "XP and leveling utilities",
	}

	cmd.AddCommand(calcCmd())
	cmd.AddCommand(tableCmd())

	return cmd
}

// osrs xp calc <level|xp> [--target <level>]
func calcCmd() *cobra.Command {
	var target int

	cmd := &cobra.Command{
		Use:   "calc <level or xp>",
		Short: "Calculate XP needed to reach next or target level",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			val, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("expected a number, got %q", args[0])
			}

			// Determine if the argument is a level or raw XP
			var currentXP int
			if val <= 126 {
				// Treat as level
				currentXP = xp.LevelToXP(val)
				fmt.Printf("Level %d = %s XP\n", val, formatInt(currentXP))
			} else {
				currentXP = val
				level := xp.XPToLevel(currentXP)
				fmt.Printf("XP %s = Level %d\n", formatInt(currentXP), level)
			}

			targetLevel := target
			if targetLevel == 0 {
				targetLevel = xp.XPToLevel(currentXP) + 1
			}

			if targetLevel > 99 {
				fmt.Println("Already at max level (99).")
				return nil
			}

			needed := xp.XPToLevel99(currentXP)
			toTarget := xp.XPBetween(currentXP, xp.LevelToXP(targetLevel))

			fmt.Printf("XP to level %d : %s\n", targetLevel, formatInt(toTarget))
			fmt.Printf("XP to level 99: %s\n", formatInt(needed))

			return nil
		},
	}

	cmd.Flags().IntVarP(&target, "target", "t", 0, "Target level (default: next level)")
	return cmd
}

// osrs xp table [--skill <name>]
func tableCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "table",
		Short: "Display the full XP table for levels 1–99",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%-8s %-14s %-14s\n", "Level", "Total XP", "XP to Next")
			fmt.Println(strings.Repeat("-", 38))
			for lvl := 1; lvl <= 99; lvl++ {
				totalXP := xp.LevelToXP(lvl)
				var toNext string
				if lvl < 99 {
					toNext = formatInt(xp.LevelToXP(lvl+1) - totalXP)
				} else {
					toNext = "—"
				}
				fmt.Printf("%-8d %-14s %-14s\n", lvl, formatInt(totalXP), toNext)
			}
			return nil
		},
	}
}

func formatInt(n int) string {
	s := strconv.Itoa(n)
	// Insert commas every 3 digits
	var b strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteRune(',')
		}
		b.WriteRune(c)
	}
	return b.String()
}
