# Monochromator

<p align="center">
  <img src="./assets/icon.png" height=250 width=250 />
</p>

<p align="center">
  <a href="https://open-vsx.org/extension/beem/monochromator"><img src="https://img.shields.io/open-vsx/v/beem/monochromator?label=Version&color=%23ffffff" alt="Version" /></a>
  <a href="https://open-vsx.org/extension/beem/monochromator"><img src="https://img.shields.io/open-vsx/dt/beem/monochromator?label=Downloads&color=%23ffffff" alt="Downloads" /></a>
  <a href="https://codeberg.org/beem/monochromator/raw/branch/master/LICENSE"><img src="https://img.shields.io/badge/License-MIT-%23ffffff" alt="License" /></a>
</p>

<p align="center">
  A simple <b>monochrome</b> theme for <b>Visual Studio Code</b>
</p>

---

A minimalist color theme that reduces visual distractions for better focus while coding. The extension includes three main variants: **Dark** and **Light**. Each variant comes with several additional accent color options: **Ruby**, **Amber**, **Sulfur**, **Emerald**, **Aquamarine**, **Amethyst**, **Plain**.

---

![Screenshots](./assets/screenshots.webp)

*Font: Fira Code*

## Installation

To ensure 100% correct behaviour, it is recommended to add the following to your `settings.json`:

```
"editor.bracketPairColorization.enabled": false
```

### Download

The extension can be downloaded from [Open VSX](https://open-vsx.org/extension/beem/monochromator), the [Actions](https://codeberg.org/beem/monochromator/actions) tab or directly in VSCodium.

### Building from source

```sh
git clone https://codeberg.org/beem/monochromator.git
cd monochromator
npm install
npx vsce package
code --install-extension monochromator-*.vsix
```

⚠️ **Note:** Depending on your [Code - OSS](https://github.com/microsoft/vscode) distribution, the command `code` could also be named ...
- `codium` (VSCodium)
- `code-insiders` (Visual Studio Code Insiders)
- `codium-insiders` (VSCodium Insiders).

## Contributing

As there are so many programming and markup languages out there, it's practically impossible to verify how the highlighting works for each one of them. If you spot something that looks off with the syntax highlighting while using this theme, I'd really appreciate your help:

- **Open an issue** on the Codeberg repo

When you do, it would be super helpful if you could include:

- Which language you were using
- What looks wrong to you
- A screenshot if you can
- How (and why) you think it should look instead

## License

Copyright © 2025 Josias Beem

This project is licensed under the **MIT License** - see the `LICENSE` file for details.
