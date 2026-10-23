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
	metaConstructor  = "meta.constructor.monochromator"  // rule 4: functions and constructors
	metaKeyword      = "meta.keyword.monochromator"      // rule 2: keywords
	metaLiteral      = "meta.literal.monochromator"      // rule 5: literals
	metaPlaceholder  = "meta.placeholder.monochromator"  // rule 6: escapes, placeholders, interpolation
	metaPunctuation  = "meta.punctuation.monochromator"  // rule 7: operators, punctuation
	metaSelf         = "meta.self.monochromator"         // rule 7: self and this
	metaName         = "meta.name.monochromator"         // rule 1: names
	metaFlag         = "meta.flag.monochromator"         // rule 1: shell flags
	metaRawString    = "meta.rawstring.monochromator"    // marker without a style
	metaTypeOperator = "meta.typeoperator.monochromator" // marker without a style
	metaEnumBody     = "meta.enumbody.monochromator"     // marker without a style
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
	// The rules that a pattern includes with #name.
	repository map[string]rule
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

// The this injections of TypeScript and JavaScript share these. this before a
// dot (also ?. and !.) stays dim, as the object of a member. this and super
// before a parenthesis are the names of calls (new this(), super()) and get the
// style of a function. In all other positions, this looks like a variable. The
// rules are copies of this-literal and super-literal with a lookahead. Types
// stay out: there this is the type of the object and keeps its dim style
// (typescript.json). The expressions of a template string are code, its text is
// not. Text in JSX is not code.
//
// An L: injection wins against an empty match of the grammar at the same
// position. The grammar starts and ends some regions with an empty match, and
// the injection must not move these borders. For this reason the first
// injection does not touch these places:
//   - this before a parenthesis: meta.function-call starts at the name. The
//     second injection marks the name only inside meta.function-call.
//   - this before a colon: the grammar marks an object key there. A this in a
//     ternary (c ? this : x) also stays dim.
//   - this directly after the => of an arrow function without a block: the
//     body starts where meta.arrow ends.
//   - parameters: there the grammar also gives this variable.parameter. The
//     theme shows this with variable.parameter as text.
const (
	scriptThisExclude = " - (comment | (string - meta.template.expression) | meta.type.annotation | meta.return.type | meta.type.parameters | meta.type.declaration | meta.type.function.return | meta.type.paren.cover | meta.type.tuple | meta.parameters | (meta.arrow - (meta.arrow meta.block)) | (meta.jsx.children - meta.embedded.expression) | (meta.embedded.expression meta.jsx.children))"
	scriptThisAlone   = `(?<!=>)(?<!=>\s)(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))this\b(?!\$)(?!\s*(?:[?!]?\.|\(|:))`
	scriptThisCall    = `(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))this\b(?!\$)(?=\s*\()`
	scriptSuperCall   = `(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))super\b(?!\$)(?=\s*\()`
)

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
	// Own enum variants look like constructors; the Rust grammar gives them
	// the scope of a type. A use: a name after X:: or Self:: in a value or a
	// pattern, not in a type position and not after a super:: path. A
	// declaration: a name at the start of an entry in the body of enum Name.
	// The copy of that { is a region with a marker, and #braces keeps the
	// inner { } of a struct variant inside it. Some, None, Ok and Err have
	// rules of their own.
	{
		file:     "rust-variants.json",
		injectTo: []string{"source.rust"},
		selector: "L:source.rust - (comment | string | meta.function.definition | meta.attribute | meta.use)",
		patterns: []rule{
			{Match: `(?<=(?:^|[^A-Za-z0-9_])(?:[A-Z][A-Za-z0-9_]+|Self)::)(?=(?!(?:Some|None|Ok|Err)\b)[A-Z][A-Za-z0-9_]*[a-z][A-Za-z0-9_]*\b\s*(?:=>|\||\)|,|;|\}|\]|\.|\{|==|!=|if\b|$))(?<!\bsuper::(?:[A-Za-z_]\w*::)*)(?<!(?:->|<|\bas|\bimpl|\bfor|\bdyn|&|\btype\s+\w+\s*=)\s*(?:&\s*|mut\s+)*(?:[A-Za-z_]\w*::)+)(?!(?:Some|None|Ok|Err)\b)[A-Z][A-Za-z0-9_]*[a-z][A-Za-z0-9_]*\b`, Name: metaConstructor + " entity.name.type.rust"},
			{Begin: `(?<=\benum\s+_?[A-Z][A-Za-z0-9_]*\s*(?:<[^{};]*>)?\s*(?:where\b[^{};]*)?)\{`, End: `\}`, ContentName: metaEnumBody, BeginCaptures: map[string]rule{
				"0": {Name: "punctuation.brackets.curly.rust"},
			}, EndCaptures: map[string]rule{
				"0": {Name: "punctuation.brackets.curly.rust"},
			}, Patterns: []rule{
				{Include: "#braces"},
				{Match: `(?:(?<=[{,\]])|^)\s*((?!(?:Some|None|Ok|Err)\b)[A-Z][A-Za-z0-9_]*[a-z][A-Za-z0-9_]*\b)(?!\s*\()`, Captures: map[string]rule{
					"1": {Name: metaConstructor + " entity.name.type.rust"},
				}},
				{Include: "source.rust"},
			}},
		},
		repository: map[string]rule{
			"braces": {Begin: `\{`, End: `\}`, BeginCaptures: map[string]rule{
				"0": {Name: "punctuation.brackets.curly.rust"},
			}, EndCaptures: map[string]rule{
				"0": {Name: "punctuation.brackets.curly.rust"},
			}, Patterns: []rule{
				{Include: "#braces"},
				{Include: "source.rust"},
			}},
		},
	},
	// self is dim only before a dot, as the object of a field; alone and as
	// parameter it looks like a variable. As path (use a::{self}, self::m) it
	// stays dim like super. Copy of lvariables with a lookahead. Attributes
	// (#[instrument(skip(self))]) and extern crate self stay out: there the
	// grammar gives self no scope of its own.
	{
		file:     "rust-self.json",
		injectTo: []string{"source.rust"},
		selector: "L:source.rust - (comment | string | meta.use | meta.import | meta.attribute)",
		patterns: []rule{
			{Match: `\bself\b(?!\s*(?:\.|::))`, Name: metaName + " variable.language.self.rust"},
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
	// this is dim only before a dot or :: (this.x, this::m), as the object of
	// a member; alone it looks like a variable. this(...) and super(...) call
	// a constructor and get the style of a function. Copies of
	// constants-and-special-vars with a lookahead. A qualified this
	// (Outer.this.x) is a property for the grammar: a copy of its property
	// rule makes this dim there too. Method headers stay out, because the
	// grammar gives the receiver parameter (Outer this) variable.parameter.
	{
		file:     "java-this.json",
		injectTo: []string{"source.java"},
		selector: "L:source.java - (comment | string | meta.method.identifier)",
		patterns: []rule{
			{Match: `\bthis\b(?=\s*\()`, Name: metaConstructor + " variable.language.this.java"},
			{Match: `\bsuper\b(?=\s*\()`, Name: metaConstructor + " variable.language.java"},
			{Match: `\bthis\b(?!\s*(?:\.|::|\())`, Name: metaName + " variable.language.this.java"},
			{Match: `(\.)\s*(this)(?=\s*\.\s*[a-zA-Z_$][\w$]*)`, Captures: map[string]rule{
				"1": {Name: "punctuation.separator.period.java"},
				"2": {Name: metaSelf + " variable.other.object.property.java"},
			}},
		},
	},
	// this is dim only before a dot (also ?. and !.), as the object of a
	// member; alone it looks like a variable. this(...) and base(...) in a
	// constructor initializer call a constructor and get the style of a
	// function. The grammar starts its constructor-initializer region at the
	// name. For this reason the injection copies that region and marks the
	// name: a plain match would win against the region and leave the
	// arguments to the rule for parameter lists. A this before a name is the
	// modifier of an extension method (this string s) and stays out; is, as,
	// switch and with after this are operators.
	{
		file:     "csharp-this.json",
		injectTo: []string{"source.cs"},
		selector: "L:source.cs - (comment | (string - meta.embedded.interpolation))",
		patterns: []rule{
			{
				Begin:         `\b(base|this)\b\s*(?=\()`,
				BeginCaptures: map[string]rule{"1": {Name: metaConstructor + " variable.language.$1.cs"}},
				End:           `(?<=\))`,
				Patterns:      []rule{{Include: "source.cs#argument-list"}},
			},
			{Match: `\bthis\b(?!\s*(?:[?!]?\.|\())(?!\s+(?!(?:is|as|switch|with)\b)[@_[:alpha:]])`, Name: metaName + " variable.language.this.cs"},
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
	// this is dim only before ->, as the object of a member; alone (return
	// *this, run(this), [this]) it looks like a variable. Copy of
	// the_this_keyword with a lookahead. The grammar starts that match at the
	// white space before this, and the copy does the same: only a match at the
	// same position wins. Parameter lists stay out: there the grammar starts
	// each parameter at its first word, and it reads this Self&& self as a
	// type.
	{
		file:     "cpp-this.json",
		injectTo: []string{"source.cpp"},
		selector: "L:source.cpp - (comment | string | meta.function.definition.parameters | meta.parameter)",
		patterns: []rule{
			{Match: `(\s*+)((?<!\w)this(?!\w))(?!\s*->)`, Captures: map[string]rule{
				"2": {Name: metaName + " variable.language.this.cpp"},
			}},
		},
	},
	// $this is dim only before -> or :: ($this->x, $this::C), as the object of
	// a member; alone it looks like a variable. Copy of var_language with a
	// lookahead. Strings stay out: in "$this->x" the grammar reads this as a
	// plain variable. VS Code gives PHP files the root text.html.php, which
	// embeds source.php.
	{
		file:     "php-this.json",
		injectTo: []string{"text.html.php"},
		selector: "L:source.php - (comment | string)",
		patterns: []rule{
			{Match: `(\$)this\b(?!\s*(?:\??->|::|\())`, Name: metaName + " variable.language.this.php", Captures: map[string]rule{
				"1": {Name: "punctuation.definition.variable.php"},
			}},
		},
	},
	// self is dim only before a dot (also &. and ::), as the receiver of a
	// method; alone it looks like a variable. Copy of the self rule with a
	// lookahead. self: is the key of a hash and stays out. Method heads stay
	// out too: the grammar reads def +@ = self as a parameter list.
	{
		file:     "ruby-self.json",
		injectTo: []string{"source.ruby"},
		selector: "L:source.ruby - (comment | (string - meta.embedded) | meta.function.method)",
		patterns: []rule{
			{Match: `\bself\b(?![?!])(?!\s*(?:&?\.|::))(?!:)`, Name: metaName + " variable.language.self.ruby"},
		},
	},
	// self is dim only before a dot or a colon (self.x, self:m()), as the
	// object of a member; alone it looks like a variable. .. joins strings and
	// does not count as a dot. Copy of the self rule with a lookahead.
	// Parameter lists stay out, because there the grammar gives self
	// variable.parameter.
	{
		file:     "lua-self.json",
		injectTo: []string{"source.lua"},
		selector: "L:source.lua - (comment | string | meta.parameter)",
		patterns: []rule{
			{Match: `(?<![^.]\.|:)\b(self)\b(?!\s*(?:\.(?!\.)|:))`, Name: metaName + " variable.language.self.lua"},
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
	// null, undefined, true, false and this share support.type.builtin.ts
	// with object, which stays a type through support.type. Copy of
	// type-builtin-literals, only in the type regions that use it. After | and
	// & the grammar opens a region without scope that ends before the next
	// type; an L: rule at that position would win against the end and leave
	// the region open. The first pattern copies that region with a marker,
	// and inside the marker this injection is off. One selector with
	// alternatives: each comma part is an injection of its own with a search
	// of its own.
	{
		file:     "typescript.json",
		injectTo: []string{"source.ts"},
		selector: "L:(meta.type.parameters.ts | meta.type.annotation.ts | meta.type.declaration.ts | meta.return.type.ts | meta.type.function.return.ts | meta.type.paren.cover.ts | meta.type.tuple.ts) - (string | comment | " + metaTypeOperator + ")",
		patterns: []rule{
			{Begin: `[&|](?!\s*\{)`, End: `(?=\S)`, ContentName: metaTypeOperator, BeginCaptures: map[string]rule{
				"0": {Name: "keyword.operator.type.ts"},
			}},
			{Match: `(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))(?:(this)|(null|undefined|true|false))(?![_$[:alnum:]])(?:(?=\.\.\.)|(?!\.))(?!\s*\??:)(?!\s*\()`, Captures: map[string]rule{
				"1": {Name: metaSelf + " support.type.builtin.ts"},
				"2": {Name: metaLiteral + " support.type.builtin.ts"},
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
			{Match: `(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))(?:(this)|(null|undefined|true|false))(?![_$[:alnum:]])(?:(?=\.\.\.)|(?!\.))(?!\s*\??:)(?!\s*\()`, Captures: map[string]rule{
				"1": {Name: metaSelf + " support.type.builtin.tsx"},
				"2": {Name: metaLiteral + " support.type.builtin.tsx"},
			}},
		},
	},
	{
		file:     "typescript-this.json",
		injectTo: []string{"source.ts"},
		selector: "L:source.ts" + scriptThisExclude,
		patterns: []rule{{Match: scriptThisAlone, Name: metaName + " variable.language.this.ts"}},
	},
	{
		file:     "typescript-this-call.json",
		injectTo: []string{"source.ts"},
		selector: "L:meta.function-call.ts",
		patterns: []rule{
			{Match: scriptThisCall, Name: metaConstructor + " variable.language.this.ts"},
			{Match: scriptSuperCall, Name: metaConstructor + " variable.language.super.ts"},
		},
	},
	{
		file:     "typescriptreact-this.json",
		injectTo: []string{"source.tsx"},
		selector: "L:source.tsx" + scriptThisExclude,
		patterns: []rule{{Match: scriptThisAlone, Name: metaName + " variable.language.this.tsx"}},
	},
	{
		file:     "typescriptreact-this-call.json",
		injectTo: []string{"source.tsx"},
		selector: "L:meta.function-call.tsx",
		patterns: []rule{
			{Match: scriptThisCall, Name: metaConstructor + " variable.language.this.tsx"},
			{Match: scriptSuperCall, Name: metaConstructor + " variable.language.super.tsx"},
		},
	},
	// VS Code also gives source.js.jsx the injections of source.js (scope
	// prefix). The two JavaScript injections leave JSX to the javascriptreact
	// injections, which keep the scope names of JSX.
	{
		file:     "javascript-this.json",
		injectTo: []string{"source.js"},
		selector: "L:source.js - source.js.jsx" + scriptThisExclude,
		patterns: []rule{{Match: scriptThisAlone, Name: metaName + " variable.language.this.js"}},
	},
	{
		file:     "javascript-this-call.json",
		injectTo: []string{"source.js"},
		selector: "L:meta.function-call.js - meta.function-call.js.jsx",
		patterns: []rule{
			{Match: scriptThisCall, Name: metaConstructor + " variable.language.this.js"},
			{Match: scriptSuperCall, Name: metaConstructor + " variable.language.super.js"},
		},
	},
	{
		file:     "javascriptreact-this.json",
		injectTo: []string{"source.js.jsx"},
		selector: "L:source.js.jsx" + scriptThisExclude,
		patterns: []rule{{Match: scriptThisAlone, Name: metaName + " variable.language.this.js.jsx"}},
	},
	{
		file:     "javascriptreact-this-call.json",
		injectTo: []string{"source.js.jsx"},
		selector: "L:meta.function-call.js.jsx",
		patterns: []rule{
			{Match: scriptThisCall, Name: metaConstructor + " variable.language.this.js.jsx"},
			{Match: scriptSuperCall, Name: metaConstructor + " variable.language.super.js.jsx"},
		},
	},
	// The dot shares constant.other.character-class.regexp with \s \d \w.
	// Copy of regex-character-class without the dot.
	{
		file:     "regexp.json",
		injectTo: []string{"source.ts", "source.tsx", "source.js", "source.js.jsx"},
		selector: "L:(string.regexp.ts | string.regexp.tsx | string.regexp.js) - comment",
		patterns: []rule{
			{Match: `\\[wWsSdDtrnvf]`, Name: metaPlaceholder + " constant.other.character-class.regexp"},
		},
	},
	// The Swift grammar gives dot and comma no scope.
	{
		file:     "swift.json",
		injectTo: []string{"source.swift"},
		selector: "source.swift - (comment | string)",
		patterns: []rule{
			{Match: `[.,]`, Name: metaPunctuation},
		},
	},
	// The same inside \( ), where the content looks like outside the string.
	{
		file:     "swift-interpolation.json",
		injectTo: []string{"source.swift"},
		selector: "meta.embedded.line.swift - (comment | meta.embedded.line.swift string)",
		patterns: []rule{
			{Match: `[.,]`, Name: metaPunctuation},
		},
	},
	// self, super and Self share variable.language.swift. self and super are
	// dim only before a dot (also ?. and !.), as the object of a member; alone
	// self looks like a variable. Self is a type, as in Rust. Copies of the
	// rule in keywords with a lookahead.
	{
		file:     "swift-self.json",
		injectTo: []string{"source.swift"},
		selector: "L:source.swift - (comment | (string - meta.embedded.line))",
		patterns: []rule{
			{Match: `(?<!\.)\b(?:self|super)\b(?=\s*[?!]?\.(?!\.))`, Name: metaSelf + " variable.language.swift"},
			{Match: `(?<!\.)\bSelf\b`, Name: metaType + " variable.language.swift"},
		},
	},
	// The $ of $name and $( ) has no scope of its own. Copies of the two
	// interpolation rules; the first one copies the name pattern of the Julia
	// grammar. Only the strings that interpolate: raw"", r"", '' and var""
	// do not.
	{
		file:     "julia.json",
		injectTo: []string{"source.julia"},
		selector: "L:(string.quoted.double.julia | string.quoted.triple.double.julia | string.docstring.julia | string.interpolated.backtick.julia) - (comment | variable.interpolation.julia)",
		patterns: []rule{
			{Match: `(\$)(?:[[:alpha:]_\p{Lu}\p{Ll}\p{Lt}\p{Lm}\p{Lo}\p{Nl}⅀-⅄∿⊾⊿⊤⊥∂∅-∇∎∏∐∑∞∟∫-∳⋀-⋃◸-◿♯⟘⟙⟀⟁⦰-⦴⨀-⨆⨉-⨖⨛⨜𝛁𝛛𝛻𝜕𝜵𝝏𝝯𝞉𝞩𝟃ⁱ-⁾₁-₎∠-∢⦛-⦯℘℮゛-゜𝟎-𝟡]|[^\P{So}←-⇿]|[^\p{^Sc}$])(?:[[:word:]_!\p{Lu}\p{Ll}\p{Lt}\p{Lm}\p{Lo}\p{Nl}⅀-⅄∿⊾⊿⊤⊥∂∅-∇∎∏∐∑∞∟∫-∳⋀-⋃◸-◿♯⟘⟙⟀⟁⦰-⦴⨀-⨆⨉-⨖⨛⨜𝛁𝛛𝛻𝜕𝜵𝝏𝝯𝞉𝞩𝟃ⁱ-⁾₁-₎∠-∢⦛-⦯℘℮゛-゜𝟎-𝟡]|[^\P{Mn}-¡]|[^\P{Mc}-¡]|[^\P{Nd}-¡]|[^\P{Pc}-¡]|[^\P{Sk}-¡]|[^\P{Me}-¡]|[^\P{No}-¡]|[′-‷⁗]|[^\P{So}←-⇿]|[^\p{^Sc}$])*`, Name: "variable.interpolation.julia", Captures: map[string]rule{
				"1": {Name: metaPlaceholder},
			}},
			{Begin: `(\$)(\()`, End: `\)`, Name: "variable.interpolation.julia", BeginCaptures: map[string]rule{
				"1": {Name: metaPlaceholder},
				"2": {Name: metaPlaceholder + " meta.bracket.julia"},
			}, EndCaptures: map[string]rule{
				"0": {Name: metaPlaceholder + " meta.bracket.julia"},
			}, Patterns: []rule{
				{Include: "source.julia#self_no_for_block"},
			}},
		},
	},
	// The delimiters share meta.embedded.expression.dart with the content.
	// Copies of the two string-interp rules; operators and brackets in ${ }
	// have no scope there. The call rule takes the ( of a call into its match
	// without a scope; a copy of it gives the ( the scope of punctuation. A
	// string in ${ } is again a string with interpolation, so the selector
	// excludes ${ } only outside such a string.
	{
		file:     "dart.json",
		injectTo: []string{"source.dart"},
		selector: "L:string.interpolated - (comment | meta.embedded.expression.dart - (meta.embedded.expression.dart string.interpolated))",
		patterns: []rule{
			{Match: `(\$)([a-zA-Z0-9_]+)`, Name: "meta.embedded.expression.dart", Captures: map[string]rule{
				"1": {Name: metaPlaceholder},
				"2": {Name: "variable.parameter.dart"},
			}},
			{Begin: `\$\{`, End: `\}`, Name: "meta.embedded.expression.dart", BeginCaptures: map[string]rule{
				"0": {Name: metaPlaceholder},
			}, EndCaptures: map[string]rule{
				"0": {Name: metaPlaceholder},
			}, Patterns: []rule{
				{Match: `([_$]*[a-z][a-zA-Z0-9_$]*)(<(?:[a-zA-Z0-9_$<>?]|,\s*|\s+extends\s+)+>)?([!?]?\()`, Captures: map[string]rule{
					"1": {Name: "entity.name.function.dart"},
					"2": {Patterns: []rule{{Include: "source.dart#type-args"}}},
					"3": {Name: metaPunctuation},
				}},
				{Include: "source.dart#expression"},
				{Match: `[-+*/%~!?:<>=&|^.,()\[\]]`, Name: metaPunctuation},
			}},
		},
	},
	// this and super are dim only before a dot (also ?. and the cascade ..),
	// as the object of a member; alone this looks like a variable. this(...)
	// and super(...) call a constructor and get the style of a function.
	// Copies of the rule in constants-and-special-vars with a lookahead.
	{
		file:     "dart-this.json",
		injectTo: []string{"source.dart"},
		selector: "L:source.dart - (comment | (string - meta.embedded.expression))",
		patterns: []rule{
			{Match: `(?<!\$)\b(?:this|super)\b(?!\$)(?=\s*\??\.)`, Name: metaSelf + " variable.language.dart"},
			{Match: `(?<!\$)\b(?:this|super)\b(?!\$)(?=\s*\()`, Name: metaConstructor + " variable.language.dart"},
		},
	},
	// The rule for $name takes $ and the first letter. Copy of that rule. A dot
	// is part of $name only before a name: in "$major.$minor" it is text of
	// the string. Strings in """ """ have a scope of their own.
	{
		file:     "groovy.json",
		injectTo: []string{"source.groovy"},
		selector: "L:(string.quoted.double.groovy | string.quoted.double.multiline.groovy) - (comment | variable.other.interpolated.groovy | source.groovy.embedded)",
		patterns: []rule{
			{Begin: `(\$)\w`, End: `(?=\W)`, Name: "variable.other.interpolated.groovy", ApplyEndPatternLast: true, BeginCaptures: map[string]rule{
				"1": {Name: metaPlaceholder},
			}, Patterns: []rule{
				{Match: `\w`, Name: "variable.other.interpolated.groovy"},
				{Match: `\.(?=[A-Za-z_])`, Name: "keyword.other.dereference.groovy"},
			}},
		},
	},
	// The Groovy grammar gives the content of ${ } no scopes.
	{
		file:     "groovy-interpolation.json",
		injectTo: []string{"source.groovy"},
		selector: "source.groovy.embedded.source - comment",
		patterns: []rule{
			{Match: `\b\d[\d_]*(?:\.\d+)?[a-zA-Z]?\b`, Name: metaLiteral},
			{Match: `[-+*/%=!<>&|?:.,()\[\]]`, Name: metaPunctuation},
		},
	},
	// self is dim only before a dot, as the object of an attribute; alone it
	// looks like a variable, and cls before a dot is dim too. Copies of
	// special-variables with a lookahead. f-strings stay out, because their
	// text and their fields share one scope, and so do parameter lists of
	// functions and lambdas, where a theme rule gives self the color of a
	// parameter. self before a parenthesis or a bracket also stays out:
	// meta.function-call and meta.item-access start there with an empty match,
	// and an L: injection would win against it. Theme rules give the name of
	// the call the style of a function and the name of the item access the
	// style of a variable.
	{
		file:     "python.json",
		injectTo: []string{"source.python"},
		selector: "L:source.python - (comment | string | meta.fstring | meta.function.parameters | meta.function.lambda.parameters)",
		patterns: []rule{
			{Match: `\b(?<!\.)(self)\b(?!\s*[.(\[])`, Captures: map[string]rule{
				"1": {Name: metaName + " variable.language.special.self.python"},
			}},
			{Match: `\b(?<!\.)(cls)\b(?=\s*\.)`, Captures: map[string]rule{
				"1": {Name: metaSelf + " variable.language.special.cls.python"},
			}},
		},
	},
}

func (in injection) scopeName() string {
	return "monochromator.injection." + strings.TrimSuffix(in.file, ".json")
}

type grammarFile struct {
	ScopeName         string          `json:"scopeName"`
	InjectionSelector string          `json:"injectionSelector"`
	Patterns          []rule          `json:"patterns"`
	Repository        map[string]rule `json:"repository,omitempty"`
}

func writeGrammar(filePath string, in injection) error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	// The patterns contain < > and &, which the default escaping would write
	// as \u003c and so on.
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(grammarFile{ScopeName: in.scopeName(), InjectionSelector: in.selector, Patterns: in.patterns, Repository: in.repository})
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
