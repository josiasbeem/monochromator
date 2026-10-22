# Changelog

This file shows all important changes to the Monochromator extension for Zed.

The format of this file follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and the version numbers follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.0] - 2026-09-22

### Changed

- Only structural keywords are bold now, so they show the code structure again
    - Types, boolean values, escape sequences, and special strings are no longer bold
    - Constructors are now italic like the functions
    - Python class names in the class definition now have the text color
    - `this`, `self`, decorators, annotations, attributes, and HTML attribute names are now gray
- Functions and methods are now brighter than variables, which makes long call chains easier to read
- Dimmer editor foreground and UI text
    - Dark: `#e7e7e7` -> `#ccccce`
    - Light: `#000000` -> `#303034`
- The icon color now matches the editor foreground

### Fixed

- Built-in constants such as `nil`, `None`, and `null` did not have the accent color
- Keyword operators such as `is`, `not`, and `or` were bold
- Rust lifetimes had the text color
- JSON keys and YAML keys had the accent color
- Markdown headings, HTML tags, and CSS selectors were not bold
- CSS units were bold
- JSDoc tags were not gray
- Ruby symbols had the accent color

## [0.3.0] - 2026-07-15

### Changed

- Darker background in the dark theme
    - Background: `#0f0f12` -> `#08080a`
- Text selections are now tinted with each variant's accent color

## [0.2.2] - 2026-04-30

### Changed

- Renamed the extension ID to `monochromator-theme` to follow the Zed registry conventions

## [0.2.1] - 2026-04-24

### Added

- `LICENSE` symlink that the Zed extension registry requires

## [0.2.0] - 2026-04-24

### Changed

- Light theme accent colors now meet WCAG AA
    - Aquamarine: `#01aacf` -> `#007b94`
    - Amber: `#e66100` -> `#bf4e00`
    - Sulfur: `#927e00` -> `#856e00`

## [0.1.0] - 2026-03-28

### Added

- First release

[0.4.0]: https://github.com/josiasbeem/monochromator/compare/zed/v0.3.0...zed/v0.4.0
[0.3.0]: https://github.com/josiasbeem/monochromator/compare/zed/v0.2.2...zed/v0.3.0
[0.2.2]: https://github.com/josiasbeem/monochromator/compare/zed/v0.2.1...zed/v0.2.2
[0.2.1]: https://github.com/josiasbeem/monochromator/compare/zed/v0.2.0...zed/v0.2.1
[0.2.0]: https://github.com/josiasbeem/monochromator/compare/zed/v0.1.0...zed/v0.2.0
[0.1.0]: https://github.com/josiasbeem/monochromator/tree/zed/v0.1.0
