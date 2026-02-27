package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// OSRS hiscores skills. Ordered to match XP.go
var osrsToLocal = [24]int{
	0,  // Attack
	6,  // Defence
	3,  // Strength
	1,  // Hitpoints
	9,  // Ranged
	12, // Prayer
	15, // Magic
	11, // Cooking
	17, // Woodcutting
	16, // Fletching
	8,  // Fishing
	14, // Firemaking
	13, // Crafting
	5,  // Smithing
	2,  // Mining
	7,  // Herblore
	4,  // Agility
	10, // Thieving
	19, // Slayer
	20, // Farming
	18, // Runecraft
	22, // Hunter
	21, // Construction
	23, // Sailing
}

type Result struct {
	XP [24]int
}

func Lookup(rsn string) (Result, error) {
	url := fmt.Sprintf(
		"https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player=%s",
		rsn,
	)
	resp, err := http.Get(url)
	if err != nil {
		return Result{}, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return Result{}, fmt.Errorf("player %q not found", strings.ToLower(rsn))
	}
	if resp.StatusCode != 200 {
		return Result{}, fmt.Errorf("hiscores error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{}, fmt.Errorf("read error: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(body)), "\n")

	var result Result
	for i, localIdx := range osrsToLocal {
		apiLine := i + 1 // skip line 0 (Overall)
		if apiLine >= len(lines) {
			break
		}
		parts := strings.Split(strings.TrimSpace(lines[apiLine]), ",")
		if len(parts) < 3 {
			continue
		}
		rawXP, err := strconv.Atoi(parts[2])
		if err != nil || rawXP < 0 {
			rawXP = 0
		}
		result.XP[localIdx] = rawXP
	}
	return result, nil
}
