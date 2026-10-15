<p align="center">
  <img src="./assets/icon.png" height=250 width=250 />
</p>

<h1 align="center">Monochromator</h1>

<p align="center">
  <a href="https://codeberg.org/beem/monochromator/issues"><img src="https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fcodeberg.org%2Fapi%2Fv1%2Frepos%2Fbeem%2Fmonochromator&query=%24.open_issues_count&style=flat-square&label=Issues&color=%23e7e7e7&labelColor=%23393941&logo=codeberg&logoColor=%23e7e7e7" alt="Issues" /></a>
  <a href="https://open-vsx.org/extension/beem/monochromator"><img src="https://img.shields.io/open-vsx/dt/beem/monochromator?style=flat-square&label=Downloads&color=%23e7e7e7&labelColor=%23393941&logo=vscodium&logoColor=%23e7e7e7" alt="Downloads" /></a>
  <a href="https://codeberg.org/beem/monochromator/raw/branch/master/LICENSE"><img src="https://img.shields.io/badge/License-MIT-%23e7e7e7?style=flat-square&labelColor=%23393941&logo=opensourceinitiative&logoColor=%23e7e7e7" alt="License" /></a>
</p>

<p align="center">
  Distraction-free <b>monochrome</b> coding theme
</p>

![Screenshots](./assets/screenshots.webp)

## Installation

