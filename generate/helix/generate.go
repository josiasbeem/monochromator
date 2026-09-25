// Package helix generates the Helix theme files from the palette.
package helix

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"generate/palette"
)

type color string

const (
	accent      color = "accent"
	bg          color = "bg"
	bgSecondary color = "bg_secondary"
	bgSelection color = "bg_selection"
	blue        color = "blue"
	border      color = "border"
	fgFaint     color = "fg_faint"
	fg          color = "fg"
	fgContrast  color = "fg_contrast"
	fgMuted     color = "fg_muted"
	green       color = "green"
	inlayHint   color = "inlay_hint"
	orange      color = "orange"
	red         color = "red"
	selection   color = "selection"
)

type modifier string

const (
	bold       modifier = "bold"
	crossedOut modifier = "crossed_out"
	italic     modifier = "italic"
)

type underlineStyle string

const (
	curl underlineStyle = "curl"
	line underlineStyle = "line"
)

type style struct {
	fg             color
	bg             color
	underline      underlineStyle
	underlineColor color
	modifiers      []modifier
}

func plainSyntax() map[string]style {
	return map[string]style{
		"comment": {fg: fgMuted, modifiers: []modifier{italic}},
		// Helix gives unused variables such as "_" this scope. Without
		// this entry, they get the comment style.
		"comment.unused": {fg: fg},
	}
}

// "punctuation.special" stays undefined. TypeScript uses it for full class
// fields, and Helix gives its modifiers to all tokens inside.
func defaultSyntax() map[string]style {
	syntax := plainSyntax()
	maps.Copy(syntax, map[string]style{
		"attribute":                 {fg: fgMuted},
		"constant":                  {fg: fg},
		"constant.builtin":          {fg: accent},
		"constant.character":        {fg: accent},
		"constant.character.escape": {fg: accent, underline: line},
		"constant.numeric":          {fg: accent},
		"constructor":               {fg: fgContrast, modifiers: []modifier{italic}},
		"function":                  {fg: fgContrast, modifiers: []modifier{italic}},
		"keyword":                   {fg: fgContrast, modifiers: []modifier{bold}},
		"keyword.operator":          {fg: fgMuted},
		// Rust gives the "&" in reference types this scope. It is an
		// operator, so it is not bold.
		"keyword.storage.modifier.ref": {fg: fgMuted},
		"label":                        {fg: fgContrast, modifiers: []modifier{bold}},
		"markup.bold":                  {fg: fgContrast, modifiers: []modifier{bold}},
		"markup.heading":               {fg: fgContrast, modifiers: []modifier{bold}},
		"markup.heading.marker":        {fg: fgMuted},
		"markup.italic":                {fg: fg, modifiers: []modifier{italic}},
		"markup.link.url":              {underline: line},
		"markup.list":                  {fg: fgMuted},
		"markup.quote":                 {fg: fgMuted, modifiers: []modifier{italic}},
		"markup.raw":                   {fg: accent},
		// Code blocks show other languages, so the base is the text color.
		"markup.raw.block":      {fg: fg},
		"markup.strikethrough":  {modifiers: []modifier{crossedOut}},
		"namespace":             {fg: fg},
		"operator":              {fg: fgMuted},
		"punctuation":           {fg: fgMuted},
		"string":                {fg: accent},
		"string.special.symbol": {fg: fg},
		"tag":                   {fg: fgContrast, modifiers: []modifier{bold}},
		"type":                  {fg: fgContrast, underline: line},
		"variable":              {fg: fg},
		"variable.builtin":      {fg: fgMuted},
	})
	return syntax
}

