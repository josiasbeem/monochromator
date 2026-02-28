package palette

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Palette struct {
	BG          string `json:"bg"`
	BGSecondary string `json:"bg_secondary"`
	BGSelection string `json:"bg_selection"`
	BGElevated  string `json:"bg_elevated"`
	FG          string `json:"fg"`
	FGContrast  string `json:"fg_contrast"`
	FGSecondary string `json:"fg_secondary"`
	FGMuted     string `json:"fg_muted"`
	Border      string `json:"border"`
	Red         string `json:"red"`
	Orange      string `json:"orange"`
	Yellow      string `json:"yellow"`
	Green       string `json:"green"`
	Cyan        string `json:"cyan"`
	Blue        string `json:"blue"`
	Magenta     string `json:"magenta"`
}

type Variant struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (p Palette) Color(name string) (string, error) {
	switch name {
	case "red":
		return p.Red, nil
	case "orange":
		return p.Orange, nil
	case "yellow":
		return p.Yellow, nil
	case "green":
		return p.Green, nil
	case "cyan":
		return p.Cyan, nil
	case "blue":
		return p.Blue, nil
	case "magenta":
		return p.Magenta, nil
	default:
		return "", fmt.Errorf("unknown color %q", name)
	}
}

type Config struct {
	Palettes map[string]Palette
	Variants []Variant
}

func (c *Config) Validate() error {
	for _, mode := range []string{"dark", "light"} {
		if _, ok := c.Palettes[mode]; !ok {
			return fmt.Errorf("missing %q palette", mode)
		}
	}
	if len(c.Variants) == 0 {
		return fmt.Errorf("no variants defined")
	}
	hasDefault := false
	for i, v := range c.Variants {
		if v.Name == "" {
			return fmt.Errorf("variant %d has empty name", i)
		}
		if v.Name == "default" {
			hasDefault = true
		}
		if _, err := c.Palettes["dark"].Color(v.Color); err != nil {
			return fmt.Errorf("variant %q: %w", v.Name, err)
		}
	}
	if !hasDefault {
		return fmt.Errorf("missing \"default\" variant")
	}
	return nil
}

func Load(paletteDir string) (*Config, error) {
	baseData, err := os.ReadFile(filepath.Join(paletteDir, "base.json"))
	if err != nil {
		return nil, fmt.Errorf("loading base palette: %w", err)
	}
	var palettes map[string]Palette
	if err := json.Unmarshal(baseData, &palettes); err != nil {
		return nil, fmt.Errorf("parsing base palette: %w", err)
	}

	varData, err := os.ReadFile(filepath.Join(paletteDir, "variants.json"))
	if err != nil {
		return nil, fmt.Errorf("loading variants: %w", err)
	}
	var variants []Variant
	if err := json.Unmarshal(varData, &variants); err != nil {
		return nil, fmt.Errorf("parsing variants: %w", err)
	}

	cfg := &Config{Palettes: palettes, Variants: variants}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return cfg, nil
}

func Pick(mode, dark, light string) string {
	if mode == "dark" {
		return dark
	}
	return light
}
