package vscode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// The injection grammars add one of these meta scopes around a token. A
// foreign theme has no rule for them and shows the code unchanged; the scope*
// lists in generate.go give them the style of their rule in the spec.
const (
	metaType         = "meta.type.monochromator"         // rule 3: types
	metaKeyword      = "meta.keyword.monochromator"      // rule 2: keywords
	metaLiteral      = "meta.literal.monochromator"      // rule 5: literals
	metaPlaceholder  = "meta.placeholder.monochromator"  // rule 6: escapes, placeholders, interpolation
	metaPunctuation  = "meta.punctuation.monochromator"  // rule 7: operators, punctuation
	metaSelf         = "meta.self.monochromator"         // rule 7: self and this
	metaFlag         = "meta.flag.monochromator"         // rule 1: shell flags
	metaRawString    = "meta.rawstring.monochromator"    // marker without a style
	metaTypeOperator = "meta.typeoperator.monochromator" // marker without a style
)

// An injection adds patterns to the grammar that VS Code uses for a language;
// it does not replace that grammar. A selector with "L:" wins against the
// grammar at the same position: then each pattern copies the match and the
// inner scopes of the grammar rule it replaces. Without "L:" a pattern only
// takes text that no rule of the grammar matches.
type injection struct {
	file     string // file name in editors/vscode/syntaxes
	injectTo []string
	selector string
	patterns []rule
}

// A rule is one TextMate pattern with the keys that the injections use.
type rule struct {
	Match               string          `json:"match,omitempty"`
	Begin               string          `json:"begin,omitempty"`
	End                 string          `json:"end,omitempty"`
	Name                string          `json:"name,omitempty"`
	ContentName         string          `json:"contentName,omitempty"`
	ApplyEndPatternLast bool            `json:"applyEndPatternLast,omitempty"`
	Include             string          `json:"include,omitempty"`
	Captures            map[string]rule `json:"captures,omitempty"`
	BeginCaptures       map[string]rule `json:"beginCaptures,omitempty"`
	EndCaptures         map[string]rule `json:"endCaptures,omitempty"`
	Patterns            []rule          `json:"patterns,omitempty"`
}

