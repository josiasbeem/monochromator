// Package vscode generates the VS Code theme files from the palette.
package vscode

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"generate/palette"
)

// The scope* lists group TextMate scopes by style. The list name gives
// the font style and the color group: Gray = FGMuted, Text = FG or
// FGContrast, Literals = the accent color. Keep the lists sorted.
var scopeGray = []string{
	"constant.character.escape.content.typst",
	"constant.other.symbol.typst",
	"entity.name.function.hash.typst",
	"entity.name.operator.cpp",
	"keyword.control.equation.align.latex",
	"keyword.control.equation.newline.latex",
	"keyword.control.flow.block-scalar.literal.yaml",
	"keyword.control.newline.tex",
	"keyword.control.table.cell.latex",
	"keyword.control.table.newline.latex",
	"keyword.control.ternary",
	"keyword.operator",
	"keyword.other.r",
	"keyword.symbol",
	"markup.content.brace.typst",
	"markup.heading.setext",
	"meta.brace.curly.typst",
	"meta.brace.round.js",
	"meta.brace.round.ts",
	"meta.brace.round.typst",
	"meta.brace.square.js",
	"meta.brace.square.ts",
	"meta.brace.square.typst",
	"meta.entry.braces.bibtex",
	"meta.separator.markdown",
	"punctuation",
	"punctuation.definition.constant.begin.markdown",
	"punctuation.definition.constant.end.markdown",
	"punctuation.definition.constant.hashkey.ruby",
	"punctuation.definition.constant.markdown",
	"punctuation.definition.constant.math.tex",
	"punctuation.definition.constant.ruby",
	"punctuation.definition.entity.function.2.nix",
	"punctuation.definition.entity.function.nix",
	"punctuation.definition.heading.typst",
	"punctuation.definition.keyword.latex",
	"punctuation.definition.keyword.tex",
	"punctuation.definition.raw.begin.typst",
	"punctuation.definition.raw.end.typst",
	"punctuation.definition.string.begin.latex",
	"punctuation.definition.string.begin.math.typst",
	"punctuation.definition.string.begin.tex",
	"punctuation.definition.string.end.latex",
	"punctuation.definition.string.end.math.typst",
	"punctuation.definition.string.end.tex",
	"storage.modifier.array",
	"storage.modifier.pointer",
	"storage.modifier.reference",
	"storage.type.function.arrow",
	"variable.parameter.function.4.nix",
}

var scopeBoldItalicGray = []string{
	"keyword.other.documentation",
	"meta.shebang.shell",
	"punctuation.definition.block.tag.jsdoc",
	"storage.type.class.doxygen",
	"storage.type.class.jsdoc",
}

var scopeItalicGray = []string{
	"comment",
	"entity.name.type.instance.jsdoc",
	"markup.quote",
	"punctuation.definition.comment",
	"variable.other.jsdoc",
}

var scopeText = []string{
	"constant",
	"constant.language.register",
	"constant.language.symbol.hashkey.ruby",
	"constant.language.symbol.ruby",
	"constant.other.bareword",
	"constant.other.caps",
	"constant.other.class",
	"constant.other.database-name",
	"constant.other.enum",
	"constant.other.haskell",
	"constant.other.key",
	"constant.other.option",
	"constant.other.reference",
	"constant.other.table-name",
	"entity.name.function.preprocessor",
	"entity.name.section.latex",
	"entity.name.type.class",
	"entity.name.type.entry-key.bibtex",
	"entity.name.type.enum",
	"entity.name.type.package.go",
	"entity.name.type.struct.rust",
	"entity.name.type.trait.rust",
	"keyword.label.assembly",
	"markup.bold.textbf.latex",
	"markup.italic.emph.latex",
	"markup.italic.textit.latex",
	"meta.interpolation.rust",
	"punctuation.definition.variable",
	"storage.modifier.import",
	"storage.modifier.package",
	"string.other.link.description.markdown",
	"string.other.link.title.markdown",
	"string.unquoted.argument.shell",
	"string.unquoted.plain.out.yaml",
	"support.class.latex",
	"support.class.math.block.environment.latex",
	"support.class.math.block.tex",
	"support.type.property-name",
	"support.type.vendored.property-name",
	"variable.language",
	"variable.other",
	"variable.parameter.positional.shell",
	"variable.parameter.ts",
}

