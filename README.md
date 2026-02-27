<p align="center">
  <img src="./assets/icon.png" height=250 width=250 />
</p>

<h1 align="center">Monochromator</h1>

<p align="center">
  <a href="https://open-vsx.org/extension/beem/monochromator"><img src="https://img.shields.io/open-vsx/v/beem/monochromator?style=flat-square&label=Version&color=%23e7e7e7&labelColor=%23393941" alt="Version" /></a>
  <a href="https://open-vsx.org/extension/beem/monochromator"><img src="https://img.shields.io/open-vsx/dt/beem/monochromator?style=flat-square&label=Downloads&color=%23e7e7e7&labelColor=%23393941&logo=vscodium&logoColor=%23e7e7e7" alt="Downloads" /></a>
  <a href="https://codeberg.org/beem/monochromator/raw/branch/master/LICENSE"><img src="https://img.shields.io/badge/License-MIT-%23e7e7e7?style=flat-square&labelColor=%23393941" alt="License" /></a>
</p>

<p align="center">
  Distraction-free <b>monochrome</b> coding theme for <b>Visual Studio Code</b>
</p>

![Screenshots](./assets/screenshots.webp)

<details>
<summary>Recommended Settings</summary>

To ensure 100% correct behaviour, it is recommended to add the following to your `settings.json`:

```json
{
  // This ensures that string interpolation delimiters look right
  "editor.bracketPairColorization.enabled": false,

  // This ensures that variables do not look similar to comments
  "editor.showUnused": false
}
```
</details>

## Installation

### Download

- [Open VSX](https://open-vsx.org/extension/beem/monochromator)
- [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=beem.monochromator)
- [Actions](https://codeberg.org/beem/monochromator/actions) (latest build)
- Directly in VSCodium or VS Code

### Building From Source

```sh
git clone https://codeberg.org/beem/monochromator.git
cd monochromator
npm install
npx vsce package
```

## Contributing

As there are many programming and markup languages (each with a lot of different extensions) out there, it's practically impossible to verify how the highlighting works for each one of them. If you spot something that looks off with the syntax highlighting while using this theme, I'd really appreciate your help :-)

You can help by [opening an issue](https://codeberg.org/beem/monochromator/issues). When you do, please include:

- Which language you were using (and extension if there's no built-in language support)
- What looks wrong to you
- A screenshot if you can
- How (and why) you think it should look instead

## License

Copyright © 2025 Josias Beem

This project is licensed under the **MIT License** - see the `LICENSE` file for details.
