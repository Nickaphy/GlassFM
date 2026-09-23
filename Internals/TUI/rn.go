package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// startRename remembers the selection and opens a prompt prefilled with its name.
func (m Model) startRename() (tea.Model, tea.Cmd) {
	if len(m.items) == 0 {
		m.mode = modeBrowse
		m.status = "nothing selected to rename"
		return m, nil
	}

	src := filepath.Join(m.cwd, m.items[m.cursor])
	if abs, err := filepath.Abs(src); err == nil {
		src = abs
	}

	m.pendingSrc = src
	m.command = "rn"
	m.input = filepath.Base(src)
	m.status = ""
	m.mode = modePrompt
	return m, nil
}

// submitRn renames pendingSrc to the new basename in the same directory.
func (m Model) submitRn(newName string) (tea.Model, tea.Cmd) {
	if m.pendingSrc == "" {
		m.status = "no source to rename"
		return m, nil
	}
	newName = strings.TrimSpace(newName)
	if newName == "" {
		m.status = "type a new name"
		return m, nil
	}
	if strings.Contains(newName, string(filepath.Separator)) {
		m.status = "name cannot contain path separators"
		return m, nil
	}

	dest := filepath.Join(filepath.Dir(m.pendingSrc), newName)
	if err := FileSystemOperations.Rename(m.pendingSrc, dest); err != nil {
		m.status = err.Error()
		return m, nil
	}

	return m.afterAction("mv " + m.pendingSrc + " " + dest)
}