// "special" stays undefined, because the pickers also use it for the
// matched characters. The fg of "ui.background" and "ui.popup" sets the
// color of the picker and popup borders.
func buildUI(jumpLabel color) map[string]style {
	return map[string]style{
		"diagnostic.deprecated":    {modifiers: []modifier{crossedOut}},
		"diagnostic.error":         {underline: curl, underlineColor: red},
		"diagnostic.hint":          {underline: curl, underlineColor: fgMuted},
		"diagnostic.info":          {underline: curl, underlineColor: blue},
		"diagnostic.warning":       {underline: curl, underlineColor: orange},
		"diff.delta":               {fg: orange},
		"diff.delta.moved":         {fg: green},
		"diff.minus":               {fg: red},
		"diff.plus":                {fg: green},
		"error":                    {fg: red},
		"hint":                     {fg: fgMuted},
		"info":                     {fg: blue},
		"ui.background":            {fg: border, bg: bg},
		"ui.bufferline":            {fg: fgMuted, bg: bgSecondary},
		"ui.bufferline.active":     {fg: fg, bg: bg},
		"ui.bufferline.background": {bg: bgSecondary},
		"ui.cursor":                {fg: bg, bg: fgMuted},
		"ui.cursor.match":          {bg: bgSelection},
		"ui.cursor.primary":        {fg: bg, bg: fg},
		"ui.cursorline.primary":    {bg: bgSecondary},
		"ui.help":                  {fg: fg, bg: bgSecondary},
		"ui.highlight":             {bg: bgSecondary},
		"ui.linenr":                {fg: fgMuted},
		"ui.linenr.selected":       {fg: fg},
		"ui.menu":                  {fg: fg, bg: bgSecondary},
		"ui.menu.scroll":           {fg: bgSelection, bg: bgSecondary},
		"ui.menu.selected":         {fg: fg, bg: bgSelection},
		"ui.popup":                 {fg: border, bg: bgSecondary},
		"ui.popup.info":            {fg: fg, bg: bgSecondary},
		"ui.selection":             {bg: selection},
		"ui.statusline":            {fg: fg, bg: bgSecondary},
		"ui.statusline.inactive":   {fg: fgMuted, bg: bgSecondary},
		"ui.text":                  {fg: fg},
		"ui.text.focus":            {fg: fg, bg: bgSelection},
		"ui.text.inactive":         {fg: fgMuted},
		"ui.virtual":               {fg: fgFaint},
		"ui.virtual.inlay-hint":    {fg: fgMuted, bg: inlayHint},
		"ui.virtual.jump-label":    {fg: bg, bg: jumpLabel, modifiers: []modifier{bold}},
		"ui.virtual.ruler":         {bg: bgSecondary},
		"ui.window":                {fg: border},
		"warning":                  {fg: orange},
	}
}

// Helix has no alpha channel. The generator mixes the transparent VS Code
// colors with the background. 0x4d is about 30%.
func buildPalette(mode string, p palette.Palette, accentHex string) (map[color]string, error) {
	colors := map[color]string{
		bg:          p.BG,
		bgSecondary: p.BGSecondary,
		bgSelection: p.BGSelection,
		blue:        p.Blue,
		border:      p.Border,
		fg:          p.FG,
		fgContrast:  p.FGContrast,
		fgMuted:     p.FGMuted,
		green:       p.Green,
		inlayHint:   p.BGSecondary,
		orange:      p.Orange,
		red:         p.Red,
		selection:   p.BGSelection,
	}

	faintHex, err := blend(p.FGMuted, p.BG, 0x50)
	if err != nil {
		return nil, fmt.Errorf("mixing faint color: %w", err)
	}
	colors[fgFaint] = faintHex

	if mode == "light" {
		inlayHex, err := blend(p.BGSecondary, p.BG, 0x80)
		if err != nil {
			return nil, fmt.Errorf("mixing inlay hint color: %w", err)
		}
		colors[inlayHint] = inlayHex
	}

	if accentHex != "" {
		selectionHex, err := blend(accentHex, p.BG, 0x4d)
		if err != nil {
			return nil, fmt.Errorf("mixing selection color: %w", err)
		}
		colors[accent] = accentHex
		colors[selection] = selectionHex
	}

	return colors, nil
}

type rgb struct{ r, g, b uint8 }

func parseHex(hex string) (rgb, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return rgb{}, fmt.Errorf("color %q does not have the form #rrggbb", hex)
	}
	v, err := strconv.ParseUint(hex[1:], 16, 32)
	if err != nil {
		return rgb{}, fmt.Errorf("color %q: %w", hex, err)
	}
	return rgb{r: uint8(v >> 16), g: uint8(v >> 8), b: uint8(v)}, nil
}

