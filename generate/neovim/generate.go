package neovim

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"generate/palette"
)

const (
	Accent          = "Accent"
	AccentBold      = "AccentBold"
	ContrastBold    = "ContrastBold"
	Muted           = "Muted"
	MutedItalic     = "MutedItalic"
	MutedBoldItalic = "MutedBoldItalic"
	Plain           = "Plain"
	PlainItalic     = "PlainItalic"
	Underlined      = "Underlined"
)

func parseHex(color, component string) int64 {
	v, err := strconv.ParseInt(component, 16, 0)
	if err != nil {
		panic(fmt.Sprintf("invalid hex color %q: %v", color, err))
	}
	return v
}

func blendColor(fg, bg string, alpha int) string {
	fgR := parseHex(fg, fg[1:3])
	fgG := parseHex(fg, fg[3:5])
	fgB := parseHex(fg, fg[5:7])
	bgR := parseHex(bg, bg[1:3])
	bgG := parseHex(bg, bg[3:5])
	bgB := parseHex(bg, bg[5:7])
	a := float64(alpha) / 255.0
	r := int(float64(fgR)*a + float64(bgR)*(1-a))
	g := int(float64(fgG)*a + float64(bgG)*(1-a))
	b := int(float64(fgB)*a + float64(bgB)*(1-a))
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func hlLine(name, fg, bg, style string) string {
	parts := []string{fmt.Sprintf("vim.api.nvim_set_hl(0, %q, {", name)}
	opts := []string{}
	if fg != "" {
		opts = append(opts, fmt.Sprintf("fg = %q", fg))
	}
	if bg != "" {
		opts = append(opts, fmt.Sprintf("bg = %q", bg))
	}
	if strings.Contains(style, "bold") {
		opts = append(opts, "bold = true")
	}
	if strings.Contains(style, "italic") {
		opts = append(opts, "italic = true")
	}
	if strings.Contains(style, "underline") {
		opts = append(opts, "underline = true")
	}
	if strings.Contains(style, "strikethrough") {
		opts = append(opts, "strikethrough = true")
	}
	parts = append(parts, strings.Join(opts, ", "))
	parts = append(parts, "})")
	return strings.Join(parts, " ")
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

func buildStyleTypes(p palette.Palette, accent string) []string {
	types := []string{
		hlLine(ContrastBold, p.FGContrast, "", "bold"),
		hlLine(Muted, p.FGMuted, "", ""),
		hlLine(MutedItalic, p.FGMuted, "", "italic"),
		hlLine(MutedBoldItalic, p.FGMuted, "", "bold italic"),
		hlLine(Plain, p.FG, "", ""),
		hlLine(PlainItalic, p.FG, "", "italic"),
		hlLine(Underlined, "", "", "underline"),
	}
	if accent != "" {
		types = append(types,
			hlLine(Accent, accent, "", ""),
			hlLine(AccentBold, accent, "", "bold"),
		)
	}
	return types
}

func buildPlainSyntax() []string {
	return []string{
		linkLine("Comment", MutedItalic),
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
		linkLine("Bold", ContrastBold),
		linkLine("Boolean", AccentBold),
		linkLine("Character", Accent),
		linkLine("Comment", MutedItalic),
		linkLine("Constant", Plain),
		linkLine("Delimiter", Muted),
		linkLine("Float", Accent),
		linkLine("Function", PlainItalic),
		linkLine("Identifier", Plain),
		linkLine("Italic", PlainItalic),
		linkLine("Keyword", ContrastBold),
		linkLine("Number", Accent),
		linkLine("Operator", Muted),
		linkLine("PreProc", ContrastBold),
		linkLine("Special", AccentBold),
		linkLine("SpecialChar", AccentBold),
		linkLine("SpecialComment", MutedBoldItalic),
		linkLine("Statement", ContrastBold),
		linkLine("String", Accent),
		linkLine("Tag", ContrastBold),
		linkLine("Type", ContrastBold),

		// Treesitter
		linkLine("@boolean", AccentBold),
		linkLine("@character", Accent),
		linkLine("@character.printf", AccentBold),
		linkLine("@character.special.java", ContrastBold),
		linkLine("@comment", MutedItalic),
		linkLine("@constant", Plain),
		linkLine("@constant.builtin", AccentBold),
		linkLine("@constant.builtin.rust", ContrastBold),
		linkLine("@constructor", ContrastBold),
		linkLine("@float", Accent),
		linkLine("@function", PlainItalic),
		linkLine("@function.builtin", PlainItalic),
		linkLine("@keyword", ContrastBold),
		linkLine("@keyword.javadoc", MutedBoldItalic),
		linkLine("@lsp.type.namespace", Plain),
		linkLine("@lsp.typemod.keyword.documentation", MutedBoldItalic),
		linkLine("@markup.math", Plain),
		linkLine("@markup.raw", Plain),
		linkLine("@module", Plain),
		linkLine("@number", Accent),
		linkLine("@operator", Muted),
		linkLine("@punctuation", Muted),
		linkLine("@punctuation.special", MutedItalic),
		linkLine("@string", Accent),
		linkLine("@string.documentation.python", Muted),
		linkLine("@string.escape", AccentBold),
		linkLine("@tag", ContrastBold),
		linkLine("@tag.attribute", Plain),
		linkLine("@tag.delimiter", Muted),
		linkLine("@type", ContrastBold),
		linkLine("@type.builtin", ContrastBold),
		linkLine("@variable", Plain),
		linkLine("@variable.builtin", ContrastBold),
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
		modeLabel := palette.Pick(mode, "dark", "light")
		editorHLs := buildEditorHighlights(p)

		plainName := "monochromator-" + modeLabel + "-plain"
		plainSyntax := append(buildStyleTypes(p, ""), buildPlainSyntax()...)
		if err := os.WriteFile(filepath.Join(outputDir, plainName+".lua"), []byte(buildThemeLua(plainName, mode, editorHLs, plainSyntax)), 0o644); err != nil {
			return fmt.Errorf("writing %s.lua: %w", plainName, err)
		}

		for _, v := range cfg.Variants {
			accent, err := p.Color(v.Color)
			if err != nil {
				return fmt.Errorf("resolving accent for variant %q: %w", v.Name, err)
			}

			var themeName string
			if v.Name == "default" {
				themeName = "monochromator-" + modeLabel
			} else {
				themeName = "monochromator-" + modeLabel + "-" + v.Name
			}

			syntax := append(buildStyleTypes(p, accent), buildDefaultSyntax()...)
			if err := os.WriteFile(filepath.Join(outputDir, themeName+".lua"), []byte(buildThemeLua(themeName, mode, editorHLs, syntax)), 0o644); err != nil {
				return fmt.Errorf("writing %s.lua: %w", themeName, err)
			}
		}
	}

	return nil
}
