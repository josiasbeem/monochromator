package vscode

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"generate/palette"
)

type packageGrammar struct {
	ScopeName string   `json:"scopeName"`
	Path      string   `json:"path"`
	InjectTo  []string `json:"injectTo"`
}

// readJSON fails the test when the file is missing or is no valid JSON.
func readJSON(t *testing.T, filePath string, value any) {
	t.Helper()
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, value); err != nil {
		t.Fatalf("%s: %v", filePath, err)
	}
}

// monochromatorScopes collects every scope with the suffix .monochromator in
// the names and content names of a grammar.
func monochromatorScopes(value any, found map[string]bool) {
	switch v := value.(type) {
	case map[string]any:
		for _, key := range []string{"name", "contentName"} {
			name, _ := v[key].(string)
			for _, scope := range strings.Fields(name) {
				if strings.HasSuffix(scope, ".monochromator") {
					found[scope] = true
				}
			}
		}
		for _, child := range v {
			monochromatorScopes(child, found)
		}
	case []any:
		for _, child := range v {
			monochromatorScopes(child, found)
		}
	}
}

func TestInjectionGrammars(t *testing.T) {
	grammarDir := t.TempDir()
	if err := GenerateInjections(grammarDir); err != nil {
		t.Fatal(err)
	}
	cfg, err := palette.Load(filepath.Join("..", "..", "palette"))
	if err != nil {
		t.Fatal(err)
	}
	themeDir := t.TempDir()
	if err := Generate(cfg, themeDir); err != nil {
		t.Fatal(err)
	}

	var manifest struct {
		Contributes struct {
			Grammars []packageGrammar `json:"grammars"`
		} `json:"contributes"`
	}
	readJSON(t, filepath.Join("..", "..", "editors", "vscode", "package.json"), &manifest)
	listed := map[string]packageGrammar{}
	for _, grammar := range manifest.Contributes.Grammars {
		listed[grammar.Path] = grammar
	}

	var theme struct {
		TokenColors []struct {
			Scope []string `json:"scope"`
		} `json:"tokenColors"`
	}
	readJSON(t, filepath.Join(themeDir, "dark", "default.json"), &theme)
	styled := map[string]bool{}
	for _, tokenRule := range theme.TokenColors {
		for _, selector := range tokenRule.Scope {
			for _, scope := range strings.Fields(selector) {
				styled[scope] = true
			}
		}
	}

	files, err := os.ReadDir(grammarDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != len(listed) {
		t.Errorf("package.json lists %d grammars, the generator writes %d", len(listed), len(files))
	}
	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			var grammar map[string]any
			readJSON(t, filepath.Join(grammarDir, file.Name()), &grammar)

			entry, ok := listed["./syntaxes/"+file.Name()]
			if !ok {
				t.Fatalf("package.json does not list ./syntaxes/%s", file.Name())
			}
			if entry.ScopeName != grammar["scopeName"] {
				t.Errorf("package.json scopeName %q, grammar %q", entry.ScopeName, grammar["scopeName"])
			}
			for _, in := range injections {
				if in.file == file.Name() && !reflect.DeepEqual(entry.InjectTo, in.injectTo) {
					t.Errorf("package.json injectTo %v, injection %v", entry.InjectTo, in.injectTo)
				}
			}

			// A scope without a theme rule would leave the token as it is;
			// the markers have no style on purpose.
			found := map[string]bool{}
			monochromatorScopes(grammar, found)
			for scope := range found {
				if scope != metaRawString && scope != metaTypeOperator && scope != metaEnumBody && !styled[scope] {
					t.Errorf("%s has no rule in the theme", scope)
				}
			}
		})
	}
}
