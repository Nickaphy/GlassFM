package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"

	tea "github.com/charmbracelet/bubbletea"
)

// applyDir changes cwd and reloads the listing. Caller decides mode/flash.
func (m Model) applyDir(path string) (Model, error) {
	if err := FileSystemOperations.ChangeDir(path); err != nil {
		return m, err
	}
	cwd, err := FileSystemOperations.GetWorkingDir()
	if err != nil {
		return m, err
	}
	m.cwd = cwd
	m.cursor = 0
	m.items = m.loadItems()
	return m, nil
}

// submitCd changes the process working directory, then refreshes the listing.
func (m Model) submitCd(path string) (tea.Model, tea.Cmd) {
	next, err := m.applyDir(path)
	if err != nil {
		m.status = err.Error()
		return m, nil
	}
	return next.afterAction("cd " + next.cwd)
}