var scopeBoldText = []string{
	"constant.character.escape.markdown",
	"constant.character.escape.tex",
	"constant.character.math.tex",
	"constant.other.general.math.tex",
	"constant.other.math.tex",
	"entity.name.command.shell",
	"entity.name.function.decorator",
	"entity.name.function.instructions",
	"entity.name.function.math.tex",
	"entity.name.section",
	"entity.name.tag",
	"entity.name.type",
	"entity.other.attribute-name.class.css",
	"entity.other.attribute-name.id.css",
	"entity.other.attribute-name.pseudo-class",
	"keyword",
	"keyword.operator.new",
	"keyword.operator.word.mnemonic",
	"markup.bold",
	"markup.heading.typst",
	"punctuation.definition.annotation",
	"punctuation.definition.attribute",
	"punctuation.definition.decorator",
	"punctuation.definition.directive",
	"punctuation.definition.entity",
	"punctuation.definition.include",
	"punctuation.definition.keyword",
	"punctuation.definition.string.heredoc.delimiter.shell",
	"storage.modifier",
	"storage.modifier.sql",
	"storage.type",
	"support.class",
	"support.constant.subsections",
	"support.function.be.latex",
	"support.function.builtin.shell",
	"support.function.emph.latex",
	"support.function.footnote.latex",
	"support.function.general.tex",
	"support.function.marginpar.latex",
	"support.function.mips",
	"support.function.mnemonic",
	"support.function.pseudo.mips",
	"support.function.riscv",
	"support.function.section.latex",
	"support.function.textbf.latex",
	"support.function.textit.latex",
	"support.function.texttt.latex",
	"support.function.verb.latex",
	"support.type",
	"support.type.property-name.table.toml",
	"variable.language.self",
	"variable.language.special.self",
	"variable.language.super",
	"variable.language.this",
	"variable.other.readwrite.hash.typst",
	"variable.other.readwrite.typst",
}

var scopeItalicText = []string{
	"entity.name.function",
	"keyword.operator.sizeof",
	"support.function",
}

var scopeLiterals = []string{
	"constant.character",
	"constant.numeric",
	"constant.other",
	"keyword.other.unit",
	"markup.inline.raw",
	"meta.key-assignment.bibtex",
	"punctuation.definition.char",
	"punctuation.definition.constant",
	"punctuation.definition.raw",
	"punctuation.definition.reference.typst",
	"punctuation.definition.string",
	"punctuation.separator.dot.decimal",
	"string",
}

var scopeBoldLiterals = []string{
	"constant.character.escape",
	"constant.character.format.placeholder",
	"constant.language",
	"constant.other.placeholder",
	"keyword.constant",
	"keyword.format.specifier",
	"punctuation.definition.group.begin.bibtex",
	"punctuation.definition.group.end.bibtex",
	"punctuation.definition.interpolation",
	"punctuation.definition.template-expression",
	"punctuation.definition.variable.shell",
	"storage.type.string.python",
	"string.quoted.byte.raw",
}

func tokenRule(scope []string, foreground, fontStyle string) map[string]any {
	settings := map[string]any{"fontStyle": fontStyle}
	if foreground != "" {
		settings["foreground"] = foreground
	}
	return map[string]any{"scope": scope, "settings": settings}
}

// Color values with eight hex digits carry an alpha channel: the last
// two digits set the opacity ("80" is about 50%).

