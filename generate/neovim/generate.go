package neovim

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"generate/palette"
)

func blendColor(fg, bg string, alpha int) string {
	fgR, _ := strconv.ParseInt(fg[1:3], 16, 0)
	fgG, _ := strconv.ParseInt(fg[3:5], 16, 0)
	fgB, _ := strconv.ParseInt(fg[5:7], 16, 0)
	bgR, _ := strconv.ParseInt(bg[1:3], 16, 0)
	bgG, _ := strconv.ParseInt(bg[3:5], 16, 0)
	bgB, _ := strconv.ParseInt(bg[5:7], 16, 0)
	a := float64(alpha) / 255.0
	r := int(float64(fgR)*a + float64(bgR)*(1-a))
	g := int(float64(fgG)*a + float64(bgG)*(1-a))
	b := int(float64(fgB)*a + float64(bgB)*(1-a))
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func luaColor(hex string) string {
	return fmt.Sprintf("%q", hex)
}

func hlLine(name, fg, bg, style string) string {
	parts := []string{fmt.Sprintf("vim.api.nvim_set_hl(0, %q, {", name)}
	opts := []string{}
	if fg != "" {
		opts = append(opts, fmt.Sprintf("fg = %s", luaColor(fg)))
	}
	if bg != "" {
		opts = append(opts, fmt.Sprintf("bg = %s", luaColor(bg)))
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
		hlLine("Pmenu", p.FG, p.BGSecondary, ""),
		hlLine("PmenuSel", p.FG, p.BGSelection, ""),
		hlLine("Search", "", p.BGSelection, ""),
		hlLine("StatusLine", p.FGSecondary, p.BGSecondary, ""),
		hlLine("TabLineFill", "", p.BGSecondary, ""),
		hlLine("Visual", "", p.BGSelection, ""),
		hlLine("WinSeparator", p.BGSecondary, "", ""),
	}
}

func buildPlainSyntax(p palette.Palette) []string {
	return []string{
		hlLine("Comment", p.FGMuted, "", "italic"),

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

func buildDefaultSyntax(p palette.Palette, accent string) []string {
	return []string{
		hlLine("Bold", p.FGContrast, "", "bold"),
		hlLine("Boolean", accent, "", "bold"),
		hlLine("Character", accent, "", ""),
		hlLine("Comment", p.FGMuted, "", "italic"),
		hlLine("Constant", p.FG, "", ""),
		hlLine("Delimiter", p.FGMuted, "", ""),
		hlLine("Float", accent, "", ""),
		hlLine("Function", p.FG, "", "italic"),
		hlLine("Identifier", p.FG, "", ""),
		hlLine("Italic", p.FG, "", "italic"),
		hlLine("Keyword", p.FGContrast, "", "bold"),
		hlLine("Number", accent, "", ""),
		hlLine("Operator", p.FGMuted, "", ""),
		hlLine("PreProc", p.FGContrast, "", "bold"),
		hlLine("Special", accent, "", "bold"),
		hlLine("SpecialChar", accent, "", "bold"),
		hlLine("Statement", p.FGContrast, "", "bold"),
		hlLine("String", accent, "", ""),
		hlLine("Tag", p.FGContrast, "", "bold"),
		hlLine("Type", p.FGContrast, "", "bold"),
		hlLine("Underlined", "", "", "underline"),

		linkLine("@boolean", "Boolean"),
		linkLine("@character", "Character"),
		linkLine("@comment", "Comment"),
		linkLine("@constant", "Constant"),
		linkLine("@constant.builtin", "Boolean"),
		linkLine("@constructor", "Type"),
		linkLine("@float", "Float"),
		linkLine("@function", "Function"),
		linkLine("@function.builtin", "Function"),
		linkLine("@keyword", "Keyword"),
		linkLine("@number", "Number"),
		linkLine("@operator", "Operator"),
		linkLine("@punctuation", "Delimiter"),
		linkLine("@string", "String"),
		linkLine("@string.escape", "SpecialChar"),
		linkLine("@tag", "Tag"),
		linkLine("@type", "Type"),
		linkLine("@type.builtin", "Type"),
		linkLine("@variable", "Identifier"),
	}
}

func buildThemeLua(name, mode string, editorHLs, syntaxHLs []string) string {
	var b strings.Builder
	b.WriteString("vim.cmd(\"highlight clear\")\n")
	b.WriteString(fmt.Sprintf("vim.g.colors_name = %q\n", name))
	b.WriteString(fmt.Sprintf("vim.o.background = %q\n\n", mode))

	for _, line := range editorHLs {
		b.WriteString(line + "\n")
	}
	b.WriteString("\n")
	for _, line := range syntaxHLs {
		b.WriteString(line + "\n")
	}

	return b.String()
}

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0o644)
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
		plainLua := buildThemeLua(plainName, mode, editorHLs, buildPlainSyntax(p))
		if err := writeFile(filepath.Join(outputDir, plainName+".lua"), plainLua); err != nil {
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

			lua := buildThemeLua(themeName, mode, editorHLs, buildDefaultSyntax(p, accent))
			if err := writeFile(filepath.Join(outputDir, themeName+".lua"), lua); err != nil {
				return fmt.Errorf("writing %s.lua: %w", themeName, err)
			}
		}
	}

	return nil
}