func blend(fore, back string, alpha uint8) (string, error) {
	f, err := parseHex(fore)
	if err != nil {
		return "", err
	}
	b, err := parseHex(back)
	if err != nil {
		return "", err
	}
	mix := func(f, b uint8) uint8 {
		return uint8((uint(f)*uint(alpha) + uint(b)*(0xff-uint(alpha)) + 0x7f) / 0xff)
	}
	return fmt.Sprintf("#%02x%02x%02x", mix(f.r, b.r), mix(f.g, b.g), mix(f.b, b.b)), nil
}

func (s style) toml() string {
	var fields []string
	if s.fg != "" {
		fields = append(fields, fmt.Sprintf("fg = %q", s.fg))
	}
	if s.bg != "" {
		fields = append(fields, fmt.Sprintf("bg = %q", s.bg))
	}
	if s.underline != "" {
		underline := fmt.Sprintf("style = %q", s.underline)
		if s.underlineColor != "" {
			underline = fmt.Sprintf("color = %q, %s", s.underlineColor, underline)
		}
		fields = append(fields, "underline = { "+underline+" }")
	}
	if len(s.modifiers) > 0 {
		quoted := make([]string, len(s.modifiers))
		for i, m := range s.modifiers {
			quoted[i] = strconv.Quote(string(m))
		}
		fields = append(fields, "modifiers = ["+strings.Join(quoted, ", ")+"]")
	}
	return "{ " + strings.Join(fields, ", ") + " }"
}

// The [palette] table must be the last table, because a TOML table
// contains all keys after its header.
func renderTheme(styles map[string]style, colors map[color]string) (string, error) {
	for c, hex := range colors {
		if _, err := parseHex(hex); err != nil {
			return "", fmt.Errorf("palette entry %q: %w", c, err)
		}
	}

	var b strings.Builder
	for _, scope := range slices.Sorted(maps.Keys(styles)) {
		s := styles[scope]
		for _, c := range []color{s.fg, s.bg, s.underlineColor} {
			if _, ok := colors[c]; c != "" && !ok {
				return "", fmt.Errorf("scope %q: color %q is not in the palette", scope, c)
			}
		}
		fmt.Fprintf(&b, "%q = %s\n", scope, s.toml())
	}

	b.WriteString("\n[palette]\n")
	for _, c := range slices.Sorted(maps.Keys(colors)) {
		fmt.Fprintf(&b, "%s = %q\n", c, colors[c])
	}
	return b.String(), nil
}

func writeTheme(path string, styles map[string]style, colors map[color]string) error {
	text, err := renderTheme(styles, colors)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(text), 0o644)
}

// Generate writes the Helix theme files for both modes into outputDir:
// one plain theme and one theme for each variant. Helix uses the file
// name as the theme name, for example "monochromator-dark-amber".
func Generate(cfg *palette.Config, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("creating directory %s: %w", outputDir, err)
	}

	for _, mode := range []string{"dark", "light"} {
		p := cfg.Palettes[mode]

		colors, err := buildPalette(mode, p, "")
		if err != nil {
			return fmt.Errorf("building %s plain palette: %w", mode, err)
		}
		styles := buildUI(fgContrast)
		maps.Copy(styles, plainSyntax())
		plainPath := filepath.Join(outputDir, "monochromator-"+mode+"-plain.toml")
		if err := writeTheme(plainPath, styles, colors); err != nil {
			return fmt.Errorf("writing %s: %w", plainPath, err)
		}

		for _, v := range cfg.Variants {
			accentHex, err := p.Color(v.Color)
			if err != nil {
				return fmt.Errorf("resolving accent for variant %q: %w", v.Name, err)
			}

			fileName := "monochromator-" + mode
			if v.Name != "default" {
				fileName += "-" + v.Name
			}

			colors, err := buildPalette(mode, p, accentHex)
			if err != nil {
				return fmt.Errorf("building %s palette: %w", fileName, err)
			}
			styles := buildUI(accent)
			maps.Copy(styles, defaultSyntax())
			path := filepath.Join(outputDir, fileName+".toml")
			if err := writeTheme(path, styles, colors); err != nil {
				return fmt.Errorf("writing %s: %w", path, err)
			}
		}
	}

	return nil
}
