# Changelog

## v0.34.0 - 2026-09-18

- Made bold exclusive to structural keywords so that the keywords show the code structure again
    - Types are now underlined instead of bold
    - Escape sequences and format placeholders are now underlined instead of bold
    - Language constants such as `true`, `false`, `nil`, and `null` are no longer bold
    - `this`, `self`, `super`, decorators, annotations, and attributes are now gray
- Made functions and methods brighter than variables to keep the overview in long call chains
- Dimmed the editor foreground
    - Dark: `#e7e7e7` -> `#ccccce`
    - Light: `#000000` -> `#303034`
- Unified the UI foreground with the editor foreground
- Fixed `delete` and the named casts in C++, which were gray while `new` was bold
- Fixed the function calls in Python, which were not italic
- Fixed the base classes and interfaces after `extends` and `implements`, which did not have the type style
- Fixed the digit separators and exponent signs in number literals, which were gray

## v0.33.1 - 2026-09-01

- Fixed the unreadable extension action buttons in the light themes

## v0.33.0 - 2026-07-15

- Darkened the dark theme background
    - Background: `#0f0f12` -> `#08080a`
    - To keep the previous value, set `"workbench.colorCustomizations": { "editor.background": "#0f0f12", "terminal.background": "#0f0f12" }` in your `settings.json`
- Tinted the text selections with each variant's accent color

## v0.32.0 - 2026-04-14

- Improved the accent colors of the light themes for WCAG AA compliance
    - Aquamarine: `#01aacf` -> `#007b94`
    - Amber: `#e66100` -> `#bf4e00`
    - Sulfur: `#927e00` -> `#856e00`

## v0.31.2 - 2026-04-12

