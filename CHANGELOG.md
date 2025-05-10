# Changelog

## v0.15.0 - 2025-05-10

- Fix LaTeX math mode
- Fix more constants
- Various QOL improvements
    - Documentation keywords bold and gray (also things like `@Override` in Java or shebangs)
    - Preprocessor directives are now bold and gray
    - Better highlighting for escaped characters
    - HTML and CSS tags are now bold

## v0.14.1 - 2025-05-08

- Fix unit highlighting (such as "px" in CSS)
- Fix variable highlighting of shell scripts and JavaScript
- Fix punctuation highlighting of some characters in Java and R
- Fix regex highlighting

## v0.14.0 - 2025-05-04

- Add the following themes:
    - `Monochromator Dark Amber`
    - `Monochromator Dark Emerald`
    - `Monochromator Dark Ruby`
    - `Monochromator Light Amber`
    - `Monochromator Light Emerald`
    - `Monochromator Light Ruby`
- Highlight timestamps as literals
- Reduce contrast of `editor.wordHighlightStrongBackground`

## v0.13.3 - 2025-05-04

- Fix highlighting of preprocessor names and hexadecimal numbers

## v0.13.2 - 2025-04-22

- Fix highlighting in Python, Typescript, LaTeX, YAML and TOML

## v0.13.1 - 2025-04-21

- Correct selection highlighting
- Fix inconsistencies
- Remove redundant code

## v0.13.0 - 2025-04-18

- Fix highlight color of selections in minimap
- Change `scrollbarSlider.background` to match `editorGroupHeader.tabsBackground`
- Change `textCodeBlock.background` to match `textBlockQuote.background`
- Change icon design
- Change colors to match https://gitlab.gnome.org/GNOME/libadwaita/-/merge_requests/1317
    - `#1e1e1e` -> `#1d1d20`
    - `#242424` -> `#222226`
    - `#303030` -> `#2e2e32`
    - `#3f3f3f` -> `#424246`
    - `#dfdfdf` -> `#d9d9da`
    - `#ebebeb` -> `#ebebed`
    - `#f5f5f5` -> `#f2f2f4`

## v0.12.0 - 2025-04-16

- Darken shade of blue in `Monochromator Dark`
- Reduce use of bold emphasis
- Integrate selections into the theme
- Fix highlighting of constants (has been wrong in e.g. Rust)
- Improve Markdown preview
- Fix lightbulb color to finally display in the right shade of gray
- Change icon design
- Add badges to `README.md`
- Remove redundant code

## v0.11.1 - 2025-04-10

- Fix tabs color
- Fix screenshots

## v0.11.0 - 2025-04-10

- General UI overhaul
- Add color for `gitDecoration.addedResourceForeground`

## v0.10.3 - 2025-03-24

- Fix dropdown background color
- Decrease contrast of inlay hints

## v0.10.2 - 2025-03-23

- Fix color of ...
    - Scrollbar slider
    - Panel title
    - Fold background

## v0.10.1 - 2025-03-21

- Replace `icon.png` with `icon128.png` and `icon1024.png`
- Merge screenshots
- Beautify `README.md`
- Fix changelog formatting

## v0.10.0 - 2025-03-21

- Add boldness solely for `constant.language`
- Fix menu selection color in light theme
- Increase overall contrast
- Improve look of ...
    - Activity bar
    - Command palette
    - Keybindings

## v0.9.0 - 2025-03-21

- Integrate more components to fit the theme such as ...
    - Checkboxes
    - Breadcrumbs
    - Git decoration for ignored files
    - Indent guides
    - Sticky scroll shadow
    - Menu (which is displayed when you right-click)
    - Widgets
    - Notifications
- Simplify changelog format
- Set semantic highlighting to true
- Fix missing hashtag for hex code
- Remove boldness of constants
- Revoke changes of `v0.8.4`

## v0.8.4 - 2025-03-17

- Fix color for inactive background
    - `titleBar.inactiveBackground` from `#ebebeb` to `#f0f0f0` in light theme
    - `titleBar.inactiveBackground` from `#303030` to `#2a2a2a` in dark theme

## v0.8.3 - 2025-03-07

- Add download section to `README.md`
- Add bugs section to `package.json`
- Change name of `logo.png` to `icon.png`
- Change structure of screenshots (created a directory for screenshots)
- Change structure of `README.md`
- Fixed formatting in `CHANGELOG.md`

## v0.8.2 - 2025-03-06

- Change info color to match theme

## v0.8.1 - 2025-03-02

- Fix screenshots in `README.md` by changing the git repository

## v0.8.0 - 2025-03-02

- Add `CHANGELOG.md`
- Add more information, a "License" section and screenshots to `README.md`
- Add theme colors to support Markdown better (in the editor)
- Change font in logo from "Fira Code" to "IBM Plex Mono"
- Make color of progress bars less aggressive
- Improve readability by changing the colors of URLs to blue
- `textBlockQuote` for better readability (in Markdown Preview)
- Fix color of `focusBorder` in the dark theme

## v0.7.0 - 2025-02-28

- Add theme colors to ...
    - Source control page
    - Code lens
    - Indent guide
    - Inlay hint
    - Suggest widget
    - Extension button
    - Git decoration
    - Input background
    - Markdown preview
- Change color of literals to blue for dark and light theme

## v0.6.0 - 2025-02-27

- Change color of borders
- Change colors of errors and warnings in the minimap

## v0.5.0 - 2025-02-24

- Add theme colors to ...
    - Icons
    - Progress bar
    - Badges
    - URLs
- Change git colors to make them more vibrant
- Change some colors to fit in better with the theme

## v0.4.0 - 2025-02-21

- Change logo to fit in with the new theme changes
- Change git decorations to make them more readable
- Change color of light theme to magenta

## v0.3.0 - 2025-02-15

- Add one color for literals
- Add terminal colors
- Change logo to fit in with the new theme changes
- Remove high contrast variants

## v0.2.0 - 2025-02-08

- Add high contrast variants
- Improve theme colors
- Fix font in logo

## v0.1.0 - 2025-02-08

- Add `LICENSE`
- Add `README.md`
- Add Monochromator Dark
- Add Monochromator Light
