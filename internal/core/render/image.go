// Package render handles converting remote images to ANSI art via chafa.
package render

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DefaultSize is the chafa output size. width x height
const DefaultSize = "20x12"

// ImageToANSI fetches an image from url, converts it via chafa
func ImageToANSI(url string, size string) (string, error) {
	if size == "" {
		size = DefaultSize
	}

	// Check cache first
	if cached, ok := globalCache.get(url); ok {
		return cached, nil
	}

	// Download to a temp file
	tmpPath, err := downloadToTemp(url)
	if err != nil {
		return "", fmt.Errorf("download failed: %w", err)
	}
	defer os.Remove(tmpPath)

	// Shell out to chafa
	ansi, err := runChafa(tmpPath, size)
	if err != nil {
		return "", fmt.Errorf("chafa error: %w", err)
	}

	globalCache.set(url, ansi)
	return ansi, nil
}

// SkillIconURL returns the OSRS wiki URL for a given skill name
func SkillIconURL(skillName string) string {
	name := strings.ReplaceAll(skillName, " ", "_")
	return fmt.Sprintf("https://oldschool.runescape.wiki/images/%s_icon.png", name)
}

// ChafaAvailable checks whether chafa is installed on the system
func ChafaAvailable() bool {
	_, err := exec.LookPath("chafa")
	return err == nil
}

// -- Internal ----------

func downloadToTemp(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d for %s", resp.StatusCode, url)
	}

	// Preserve extension so chafa can detect format
	ext := filepath.Ext(url)
	if ext == "" {
		ext = ".png"
	}

	tmp, err := os.CreateTemp("", "osrs-img-*"+ext)
	if err != nil {
		return "", err
	}
	defer tmp.Close()

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		return "", err
	}
	return tmp.Name(), nil
}

func runChafa(path, size string) (string, error) {
	cmd := exec.Command("chafa",
		"--size", size,
		"--format", "symbols", // ANSI symbols mode
		path,
	)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
