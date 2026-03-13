package neovim

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"generate/palette"
)

const (
	accent          = "Accent"
	accentBold      = "AccentBold"
	contrastBold    = "ContrastBold"
	muted           = "Muted"
	mutedItalic     = "MutedItalic"
	mutedBoldItalic = "MutedBoldItalic"
	plain           = "Plain"
	plainItalic     = "PlainItalic"
	underlined      = "Underlined"
)

func parseHex(color, component string) int64 {
	v, err := strconv.ParseInt(component, 16, 0)
	if err != nil {
		panic(fmt.Sprintf("invalid hex color %q: %v", color, err))
	}
	return v
}

func blendColor(fg, bg string, alpha int) string {
	a := float64(alpha) / 255.0
	var c [3]int
	for i := range 3 {
		off := 1 + i*2
		f := parseHex(fg, fg[off:off+2])
		b := parseHex(bg, bg[off:off+2])
		c[i] = int(float64(f)*a + float64(b)*(1-a))
	}
	return fmt.Sprintf("#%02x%02x%02x", c[0], c[1], c[2])
}

func hlLine(name, fg, bg, style string) string {
	var opts []string
	if fg != "" {
		opts = append(opts, fmt.Sprintf("fg = %q", fg))
	}
	if bg != "" {
		opts = append(opts, fmt.Sprintf("bg = %q", bg))
	}
	for _, s := range []string{"bold", "italic", "underline", "strikethrough"} {
		if strings.Contains(style, s) {
			opts = append(opts, s+" = true")
		}
	}
	return fmt.Sprintf("vim.api.nvim_set_hl(0, %q, { %s })", name, strings.Join(opts, ", "))
}

func linkLine(from, to string) string {
	return fmt.Sprintf("vim.api.nvim_set_hl(0, %q, { link = %q })", from, to)
}

func buildEditorHighlights(p palette.Palette) []string {
	return []string{
		hlLine("Cursor", p.BG, p.FG, ""),
		hlLine("CursorLine", "", p.BGSecondary, ""),
		hlLine("CursorLineNr", p.FG, "", ""),
		hlLine("DiagnosticError", p.Red, "", ""),
		hlLine("DiagnosticInfo", p.Blue, "", ""),
		hlLine("DiagnosticWarn", p.Orange, "", ""),
		hlLine("DiffAdd", "", blendColor(p.Green, p.BG, 0x26), ""),
		hlLine("DiffDelete", "", blendColor(p.Red, p.BG, 0x26), ""),
		hlLine("FloatBorder", p.Border, p.BGSecondary, ""),
		hlLine("IblIndent", blendColor(p.FGMuted, p.BG, 0x50), "", ""),
		hlLine("IblScope", p.FGMuted, "", ""),
		hlLine("LineNr", p.FGMuted, "", ""),
		hlLine("Normal", p.FG, p.BG, ""),
		hlLine("NormalFloat", p.FG, p.BGSecondary, ""),
		hlLine("NvimTreeCursorLine", "", p.BGSelection, ""),
		hlLine("NvimTreeEndOfBuffer", "", p.BGSecondary, ""),
		hlLine("NvimTreeNormal", p.FG, p.BGSecondary, ""),
		hlLine("Pmenu", p.FG, p.BGSecondary, ""),
		hlLine("PmenuSel", p.FG, p.BGSelection, ""),
		hlLine("Search", "", p.BGSelection, ""),
		hlLine("StatusLine", p.FGSecondary, p.BGSecondary, ""),
		hlLine("TabLineFill", "", p.BGSecondary, ""),
		hlLine("Visual", "", p.BGSelection, ""),
		hlLine("WinSeparator", p.BGSecondary, "", ""),
	}
}

func buildStyleTypes(p palette.Palette, accentColor string) []string {
	types := []string{
		hlLine(contrastBold, p.FGContrast, "", "bold"),
		hlLine(muted, p.FGMuted, "", ""),
		hlLine(mutedItalic, p.FGMuted, "", "italic"),
		hlLine(mutedBoldItalic, p.FGMuted, "", "bold italic"),
		hlLine(plain, p.FG, "", ""),
		hlLine(plainItalic, p.FG, "", "italic"),
		hlLine(underlined, "", "", "underline"),
	}
	if accentColor != "" {
		types = append(types,
			hlLine(accent, accentColor, "", ""),
			hlLine(accentBold, accentColor, "", "bold"),
		)
	}
	return types
}

