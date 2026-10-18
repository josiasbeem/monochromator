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