- Fixed the file highlight issue in the file explorer ([#8](https://codeberg.org/beem/monochromator/issues/8))

## v0.31.1 - 2026-02-28

- Made the scrollbar transparent so that it no longer covers the breadcrumbs ([#5](https://codeberg.org/beem/monochromator/issues/5))
- Replaced the rarely used color values with similar existing values for consistency
- Removed the unnecessary migration code from v0.27.0

## v0.31.0 - 2026-02-26

- Dimmed the white foreground value in the dark theme
- Improved support for Typst
- Fixed the border colors of the editor group, the panel, and the overview ruler

## v0.30.0 - 2026-01-14

- Made functions and methods italic
- Made the gray shade darker in the dark theme
- Fixed the highlighting of Go package names
- Fixed the highlighting of JavaScript braces

## v0.29.0 - 2025-11-19

- Made the border color of the command center more subtle

## v0.28.1 - 2025-11-18

- Fixed the editor gutter colors for staged code

## v0.28.0 - 2025-11-04

- Various UI improvements:
    - Removed the visible UI borders for a cleaner appearance
    - Improved the button contrast
    - Increased the contrast between the editor and the UI elements for better focus on the code
    - Improved the contrast of comments in the light theme for better readability
    - Improved the look of the keybinding labels
    - Removed the focus borders

## v0.27.1 - 2025-10-21

- Fixed the syntax highlighting of built-in shell commands

## v0.27.0 - 2025-10-18

- Migrated the `Dark+` themes to `Dark`
- Various UI improvements
- Redesigned the light theme
- Improved the cyan and magenta colors (also the terminal colors)
- Improved the inlay hints
- Fixed the highlighting of TypeScript braces

## v0.26.1 - 2025-10-09

- Fixed the syntax highlighting of CSS ID selectors

## v0.26.0 - 2025-10-06

- Reduced the blue tint in the Dark+ background colors

## v0.25.4 - 2025-09-29

- Fixed the syntax highlighting for BibTeX files

## v0.25.3 - 2025-09-16

- Fixed the command center colors
- Fixed the blue tint in the Dark+ themes (added more blue)
- Fixed the border and foreground colors of the hover widgets
- Fixed the header color of the multi diff editor
- Fixed the icon colors of the activity bar so that they fit the status bar colors better

## v0.25.2 - 2025-09-15

- Fixed the colors of the status bar elements

## v0.25.1 - 2025-08-20

- Improved the syntax highlighting for Nix function arguments (the `{`, `}`, and `_` characters)

## v0.25.0 - 2025-08-11

- Added a new theme variant with a cyan accent color: `Aquamarine`
- Integrated the colors of the git diff editor into the themes
- Changed the orange terminal color to yellow
- Changed the cyan terminal color to fit the new `Aquamarine` variant
- Fixed the purple terminal color in the light themes

## v0.24.0 - 2025-08-10

- Added the `Dark+` theme variants (the `Dark` theme with improved contrast)
- Updated the terminal color scheme for better theme integration

## v0.23.3 - 2025-07-22

- Fixed the string storage type in Go

## v0.23.2 - 2025-07-22

- Fixed the highlighting of the assignment operator

## v0.23.1 - 2025-07-20

- Fixed the syntax highlighting for char literals (single quotes) in Rust
- Fixed the highlighting of the assignment operator in Rust code

## v0.23.0 - 2025-07-18

- Added a new variant with a yellow accent color:
    - `Monochromator Dark Sulfur`
    - `Monochromator Light Sulfur`
- Added the background colors for status bar errors and warnings to the theme

## v0.22.0 - 2025-07-15

- Improved the color readability:
    - Changed the purple in the dark themes (`#c61aff` -> `#ee40ff`)
    - Changed the green in the light themes (`#26a269` -> `#017f45`)
- Fixed the "Plain" theme variants: corrected the values that `v0.18.0` did not change

## v0.21.0 - 2025-07-04

- Added x86_64, ARM, RISC-V, and MIPS support for these OpenVSX extensions:
    - x86_64: `13xforever.language-x86-64-assembly`, `basdp.language-gas-x86`
    - ARM: `dan-c-underwood.arm`
    - RISC-V: `sunshaoce.RISC-V`
    - MIPS: `kdarkhan.mips`

## v0.20.1 - 2025-06-16

- Fixed the backtick symbol in Markdown for the accent color variants

## v0.20.0 - 2025-06-13

- Improved the Markdown and LaTeX highlighting
- Improved the JavaScript string interpolation

## v0.19.0 - 2025-06-12

- Added a new variant with a purple accent color:
    - `Monochromator Dark Amethyst`
    - `Monochromator Light Amethyst`

## v0.18.0 - 2025-06-11

- Made the shade of blue more visible
- Improved the string highlighting
- Fixed C++ pointers
- Fixed `self` and entity names in Rust
- Fixed enums and function arrows in Java
- Fixed constants in Haskell, Perl, Ruby, PHP, and other languages
- Fixed support functions in Lua, Julia, and other languages
- Fixed strings and symbols in F#
- Fixed null and booleans in Zig
- Decreased the minimum VS Code engine version

## v0.17.2 - 2025-06-05

- Fixed Doxygen comments for languages other than C++

## v0.17.1 - 2025-06-03

- Fixed the Code Actions menu

## v0.17.0 - 2025-05-25

- Added `Monochromator Dark Plain` and `Monochromator Light Plain`
- Fixed the syntax highlighting for backslashes that precede LaTeX keywords
- Redesigned the extension icon

## v0.16.0 - 2025-05-24

- Limited the bold format to keywords only (functions are no longer bold)
- String quotation marks now match the color of literal text
- Reverted the preprocessor directive style to the pre-`0.15.0` appearance
- Improved the appearance of YAML files
- Documentation keywords are now italic
- Fixed the highlighting of hex codes in HTML

## v0.15.2 - 2025-05-19

- Fixed the highlighting of SQL database names and table names

## v0.15.1 - 2025-05-17

- Fixed the highlighting of `storage.modifier.sql`

## v0.15.0 - 2025-05-10

- Fixed the LaTeX math mode
- Fixed more constants
- Various quality improvements:
    - Made documentation keywords bold and gray (for example, `@Override` in Java, and shebangs)
    - Made preprocessor directives bold and gray
    - Improved the highlighting of escaped characters
    - Made HTML and CSS tags bold

## v0.14.1 - 2025-05-08

- Fixed the highlighting of units (such as "px" in CSS)
- Fixed the highlighting of variables in shell scripts and JavaScript
- Fixed the highlighting of some punctuation characters in Java and R
- Fixed the regex highlighting

## v0.14.0 - 2025-05-04

- Added these themes:
    - `Monochromator Dark Amber`
    - `Monochromator Dark Emerald`
    - `Monochromator Dark Ruby`
    - `Monochromator Light Amber`
    - `Monochromator Light Emerald`
    - `Monochromator Light Ruby`
- Highlighted timestamps as literals
- Reduced the contrast of `editor.wordHighlightStrongBackground`

## v0.13.3 - 2025-05-04

- Fixed the highlighting of preprocessor names and hexadecimal numbers

## v0.13.2 - 2025-04-22

- Fixed the highlighting in Python, TypeScript, LaTeX, YAML, and TOML

## v0.13.1 - 2025-04-21

- Fixed the selection highlighting
- Fixed inconsistencies
- Removed redundant code

## v0.13.0 - 2025-04-18

- Fixed the highlight color of selections in the minimap
- Changed `scrollbarSlider.background` to match `editorGroupHeader.tabsBackground`
- Changed `textCodeBlock.background` to match `textBlockQuote.background`
- Changed the icon design
- Changed the colors to match https://gitlab.gnome.org/GNOME/libadwaita/-/merge_requests/1317
    - `#1e1e1e` -> `#1d1d20`
    - `#242424` -> `#222226`
    - `#303030` -> `#2e2e32`
    - `#3f3f3f` -> `#424246`
    - `#dfdfdf` -> `#d9d9da`
    - `#ebebeb` -> `#ebebed`
    - `#f5f5f5` -> `#f2f2f4`

## v0.12.0 - 2025-04-16

- Darkened the shade of blue in `Monochromator Dark`
- Reduced the use of bold text
- Integrated selections into the theme
- Fixed the highlighting of constants (this was wrong in Rust, for example)
- Improved the Markdown preview
- Fixed the lightbulb color to show the correct shade of gray
- Changed the icon design
- Added badges to `README.md`
- Removed redundant code

## v0.11.1 - 2025-04-10

- Fixed the color of the tabs
- Fixed the screenshots

## v0.11.0 - 2025-04-10

- Redesigned the general UI
- Added a color for `gitDecoration.addedResourceForeground`

## v0.10.3 - 2025-03-24

- Fixed the dropdown background color
- Decreased the contrast of inlay hints

## v0.10.2 - 2025-03-23

- Fixed the colors of:
    - Scrollbar slider
    - Panel title
    - Fold background

## v0.10.1 - 2025-03-21

- Replaced `icon.png` with `icon128.png` and `icon1024.png`
- Merged the screenshots
- Improved the look of `README.md`
- Fixed the changelog format

## v0.10.0 - 2025-03-21

- Added bold text only for `constant.language`
- Fixed the menu selection color in the light theme
- Increased the overall contrast
- Improved the look of:
    - Activity bar
    - Command palette
    - Keybindings

## v0.9.0 - 2025-03-21

- Integrated more components into the theme:
    - Checkboxes
    - Breadcrumbs
    - Git decoration for ignored files
    - Indent guides
    - Sticky scroll shadow
    - Menu (it opens when you right-click)
    - Widgets
    - Notifications
- Simplified the changelog format
- Set semantic highlighting to true
- Fixed a missing hashtag in a hex code
- Removed the bold text from constants
- Reverted the changes of `v0.8.4`

## v0.8.4 - 2025-03-17

- Fixed the color of the inactive background
    - `titleBar.inactiveBackground` from `#ebebeb` to `#f0f0f0` in the light theme
    - `titleBar.inactiveBackground` from `#303030` to `#2a2a2a` in the dark theme

## v0.8.3 - 2025-03-07

- Added a download section to `README.md`
- Added a bugs section to `package.json`
- Renamed `logo.png` to `icon.png`
- Changed the structure of the screenshots (created a screenshot directory)
- Changed the structure of `README.md`
- Fixed the format of `CHANGELOG.md`

## v0.8.2 - 2025-03-06

- Changed the info color to match the theme

## v0.8.1 - 2025-03-02

- Fixed the screenshots in `README.md`: changed the git repository

## v0.8.0 - 2025-03-02

- Added `CHANGELOG.md`
- Added more information, a "License" section, and screenshots to `README.md`
- Added theme colors to support Markdown better (in the editor)
- Changed the font in the logo from "Fira Code" to "IBM Plex Mono"
- Made the color of the progress bars more subtle
- Improved the readability: changed the color of URLs to blue
- Changed `textBlockQuote` for better readability (in the Markdown preview)
- Fixed the color of `focusBorder` in the dark theme

## v0.7.0 - 2025-02-28

- Added theme colors to:
    - Source control page
    - Code lens
    - Indent guide
    - Inlay hint
    - Suggest widget
    - Extension button
    - Git decoration
    - Input background
    - Markdown preview
- Changed the color of literals to blue for the dark and light themes

## v0.6.0 - 2025-02-27

- Changed the color of the borders
- Changed the colors of errors and warnings in the minimap

## v0.5.0 - 2025-02-24

- Added theme colors to:
    - Icons
    - Progress bar
    - Badges
    - URLs
- Changed the git colors to make them more vibrant
- Changed some colors to fit the theme better

## v0.4.0 - 2025-02-21

- Changed the logo to fit the new theme changes
- Changed the git decorations to make them more readable
- Changed the color of the light theme to magenta

## v0.3.0 - 2025-02-15

- Added one color for literals
- Added terminal colors
- Changed the logo to fit the new theme changes
- Removed the high contrast variants

## v0.2.0 - 2025-02-08

- Added high contrast variants
- Improved the theme colors
- Fixed the font in the logo

## v0.1.0 - 2025-02-08

- Added `LICENSE`
- Added `README.md`
- Added Monochromator Dark
- Added Monochromator Light
