// Package palette loads and validates the shared color palette from
// the palette/ directory.
package palette

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Palette is the full color set for one mode (dark or light). BGLine is
// the background of the line with the cursor. It has its own key because
// BGSecondary also colors the side bar, and in the dark mode that step is
// too strong for a line inside the code.
type Palette struct {
	BG          string `json:"bg"`
	BGSecondary string `json:"bg_secondary"`
	BGSelection string `json:"bg_selection"`
	BGElevated  string `json:"bg_elevated"`
	BGLine      string `json:"bg_line"`
	FG          string `json:"fg"`
	FGContrast  string `json:"fg_contrast"`
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

// Variant is one accent-color variant of the theme.
type Variant struct {
	// Name is the lowercase variant name. The name "default" is special:
	// it produces the base theme without a name suffix.
	Name string `json:"name"`
	// Color is the name of a Palette accent color, for example "cyan".
	Color string `json:"color"`
}

// Color returns the hex value of the named accent color. Valid names
// are red, orange, yellow, green, cyan, blue, and magenta.
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

// Config is the parsed content of the palette directory.
type Config struct {
	// Palettes maps a mode ("dark" or "light") to its color set.
	Palettes map[string]Palette
	Variants []Variant
}

// Validate makes sure that both mode palettes exist and have a line
// background, that a "default" variant exists, and that each variant has
// a name and a known color.
func (c *Config) Validate() error {
	for _, mode := range []string{"dark", "light"} {
		p, ok := c.Palettes[mode]
		if !ok {
			return fmt.Errorf("missing %q palette", mode)
		}
		if p.BGLine == "" {
			return fmt.Errorf("%q palette has no bg_line", mode)
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

// Load reads base.json and variants.json from paletteDir and
// validates the combined config.
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

// Pick returns dark if mode is "dark". For all other modes, it
// returns light.
func Pick(mode, dark, light string) string {
	if mode == "dark" {
		return dark
	}
	return light
}
