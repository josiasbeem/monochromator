package main

import (
	"fmt"
	"os"
	"path/filepath"

	"generate/helix"
	"generate/neovim"
	"generate/palette"
	"generate/vscode"
	"generate/zed"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// All paths assume that the generator runs inside the generate/ directory.
	repoRoot := ".."

	cfg, err := palette.Load(filepath.Join(repoRoot, "palette"))
	if err != nil {
		return err
	}

	if err := vscode.Generate(cfg, filepath.Join(repoRoot, "editors", "vscode", "themes")); err != nil {
		return err
	}
	fmt.Println("Generated VS Code themes")

	if err := vscode.GenerateInjections(filepath.Join(repoRoot, "editors", "vscode", "syntaxes")); err != nil {
		return err
	}
	fmt.Println("Generated VS Code injection grammars")

	if err := zed.Generate(cfg, filepath.Join(repoRoot, "editors", "zed", "themes")); err != nil {
		return err
	}
	fmt.Println("Generated Zed themes")

	if err := helix.Generate(cfg, filepath.Join(repoRoot, "editors", "helix", "themes")); err != nil {
		return err
	}
	fmt.Println("Generated Helix themes")

	if err := neovim.Generate(cfg, filepath.Join(repoRoot, "editors", "neovim", "colors")); err != nil {
		return err
	}
	fmt.Println("Generated Neovim color schemes")

	return nil
}
