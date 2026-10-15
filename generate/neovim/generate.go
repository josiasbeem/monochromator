// Package neovim generates the Neovim color scheme files from the palette.
package neovim

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
	accent         color = "accent"
	bg             color = "bg"
	bgSecondary    color = "bg_secondary"
	bgSelection    color = "bg_selection"
	blue           color = "blue"
	border         color = "border"
	cyan           color = "cyan"
	diffAdd        color = "diff_add"
	diffChange     color = "diff_change"
	diffDelete     color = "diff_delete"
	diffText       color = "diff_text"
	diffTextAdd    color = "diff_text_add"
	fg             color = "fg"
	fgContrast     color = "fg_contrast"
	fgFaint        color = "fg_faint"
	fgMuted        color = "fg_muted"
	fold           color = "fold"
	green          color = "green"
	inlayHint      color = "inlay_hint"
	magenta        color = "magenta"
	orange         color = "orange"
	red            color = "red"
	reference      color = "reference"
	referenceWrite color = "reference_write"
	search         color = "search"
	selection      color = "selection"
	terminalBlack  color = "terminal_black"
	terminalWhite  color = "terminal_white"
	yellow         color = "yellow"
)

type attr string

const (
	bold          attr = "bold"
	italic        attr = "italic"
	nocombine     attr = "nocombine"
	strikethrough attr = "strikethrough"
	undercurl     attr = "undercurl"
	underline     attr = "underline"
)

type style struct {
	fg    color
	bg    color
	sp    color
	attrs []attr
}

// Neovim adds the attributes of all highlights on a character: a function
// name in a bold heading becomes bold and italic. VS Code uses only the
// most specific scope. The "nocombine" attribute gives the VS Code
// behavior.
func syntaxStyle(fg color, attrs ...attr) style {
	return style{fg: fg, attrs: append(attrs, nocombine)}
}

// Neovim injects the comment languages into comments and highlights their
// parts like code. VS Code shows these parts like the comment around them.
// Only the tags, for example "@param", are bold, and the punctuation of
// JSDoc is not italic.
func commentLanguageSyntax(tag, text, punctuation style) map[string]style {
	syntax := map[string]style{}
	for lang, captures := range map[string][]string{
		"comment": {
			"@constant", "@number", "@punctuation.bracket", "@punctuation.delimiter",
			"@string.special.url",
		},
		"doxygen": {
			"@function", "@keyword.modifier", "@label", "@markup.italic", "@operator",
			"@punctuation.bracket", "@punctuation.delimiter", "@type", "@variable",
			"@variable.parameter",
		},
		"javadoc": {
			"@function", "@markup.link.url", "@markup.raw", "@module", "@number",
			"@operator", "@property", "@punctuation.bracket", "@punctuation.delimiter",
			"@string", "@string.special", "@type", "@type.builtin", "@variable",
			"@variable.member", "@variable.parameter",
		},
		"jsdoc": {
			"@function", "@property", "@type", "@variable", "@variable.parameter",
		},
		"luadoc": {
			"@comment", "@constant.builtin", "@function.macro", "@keyword.function",
			"@module", "@number", "@operator", "@punctuation.bracket",
			"@punctuation.delimiter", "@punctuation.special", "@string",
			"@string.special.path", "@type", "@type.builtin", "@variable",
			"@variable.builtin", "@variable.member", "@variable.parameter",
		},
		"phpdoc": {
			"@function.method", "@property", "@string.special.url", "@type", "@variable",
			"@variable.member", "@variable.parameter",
		},
	} {
		for _, capture := range captures {
			syntax[capture+"."+lang] = text
		}
	}
	for _, group := range []string{
		"@attribute.phpdoc", "@keyword.coroutine.luadoc", "@keyword.doxygen",
		"@keyword.import.luadoc", "@keyword.javadoc", "@keyword.jsdoc", "@keyword.luadoc",
		"@keyword.modifier.luadoc", "@keyword.phpdoc", "@keyword.return.luadoc",
		"@tag.doxygen",
	} {
		syntax[group] = tag
	}
	for _, group := range []string{"@punctuation.bracket.jsdoc", "@punctuation.delimiter.jsdoc"} {
		syntax[group] = punctuation
	}
	return syntax
}

