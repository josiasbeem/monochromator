package zed

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"generate/palette"
)

func plainSyntax(p palette.Palette) map[string]any {
	return map[string]any{
		"comment": map[string]any{"color": p.FGMuted, "font_style": "italic"},
	}
}

func defaultSyntax(p palette.Palette, accent string) map[string]any {
	syntax := plainSyntax(p)
	maps.Copy(syntax, map[string]any{
		"boolean":         map[string]any{"color": accent, "font_weight": 700},
		"constant":        map[string]any{"color": p.FG},
		"constructor":     map[string]any{"color": p.FGContrast, "font_weight": 700},
		"emphasis":        map[string]any{"color": p.FG, "font_style": "italic"},
		"emphasis.strong": map[string]any{"color": p.FGContrast, "font_weight": 700},
		"function":        map[string]any{"color": p.FG, "font_style": "italic"},
		"keyword":         map[string]any{"color": p.FGContrast, "font_weight": 700},
		"number":          map[string]any{"color": accent},
		"operator":        map[string]any{"color": p.FGMuted},
		"punctuation":     map[string]any{"color": p.FGMuted},
		"string":          map[string]any{"color": accent},
		"string.escape":   map[string]any{"color": accent, "font_weight": 700},
		"string.special":  map[string]any{"color": accent, "font_weight": 700},
		"text.literal":    map[string]any{"color": accent},
		"type":            map[string]any{"color": p.FGContrast, "font_weight": 700},
		"variable":        map[string]any{"color": p.FG},
	})
	return syntax
}

func buildStyle(mode string, p palette.Palette, syntax map[string]any) map[string]any {
	return map[string]any{
		"border":                        p.BGSecondary,
		"border.variant":                "#00000000",
		"created":                       p.Green,
		"deleted":                       p.Red,
		"editor.active_line.background": p.BGSecondary,
		"editor.background":             p.BG,
		"editor.document_highlight.bracket_background": p.BGSelection,
		"editor.foreground":                            p.FG,
		"editor.gutter.background":                     p.BG,
		"ghost_element.active":                         p.BGSelection,
		"ghost_element.hover":                          p.BGSelection,
		"ghost_element.background":                     "#00000000",
		"element.background":                           p.BGSelection,
		"elevated_surface.background":                  p.BGSecondary,
		"error":                                        p.Red,
		"hint":                                         p.FGMuted,
		"info":                                         p.Blue,
		"info.background":                              p.BGSelection,
		"modified":                                     p.Orange,
		"panel.background":                             p.BGSecondary,
		"players": []any{
			map[string]any{"cursor": p.FG, "selection": p.BGSelection, "background": p.FG},
			map[string]any{"cursor": p.Blue, "selection": p.Blue + "30", "background": p.Blue},
			map[string]any{"cursor": p.Red, "selection": p.Red + "30", "background": p.Red},
			map[string]any{"cursor": p.Green, "selection": p.Green + "30", "background": p.Green},
			map[string]any{"cursor": p.Yellow, "selection": p.Yellow + "30", "background": p.Yellow},
			map[string]any{"cursor": p.Cyan, "selection": p.Cyan + "30", "background": p.Cyan},
			map[string]any{"cursor": p.Magenta, "selection": p.Magenta + "30", "background": p.Magenta},
			map[string]any{"cursor": p.Orange, "selection": p.Orange + "30", "background": p.Orange},
		},
		"scrollbar.thumb.active_background": palette.Pick(mode, p.BGElevated+"90", p.FGMuted+"70"),
		"scrollbar.thumb.background":        palette.Pick(mode, p.BGElevated+"43", p.BGSelection+"60"),
		"scrollbar.thumb.hover_background":  palette.Pick(mode, p.BGElevated+"70", p.FGMuted+"50"),
		"scrollbar.track.border":            p.BGSecondary,
		"status_bar.background":             p.BGSecondary,
		"syntax":                            syntax,
		"tab.active_background":             p.BG,
		"tab.inactive_background":           p.BGSecondary,
		"tab_bar.background":                p.BGSecondary,
		"terminal.ansi.black":               palette.Pick(mode, p.BG, p.FG),
		"terminal.ansi.blue":                p.Blue,
		"terminal.ansi.bright_black":        p.FGMuted,
		"terminal.ansi.bright_blue":         p.Blue,
		"terminal.ansi.bright_cyan":         p.Cyan,
		"terminal.ansi.bright_green":        p.Green,
		"terminal.ansi.bright_magenta":      p.Magenta,
		"terminal.ansi.bright_red":          p.Red,
		"terminal.ansi.bright_white":        palette.Pick(mode, p.FG, "#ffffff"),
		"terminal.ansi.bright_yellow":       p.Yellow,
		"terminal.ansi.cyan":                p.Cyan,
		"terminal.ansi.green":               p.Green,
		"terminal.ansi.magenta":             p.Magenta,
		"terminal.ansi.red":                 p.Red,
		"terminal.ansi.white":               p.FGMuted,
		"terminal.ansi.yellow":              p.Yellow,
		"terminal.background":               p.BG,
		"title_bar.background":              p.BGSecondary,
		"version_control.added":             p.Green,
		"version_control.deleted":           p.Red,
		"version_control.modified":          p.Orange,
		"warning":                           p.Orange,
	}
}

func buildThemeFamily(familyName string, theme map[string]any) map[string]any {
	return map[string]any{
		"$schema": "https://zed.dev/schema/themes/v0.2.0.json",
		"name":    familyName,
		"author":  "Josias Beem <mail@beem.software>",
		"themes":  []any{theme},
	}
}

func buildTheme(name, mode string, style map[string]any) map[string]any {
	return map[string]any{
		"name":       name,
		"appearance": mode,
		"style":      style,
	}
}

func writeJSON(path string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func Generate(cfg *palette.Config, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("creating directory %s: %w", outputDir, err)
	}

	for _, mode := range []string{"dark", "light"} {
		p := cfg.Palettes[mode]
		modeLabel := palette.Pick(mode, "Dark", "Light")

		if err := writeJSON(filepath.Join(outputDir, mode+"-plain.json"), buildThemeFamily(
			"Monochromator "+modeLabel+" Plain",
			buildTheme(
				"Monochromator "+modeLabel+" Plain",
				mode,
				buildStyle(mode, p, plainSyntax(p)),
			),
		)); err != nil {
			return fmt.Errorf("writing %s-plain.json: %w", mode, err)
		}

		for _, v := range cfg.Variants {
			accent, err := p.Color(v.Color)
			if err != nil {
				return fmt.Errorf("resolving accent for variant %q: %w", v.Name, err)
			}

			var themeName string
			if v.Name == "default" {
				themeName = "Monochromator " + modeLabel
			} else {
				label := strings.ToUpper(v.Name[:1]) + v.Name[1:]
				themeName = "Monochromator " + modeLabel + " " + label
			}

			if err := writeJSON(filepath.Join(outputDir, mode+"-"+v.Name+".json"), buildThemeFamily(
				themeName,
				buildTheme(
					themeName,
					mode,
					buildStyle(mode, p, defaultSyntax(p, accent)),
				),
			)); err != nil {
				return fmt.Errorf("writing %s-%s.json: %w", mode, v.Name, err)
			}
		}
	}

	return nil
}
