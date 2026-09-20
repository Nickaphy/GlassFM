package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"

	tea "github.com/charmbracelet/bubbletea"
)

// submitMk creates a directory. CreateDir uses os.Mkdir: error if the path
// already exists (no overwrite).
func (m Model) submitMk(path string) (tea.Model, tea.Cmd) {
	if err := FileSystemOperations.CreateDir(path); err != nil {
		m.status = err.Error()
		return m, nil
	}
	return m.afterAction()
}