// The plain variant styles only the comments, like the VS Code plain
// theme. It must set all other syntax groups to the text color, because
// Neovim has default colors for them.
func plainSyntax() map[string]style {
	syntax := map[string]style{}
	for group := range defaultSyntax() {
		syntax[group] = syntaxStyle(fg)
	}
	comment := syntaxStyle(fgMuted, italic)
	for _, group := range []string{
		"@comment", "@comment.error", "@comment.note", "@comment.todo", "@comment.warning",
		"Comment", "SpecialComment", "Todo",
		// The shebang lines of shell scripts and Python scripts are comments
		// in VS Code.
		"@keyword.directive.bash", "@keyword.directive.python",
	} {
		syntax[group] = comment
	}
	maps.Copy(syntax, commentLanguageSyntax(comment, comment, comment))
	return syntax
}

// The names are the same as in the scope lists of the VS Code generator:
// Gray = FGMuted, Text = FG or FGContrast, Literal = the accent color.
// The groups without a dot, for example "Statement", are the Vim syntax
// groups. Neovim uses them for languages without a tree-sitter parser.
func defaultSyntax() map[string]style {
	var (
		gray             = syntaxStyle(fgMuted)
		boldItalicGray   = syntaxStyle(fgMuted, bold, italic)
		italicGray       = syntaxStyle(fgMuted, italic)
		text             = syntaxStyle(fg)
		boldText         = syntaxStyle(fgContrast, bold)
		italicText       = syntaxStyle(fgContrast, italic)
		underlineText    = syntaxStyle(fgContrast, underline)
		literal          = syntaxStyle(accent)
		underlineLiteral = syntaxStyle(accent, underline)
	)
	syntax := map[string]style{
		"@attribute":                   gray,
		"@attribute.builtin":           gray,
		"@boolean":                     literal,
		"@character":                   literal,
		"@character.special":           literal,
		"@comment":                     italicGray,
		"@comment.error":               italicGray,
		"@comment.note":                italicGray,
		"@comment.todo":                italicGray,
		"@comment.warning":             italicGray,
		"@constant":                    text,
		"@constant.builtin":            literal,
		"@constructor":                 italicText,
		"@function":                    italicText,
		"@function.builtin":            italicText,
		"@keyword":                     boldText,
		"@keyword.conditional.ternary": gray,
		"@keyword.operator":            gray,
		"@label":                       text,
		"@markup":                      text,
		"@markup.heading":              boldText,
		"@markup.italic":               syntaxStyle(fg, italic),
		"@markup.link":                 text,
		"@markup.link.url":             syntaxStyle(fg, underline),
		"@markup.list":                 gray,
		"@markup.quote":                italicGray,
		"@markup.raw":                  literal,
		"@markup.raw.block":            text,
		"@markup.strikethrough":        syntaxStyle("", strikethrough),
		"@markup.strong":               boldText,
		"@markup.underline":            syntaxStyle("", underline),
		"@module":                      text,
		"@module.builtin":              text,
		"@number":                      literal,
		"@number.float":                literal,
		"@operator":                    gray,
		"@property":                    text,
		"@punctuation":                 gray,
		"@punctuation.special":         gray,
		"@string":                      literal,
		"@string.escape":               underlineLiteral,
		"@string.regexp":               literal,
		"@string.special":              literal,
		"@string.special.path":         text,
		"@string.special.symbol":       text,
		// URLs in strings keep the string style. Markdown links use
		// "@markup.link.url".
		"@string.special.url":         syntaxStyle(""),
		"@tag":                        boldText,
		"@tag.attribute":              text,
		"@tag.builtin":                boldText,
		"@tag.delimiter":              gray,
		"@type":                       underlineText,
		"@type.builtin":               underlineText,
		"@variable":                   text,
		"@variable.builtin":           gray,
		"@variable.parameter.builtin": text,
		"Comment":                     italicGray,
		"Constant":                    literal,
		"Delimiter":                   gray,
		"Function":                    italicText,
		"Identifier":                  text,
		"Operator":                    gray,
		"PreProc":                     boldText,
		"Special":                     text,
		"SpecialChar":                 underlineLiteral,
		"SpecialComment":              italicGray,
		"Statement":                   boldText,
		"StorageClass":                boldText,
		"String":                      literal,
		"Structure":                   boldText,
		"Tag":                         boldText,
		"Title":                       boldText,
		"Todo":                        italicGray,
		"Type":                        underlineText,
		"Typedef":                     boldText,
		"Underlined":                  syntaxStyle("", underline),
	}

	// A group with a language name at the end changes only that language.
	// These groups copy the VS Code style where a tree-sitter capture puts
	// tokens together that the TextMate grammar keeps apart.
	maps.Copy(syntax, map[string]style{
		// Bash: commands and built-in commands are bold, the shebang line
		// is gray, bold, and italic.
		"@function.builtin.bash":  boldText,
		"@function.call.bash":     boldText,
		"@keyword.directive.bash": boldItalicGray,
		"@label.bash":             boldText,
		"@variable.builtin.bash":  text,
		// The query uses "@none" to remove the string style from an
		// expansion such as "${name:-default}". In other languages,
		// "@none" stays undefined, because HTML uses it for the text in
		// Javadoc comments.
		"@none.bash": text,

		// C and C++: macro names in "#define" have the text color, and
		// "sizeof" is italic like a function.
		"@function.macro.c":   text,
		"@function.macro.cpp": text,
		"@keyword.operator.c": italicText,

		// C# and Java: "new" is bold.
		"@keyword.operator.c_sharp": boldText,
		"@keyword.operator.java":    boldText,

		// CSS and SCSS: the selectors are bold.
		"@attribute.css":  boldText,
		"@attribute.scss": boldText,
		"@constant.css":   boldText,
		"@constant.scss":  boldText,
		"@type.css":       boldText,
		"@type.scss":      boldText,

		// The format placeholders, for example "%s", are underlined.
		"@character.printf": underlineLiteral,

		// Haskell: data constructors have the text color.
		"@constructor.haskell": text,

		// The delimiters of string interpolation, for example "${" and "}",
		// are underlined.
		"@punctuation.special.c_sharp":    underlineLiteral,
		"@punctuation.special.dart":       underlineLiteral,
		"@punctuation.special.javascript": underlineLiteral,
		"@punctuation.special.python":     underlineLiteral,
		"@punctuation.special.tsx":        underlineLiteral,
		"@punctuation.special.typescript": underlineLiteral,

		// LaTeX: commands are bold. Their arguments, titles, and formatted
		// text have the text color.
		"@function.latex":            boldText,
		"@function.macro.latex":      boldText,
		"@markup.heading.1.latex":    text,
		"@markup.heading.2.latex":    text,
		"@markup.heading.3.latex":    text,
		"@markup.heading.4.latex":    text,
		"@markup.heading.5.latex":    text,
		"@markup.heading.6.latex":    text,
		"@markup.italic.latex":       text,
		"@markup.strong.latex":       text,
		"@module.latex":              boldText,
		"@punctuation.special.latex": boldText,
		"@string.latex":              text,

		// Lua: the braces of a table are punctuation.
		"@constructor.lua": gray,

		// Make: variable names have the text color.
		"@string.make": text,

		// Markdown: the headings use "@markup.heading.1" and the other
		// levels. For this reason, "@markup.heading.markdown" changes only
		// the header row of a table. The brackets of links are gray.
		"@markup.heading.markdown":     text,
		"@markup.link.markdown_inline": gray,

		// Nix: paths are literals.
		"@string.special.path.nix": literal,

		// PHP: the class names after "new" are types.
		"@constructor.php": underlineText,

		// Python: the shebang line is a comment.
		"@keyword.directive.python": italicGray,

		// Ruby: "include", "extend", and "attr_reader" are bold like keywords.
		"@function.builtin.ruby": boldText,

		// Rust: lifetimes are underlined like types. "Some", "None", "Ok",
		// "Err", and the macros "assert!" and "panic!" are italic.
		"@attribute.builtin.rust": underlineText,
		"@attribute.rust":         underlineText,
		"@constant.builtin.rust":  italicText,
		"@keyword.exception.rust": italicText,

		// SQL: table names have the text color.
		"@type.sql": text,

		// Typst: the "#" before code is bold.
		"@punctuation.special.typst": boldText,
	})

	maps.Copy(syntax, commentLanguageSyntax(boldItalicGray, italicGray, gray))
	return syntax
}

