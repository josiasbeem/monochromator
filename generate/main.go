package main

import (
	"fmt"
	"os"
	"path/filepath"

	"generate/palette"
	"generate/vscode"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	repoRoot := ".."

	cfg, err := palette.Load(filepath.Join(repoRoot, "palette"))
	if err != nil {
		return err
	}

	if err := vscode.Generate(cfg, filepath.Join(repoRoot, "editors", "vscode", "themes")); err != nil {
		return err
	}
	fmt.Println("Generated VS Code themes")

	return nil
}
