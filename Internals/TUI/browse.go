package TUI

import tea "github.com/charmbracelet/bubbletea"

// handleBrowse: moving around the list. Space leaves browse and starts the leader.
func (m Model) handleBrowse(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return m, tea.Quit
	case " ":
		m.mode = modeLeader
		m.leader = ""
		m.status = ""
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.items)-1 {
			m.cursor++
		}
	}
	return m, nil
}
