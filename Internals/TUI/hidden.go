package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// toggleHidden flips whether dotfiles (names starting with ".") are listed.
func (m Model) toggleHidden() (tea.Model, tea.Cmd) {
	m.showHidden = !m.showHidden
	m.mode = modeBrowse
	m.leader = ""
	m.input = ""
	m.command = ""
	m.status = ""
	m.items = m.loadItems()
	m = m.withClampedCursor()

	shown := "hide hidden"
	if m.showHidden {
		shown = "show hidden"
	}
	m.flash = shown
	m.flashID++
	id := m.flashID
	return m, clearFlashAfter(id, flashDuration)
}

// loadItems lists cwd, optionally dropping names that start with ".".
func (m Model) loadItems() []string {
	items, err := FileSystemOperations.ListDir(m.cwd)
	if err != nil {
		return nil
	}
	if m.showHidden {
		return items
	}
	visible := make([]string, 0, len(items))
	for _, name := range items {
		if strings.HasPrefix(name, ".") {
			continue
		}
		visible = append(visible, name)
	}
	return visible
}

func (m Model) withClampedCursor() Model {
	if len(m.items) == 0 {
		m.cursor = 0
		return m
	}
	if m.cursor >= len(m.items) {
		m.cursor = len(m.items) - 1
	}
	return m
}
