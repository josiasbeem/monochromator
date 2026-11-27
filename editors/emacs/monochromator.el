(deftheme monochromator-dark
  "Distraction-free monochromatic dark theme with an electric-blue accent."
  :background-mode 'dark
  :kind 'color-scheme
  :family 'monochromator)

(let* ((bg             "#08080a")
       (bg-raised      "#111114")
       (bg-line        "#161619")
       (bg-surface     "#202024")
       (bg-selection  "#063954") ; #02a9ff at ~30% over the editor background
       (bg-selection-2 "#07202e")
       (border         "#505058")
       (divider        "#26262c")
       (blue           "#02a9ff")
       (blue-light     "#9fdbff")
       (blue-hover     "#39bfff")
       (blue-fg        bg)
       (white          "#ffffff")
       (fg             "#cbccd0")
       (muted          "#808087")
       (green          "#30df81")
       (orange         "#ff8c24")
       (red            "#ff4a4b")
       (green-bg       "#0e281c")
       (orange-bg      "#2d1b0c")
       (red-bg         "#2d1214"))

  (custom-theme-set-faces
   'monochromator-dark

   ;; General surfaces and window chrome.
   `(default ((t (:foreground ,fg :background ,bg))))
   `(cursor ((t (:background ,blue))))
   `(region ((t (:background ,bg-selection :extend t))))
   `(hl-line ((t (:background ,bg-line :extend t))))
   `(fringe ((t (:foreground ,muted :background ,bg))))
   `(internal-border ((t (:background ,bg))))
   `(child-frame-border ((t (:foreground ,border :background ,bg))))
   `(vertical-border ((t (:foreground ,divider :background ,bg))))
   `(window-divider ((t (:foreground ,divider))))
   `(window-divider-first-pixel ((t (:foreground ,divider))))
   `(window-divider-last-pixel ((t (:foreground ,divider))))
   `(shadow ((t (:foreground ,muted))))
   `(highlight ((t (:foreground ,white :background ,bg-surface))))
   `(secondary-selection ((t (:background ,bg-surface :extend t))))
   `(trailing-whitespace ((t (:background ,orange-bg :underline (:color ,orange)))))
   `(escape-glyph ((t (:foreground ,blue))))
   `(homoglyph ((t (:foreground ,blue))))
   `(nobreak-space ((t (:foreground ,blue :underline t))))

   ;; Basic typographic faces.
   `(bold ((t (:weight bold))))
   `(italic ((t (:slant italic))))
   `(bold-italic ((t (:weight bold :slant italic))))
   `(underline ((t (:underline t))))

   ;; Line numbers, where display-line-numbers-mode is used.
   `(line-number ((t (:foreground ,muted :background ,bg :weight normal))))
   `(line-number-current-line
     ((t (:foreground ,fg :background ,bg-line :weight bold))))

   ;; Links and standard widgets.
   `(link ((t (:foreground ,blue :underline t))))
   `(link-visited ((t (:foreground ,muted :underline t))))
   `(button ((t (:foreground ,fg :background ,bg-surface
                             :box (:line-width -1 :color ,border)))))
   `(widget-field ((t (:foreground ,fg :background ,bg-surface
                                   :box (:line-width -1 :color ,border)))))
   `(widget-button ((t (:foreground ,white :weight bold))))

   ;; Standard status faces.  These colours are for UI state, not syntax.
   `(error ((t (:foreground ,red :weight bold))))
   `(warning ((t (:foreground ,orange :weight bold))))
   `(success ((t (:foreground ,green :weight bold))))

   ;; Search, matching, and delimiter feedback.
   `(match ((t (:foreground ,white :background ,bg-selection :weight bold))))
   `(isearch ((t (:foreground ,blue-fg :background ,blue :weight bold))))
   `(isearch-fail ((t (:foreground ,red :background ,red-bg :weight bold))))
   `(lazy-highlight ((t (:foreground ,white :background ,bg-selection-2))))
   `(show-paren-match
     ((t (:foreground ,blue :background ,bg-selection-2 :weight bold
                       :box (:line-width -1 :color ,blue)))))
   `(show-paren-match-expression ((t (:background ,bg-selection-2))))
   `(show-paren-mismatch
     ((t (:foreground ,blue-fg :background ,red :weight bold))))

   ;; Minibuffer and auxiliary display areas.
   `(minibuffer-prompt ((t (:foreground ,blue :weight bold))))
   `(tooltip ((t (:foreground ,fg :background ,bg-surface))))
   `(header-line ((t (:foreground ,fg :background ,bg-raised))))
   `(header-line-highlight ((t (:foreground ,white :background ,bg-surface))))

   ;; Built-in tab interfaces.
   `(tab-bar ((t (:foreground ,muted :background ,bg-raised))))
   `(tab-bar-tab ((t (:foreground ,white :background ,bg :weight bold))))
   `(tab-bar-tab-inactive ((t (:foreground ,muted :background ,bg-raised))))
   `(tab-line ((t (:inherit tab-bar))))
   `(tab-line-tab ((t (:inherit tab-bar-tab))))
   `(tab-line-tab-current ((t (:inherit tab-bar-tab))))
   `(tab-line-tab-inactive ((t (:inherit tab-bar-tab-inactive))))
   `(tab-line-highlight ((t (:foreground ,white :background ,bg-surface))))

   ;; Active mode line deliberately carries the Monochromator accent.
   `(mode-line
     ((t (:foreground ,blue-fg :background ,blue
                      :box (:line-width 3 :color ,blue)
                      :weight normal :slant normal))))
   `(mode-line-inactive
     ((t (:foreground ,muted :background ,bg-line
                      :box (:line-width 3 :color ,bg-line)
                      :weight normal :slant normal))))
   `(mode-line-buffer-id ((t (:foreground ,blue-fg :weight bold))))
   `(mode-line-emphasis ((t (:foreground ,blue-fg :weight bold))))
   `(mode-line-highlight ((t (:foreground ,blue-fg :background ,blue-hover))))

   ;; ----------------------------------------------------------------------
   ;; Shared Font Lock faces: the language-agnostic syntax palette.
   ;; ----------------------------------------------------------------------
   `(font-lock-comment-face
     ((t (:foreground ,muted :slant italic :weight normal))))
   `(font-lock-comment-delimiter-face
     ((t (:inherit font-lock-comment-face))))
   `(font-lock-doc-face
     ((t (:foreground ,muted :weight bold :slant italic))))
   `(font-lock-keyword-face
     ((t (:foreground ,white :weight bold))))
   `(font-lock-builtin-face
     ((t (:foreground ,white :slant italic :weight normal))))
   `(font-lock-preprocessor-face
     ((t (:foreground ,white :weight bold))))
   `(font-lock-function-name-face
     ((t (:foreground ,white :slant italic :weight normal))))
   `(font-lock-type-face
     ((t (:foreground ,white :underline t :weight normal))))
   `(font-lock-variable-name-face
     ((t (:foreground ,fg :weight normal :slant normal))))
   `(font-lock-constant-face
     ((t (:foreground ,blue :weight normal))))
   `(font-lock-string-face
     ((t (:foreground ,blue :weight normal :slant normal))))
   `(font-lock-warning-face
     ((t (:foreground ,orange :weight bold))))
   `(font-lock-negation-char-face
     ((t (:foreground ,muted :weight normal))))
   `(font-lock-regexp-grouping-backslash
     ((t (:foreground ,blue :underline t))))
   `(font-lock-regexp-grouping-construct
     ((t (:foreground ,muted))))

   ;; Outline is part of Emacs; keep its hierarchy typographic, not colourful.
   `(outline-1 ((t (:foreground ,white :weight bold))))
   `(outline-2 ((t (:foreground ,fg :weight bold))))
   `(outline-3 ((t (:foreground ,white :weight bold))))
   `(outline-4 ((t (:foreground ,fg :weight normal))))
   `(outline-5 ((t (:foreground ,muted :weight bold))))
   `(outline-6 ((t (:foreground ,muted :weight normal))))
   `(outline-7 ((t (:foreground ,muted :weight normal))))
   `(outline-8 ((t (:foreground ,muted :weight normal))))

   ;; Diff faces are shared by Emacs's diff interface.
   `(diff-added ((t (:foreground ,green :background ,green-bg :extend t))))
   `(diff-removed ((t (:foreground ,red :background ,red-bg :extend t))))
   `(diff-changed ((t (:foreground ,orange :background ,orange-bg :extend t))))
   `(diff-refine-added ((t (:foreground ,green :background ,green-bg :weight bold))))
   `(diff-refine-removed ((t (:foreground ,red :background ,red-bg :weight bold))))
   `(diff-refine-changed ((t (:foreground ,orange :background ,orange-bg :weight bold))))
   `(diff-header ((t (:foreground ,muted :background ,bg-line :extend t))))
   `(diff-file-header ((t (:foreground ,white :background ,bg-surface :weight bold :extend t))))
   `(diff-hunk-header ((t (:inherit diff-header))))
   `(diff-context ((t (:foreground ,muted :extend t))))

   ;; Dired is built into Emacs; these faces keep file listings restrained.
   `(dired-directory ((t (:foreground ,white :weight bold))))
   `(dired-symlink ((t (:foreground ,blue))))
   `(dired-header ((t (:foreground ,white :weight bold))))
   `(dired-mark ((t (:foreground ,blue :weight bold))))
   `(dired-marked ((t (:foreground ,blue :weight bold))))
   `(dired-flagged ((t (:foreground ,orange :weight bold))))
   `(dired-warning ((t (:foreground ,orange))))

   ;; ANSI colours used by Emacs's built-in ansi-color package and terminals.
   `(ansi-color-black ((t (:foreground ,bg :background ,bg))))
   `(ansi-color-bright-black ((t (:foreground ,muted :background ,muted))))
   `(ansi-color-red ((t (:foreground ,red :background ,red))))
   `(ansi-color-bright-red ((t (:foreground ,red :background ,red))))
   `(ansi-color-green ((t (:foreground ,green :background ,green))))
   `(ansi-color-bright-green ((t (:foreground ,green :background ,green))))
   `(ansi-color-yellow ((t (:foreground ,orange :background ,orange))))
   `(ansi-color-bright-yellow ((t (:foreground ,orange :background ,orange))))
   `(ansi-color-blue ((t (:foreground ,blue :background ,blue))))
   `(ansi-color-bright-blue ((t (:foreground ,blue :background ,blue))))
   `(ansi-color-magenta ((t (:foreground ,fg :background ,fg))))
   `(ansi-color-bright-magenta ((t (:foreground ,white :background ,white))))
   `(ansi-color-cyan ((t (:foreground ,blue-light :background ,blue-light))))
   `(ansi-color-bright-cyan ((t (:foreground ,blue-light :background ,blue-light))))
   `(ansi-color-white ((t (:foreground ,fg :background ,fg))))
   `(ansi-color-bright-white ((t (:foreground ,white :background ,white)))))

  ;; Emacs 29+ adds shared Font Lock faces for parser-based fontification.
  ;; Apply these only when the running Emacs defines them; no language package
  ;; or tree-sitter grammar is loaded by this theme.
  (dolist
      (spec
       `((font-lock-function-call-face
          ((t (:foreground ,white :slant italic :weight normal))))
         (font-lock-variable-use-face
          ((t (:foreground ,fg :weight normal))))
         (font-lock-property-name-face
          ((t (:foreground ,fg :weight normal))))
         (font-lock-property-use-face
          ((t (:foreground ,fg :weight normal))))
         (font-lock-number-face
          ((t (:foreground ,blue :weight normal))))
         (font-lock-escape-face
          ((t (:foreground ,blue :underline t :weight normal))))
         (font-lock-regexp-face
          ((t (:foreground ,blue :weight normal))))
         (font-lock-operator-face
          ((t (:foreground ,muted :weight normal))))
         (font-lock-punctuation-face
          ((t (:foreground ,muted :weight normal))))
         (font-lock-bracket-face
          ((t (:foreground ,muted :weight normal))))
         (font-lock-delimiter-face
          ((t (:foreground ,muted :weight normal))))
         (font-lock-misc-punctuation-face
          ((t (:foreground ,muted :weight normal)))))
    (when (facep (car spec))
      (apply #'custom-theme-set-faces 'monochromator-dark (list spec))))))

;;;###autoload
(when load-file-name
  (add-to-list 'custom-theme-load-path
               (file-name-as-directory (file-name-directory load-file-name))))

(provide-theme 'monochromator-dark)

 xyen@nonq  …/monochromator/editors/emacs  master? ❯ cd ..                                                                                                                           
 xyen@nonq  …/varOSS/monochromator/editors  master? ❯ cd ..                                                                                                                          
 xyen@nonq  ~/varOSS/monochromator  master? ❯                                                                                                                                        






































