package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// startMove remembers the selected item and opens the destination prompt.
// From there: type a path, or tab to pick a folder by browsing.
func (m Model) startMove() (tea.Model, tea.Cmd) {
	if len(m.items) == 0 {
		m.mode = modeBrowse
		m.status = "nothing selected to move"
		return m, nil
	}

	src := filepath.Join(m.cwd, m.items[m.cursor])
	if abs, err := filepath.Abs(src); err == nil {
		src = abs
	}

	m.pendingSrc = src
	m.command = "mv"
	m.input = ""
	m.status = ""
	m.mode = modePrompt
	return m, nil
}

// submitMv moves pendingSrc to dst (file path, or into a directory).
func (m Model) submitMv(dst string) (tea.Model, tea.Cmd) {
	if m.pendingSrc == "" {
		m.status = "no source to move"
		return m, nil
	}

	dest := resolvePendingDest(m.pendingSrc, dst)
	if err := FileSystemOperations.Move(m.pendingSrc, dest); err != nil {
		m.status = err.Error()
		return m, nil
	}

	if abs, err := filepath.Abs(dest); err == nil {
		dest = abs
	}
	return m.afterAction("mv " + m.pendingSrc + " " + dest)
}
