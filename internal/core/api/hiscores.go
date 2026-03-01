package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL    = "https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player="
	maxRetries = 1
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

type Service struct {
	client *Client
}

func New(client *Client) *Service {
	return &Service{client: client}
}

func (s *Service) Lookup(rsn string) (Result, error) {
	encoded := url.QueryEscape(strings.TrimSpace(rsn))
	fullURL := baseURL + encoded

	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return Result{}, fmt.Errorf("request build error: %w", err)
		}

		body, status, err := s.client.do(req)
		if err != nil {
			lastErr = fmt.Errorf("network error: %w", err)
			continue
		}

		switch status {
		case 200:
			return parseHiscores(body)
		case 404:
			return Result{}, fmt.Errorf("player %q not found", strings.ToLower(rsn))
		case 429, 500, 502, 503, 504:
			lastErr = fmt.Errorf("temporary hiscores error: %d", status)
			time.Sleep(2 * time.Second)
			continue
		default:
			return Result{}, fmt.Errorf("hiscores error: %d", status)
		}
	}

	return Result{}, lastErr
}

func parseHiscores(body []byte) (Result, error) {
	lines := strings.Split(strings.TrimSpace(string(body)), "\n")
	if len(lines) < 25 {
		return Result{}, errors.New("unexpected hiscores format")
	}

	var result Result
	for i, localIdx := range osrsToLocal {
		apiLine := i + 1 // skip overall
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
