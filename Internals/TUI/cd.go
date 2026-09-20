package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"

	tea "github.com/charmbracelet/bubbletea"
)

// submitCd changes the process working directory, then refreshes the listing.
func (m Model) submitCd(path string) (tea.Model, tea.Cmd) {
	if err := FileSystemOperations.ChangeDir(path); err != nil {
		m.status = err.Error()
		return m, nil
	}
	cwd, err := FileSystemOperations.GetWorkingDir()
	if err != nil {
		m.status = err.Error()
		return m, nil
	}
	m.cwd = cwd
	m.cursor = 0
	return m.afterAction()
}