// VS Code does not use semantic tokens with this theme, because the theme
// does not turn them on. Neovim puts the semantic tokens of the language
// server on top of the tree-sitter highlights. Empty groups keep the
// tree-sitter highlights, like in VS Code. The deprecated modifier keeps
// its strikethrough, because VS Code shows deprecated code the same way.
func semanticTokenGroups() map[string]style {
	groups := map[string]style{}
	for _, tokenType := range []string{
		"class", "comment", "decorator", "enum", "enumMember", "event", "function",
		"interface", "keyword", "macro", "method", "modifier", "namespace", "number",
		"operator", "parameter", "property", "regexp", "string", "struct", "type",
		"typeParameter", "variable",
	} {
		groups["@lsp.type."+tokenType] = style{}
	}
	return groups
}

func buildUI() map[string]style {
	return map[string]style{
		"@diff.delta":              {fg: orange},
		"@diff.minus":              {fg: red},
		"@diff.plus":               {fg: green},
		"Added":                    {fg: green},
		"Changed":                  {fg: orange},
		"ColorColumn":              {bg: bgSecondary},
		"Conceal":                  {fg: fgMuted},
		"CurSearch":                {bg: bgSelection},
		"Cursor":                   {fg: bg, bg: fg},
		"CursorColumn":             {bg: bgSecondary},
		"CursorLine":               {bg: bgSecondary},
		"CursorLineNr":             {fg: fg},
		"DiagnosticDeprecated":     {attrs: []attr{strikethrough}},
		"DiagnosticError":          {fg: red},
		"DiagnosticHint":           {fg: fgMuted},
		"DiagnosticInfo":           {fg: blue},
		"DiagnosticOk":             {fg: green},
		"DiagnosticUnderlineError": {sp: red, attrs: []attr{undercurl}},
		"DiagnosticUnderlineHint":  {sp: fgMuted, attrs: []attr{undercurl}},
		"DiagnosticUnderlineInfo":  {sp: blue, attrs: []attr{undercurl}},
		"DiagnosticUnderlineOk":    {sp: green, attrs: []attr{undercurl}},
		"DiagnosticUnderlineWarn":  {sp: orange, attrs: []attr{undercurl}},
		// Neovim links unused code to "Comment". The VS Code README turns
		// off the same feature, because variables must not look similar to
		// comments.
		"DiagnosticUnnecessary": {},
		"DiagnosticWarn":        {fg: orange},
		"DiffAdd":               {bg: diffAdd},
		"DiffChange":            {bg: diffChange},
		"DiffDelete":            {fg: fgMuted, bg: diffDelete},
		"DiffText":              {bg: diffText},
		"DiffTextAdd":           {bg: diffTextAdd},
		"Directory":             {fg: fgContrast},
		"Error":                 {fg: red},
		"ErrorMsg":              {fg: red},
		"FloatBorder":           {fg: border, bg: bgSecondary},
		"FloatTitle":            {fg: fg, bg: bgSecondary},
		"Folded":                {fg: fgMuted, bg: fold},
		"LineNr":                {fg: fgMuted},
		"LspCodeLens":           {fg: fgMuted},
		"LspInlayHint":          {fg: fgMuted, bg: inlayHint},
		"LspReferenceRead":      {bg: reference},
		"LspReferenceText":      {bg: reference},
		"LspReferenceWrite":     {bg: referenceWrite},
		"MatchParen":            {bg: bgSelection},
		"ModeMsg":               {fg: fg},
		"MoreMsg":               {fg: fg},
		"NonText":               {fg: fgFaint},
		"Normal":                {fg: fg, bg: bg},
		"NormalFloat":           {fg: fg, bg: bgSecondary},
		"OkMsg":                 {fg: green},
		"Pmenu":                 {fg: fg, bg: bgSecondary},
		"PmenuBorder":           {fg: border, bg: bgSecondary},
		"PmenuExtra":            {fg: fgMuted, bg: bgSecondary},
		"PmenuExtraSel":         {fg: fgMuted, bg: bgSelection},
		"PmenuKind":             {fg: fgMuted, bg: bgSecondary},
		"PmenuKindSel":          {fg: fgMuted, bg: bgSelection},
		"PmenuSbar":             {bg: bgSecondary},
		"PmenuSel":              {fg: fg, bg: bgSelection},
		"PmenuThumb":            {bg: bgSelection},
		"Question":              {fg: fg},
		"QuickFixLine":          {bg: bgSelection},
		"Removed":               {fg: red},
		"Search":                {bg: search},
		"SignColumn":            {fg: fgMuted},
		"SpecialKey":            {fg: fgMuted},
		"SpellBad":              {sp: red, attrs: []attr{undercurl}},
		"SpellCap":              {sp: orange, attrs: []attr{undercurl}},
		"SpellLocal":            {sp: green, attrs: []attr{undercurl}},
		"SpellRare":             {sp: blue, attrs: []attr{undercurl}},
		"StatusLine":            {fg: fg, bg: bgSecondary},
		"StatusLineNC":          {fg: fgMuted, bg: bgSecondary},
		"TabLine":               {fg: fgMuted, bg: bgSecondary},
		"TabLineFill":           {bg: bgSecondary},
		"TabLineSel":            {fg: fg, bg: bg},
		"Visual":                {bg: selection},
		"WarningMsg":            {fg: orange},
		"WinBar":                {fg: fgMuted, bg: bg},
		"WinBarNC":              {fg: fgMuted, bg: bg},
		"WinSeparator":          {fg: bgSecondary},
	}
}