func buildPlainColors(p palette.Palette, mode string) map[string]string {
	return map[string]string{
		"activityBar.activeForeground":                        p.FGSecondary,
		"activityBar.background":                              p.BGSecondary,
		"activityBar.foreground":                              p.FG,
		"activityBar.inactiveForeground":                      p.FGSecondary,
		"activityBarBadge.background":                         p.FGSecondary,
		"activityBarBadge.foreground":                         p.BG,
		"badge.background":                                    "#00000000",
		"badge.foreground":                                    p.FG,
		"breadcrumb.foreground":                               p.FGMuted,
		"button.background":                                   p.BGElevated,
		"button.foreground":                                   palette.Pick(mode, p.FG, "#ffffff"),
		"button.secondaryBackground":                          p.BGElevated + "80",
		"button.secondaryForeground":                          p.FG,
		"checkbox.background":                                 p.BG,
		"commandCenter.activeBackground":                      p.BG,
		"commandCenter.activeBorder":                          p.Border + "80",
		"commandCenter.activeForeground":                      p.FG,
		"commandCenter.background":                            p.BG,
		"commandCenter.border":                                p.Border + "80",
		"commandCenter.foreground":                            p.FGSecondary,
		"diffEditor.insertedLineBackground":                   p.Green + "26",
		"diffEditor.insertedTextBackground":                   p.Green + "4d",
		"diffEditor.removedLineBackground":                    p.Red + "26",
		"diffEditor.removedTextBackground":                    p.Red + "4d",
		"dropdown.background":                                 p.BG,
		"dropdown.border":                                     p.Border,
		"editor.background":                                   p.BG,
		"editor.findMatchBackground":                          p.BGSelection,
		"editor.findMatchHighlightBackground":                 p.BGSelection + "bf",
		"editor.foldBackground":                               p.BGSecondary + "80",
		"editor.foreground":                                   p.FG,
		"editor.hoverHighlightBackground":                     p.BGSelection + "80",
		"editor.lineHighlightBackground":                      p.BGSecondary,
		"editor.selectionBackground":                          p.BGSelection,
		"editor.selectionHighlightBackground":                 p.BGSelection,
		"editor.wordHighlightBackground":                      p.BGSelection + "80",
		"editor.wordHighlightStrongBackground":                p.Blue + "40",
		"editor.wordHighlightTextBackground":                  p.BGSelection + "80",
		"editorBracketHighlight.foreground1":                  p.FG,
		"editorBracketHighlight.foreground2":                  p.FG,
		"editorBracketHighlight.foreground3":                  p.FG,
		"editorBracketHighlight.foreground4":                  p.FG,
		"editorBracketHighlight.foreground5":                  p.FG,
		"editorBracketHighlight.foreground6":                  p.FG,
		"editorBracketHighlight.unexpectedBracket.foreground": p.FG,
		"editorBracketMatch.background":                       p.BG,
		"editorBracketMatch.border":                           p.FG,
		"editorCodeLens.foreground":                           p.FGMuted,
		"editorCursor.foreground":                             p.FG,
		"editorError.foreground":                              p.Red,
		"editorGroup.border":                                  p.BGSecondary,
		"editorGroupHeader.tabsBackground":                    p.BGSecondary,
		"editorGutter.addedBackground":                        p.Green,
		"editorGutter.addedSecondaryBackground":               p.Green + "80",
		"editorGutter.deletedBackground":                      p.Red,
		"editorGutter.deletedSecondaryBackground":             p.Red + "80",
		"editorGutter.modifiedBackground":                     p.Orange,
		"editorGutter.modifiedSecondaryBackground":            p.Orange + "80",
		"editorHoverWidget.border":                            p.Border,
		"editorHoverWidget.foreground":                        p.FG,
		"editorIndentGuide.activeBackground":                  p.FGMuted,
		"editorIndentGuide.background":                        p.FGMuted + "50",
		"editorInfo.foreground":                               p.Blue,
		"editorInlayHint.background":                          palette.Pick(mode, p.BGSecondary, p.BGSecondary+"80"),
		"editorInlayHint.foreground":                          p.FGMuted,
		"editorLightBulb.foreground":                          p.FGMuted,
		"editorLightBulbAi.foreground":                        p.FGMuted,
		"editorLightBulbAutoFix.foreground":                   p.FGMuted,
		"editorLineNumber.activeForeground":                   p.FG,
		"editorLineNumber.foreground":                         p.FGMuted,
		"editorLink.activeForeground":                         p.FG,
		"editorOverviewRuler.border":                          "#00000000",
		"editorOverviewRuler.errorForeground":                 p.Red,
		"editorOverviewRuler.infoForeground":                  p.Blue,
		"editorOverviewRuler.warningForeground":               p.Orange,
		"editorStickyScrollHover.background":                  p.BGSecondary,
		"editorSuggestWidget.background":                      p.BGSecondary,
		"editorSuggestWidget.focusHighlightForeground":        p.FG,
		"editorSuggestWidget.highlightForeground":             p.FG,
		"editorSuggestWidget.selectionBackground":             p.BGSelection,
		"editorWarning.foreground":                            p.Orange,
		"editorWidget.background":                             p.BGSecondary,
		"extensionButton.background":                          p.BGElevated,
		"extensionButton.foreground":                          palette.Pick(mode, p.FG, "#ffffff"),
		"focusBorder":                                         "#00000000",
		"gitDecoration.addedResourceForeground":               p.Green,
		"gitDecoration.deletedResourceForeground":             p.Red,
		"gitDecoration.ignoredResourceForeground":             p.FGMuted,
		"gitDecoration.modifiedResourceForeground":            p.Orange,
		"gitDecoration.renamedResourceForeground":             p.Green,
		"gitDecoration.stageDeletedResourceForeground":        p.Red,
		"gitDecoration.stageModifiedResourceForeground":       p.Orange,
		"gitDecoration.untrackedResourceForeground":           p.Green,
		"icon.foreground":                                     p.FGSecondary,
		"input.background":                                    p.BG,
		"input.border":                                        p.Border,
		"inputValidation.errorBackground":                     p.BGSecondary,
		"inputValidation.errorBorder":                         p.Red,
		"inputValidation.infoBackground":                      p.BGSecondary,
		"inputValidation.infoBorder":                          p.Border,
		"inputValidation.warningBackground":                   p.BGSecondary,
		"inputValidation.warningBorder":                       p.Orange,
		"keybindingLabel.background":                          p.BG,
		"keybindingLabel.border":                              p.FGMuted,
		"keybindingLabel.bottomBorder":                        p.FGMuted,
		"keybindingLabel.foreground":                          p.FGSecondary,
		"list.activeSelectionBackground":                      p.BGSelection,
		"list.activeSelectionForeground":                      p.FG,
		"list.errorForeground":                                p.Red,
		"list.focusBackground":                                p.BGSelection,
		"list.highlightForeground":                            p.FG,
		"list.hoverBackground":                                p.BGSelection,
		"list.inactiveSelectionBackground":                    p.BGSelection + "80",
		"list.warningForeground":                              p.Orange,
		"menu.background":                                     p.BGSecondary,
		"menu.border":                                         p.Border,
		"menu.selectionBackground":                            p.BGSelection,
		"menu.selectionForeground":                            p.FGSecondary,
		"menu.separatorBackground":                            p.BGSelection,
		"minimap.errorHighlight":                              p.Red,
		"minimap.findMatchHighlight":                          p.BGSelection,
		"minimap.infoHighlight":                               p.Blue,
		"minimap.selectionHighlight":                          p.BGSelection,
		"minimap.selectionOccurrenceHighlight":                p.BGSelection,
		"minimap.warningHighlight":                            p.Orange,
		"multiDiffEditor.border":                              p.Border,
		"multiDiffEditor.headerBackground":                    p.BGSecondary,
		"notifications.background":                            p.BGSecondary,
		"notifications.foreground":                            p.FGSecondary,
		"notificationsErrorIcon.foreground":                   p.FGSecondary,
		"notificationsInfoIcon.foreground":                    p.FGSecondary,
		"notificationsWarningIcon.foreground":                 p.FGSecondary,
		"panel.border":                                        p.BGSecondary,
		"panelTitle.activeForeground":                         p.FG,
		"panelTitle.inactiveForeground":                       p.FGSecondary,
		"pickerGroup.border":                                  p.Border,
		"pickerGroup.foreground":                              p.FGSecondary,
		"progressBar.background":                              p.FGSecondary,
		"quickInput.background":                               p.BGSecondary,
		"quickInput.foreground":                               p.FGSecondary,
		"quickInputList.focusBackground":                      p.BGSelection,
		"quickInputList.focusForeground":                      p.FG,
		"scmGraph.historyItemHoverAdditionsForeground":        p.Green,
		"scmGraph.historyItemHoverDeletionsForeground":        p.Red,
		"scmGraph.historyItemHoverLabelForeground":            palette.Pick(mode, "#000000", "#ffffff"),
		"scmGraph.historyItemRefColor":                        p.FGSecondary,
		"scmGraph.historyItemRemoteRefColor":                  p.FGMuted,
		"scrollbarSlider.background":                          palette.Pick(mode, p.BGElevated+"43", p.BGSelection+"60"),
		"scrollbarSlider.hoverBackground":                     palette.Pick(mode, p.BGElevated+"70", p.FGMuted+"50"),
		"scrollbarSlider.activeBackground":                    palette.Pick(mode, p.BGElevated+"90", p.FGMuted+"70"),
		"selection.background":                                p.BGSelection,
		"sideBar.background":                                  p.BGSecondary,
		"sideBar.border":                                      p.BGSecondary,
		"sideBar.foreground":                                  p.FGSecondary,
		"sideBarSectionHeader.background":                     p.BGSecondary,
		"sideBarSectionHeader.border":                         p.BGSecondary,
		"statusBar.background":                                p.BGSecondary,
		"statusBar.border":                                    "#00000000",
		"statusBar.debuggingBackground":                       p.BGSecondary,
		"statusBar.foreground":                                p.FGSecondary,
		"statusBar.noFolderBackground":                        p.BGSecondary,
		"statusBarItem.errorBackground":                       p.Red,
		"statusBarItem.hoverBackground":                       p.BGSelection,
		"statusBarItem.prominentBackground":                   p.BGSelection,
		"statusBarItem.remoteBackground":                      p.BGSecondary,
		"statusBarItem.remoteForeground":                      p.FGSecondary,
		"statusBarItem.warningBackground":                     p.Orange,
		"tab.activeForeground":                                p.FGSecondary,
		"tab.border":                                          "#00000000",
		"tab.hoverBackground":                                 p.BG,
		"tab.inactiveBackground":                              p.BGSecondary,
		"terminal.ansiBlack":                                  palette.Pick(mode, p.BG, p.FG),
		"terminal.ansiBlue":                                   p.Blue,
		"terminal.ansiBrightBlack":                            p.FGMuted,
		"terminal.ansiBrightBlue":                             p.Blue,
		"terminal.ansiBrightCyan":                             p.Cyan,
		"terminal.ansiBrightGreen":                            p.Green,
		"terminal.ansiBrightMagenta":                          p.Magenta,
		"terminal.ansiBrightRed":                              p.Red,
		"terminal.ansiBrightWhite":                            palette.Pick(mode, p.FG, "#ffffff"),
		"terminal.ansiBrightYellow":                           p.Yellow,
		"terminal.ansiCyan":                                   p.Cyan,
		"terminal.ansiGreen":                                  p.Green,
		"terminal.ansiMagenta":                                p.Magenta,
		"terminal.ansiRed":                                    p.Red,
		"terminal.ansiWhite":                                  p.FGMuted,
		"terminal.ansiYellow":                                 p.Yellow,
		"terminal.background":                                 p.BG,
		"terminal.foreground":                                 p.FG,
		"terminal.selectionBackground":                        p.FG,
		"terminal.selectionForeground":                        p.BG,
		"terminalCommandDecoration.defaultBackground":         p.FGMuted,
		"terminalCommandDecoration.errorBackground":           p.Red,
		"terminalCommandDecoration.successBackground":         p.Green,
		"textBlockQuote.background":                           p.BGSecondary,
		"textBlockQuote.border":                               p.Border,
		"textCodeBlock.background":                            p.BGSecondary,
		"textCodeBlock.border":                                p.Border,
		"textLink.activeForeground":                           p.Blue,
		"textLink.foreground":                                 p.Blue,
		"textPreformat.background":                            p.BGSelection,
		"textPreformat.foreground":                            p.FG,
		"textSeparator.foreground":                            p.BGSelection,
		"titleBar.activeBackground":                           p.BGSecondary,
		"titleBar.activeForeground":                           p.FGSecondary,
		"titleBar.border":                                     "#00000000",
		"titleBar.inactiveBackground":                         p.BGSecondary,
		"titleBar.inactiveForeground":                         p.FGMuted,
		"tree.indentGuidesStroke":                             p.FGMuted + "50",
		"welcomePage.progress.foreground":                     p.FGSecondary,
		"widget.border":                                       "#00000000",
	}
}

