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

## Font

The screenshots use [Monochromator Mono](https://codeberg.org/beem/monochromator-mono), the matching code font. The font is a fork of Recursive with the code ligatures on by default. The theme works with every font. This font is optional.

Install the font with the font manager of your system. Then set it in your editor.

VS Code (`settings.json`):

```json
{
  "editor.fontFamily": "Monochromator Mono",

  // This makes sure that the code ligatures stay on
  "editor.fontLigatures": true
}
```

Zed (`settings.json`):

```json
{
  "buffer_font_family": "Monochromator Mono"
}
```

## How to contribute

If the syntax highlighting is not correct, [open an issue](https://codeberg.org/beem/monochromator/issues). Include the editor, the language, a screenshot, and a description of the correct appearance.

My main focus is the Visual Studio Code variant. The quality in other editors can vary.

## License

Copyright © 2025 Josias Beem

This project uses the **MIT License**
