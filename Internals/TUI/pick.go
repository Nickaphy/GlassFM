package TUI

import (
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// handlePickDest: navigating while placing a copy/move. Enter pastes into cwd.
func (m Model) handlePickDest(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		return m.cancelPending()
	case "enter":
		return m.confirmPendingHere()
	case ":":
		m.mode = modePrompt
		m.input = ""
		m.status = ""
		return m, nil
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.items)-1 {
			m.cursor++
		}
	case "left", "h":
		return m.navigateKeepMode("..")
	case "right", "l":
		if len(m.items) == 0 {
			return m, nil
		}
		return m.navigateKeepMode(m.items[m.cursor])
	}
	return m, nil
}

// confirmPendingHere finishes cp/mv into the current directory.
func (m Model) confirmPendingHere() (tea.Model, tea.Cmd) {
	switch m.command {
	case "cp":
		return m.submitCp(m.cwd)
	case "mv":
		return m.submitMv(m.cwd)
	default:
		return m.cancelPending()
	}
}

// navigateKeepMode changes directory without leaving pick mode or flashing.
func (m Model) navigateKeepMode(path string) (tea.Model, tea.Cmd) {
	next, err := m.applyDir(path)
	if err != nil {
		m.status = err.Error()
		return m, nil
	}
	next.status = ""
	return next, nil
}

func (m Model) pickFooter() string {
	name := filepath.Base(m.pendingSrc)
	verb := "copying"
	if m.command == "mv" {
		verb = "moving"
	}
	line := flashStyle.Render(verb + " " + name)
	line += "\n" + hintStyle.Render("enter paste here · : path · h/l navigate · esc cancel")
	if m.status != "" {
		line += "\n" + m.status
	}
	return m.centerBlock(line)
}