// The ANSI colors 0 to 15 for the terminal buffers, the same as in VS Code.
var terminalColors = []color{
	terminalBlack, red, green, yellow, blue, magenta, cyan, fgMuted,
	fgMuted, red, green, yellow, blue, magenta, cyan, terminalWhite,
}

// Neovim has no alpha channel. The generator mixes the transparent VS Code
// colors with the background. 0x4d is about 30%.
func buildPalette(mode string, p palette.Palette, accentHex string) (map[color]string, error) {
	colors := map[color]string{
		bg:            p.BG,
		bgSecondary:   p.BGSecondary,
		bgSelection:   p.BGSelection,
		blue:          p.Blue,
		border:        p.Border,
		cyan:          p.Cyan,
		fg:            p.FG,
		fgContrast:    p.FGContrast,
		fgMuted:       p.FGMuted,
		green:         p.Green,
		inlayHint:     p.BGSecondary,
		magenta:       p.Magenta,
		orange:        p.Orange,
		red:           p.Red,
		selection:     p.BGSelection,
		terminalBlack: palette.Pick(mode, p.BG, p.FG),
		terminalWhite: palette.Pick(mode, p.FG, "#ffffff"),
		yellow:        p.Yellow,
	}

	type mix struct {
		name  color
		fore  string
		alpha uint8
	}
	mixes := []mix{
		{diffAdd, p.Green, 0x26},
		{diffChange, p.Orange, 0x26},
		{diffDelete, p.Red, 0x26},
		{diffText, p.Orange, 0x4d},
		{diffTextAdd, p.Green, 0x4d},
		{fgFaint, p.FGMuted, 0x50},
		{fold, p.BGSecondary, 0x80},
		{reference, p.BGSelection, 0x80},
		{referenceWrite, p.Blue, 0x40},
		{search, p.BGSelection, 0xbf},
	}
	if mode == "light" {
		mixes = append(mixes, mix{inlayHint, p.BGSecondary, 0x80})
	}
	if accentHex != "" {
		colors[accent] = accentHex
		mixes = append(mixes, mix{referenceWrite, accentHex, 0x40}, mix{selection, accentHex, 0x4d})
	}

	for _, m := range mixes {
		hex, err := blend(m.fore, p.BG, m.alpha)
		if err != nil {
			return nil, fmt.Errorf("mixing %s: %w", m.name, err)
		}
		colors[m.name] = hex
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

func (s style) lua() string {
	var fields []string
	for _, f := range []struct {
		key   string
		value color
	}{{"fg", s.fg}, {"bg", s.bg}, {"sp", s.sp}} {
		if f.value != "" {
			fields = append(fields, fmt.Sprintf("%s = c.%s", f.key, f.value))
		}
	}
	for _, a := range s.attrs {
		fields = append(fields, string(a)+" = true")
	}
	if len(fields) == 0 {
		return "{}"
	}
	return "{ " + strings.Join(fields, ", ") + " }"
}

func renderScheme(name, mode string, groups map[string]style, colors map[color]string) (string, error) {
	for c, hex := range colors {
		if _, err := parseHex(hex); err != nil {
			return "", fmt.Errorf("palette entry %q: %w", c, err)
		}
	}
	checkColor := func(where string, c color) error {
		if _, ok := colors[c]; c != "" && !ok {
			return fmt.Errorf("%s: color %q is not in the palette", where, c)
		}
		return nil
	}

	var b strings.Builder
	b.WriteString("vim.cmd.highlight(\"clear\")\n")
	fmt.Fprintf(&b, "vim.o.background = %q\n", mode)
	fmt.Fprintf(&b, "vim.g.colors_name = %q\n", name)

	b.WriteString("\nlocal c = {\n")
	for _, c := range slices.Sorted(maps.Keys(colors)) {
		fmt.Fprintf(&b, "  %s = %q,\n", c, colors[c])
	}
	b.WriteString("}\n")

	b.WriteString("\nlocal set = vim.api.nvim_set_hl\n")
	for _, group := range slices.Sorted(maps.Keys(groups)) {
		s := groups[group]
		for _, c := range []color{s.fg, s.bg, s.sp} {
			if err := checkColor("group "+group, c); err != nil {
				return "", err
			}
		}
		fmt.Fprintf(&b, "set(0, %q, %s)\n", group, s.lua())
	}

	b.WriteString("\n")
	for i, c := range terminalColors {
		if err := checkColor("terminal color "+strconv.Itoa(i), c); err != nil {
			return "", err
		}
		fmt.Fprintf(&b, "vim.g.terminal_color_%d = c.%s\n", i, c)
	}
	return b.String(), nil
}

func writeScheme(path, name, mode string, groups map[string]style, colors map[color]string) error {
	text, err := renderScheme(name, mode, groups, colors)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(text), 0o644)
}

// Generate writes the Neovim color scheme files for both modes into
// outputDir: one plain scheme and one scheme for each variant. Neovim
// uses the file name as the scheme name, for example
// "monochromator-dark-amber".
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
		groups := buildUI()
		maps.Copy(groups, semanticTokenGroups())
		maps.Copy(groups, plainSyntax())
		plainName := "monochromator-" + mode + "-plain"
		plainPath := filepath.Join(outputDir, plainName+".lua")
		if err := writeScheme(plainPath, plainName, mode, groups, colors); err != nil {
			return fmt.Errorf("writing %s: %w", plainPath, err)
		}

		for _, v := range cfg.Variants {
			accentHex, err := p.Color(v.Color)
			if err != nil {
				return fmt.Errorf("resolving accent for variant %q: %w", v.Name, err)
			}

			name := "monochromator-" + mode
			if v.Name != "default" {
				name += "-" + v.Name
			}

			colors, err := buildPalette(mode, p, accentHex)
			if err != nil {
				return fmt.Errorf("building %s palette: %w", name, err)
			}
			groups := buildUI()
			maps.Copy(groups, semanticTokenGroups())
			maps.Copy(groups, defaultSyntax())
			path := filepath.Join(outputDir, name+".lua")
			if err := writeScheme(path, name, mode, groups, colors); err != nil {
				return fmt.Errorf("writing %s: %w", path, err)
			}
		}
	}

	return nil
}
