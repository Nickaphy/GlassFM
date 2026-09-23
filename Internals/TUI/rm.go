package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// startRemove remembers the selection and opens a confirm prompt.
func (m Model) startRemove() (tea.Model, tea.Cmd) {
	if len(m.items) == 0 {
		m.mode = modeBrowse
		m.status = "nothing selected to delete"
		return m, nil
	}

	src := filepath.Join(m.cwd, m.items[m.cursor])
	if abs, err := filepath.Abs(src); err == nil {
		src = abs
	}

	m.pendingSrc = src
	m.command = "rm"
	m.input = ""
	m.status = ""
	m.mode = modePrompt
	return m, nil
}

// submitRm deletes pendingSrc after confirm (enter). Uses DeleteDir for
// directories (RemoveAll) and DeleteFile for files.
func (m Model) submitRm() (tea.Model, tea.Cmd) {
	if m.pendingSrc == "" {
		m.status = "no source to delete"
		return m, nil
	}

	info, err := os.Stat(m.pendingSrc)
	if err != nil {
		m.status = err.Error()
		return m, nil
	}

	if info.IsDir() {
		err = FileSystemOperations.DeleteDir(m.pendingSrc)
	} else {
		err = FileSystemOperations.DeleteFile(m.pendingSrc)
	}
	if err != nil {
		m.status = err.Error()
		return m, nil
	}

	return m.afterAction("rm " + m.pendingSrc)
}
