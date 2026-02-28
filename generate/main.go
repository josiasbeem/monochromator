package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"generate/palette"
	"generate/vscode"
)

func main() {
	root := flag.String("root", "..", "Path to the repository root")
	flag.Parse()

	repoRoot := *root
	paletteDir := filepath.Join(repoRoot, "palette")

	cfg, err := palette.Load(paletteDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading palette: %v\n", err)
		os.Exit(1)
	}

	vsDir := filepath.Join(repoRoot, "editors", "vscode", "themes")
	if err := vscode.Generate(cfg, vsDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating VS Code themes: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Generated VS Code themes.")
}
