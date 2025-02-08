# Monochromator

A simple **monochrome** theme for VS Code. It is designed so that it also looks **harmonious** with the **GNOME** desktop environment. Following **two variants** are included:

- `Monochromator Dark`
- `Monochromator Light`

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