You can also get all the themes from [Actions](https://codeberg.org/beem/monochromator/actions).

<details>
<summary><img src="https://custom-icon-badges.demolab.com/badge/VS%20Code-0078d7.svg?logo=vsc&logoColor=white&style=flat-square" alt="VS Code" /></summary>

<br>

Install the extension from [Open VSX](https://open-vsx.org/extension/beem/monochromator), the [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=beem.monochromator), or directly in your editor.

To make sure that the theme works correctly, add these settings to your `settings.json`:

```json
{
  // This makes sure that string interpolation delimiters look correct
  "editor.bracketPairColorization.enabled": false,

  // This makes sure that variables do not look similar to comments
  "editor.showUnused": false
}
```

</details>

<details>
<summary><img src="https://img.shields.io/badge/Zed-white?logo=zedindustries&logoColor=084CCF&style=flat-square" alt="Zed" /></summary>

<br>

Install the theme from the [Zed extension registry](https://zed.dev/extensions/monochromator-theme). As an alternative, open the extensions view (`zed: extensions`) in your editor and search for **Monochromator**.

</details>

<details>
<summary><img src="https://img.shields.io/badge/Helix-706bc8?logo=helix&logoColor=fff&style=flat-square" alt="Helix" /></summary>

<br>

Copy the theme files from [`editors/helix/themes`](./editors/helix/themes) into the `themes` directory of your Helix configuration:

- Linux and macOS: `~/.config/helix/themes`
- Windows: `%AppData%\helix\themes`

Then set the theme in your `config.toml`. Helix uses the file name as the theme name, for example `monochromator-light` or `monochromator-dark-amber`.

```toml
theme = "monochromator-dark"

[editor]
# This makes sure that Helix uses the theme if it does not detect the true color support of your terminal
true-color = true
```

</details>

<details>
<summary><img src="https://img.shields.io/badge/Neovim-57A143?logo=neovim&logoColor=fff&style=flat-square" alt="Neovim" /></summary>

<br>

**You must install the tree-sitter parsers for your languages.**

Neovim includes parsers for only a small number of languages, for example Lua and Markdown. For all other languages, Neovim uses its old syntax highlighting. This highlighting does not show functions, types, and punctuation. For this reason, the color schemes cannot show the structure of your code.

The parsers come from [nvim-treesitter](https://github.com/nvim-treesitter/nvim-treesitter). nvim-treesitter needs the `tree-sitter` command and a C compiler.

This example `init.lua` uses `vim.pack`, the built-in plugin manager of Neovim 0.12 and later. It installs the color schemes, nvim-treesitter, and the parsers:

```lua
-- The color schemes are in a subdirectory of the repository
vim.pack.add({ "https://codeberg.org/beem/monochromator" }, {
  load = function(plugin)
    vim.opt.rtp:append(plugin.path .. "/editors/neovim")
  end,
})

-- This makes sure that Neovim uses the theme if it does not detect the true color support of your terminal
vim.o.termguicolors = true
vim.cmd.colorscheme("monochromator-dark")

vim.pack.add({ "https://github.com/nvim-treesitter/nvim-treesitter" })
-- Put the languages that you use into the list
require("nvim-treesitter").install({ "lua", "go", "python", "rust" })
vim.api.nvim_create_autocmd("FileType", {
  callback = function(args)
    pcall(vim.treesitter.start, args.buf)
  end,
})
```

Neovim uses the file name as the color scheme name, for example `monochromator-light` or `monochromator-dark-amber`.

With [lazy.nvim](https://github.com/folke/lazy.nvim), install nvim-treesitter as its README tells you. Then use this plugin specification for the color schemes:

```lua
{
  url = "https://codeberg.org/beem/monochromator",
  lazy = false,
  priority = 1000,
  config = function(plugin)
    vim.opt.rtp:append(plugin.dir .. "/editors/neovim")
    vim.o.termguicolors = true
    vim.cmd.colorscheme("monochromator-dark")
  end,
}
```

As an alternative, copy the color scheme files from [`editors/neovim/colors`](./editors/neovim/colors) into the `colors` directory of your Neovim configuration:

- Linux and macOS: `~/.config/nvim/colors`
- Windows: `%LocalAppData%\nvim\colors`

Then add the `termguicolors` line and the `colorscheme` line of the example to your `init.lua`.

</details>

## Font

[Monochromator Mono](https://codeberg.org/beem/monochromator-mono) is the matching code font, and the screenshots use it. You do not need it, because the theme works with every font. But a monochrome theme shows the structure of the code with font styles, not with colors. For this reason, this font makes the difference between the font styles much larger than in most code fonts:

- Its bold cuts are much heavier than the bold cuts of most code fonts.
- Its italic cuts look much more different from the upright text than in most code fonts.

<details>
<summary><img src="https://custom-icon-badges.demolab.com/badge/VS%20Code-0078d7.svg?logo=vsc&logoColor=white&style=flat-square" alt="VS Code" /></summary>

<br>

Add these settings to your `settings.json`:

```json
{
  "editor.fontFamily": "Monochromator Mono",

  // This makes sure that the code ligatures stay on
  "editor.fontLigatures": true
}
```

</details>

<details>
<summary><img src="https://img.shields.io/badge/Zed-white?logo=zedindustries&logoColor=084CCF&style=flat-square" alt="Zed" /></summary>

<br>

Add this setting to your `settings.json`:

```json
{
  "buffer_font_family": "Monochromator Mono"
}
```

</details>

<details>
<summary><img src="https://img.shields.io/badge/Helix-706bc8?logo=helix&logoColor=fff&style=flat-square" alt="Helix" /></summary>

<br>

Helix uses the font of your terminal. Set **Monochromator Mono** as the font in the settings of your terminal.

</details>

<details>
<summary><img src="https://img.shields.io/badge/Neovim-57A143?logo=neovim&logoColor=fff&style=flat-square" alt="Neovim" /></summary>

<br>

Neovim uses the font of your terminal. Set **Monochromator Mono** as the font in the settings of your terminal.

</details>

## How to contribute

If the syntax highlighting is not correct, [open an issue](https://codeberg.org/beem/monochromator/issues). Include the editor, the language, a screenshot, and a description of the correct appearance.

My main focus is the Visual Studio Code variant. The quality in other editors can vary.

## License

Copyright © 2025 Josias Beem

This project uses the **MIT License**
