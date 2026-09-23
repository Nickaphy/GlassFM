package TUI

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// handlePrompt: editing the command line. Does not touch the filesystem;
// Enter delegates that to submitPrompt.
func (m Model) handlePrompt(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		return m.cancelPending()
	case "tab":
		// cp/mv: leave the path box and pick a folder by browsing
		if (m.command == "cp" || m.command == "mv") && m.pendingSrc != "" {
			m.mode = modePickDest
			m.input = ""
			m.status = ""
			return m, nil
		}
		return m, nil
	case "enter":
		return m.submitPrompt()
	case "backspace":
		if m.input != "" {
			runes := []rune(m.input)
			m.input = string(runes[:len(runes)-1])
		}
		return m, nil
	}

	if msg.Type == tea.KeyRunes {
		m.input += string(msg.Runes)
	}
	return m, nil
}

// submitPrompt routes to the command-specific submitter. Failures stay in
// the prompt with status set; success refreshes via afterAction.
func (m Model) submitPrompt() (tea.Model, tea.Cmd) {
	path := strings.TrimSpace(m.input)
	if path == "" && m.command != "rm" {
		m.status = "type a path, then enter"
		return m, nil
	}

	switch m.command {
	case "cd":
		return m.submitCd(path)
	case "mk":
		return m.submitMk(path)
	case "cp":
		return m.submitCp(path)
	case "mv":
		return m.submitMv(path)
	case "rn":
		return m.submitRn(path)
	case "rm":
		return m.submitRm()
	default:
		m.mode = modeBrowse
		m.input = ""
		m.command = ""
		return m, nil
	}
}

// promptBox is shared chrome for any command prompt, centered in the middle.
func (m Model) promptBox(inner string) string {
	style := promptStyle
	col := m.contentColWidth()
	innerWidth := col - style.GetHorizontalFrameSize()
	if innerWidth > 0 {
		style = style.Width(innerWidth)
	}
	return m.centerBlock(style.Render(inner))
}
