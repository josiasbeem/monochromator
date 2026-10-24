# Changelog

This file shows all important changes to the Monochromator extension for VS Code.

The format of this file follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and the version numbers follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.36.0] - 2026-10-24

### Added

- Injection grammars for 19 languages, to fix scopes that the built-in grammars get wrong or do not set: Bash, C, C#, C++, Dart, Groovy, Java, JavaScript, Julia, LaTeX, Lua, PHP, PowerShell, Python, Ruby, Rust, Swift, TypeScript, and Zig

### Changed

- Darker accent colors in the light themes. Literals now have at least 6.8:1 contrast on the cursor line, and color-blind readers can tell them apart from comments
- The light Amethyst theme now uses the violet of the dark theme
- Brighter red in the dark themes (`#fe303a` -> `#ff4a4b`)
- Quieter cursor line in the dark themes (`#202024` -> `#161619`)
- More even blue tint in the grays, with the same lightness
- `this` and `self` are now gray only before a dot or an arrow (`this.name`, `self.x`, `this->x`). Alone or as a parameter, they look like a variable
    - Calls such as `super()`, `this()`, `base()`, `cls()`, and `new this()` are now italic like other calls
    - `Self` in Swift is now underlined like a type, as in Rust
- Word operators such as `and`, `or`, `not`, `in`, `is`, `as`, and `instanceof` are now bold like other keywords
- Words written like a call are now italic: `sizeof`, `typeof`, `nameof`, `alignof`, `alignas`, `typeid`, and `import()`
- Names in string interpolation now look like the same names outside the string
- The delimiters and format specs of interpolation, such as `${`, `#{`, `$(`, and `:>8`, are now underlined in the accent color, also the `$` before a name
- Type names are now underlined also in their declaration, for example `class Box`, `enum Kind`, and Rust structs and traits
- `null`, `undefined`, `true`, and `false` in TypeScript types now look like literals, and `this` in a type is gray
- Rust enum variants such as `Some`, `Ok`, and your own variants are now italic like constructors
- Python docstrings now look like comments. Doctest prompts (`>>>`, `...`) and code tags such as `TODO` and `FIXME` are bold, like doc tags
- Diff lines are now green and red
- Prefixes and suffixes of numbers, such as `0x` and `j` in Python and `n` in JavaScript, now have the literal color
- Regex escapes such as `\d` in JavaScript and TypeScript are now underlined in the accent color
- Test operators in Bash and PowerShell, such as `-f`, `-eq`, and `-and`, now use the text color, like other flags
- The extension now links to the repository on GitHub

### Fixed

- Keys in YAML and INI, namespaces, and variables in Typst were not in the text color
- Some parts of doc comments did not look like comments: type names in Javadoc, type braces in JSDoc, attribute values in C# doc comments, and types in LuaDoc
- Tags in C# doc comments and LuaDoc did not look like doc tags
- Brackets in Julia, TSX, and Zig, regex anchors, the YAML folded block indicator, the Typst `#`, the Makefile `@`, the PowerShell `@` in `@(` and `@{`, and Bash heredoc delimiters were not gray
- String prefixes in Python and Rust (`f`, `r`, `b`) were underlined
- The Markdown escape (`\*`) was not underlined in the accent color
- CSS pseudo-elements, the C# `event` keyword, and TOML array table names were not bold
- Integer types, C types, and type keywords in Zig were not underlined
- Built-in Python decorators such as `@property` were not gray like other decorators
- Automatic variables in PowerShell strings did not have the text color

## [0.35.0] - 2026-10-02

### Changed

- Replaced more default colors with palette colors, for example in the symbol icons, the peek view, and the debugger
- The hover in lists and tabs now differs from the selection
- The current find match and the resize edges are now highlighted

### Fixed

- Unreadable status bar warnings and errors in the dark themes

## [0.34.0] - 2026-09-18

### Changed

- Only structural keywords are bold now, so they show the code structure again
    - Types are now underlined instead of bold
    - Escape sequences and format placeholders are now underlined instead of bold
    - Language constants such as `true`, `false`, `nil`, and `null` are no longer bold
    - `this`, `self`, `super`, decorators, annotations, and attributes are now gray