func buildPlainTheme(mode string, p palette.Palette) map[string]any {
	modeLabel := palette.Pick(mode, "Dark", "Light")
	return map[string]any{
		"colors": buildPlainColors(p, mode),
		"name":   "Monochromator " + modeLabel + " Plain",
		"tokenColors": []any{
			tokenRule([]string{"comment"}, p.FGMuted, "italic"),
		},
	}
}

func buildDefaultTheme(mode string, p palette.Palette, accent string) map[string]any {
	modeLabel := palette.Pick(mode, "Dark", "Light")
	return map[string]any{
		"colors": map[string]string{
			"editor.selectionBackground":                          accent + "4d",
			"editorBracketHighlight.foreground1":                  p.FGMuted,
			"editorBracketHighlight.foreground2":                  p.FGMuted,
			"editorBracketHighlight.foreground3":                  p.FGMuted,
			"editorBracketHighlight.foreground4":                  p.FGMuted,
			"editorBracketHighlight.foreground5":                  p.FGMuted,
			"editorBracketHighlight.foreground6":                  p.FGMuted,
			"editorBracketHighlight.unexpectedBracket.foreground": p.FGMuted,
			"editorBracketMatch.border":                           p.FGMuted,
			"selection.background":                                accent + "4d",
		},
		"include": "./plain.json",
		"name":    "Monochromator " + modeLabel,
		"tokenColors": []any{
			tokenRule(scopeGray, p.FGMuted, "normal"),
			tokenRule(scopeBoldItalicGray, p.FGMuted, "bold italic"),
			tokenRule(scopeItalicGray, p.FGMuted, "italic"),
			tokenRule(scopeText, p.FG, "normal"),
			tokenRule(scopeBoldText, p.FGContrast, "bold"),
			tokenRule(scopeItalicText, p.FG, "italic"),
			tokenRule(scopeLiterals, accent, "normal"),
			tokenRule(scopeBoldLiterals, accent, "bold"),
			tokenRule([]string{"markup.italic"}, p.FG, "italic"),
			tokenRule([]string{"markup.underline"}, "", "underline"),
			tokenRule([]string{"markup.strikethrough"}, "", "strikethrough"),
		},
	}
}

