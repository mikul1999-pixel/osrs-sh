package themes

import (
	"embed"
	"encoding/json"
)

// -- Import Theme Definitions ----------

//go:embed *.json
var themeFiles embed.FS

func LoadThemeJson(name string) (map[string]string, error) {
	data, err := themeFiles.ReadFile(name + ".json")
	if err != nil {
		return nil, err
	}

	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return m, nil
}
