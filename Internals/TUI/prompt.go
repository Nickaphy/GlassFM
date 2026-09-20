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
		m.mode = modeBrowse
		m.input = ""
		m.command = ""
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
	if path == "" {
		m.status = "type a path, then enter"
		return m, nil
	}

	switch m.command {
	case "cd":
		return m.submitCd(path)
	case "mk":
		return m.submitMk(path)
	default:
		m.mode = modeBrowse
		m.input = ""
		m.command = ""
		return m, nil
	}
}

// promptBox is shared chrome for any command prompt.
func (m Model) promptBox(inner string) string {
	style := promptStyle
	if m.width > 0 {
		innerWidth := m.width - style.GetHorizontalFrameSize()
		if innerWidth > 0 {
			style = style.Width(innerWidth)
		}
	}
	return style.Render(inner)
}