func buildPlainSyntax() []string {
	return []string{
		linkLine("Comment", mutedItalic),
		linkLine("Constant", "Normal"),
		linkLine("Delimiter", "Normal"),
		linkLine("Function", "Normal"),
		linkLine("Identifier", "Normal"),
		linkLine("Keyword", "Normal"),
		linkLine("Number", "Normal"),
		linkLine("Operator", "Normal"),
		linkLine("PreProc", "Normal"),
		linkLine("Special", "Normal"),
		linkLine("Statement", "Normal"),
		linkLine("String", "Normal"),
		linkLine("Type", "Normal"),
	}
}

func buildDefaultSyntax() []string {
	return []string{
		// Semantic groups
		linkLine("Bold", contrastBold),
		linkLine("Boolean", accentBold),
		linkLine("Character", accent),
		linkLine("Comment", mutedItalic),
		linkLine("Constant", plain),
		linkLine("Delimiter", muted),
		linkLine("Float", accent),
		linkLine("Function", plainItalic),
		linkLine("Identifier", plain),
		linkLine("Italic", plainItalic),
		linkLine("Keyword", contrastBold),
		linkLine("Number", accent),
		linkLine("Operator", muted),
		linkLine("PreProc", contrastBold),
		linkLine("Special", accentBold),
		linkLine("SpecialChar", accentBold),
		linkLine("SpecialComment", mutedBoldItalic),
		linkLine("Statement", contrastBold),
		linkLine("String", accent),
		linkLine("Tag", contrastBold),
		linkLine("Type", contrastBold),

		// Treesitter
		linkLine("@boolean", accentBold),
		linkLine("@character", accent),
		linkLine("@character.printf", accentBold),
		linkLine("@character.special.java", contrastBold),
		linkLine("@comment", mutedItalic),
		linkLine("@constant", plain),
		linkLine("@constant.builtin", accentBold),
		linkLine("@constant.builtin.rust", contrastBold),
		linkLine("@constructor", contrastBold),
		linkLine("@float", accent),
		linkLine("@function", plainItalic),
		linkLine("@function.builtin", plainItalic),
		linkLine("@keyword", contrastBold),
		linkLine("@keyword.javadoc", mutedBoldItalic),
		linkLine("@lsp.type.namespace", plain),
		linkLine("@lsp.typemod.keyword.documentation", mutedBoldItalic),
		linkLine("@markup.math", plain),
		linkLine("@markup.raw", plain),
		linkLine("@module", plain),
		linkLine("@number", accent),
		linkLine("@operator", muted),
		linkLine("@punctuation", muted),
		linkLine("@punctuation.special", mutedItalic),
		linkLine("@string", accent),
		linkLine("@string.documentation.python", muted),
		linkLine("@string.escape", accentBold),
		linkLine("@tag", contrastBold),
		linkLine("@tag.attribute", plain),
		linkLine("@tag.delimiter", muted),
		linkLine("@type", contrastBold),
		linkLine("@type.builtin", contrastBold),
		linkLine("@variable", plain),
		linkLine("@variable.builtin", contrastBold),
	}
}

func buildThemeLua(name, mode string, editorHLs, syntaxHLs []string) string {
	var b strings.Builder
	b.WriteString("vim.cmd(\"highlight clear\")\n")
	fmt.Fprintf(&b, "vim.g.colors_name = %q\n", name)
	fmt.Fprintf(&b, "vim.o.background = %q\n\n", mode)

	for _, line := range editorHLs {
		b.WriteString(line + "\n")
	}
	b.WriteString("\n")
	for _, line := range syntaxHLs {
		b.WriteString(line + "\n")
	}

	return b.String()
}

func Generate(cfg *palette.Config, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("creating directory %s: %w", outputDir, err)
	}

	for _, mode := range []string{"dark", "light"} {
		p := cfg.Palettes[mode]
		editorHLs := buildEditorHighlights(p)

		writeTheme := func(name string, syntaxHLs []string) error {
			data := buildThemeLua(name, mode, editorHLs, syntaxHLs)
			return os.WriteFile(filepath.Join(outputDir, name+".lua"), []byte(data), 0o644)
		}

		plainName := "monochromator-" + mode + "-plain"
		if err := writeTheme(plainName, slices.Concat(buildStyleTypes(p, ""), buildPlainSyntax())); err != nil {
			return fmt.Errorf("writing %s.lua: %w", plainName, err)
		}

		for _, v := range cfg.Variants {
			accentColor, err := p.Color(v.Color)
			if err != nil {
				return fmt.Errorf("resolving accent for variant %q: %w", v.Name, err)
			}

			themeName := "monochromator-" + mode
			if v.Name != "default" {
				themeName += "-" + v.Name
			}

			if err := writeTheme(themeName, slices.Concat(buildStyleTypes(p, accentColor), buildDefaultSyntax())); err != nil {
				return fmt.Errorf("writing %s.lua: %w", themeName, err)
			}
		}
	}

	return nil
}
