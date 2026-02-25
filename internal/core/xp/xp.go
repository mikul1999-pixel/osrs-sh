// Package xp implements osrs xp -> level calcs
// Formula: https://oldschool.runescape.wiki/w/Experience
package xp

import "math"

// LevelToXP returns the total XP required to reach given level
func LevelToXP(level int) int {
	if level <= 1 {
		return 0
	}
	total := 0
	for l := 1; l < level; l++ {
		total += int(math.Floor(float64(l) + 300*math.Pow(2, float64(l)/7.0)))
	}
	return total / 4
}

// XPToLevel returns the level for a given XP
func XPToLevel(totalXP int) int {
	for lvl := 98; lvl >= 1; lvl-- {
		if totalXP >= LevelToXP(lvl) {
			return lvl
		}
	}
	return 1
}

// XPBetween returns the XP difference between current and target
func XPBetween(currentXP, targetXP int) int {
	diff := targetXP - currentXP
	if diff < 0 {
		return 0
	}
	return diff
}

// XPToLevel99 returns XP needed from current XP to reach level 99
func XPToLevel99(currentXP int) int {
	return XPBetween(currentXP, LevelToXP(99))
}

// XPToNextLevel returns XP needed to reach the next level
func XPToNextLevel(currentXP int) int {
	currentLevel := XPToLevel(currentXP)
	if currentLevel >= 99 {
		return 0
	}
	return XPBetween(currentXP, LevelToXP(currentLevel+1))
}
