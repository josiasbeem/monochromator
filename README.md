# Monochromator

A simple **monochrome** theme for VS Code. It is designed so that it also looks **harmonious** with the **GNOME** desktop environment. Following **four variants** are included:

- `Monochromator Light`
- `Monochromator Dark`
- `Monochromator Light (High Contrast)`
- `Monochromator Dark (High Contrast)`

## Installation

In order to build the `.vsix` file, you'll need to make sure that the **npm package** `vsce` is installed. If not, do the following:

```sh
npm install -g vsce
```

Afterwards ...

```sh
git clone https://codeberg.org/beem/monochromator.git
cd monochromator
vsce package
```
