# Changelog

## v0.4.0 - 2026-09-22

- Made bold exclusive to structural keywords so that the keywords show the code structure again
    - Types, boolean values, escape sequences, and special strings are no longer bold
    - Constructors are now italic like the functions
    - Python class names in the class definition now have the text color
    - `this`, `self`, decorators, annotations, attributes, and HTML attribute names are now gray
- Made functions and methods brighter than variables to keep the overview in long call chains
- Dimmed the editor foreground and the UI text
    - Dark: `#e7e7e7` -> `#ccccce`
    - Light: `#000000` -> `#303034`
- Unified the icon color with the editor foreground
- Fixed the built-in constants such as `nil`, `None`, and `null`, which did not have the accent color
- Fixed the keyword operators such as `is`, `not`, and `or`, which were bold
- Fixed the Rust lifetimes, which had the text color
- Fixed the JSON keys and the YAML keys, which had the accent color
- Fixed the Markdown headings, the HTML tags, and the CSS selectors, which were not bold
- Fixed the CSS units, which were bold
- Fixed the JSDoc tags, which were not gray
- Fixed the Ruby symbols, which had the accent color

## v0.3.0 - 2026-07-15

- Darkened the dark theme background
    - Background: `#0f0f12` -> `#08080a`
- Tinted the text selections with each variant's accent color

## v0.2.2 - 2026-04-30

- Renamed the extension ID to `monochromator-theme` to follow the Zed registry conventions

## v0.2.1 - 2026-04-24

- Added the `LICENSE` symlink that the Zed extension registry requires

## v0.2.0 - 2026-04-24

- Improved the accent colors of the light themes for WCAG AA compliance
    - Aquamarine: `#01aacf` -> `#007b94`
    - Amber: `#e66100` -> `#bf4e00`
    - Sulfur: `#927e00` -> `#856e00`

## v0.1.0 - 2026-03-28

- First release