- Functions and methods are now brighter than variables, which makes long call chains easier to read
- Dimmer editor foreground
    - Dark: `#e7e7e7` -> `#ccccce`
    - Light: `#000000` -> `#303034`
- The UI foreground now matches the editor foreground

### Fixed

- `delete` and the named casts in C++ were gray while `new` was bold
- Function calls in Python were not italic
- Base classes and interfaces after `extends` and `implements` did not have the type style
- Digit separators and exponent signs in number literals were gray

## [0.33.1] - 2026-09-01

### Fixed

- Unreadable extension action buttons in the light themes

## [0.33.0] - 2026-07-15

### Changed

- Darker background in the dark theme
    - Background: `#0f0f12` -> `#08080a`
    - To keep the previous value, set `"workbench.colorCustomizations": { "editor.background": "#0f0f12", "terminal.background": "#0f0f12" }` in your `settings.json`
- Text selections are now tinted with each variant's accent color

## [0.32.0] - 2026-04-14

### Changed

- Light theme accent colors now meet WCAG AA
    - Aquamarine: `#01aacf` -> `#007b94`
    - Amber: `#e66100` -> `#bf4e00`
    - Sulfur: `#927e00` -> `#856e00`

## [0.31.2] - 2026-04-12

### Fixed

- File highlight issue in the file explorer ([#8](https://codeberg.org/beem/monochromator/issues/8))

## [0.31.1] - 2026-02-28

### Changed

- The scrollbar is now transparent, so it no longer covers the breadcrumbs ([#5](https://codeberg.org/beem/monochromator/issues/5))
- Replaced rarely used color values with similar existing ones for consistency

### Removed

- Unnecessary migration code from v0.27.0

## [0.31.0] - 2026-02-26

### Changed

- Dimmer white foreground in the dark theme
- Improved Typst support

### Fixed

- Border colors of the editor group, the panel, and the overview ruler

## [0.30.0] - 2026-01-14

### Changed

- Functions and methods are now italic
- Darker gray in the dark theme

### Fixed

- Highlighting of Go package names
- Highlighting of JavaScript braces

## [0.29.0] - 2025-11-19

### Changed

- Subtler command center border

## [0.28.1] - 2025-11-18

### Fixed

- Editor gutter colors for staged code

## [0.28.0] - 2025-11-04

### Changed

- Various UI improvements:
    - No more visible UI borders, for a cleaner look
    - Improved button contrast
    - More contrast between the editor and the UI, for better focus on the code
    - Higher contrast for comments in the light theme
    - Improved keybinding labels
    - No more focus borders

## [0.27.1] - 2025-10-21

### Fixed

- Highlighting of built-in shell commands

## [0.27.0] - 2025-10-18

### Changed

- Renamed the `Dark+` themes to `Dark`
- Various UI improvements
- Redesigned light theme
- Improved cyan and magenta, also in the terminal
- Improved inlay hints

### Fixed

- Highlighting of TypeScript braces

## [0.26.1] - 2025-10-09

### Fixed

- Highlighting of CSS ID selectors

## [0.26.0] - 2025-10-06

### Changed

- Less blue tint in the Dark+ backgrounds

## [0.25.4] - 2025-09-29

### Fixed

- Highlighting of BibTeX files

## [0.25.3] - 2025-09-16

### Fixed

- Command center colors
- Too weak blue tint in the Dark+ themes
- Border and foreground colors of hover widgets
- Header color of the multi diff editor
- Activity bar icon colors that did not fit the status bar well

## [0.25.2] - 2025-09-15

### Fixed

- Status bar colors

## [0.25.1] - 2025-08-20

### Changed

- Improved highlighting of Nix function arguments (the `{`, `}`, and `_` characters)

## [0.25.0] - 2025-08-11

### Added

- `Aquamarine` theme variant with a cyan accent color

### Changed

- Git diff editor colors are now part of the themes
- The orange terminal color is now yellow
- The cyan terminal color now fits the new `Aquamarine` variant

### Fixed

- Purple terminal color in the light themes

## [0.24.0] - 2025-08-10

### Added

- `Dark+` theme variants (the `Dark` theme with improved contrast)

### Changed

- Terminal colors now fit the themes better

## [0.23.3] - 2025-07-22

### Fixed

- String storage type in Go

## [0.23.2] - 2025-07-22

### Fixed

- Highlighting of the assignment operator

## [0.23.1] - 2025-07-20

### Fixed

- Highlighting of char literals (single quotes) in Rust
- Highlighting of the assignment operator in Rust

## [0.23.0] - 2025-07-18

### Added

- Sulfur variant with a yellow accent color:
    - `Monochromator Dark Sulfur`
    - `Monochromator Light Sulfur`
- Background colors for status bar errors and warnings

## [0.22.0] - 2025-07-15

### Changed

- More readable colors:
    - Purple in the dark themes: `#c61aff` -> `#ee40ff`
    - Green in the light themes: `#26a269` -> `#017f45`

### Fixed

- "Plain" variants missed the changes from `v0.18.0`

## [0.21.0] - 2025-07-04

### Added

- x86_64, ARM, RISC-V, and MIPS support for these OpenVSX extensions:
    - x86_64: `13xforever.language-x86-64-assembly`, `basdp.language-gas-x86`
    - ARM: `dan-c-underwood.arm`
    - RISC-V: `sunshaoce.RISC-V`
    - MIPS: `kdarkhan.mips`

## [0.20.1] - 2025-06-16

### Fixed

- Markdown backticks in the accent color variants

## [0.20.0] - 2025-06-13

### Changed

- Improved Markdown and LaTeX highlighting
- Improved highlighting of JavaScript string interpolation

## [0.19.0] - 2025-06-12

### Added

- Amethyst variant with a purple accent color:
    - `Monochromator Dark Amethyst`
    - `Monochromator Light Amethyst`

## [0.18.0] - 2025-06-11

### Changed

- More visible blue
- Improved string highlighting
- Lower minimum VS Code engine version

### Fixed

- C++ pointers
- `self` and entity names in Rust
- Enums and function arrows in Java
- Constants in Haskell, Perl, Ruby, PHP, and other languages
- Support functions in Lua, Julia, and other languages
- Strings and symbols in F#
- Null and booleans in Zig

## [0.17.2] - 2025-06-05

### Fixed

- Doxygen comments for languages other than C++

## [0.17.1] - 2025-06-03

### Fixed

- Code Actions menu

## [0.17.0] - 2025-05-25

### Added

- `Monochromator Dark Plain` and `Monochromator Light Plain`

### Changed

- Redesigned extension icon

### Fixed

- Highlighting of backslashes before LaTeX keywords

## [0.16.0] - 2025-05-24

### Changed

- Only keywords are bold now (functions are no longer bold)
- String quotation marks now match the color of literal text
- Preprocessor directives look like before `0.15.0` again
- YAML files look better
- Documentation keywords are now italic

### Fixed

- Highlighting of hex codes in HTML

## [0.15.2] - 2025-05-19

### Fixed

- Highlighting of SQL database and table names

## [0.15.1] - 2025-05-17

### Fixed

- Highlighting of `storage.modifier.sql`

## [0.15.0] - 2025-05-10

### Changed

- Various quality improvements:
    - Documentation keywords are now bold and gray (for example `@Override` in Java and shebangs)
    - Preprocessor directives are now bold and gray
    - Improved highlighting of escaped characters
    - HTML and CSS tags are now bold

### Fixed

- LaTeX math mode
- Highlighting of more constants

## [0.14.1] - 2025-05-08

### Fixed

- Highlighting of units (such as "px" in CSS)
- Highlighting of variables in shell scripts and JavaScript
- Highlighting of some punctuation in Java and R
- Regex highlighting

## [0.14.0] - 2025-05-04

### Added

- Amber, Emerald, and Ruby variants:
    - `Monochromator Dark Amber`
    - `Monochromator Dark Emerald`
    - `Monochromator Dark Ruby`
    - `Monochromator Light Amber`
    - `Monochromator Light Emerald`
    - `Monochromator Light Ruby`

### Changed

- Timestamps now look like literals
- Less contrast for `editor.wordHighlightStrongBackground`

## [0.13.3] - 2025-05-04

### Fixed

- Highlighting of preprocessor names and hexadecimal numbers

## [0.13.2] - 2025-04-22

### Fixed

- Highlighting in Python, TypeScript, LaTeX, YAML, and TOML

## [0.13.1] - 2025-04-21

### Removed

- Redundant code

### Fixed

- Selection highlighting
- Various inconsistencies

## [0.13.0] - 2025-04-18

### Changed

- `scrollbarSlider.background` now matches `editorGroupHeader.tabsBackground`
- `textCodeBlock.background` now matches `textBlockQuote.background`
- New icon design
- Colors now match https://gitlab.gnome.org/GNOME/libadwaita/-/merge_requests/1317
    - `#1e1e1e` -> `#1d1d20`
    - `#242424` -> `#222226`
    - `#303030` -> `#2e2e32`
    - `#3f3f3f` -> `#424246`
    - `#dfdfdf` -> `#d9d9da`
    - `#ebebeb` -> `#ebebed`
    - `#f5f5f5` -> `#f2f2f4`

### Fixed

- Selection highlight color in the minimap

## [0.12.0] - 2025-04-16

### Added

- Badges in `README.md`

### Changed

- Darker blue in `Monochromator Dark`
- Less bold text
- Selections are now part of the theme
- Improved Markdown preview
- New icon design

### Removed

- Redundant code

### Fixed

- Highlighting of constants, for example in Rust
- Lightbulb icon showed the wrong shade of gray

## [0.11.1] - 2025-04-10

### Fixed

- Tab color
- Screenshots

## [0.11.0] - 2025-04-10

### Added

- Color for `gitDecoration.addedResourceForeground`

### Changed

- Redesigned UI

## [0.10.3] - 2025-03-24

### Changed

- Less contrast for inlay hints

### Fixed

- Dropdown background color

## [0.10.2] - 2025-03-23

### Fixed

- Colors of:
    - Scrollbar slider
    - Panel title
    - Fold background

## [0.10.1] - 2025-03-21

### Changed

- Replaced `icon.png` with `icon128.png` and `icon1024.png`
- Screenshots are now in one file
- Improved `README.md`

### Fixed

- Changelog format

## [0.10.0] - 2025-03-21

### Added

- Bold for `constant.language` only

### Changed

- Higher overall contrast
- Improved look of:
    - Activity bar
    - Command palette
    - Keybindings

### Fixed

- Menu selection color in the light theme

## [0.9.0] - 2025-03-21

### Changed

- Theme colors for more components:
    - Checkboxes
    - Breadcrumbs
    - Git decoration for ignored files
    - Indent guides
    - Sticky scroll shadow
    - Menu (it opens when you right-click)
    - Widgets
    - Notifications
- Simpler changelog format
- Semantic highlighting is now enabled
- Reverted the `v0.8.4` changes

### Removed

- Bold for constants

### Fixed

- Missing hashtag in a hex code

## [0.8.4] - 2025-03-17

### Fixed

- Inactive title bar background color
    - `titleBar.inactiveBackground` from `#ebebeb` to `#f0f0f0` in the light theme
    - `titleBar.inactiveBackground` from `#303030` to `#2a2a2a` in the dark theme

## [0.8.3] - 2025-03-07

### Added

- Download section in `README.md`
- Bugs section in `package.json`

### Changed

- Renamed `logo.png` to `icon.png`
- Screenshots are now in their own directory
- Restructured `README.md`

### Fixed

- Format of `CHANGELOG.md`

## [0.8.2] - 2025-03-06

### Changed

- The info color now matches the theme

## [0.8.1] - 2025-03-02

### Fixed

- Broken screenshots in `README.md`

## [0.8.0] - 2025-03-02

### Added

- `CHANGELOG.md`
- More information, a "License" section, and screenshots in `README.md`
- Theme colors for Markdown in the editor

### Changed

- The logo font is now "IBM Plex Mono" instead of "Fira Code"
- Subtler progress bars
- URLs are now blue for better readability
- More readable `textBlockQuote` in the Markdown preview

### Fixed

- `focusBorder` color in the dark theme

## [0.7.0] - 2025-02-28

### Added

- Theme colors for:
    - Source control page
    - Code lens
    - Indent guide
    - Inlay hint
    - Suggest widget
    - Extension button
    - Git decoration
    - Input background
    - Markdown preview

### Changed

- Literals are now blue in the dark and light themes

## [0.6.0] - 2025-02-27

### Changed

- New border color
- New minimap colors for errors and warnings

## [0.5.0] - 2025-02-24

### Added

- Theme colors for:
    - Icons
    - Progress bar
    - Badges
    - URLs

### Changed

- More vibrant git colors
- Some colors now fit the theme better

## [0.4.0] - 2025-02-21

### Changed

- The logo now matches the theme
- More readable git decorations
- The light theme color is now magenta

## [0.3.0] - 2025-02-15

### Added

- A color for literals
- Terminal colors

### Changed

- The logo now matches the theme

### Removed

- High contrast variants

## [0.2.0] - 2025-02-08

### Added

- High contrast variants

### Changed

- Improved theme colors

### Fixed

- Logo font

## [0.1.0] - 2025-02-08

### Added

- `LICENSE`
- `README.md`
- Monochromator Dark
- Monochromator Light

[0.36.0]: https://github.com/josiasbeem/monochromator/compare/vscode/v0.35.0...vscode/v0.36.0
[0.35.0]: https://github.com/josiasbeem/monochromator/compare/vscode/v0.34.0...vscode/v0.35.0
[0.34.0]: https://github.com/josiasbeem/monochromator/compare/vscode/v0.33.1...vscode/v0.34.0
[0.33.1]: https://github.com/josiasbeem/monochromator/compare/vscode/v0.33.0...vscode/v0.33.1
[0.33.0]: https://github.com/josiasbeem/monochromator/compare/vscode/v0.32.0...vscode/v0.33.0
[0.32.0]: https://github.com/josiasbeem/monochromator/compare/vscode/v0.31.2...vscode/v0.32.0
[0.31.2]: https://github.com/josiasbeem/monochromator/compare/vscode/v0.31.1...vscode/v0.31.2
[0.31.1]: https://github.com/josiasbeem/monochromator/compare/v0.31.0...vscode/v0.31.1
[0.31.0]: https://github.com/josiasbeem/monochromator/compare/v0.30.0...v0.31.0
[0.30.0]: https://github.com/josiasbeem/monochromator/compare/v0.29.0...v0.30.0
[0.29.0]: https://github.com/josiasbeem/monochromator/compare/v0.28.1...v0.29.0
[0.28.1]: https://github.com/josiasbeem/monochromator/compare/v0.28.0...v0.28.1
[0.28.0]: https://github.com/josiasbeem/monochromator/compare/v0.27.1...v0.28.0
[0.27.1]: https://github.com/josiasbeem/monochromator/compare/v0.27.0...v0.27.1
[0.27.0]: https://github.com/josiasbeem/monochromator/compare/v0.26.1...v0.27.0
[0.26.1]: https://github.com/josiasbeem/monochromator/compare/v0.26.0...v0.26.1
[0.26.0]: https://github.com/josiasbeem/monochromator/compare/v0.25.4...v0.26.0
[0.25.4]: https://github.com/josiasbeem/monochromator/compare/v0.25.3...v0.25.4
[0.25.3]: https://github.com/josiasbeem/monochromator/compare/v0.25.2...v0.25.3
[0.25.2]: https://github.com/josiasbeem/monochromator/compare/v0.25.1...v0.25.2
[0.25.1]: https://github.com/josiasbeem/monochromator/compare/v0.25.0...v0.25.1
[0.25.0]: https://github.com/josiasbeem/monochromator/compare/v0.24.0...v0.25.0
[0.24.0]: https://github.com/josiasbeem/monochromator/compare/v0.23.3...v0.24.0
[0.23.3]: https://github.com/josiasbeem/monochromator/compare/v0.23.2...v0.23.3
[0.23.2]: https://github.com/josiasbeem/monochromator/compare/v0.23.1...v0.23.2
[0.23.1]: https://github.com/josiasbeem/monochromator/compare/v0.23.0...v0.23.1
[0.23.0]: https://github.com/josiasbeem/monochromator/compare/v0.22.0...v0.23.0
[0.22.0]: https://github.com/josiasbeem/monochromator/compare/v0.21.0...v0.22.0
[0.21.0]: https://github.com/josiasbeem/monochromator/compare/v0.20.1...v0.21.0
[0.20.1]: https://github.com/josiasbeem/monochromator/compare/v0.20.0...v0.20.1
[0.20.0]: https://github.com/josiasbeem/monochromator/compare/v0.19.0...v0.20.0
[0.19.0]: https://github.com/josiasbeem/monochromator/compare/v0.18.0...v0.19.0
[0.18.0]: https://github.com/josiasbeem/monochromator/compare/v0.17.2...v0.18.0
[0.17.2]: https://github.com/josiasbeem/monochromator/compare/v0.17.1...v0.17.2
[0.17.1]: https://github.com/josiasbeem/monochromator/compare/v0.17.0...v0.17.1
[0.17.0]: https://github.com/josiasbeem/monochromator/compare/v0.16.0...v0.17.0
[0.16.0]: https://github.com/josiasbeem/monochromator/compare/v0.15.2...v0.16.0
[0.15.2]: https://github.com/josiasbeem/monochromator/compare/v0.15.1...v0.15.2
[0.15.1]: https://github.com/josiasbeem/monochromator/compare/v0.15.0...v0.15.1
[0.15.0]: https://github.com/josiasbeem/monochromator/compare/v0.14.1...v0.15.0
[0.14.1]: https://github.com/josiasbeem/monochromator/compare/v0.14.0...v0.14.1
[0.14.0]: https://github.com/josiasbeem/monochromator/compare/v0.13.3...v0.14.0
[0.13.3]: https://github.com/josiasbeem/monochromator/compare/v0.13.2...v0.13.3
[0.13.2]: https://github.com/josiasbeem/monochromator/compare/v0.13.1...v0.13.2
[0.13.1]: https://github.com/josiasbeem/monochromator/compare/v0.13.0...v0.13.1
[0.13.0]: https://github.com/josiasbeem/monochromator/compare/v0.12.0...v0.13.0
[0.12.0]: https://github.com/josiasbeem/monochromator/compare/v0.11.1...v0.12.0
[0.11.1]: https://github.com/josiasbeem/monochromator/compare/v0.11.0...v0.11.1
[0.11.0]: https://github.com/josiasbeem/monochromator/compare/v0.10.3...v0.11.0
[0.10.3]: https://github.com/josiasbeem/monochromator/compare/v0.10.2...v0.10.3
[0.10.2]: https://github.com/josiasbeem/monochromator/compare/v0.10.1...v0.10.2
[0.10.1]: https://github.com/josiasbeem/monochromator/compare/v0.10.0...v0.10.1
[0.10.0]: https://github.com/josiasbeem/monochromator/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/josiasbeem/monochromator/compare/v0.8.4...v0.9.0
[0.8.4]: https://github.com/josiasbeem/monochromator/compare/v0.8.3...v0.8.4
[0.8.3]: https://github.com/josiasbeem/monochromator/compare/v0.8.2...v0.8.3
[0.8.2]: https://github.com/josiasbeem/monochromator/compare/v0.8.1...v0.8.2
[0.8.1]: https://github.com/josiasbeem/monochromator/compare/v0.8.0...v0.8.1
[0.8.0]: https://github.com/josiasbeem/monochromator/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/josiasbeem/monochromator/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/josiasbeem/monochromator/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/josiasbeem/monochromator/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/josiasbeem/monochromator/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/josiasbeem/monochromator/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/josiasbeem/monochromator/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/josiasbeem/monochromator/tree/v0.1.0
