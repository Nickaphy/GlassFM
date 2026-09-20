# Documentation

Simple map of what each file is for. Disk work lives in `FileSystemOperations`; the screen lives in `TUI`.

---

## Root

### `main.go`
Starts the app: clears the terminal, builds the Bubble Tea program from `TUI.NewModel()`, and runs it until quit.

### `go.mod` / `go.sum`
Go module name and dependency versions (Bubble Tea, Lipgloss, etc.).

### `README.md`
Project overview for humans cloning the repo.

---

## `Internals/Docs/`

### `Keybindings.md`
Intended keys: navigation (`h`/`l`/`j`/`k`) and leader commands after Space (`cd`, `mk`, …).

### `Progress.md`
Scratch checklist of what is done so far.

### `Documentation.md`
This file — short descriptions of the codebase layout.

---

## `Internals/FileSystemOperations/`

Thin wrappers around Go’s `os` package. The TUI calls these; they do not draw UI.

### `GetWorkingDir.go`
Returns the process’s current working directory (`os.Getwd`).

### `ChangeDir.go`
Changes the process working directory (`os.Chdir`) — used by `SPC cd`.

### `ReadDir.go`
Lists names in a directory (`ListDir` → `os.ReadDir`).

### `CreateDir.go`
Creates a directory (`os.Mkdir`). Fails if the path already exists — no overwrite.

### `CreateFile.go`
Creates (or truncates) a file at a path.

### `Rename.go`
Renames a file or directory.

### `DeleteFile.go`
Deletes a file.

### `DeleteDir.go`
Deletes a directory.

### `Copy.go`
Copies a file or directory to a destination.

### `Move.go`
Moves a file or directory to a destination.

---

## `Internals/TUI/`

All files are `package TUI`. One Bubble Tea `Model`; files are split by responsibility.

### `model.go`
Core state (`Model`), startup (`NewModel`), `Init` / `Update` routing, and `afterAction` (reload list + start flash toast).

### `styles.go`
Shared Lipgloss styles: title, dim hints, prompt box, teal flash text.

### `browse.go`
Normal browsing keys: `j`/`k` move, Space opens leader, `q` quits.
Dir forward and back: `h`/`l` 

### `leader.go`
After Space: collect a command name (`cd`, `mk`, …), match against the menu, open a prompt when complete.

### `prompt.go`
Typing in the command line (enter / esc / backspace). `submitPrompt` dispatches to the right command file. Shared `promptBox` chrome.

### `view.go`
Draws the screen: title, cwd, file list, mode-specific footer (leader menu, prompt, or hints/flash).

### `flash.go`
Timer that clears the success toast after a few seconds (`flashExpiredMsg` + `flashID`).

### `cd.go`
Runs change-directory, then shows a flash like `cd /absolute/path`.

### `mkdir.go`
Runs create-directory, then shows a flash like `mkdir /absolute/path`.
