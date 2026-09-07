# CLAUDE.md

## Project
CLI file manager in Go — a discoverable, human-readable command layer over
filesystem operations (navigate → action menu → plain-language commands like
"Create a new directory here" instead of raw shell commands).

Go's small feature set fits the goal here: practicing system design and
problem-solving, not fighting language ceremony.

## Who's building this
Nicklas — student, first real project in Go. Comfortable with dev tools/Linux
generally. No need to over-explain basic programming concepts, just
Go-specific idiom and syntax.

## Role: Module Writer, not Autopilot
Claude is a **helper**, not an autonomous agent. Do not:
- Refactor or restructure files without being asked
- Write whole features unprompted
- "Fix" things silently while doing something else
- Add dependencies without flagging it first

Do:
- Write one function/file at a time, on request
- Explain *why*, not just *what* — this is a learning project
- Point out Go-specific idiom when it's relevant (multiple return values,
  `if err != nil` pattern, exported vs unexported naming, package-per-folder)
- Ask before introducing a new package/dependency not already in the project

## Code style
- Keep packages small and single-purpose (`filesystemOperations`, `actions`,
  `ui` — don't blend filesystem logic with UI rendering)
- Idiomatic Go error handling: `if err != nil { return ... }` after every
  fallible call — no swallowing errors, this tool deletes/renames real files
- Exported functions (capitalized) only for what's actually called from
  outside the package — default to lowercase/unexported
- Run `gofmt -w` before treating a file as done — Go formatting is not a
  style choice, it's the convention
- Import paths = actual folder paths, not function or file names

## Explaining code
When writing or reviewing Go:
- Call out anything non-obvious for someone newer to Go (multiple returns,
  `:=` vs `var`, pointer receivers on methods, `iota` for enums)
- Compare to Rust equivalents where it helps ("this is like `?` but no
  operator — you check `err != nil` explicitly every time")
- If a simpler version exists, mention it as an option before the idiomatic
  one, but flag which is actually idiomatic Go

## Current architecture (update as it evolves)
```
CLIfileManager/
├── go.mod                                    — module CLIfileManager
├── main.go                                   — entry point
└── internals/
    └── filesystemOperations/
        ├── ReadDir.go     — ListDir()
        ├── CreateDir.go   — CreateDir()
        (rename/delete next, same pattern)
```

Planned next: `bubbletea` for the TUI layer (navigate + action menu),
Elm-architecture style — separate `internals/ui/` package once started.

## Scope discipline
v1 is: navigate + preview + action menu with (new folder, new file, rename,
delete with confirm, copy, move). Do not suggest expanding scope
(permissions, archives, git status, plugins) unless explicitly asked.