// A variant theme only overrides the accent rules. It gets all other
// rules from default.json through the "include" chain
// (variant -> default -> plain).
func buildVariantTheme(mode, variantName, accent string) map[string]any {
	modeLabel := palette.Pick(mode, "Dark", "Light")
	label := strings.ToUpper(variantName[:1]) + variantName[1:]
	return map[string]any{
		"colors": map[string]string{
			"editor.selectionBackground":           accent + "4d",
			"editor.wordHighlightStrongBackground": accent + "40",
			"selection.background":                 accent + "4d",
		},
		"include": "./default.json",
		"name":    "Monochromator " + modeLabel + " " + label,
		"tokenColors": []any{
			tokenRule(scopeLiterals, accent, "normal"),
			tokenRule(scopeBoldLiterals, accent, "bold"),
		},
	}
}

func writeTheme(filePath string, theme map[string]any) error {
	data, err := json.Marshal(theme)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0o644)
}

// Generate writes the VS Code theme files for both modes into
// outputDir: one plain theme, the default theme, and one theme for
// each variant.
func Generate(cfg *palette.Config, outputDir string) error {
	for _, mode := range []string{"dark", "light"} {
		dir := filepath.Join(outputDir, mode)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}

		p := cfg.Palettes[mode]

		if err := writeTheme(filepath.Join(dir, "plain.json"), buildPlainTheme(mode, p)); err != nil {
			return fmt.Errorf("writing plain.json: %w", err)
		}

		for _, v := range cfg.Variants {
			accent, err := p.Color(v.Color)
			if err != nil {
				return fmt.Errorf("resolving accent for variant %q: %w", v.Name, err)
			}

			var theme map[string]any
			if v.Name == "default" {
				theme = buildDefaultTheme(mode, p, accent)
			} else {
				theme = buildVariantTheme(mode, v.Name, accent)
			}

			if err := writeTheme(filepath.Join(dir, v.Name+".json"), theme); err != nil {
				return fmt.Errorf("writing %s.json: %w", v.Name, err)
			}
		}
	}

	return nil
}