var injections = []injection{
	// Self and self share variable.language.self.rust, the anonymous lifetime
	// '_ shares the scope of 'a. Copies of lvariables and lifetimes. Raw strings
	// share the scope of normal strings but have no interpolation: the marker
	// keeps rust-string.json out of them. Inside generics the Rust grammar has
	// no strings, so the copy does not start one after a name (Encode<"x">).
	{
		file:     "rust.json",
		injectTo: []string{"source.rust"},
		selector: "L:source.rust - (comment | string)",
		patterns: []rule{
			{Match: `\bSelf\b`, Name: metaType + " variable.language.self.rust"},
			{Match: `(')(_)(?!')\b`, Captures: map[string]rule{
				"1": {Name: "punctuation.definition.lifetime.rust"},
				"2": {Name: metaPunctuation + " entity.name.type.lifetime.rust"},
			}},
			{Begin: `(?<![A-Za-z0-9_])(b?r)(#*)(")`, End: `(")(\2)`, Name: metaRawString + " string.quoted.double.rust", BeginCaptures: map[string]rule{
				"1": {Name: "string.quoted.byte.raw.rust"},
				"2": {Name: "punctuation.definition.string.raw.rust"},
				"3": {Name: "punctuation.definition.string.rust"},
			}, EndCaptures: map[string]rule{
				"1": {Name: "punctuation.definition.string.rust"},
				"2": {Name: "punctuation.definition.string.raw.rust"},
			}},
		},
	},
	// The format spec in {c:?} sits in one token with the name. Copy of
	// interpolations, without {{ and without spaces in the spec.
	{
		file:     "rust-string.json",
		injectTo: []string{"source.rust"},
		selector: "L:string.quoted.double.rust - " + metaRawString,
		patterns: []rule{
			{Match: `(?<!\{)({)\w*(:[^"{}\s]*)?(})`, Name: "meta.interpolation.rust", Captures: map[string]rule{
				"1": {Name: "punctuation.definition.interpolation.rust"},
				"2": {Name: metaPlaceholder},
				"3": {Name: "punctuation.definition.interpolation.rust"},
			}},
		},
	},
	// The C grammar gives the name after struct, union and enum, the library
	// types without _t and the dot of a designator no scope.
	{
		file:     "c.json",
		injectTo: []string{"source.c"},
		selector: "source.c - (comment | string)",
		patterns: []rule{
			{Match: `(?<=\bstruct\s|\bunion\s|\benum\s)[A-Za-z_]\w*`, Name: metaType},
			{Match: `\b(?:FILE|va_list|jmp_buf)\b`, Name: metaType},
			{Match: `\.(?=[A-Za-z_])`, Name: metaPunctuation},
		},
	},
	// A class as qualifier (Objects.requireNonNull) has the scope of a variable.
	// Copy of #objects for capitalized names with a lowercase letter; a
	// mixed-case member (Map.Entry) is a nested type that other rules own.
	{
		file:     "java.json",
		injectTo: []string{"source.java"},
		selector: "L:source.java - (comment | string | meta.import.java)",
		patterns: []rule{
			{Match: `(?<![\w$])(?<!\bnew\s)[A-Z][A-Z0-9_$]*[a-z][\w$]*(?=\s*\.\s*(?:[a-z_$][\w$]*|[A-Z][A-Z0-9_$]*)\b)`, Name: metaType + " variable.other.object.java"},
		},
	},
	// Only the whole import path has a scope.
	{
		file:     "java-import.json",
		injectTo: []string{"source.java"},
		selector: "storage.modifier.import.java - comment",
		patterns: []rule{
			{Match: `\b[A-Z][A-Z0-9_$]*[a-z][\w$]*\b`, Name: metaType},
		},
	},
	// The escape rule \. covers only \u of \u2014.
	{
		file:     "java-string.json",
		injectTo: []string{"source.java"},
		selector: "L:(string.quoted.double.java | string.quoted.single.java | string.quoted.triple.java)",
		patterns: []rule{
			{Match: `(\\u)([0-9a-fA-F]{4})`, Captures: map[string]rule{
				"1": {Name: "constant.character.escape.java"},
				"2": {Name: metaPlaceholder},
			}},
		},
	},
	// Braces after commands like \title have no scope.
	{
		file:     "latex.json",
		injectTo: []string{"text.tex.latex", "text.tex"},
		selector: "text.tex - (comment | markup.raw | meta.embedded | source | string)",
		patterns: []rule{
			{Match: `[{}]`, Name: metaPunctuation},
		},
	},
	// = < > in math have no scope. The spaces keep the rule out of files with
	// an unpaired $, where the grammar reads prose and XML as math.
	{
		file:     "latex-math.json",
		injectTo: []string{"text.tex.latex", "text.tex"},
		selector: "meta.math.block - (comment | markup.raw | meta.embedded | source | string)",
		patterns: []rule{
			{Match: `(?<=\s)[=<>](?=\s)`, Name: metaPunctuation},
		},
	},
	// The braces of an initializer have no scope.
	{
		file:     "cpp.json",
		injectTo: []string{"source.cpp"},
		selector: "source.cpp - (comment | string)",
		patterns: []rule{
			{Match: `[{}]`, Name: metaPunctuation},
		},
	},
	// Test flags share keyword.operator.logical.shell with && == !. Copy of the
	// flag rule.
	{
		file:     "shell.json",
		injectTo: []string{"source.shell"},
		selector: "L:meta.scope.logical-expression.shell - (string | meta.scope.subshell | meta.parameter-expansion | meta.arithmetic | comment)",
		patterns: []rule{
			{Match: `(?<!\S)-(nt|ot|ef|eq|ne|l[te]|g[te]|[a-hknoprstuwxzOGLSN])\b`, Name: metaFlag + " keyword.operator.logical.shell"},
		},
	},
	// error shares keyword.type.zig with type, void and bool. Copy of the error
	// rule; a quoted name @"error" has no keywords.
	{
		file:     "zig.json",
		injectTo: []string{"source.zig"},
		selector: "L:source.zig - (comment | string | variable.string)",
		patterns: []rule{
			{Match: `\b(error)(?:(\.)(\w+))?\b`, Captures: map[string]rule{
				"1": {Name: metaKeyword + " keyword.type.zig"},
				"2": {Name: "punctuation.accessor.zig"},
				"3": {Name: "entity.name.error.zig"},
			}},
		},
	},
	// this shares support.type.builtin.ts with object, which stays a type
	// through support.type. Copy of type-builtin-literals, only in the type
	// regions that use it. After | and & the grammar opens a region without
	// scope that ends before the next type; an L: rule at that position would
	// win against the end and leave the region open. The first pattern copies
	// that region with a marker, and inside the marker this injection is off.
	// One selector with alternatives: each comma part is an injection of its
	// own with a search of its own.
	{
		file:     "typescript.json",
		injectTo: []string{"source.ts"},
		selector: "L:(meta.type.parameters.ts | meta.type.annotation.ts | meta.type.declaration.ts | meta.return.type.ts | meta.type.function.return.ts | meta.type.paren.cover.ts | meta.type.tuple.ts) - (string | comment | " + metaTypeOperator + ")",
		patterns: []rule{
			{Begin: `[&|](?!\s*\{)`, End: `(?=\S)`, ContentName: metaTypeOperator, BeginCaptures: map[string]rule{
				"0": {Name: "keyword.operator.type.ts"},
			}},
			{Match: `(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))(this)(?![_$[:alnum:]])(?:(?=\.\.\.)|(?!\.))(?!\s*\??:)(?!\s*\()`, Captures: map[string]rule{
				"1": {Name: metaSelf + " support.type.builtin.ts"},
			}},
		},
	},
	// The same for TSX.
	{
		file:     "typescriptreact.json",
		injectTo: []string{"source.tsx"},
		selector: "L:(meta.type.parameters.tsx | meta.type.annotation.tsx | meta.type.declaration.tsx | meta.return.type.tsx | meta.type.function.return.tsx | meta.type.paren.cover.tsx | meta.type.tuple.tsx) - (string | comment | " + metaTypeOperator + ")",
		patterns: []rule{
			{Begin: `[&|](?!\s*\{)`, End: `(?=\S)`, ContentName: metaTypeOperator, BeginCaptures: map[string]rule{
				"0": {Name: "keyword.operator.type.tsx"},
			}},
			{Match: `(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))(this)(?![_$[:alnum:]])(?:(?=\.\.\.)|(?!\.))(?!\s*\??:)(?!\s*\()`, Captures: map[string]rule{
				"1": {Name: metaSelf + " support.type.builtin.tsx"},
			}},
		},
	},
}

func (in injection) scopeName() string {
	return "monochromator.injection." + strings.TrimSuffix(in.file, ".json")
}

type grammarFile struct {
	ScopeName         string `json:"scopeName"`
	InjectionSelector string `json:"injectionSelector"`
	Patterns          []rule `json:"patterns"`
}

func writeGrammar(filePath string, in injection) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	// The patterns contain < > and &, which the default escaping would write
	// as \u003c and so on.
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(grammarFile{ScopeName: in.scopeName(), InjectionSelector: in.selector, Patterns: in.patterns})
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, buf.Bytes(), 0o644)
}

// GenerateInjections writes one TextMate injection grammar per entry of
// injections into outputDir. editors/vscode/package.json lists each file
// under contributes.grammars.
func GenerateInjections(outputDir string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("creating directory %s: %w", outputDir, err)
	}
	for _, in := range injections {
		if err := writeGrammar(filepath.Join(outputDir, in.file), in); err != nil {
			return fmt.Errorf("writing %s: %w", in.file, err)
		}
	}
	return nil
}
