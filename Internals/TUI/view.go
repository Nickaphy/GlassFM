package TUI

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View draws the current Model as text. No state changes here.
func (m Model) View() string {
	title := titleStyle.Width(m.width).Align(lipgloss.Center).Render("GlassFM")
	cwdLine := lipgloss.NewStyle().Width(m.width).Align(lipgloss.Center).Render(m.cwd)

	s := title + "\n" + cwdLine + "\n\n"

	if len(m.items) == 0 {
		s += "  (empty)\n"
	}
	for i, item := range m.items {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		s += cursor + item + "\n"
	}

	s += "\n" + m.footer()
	return s
}

// footer is the mode-specific help line so the user can see what to type next.
func (m Model) footer() string {
	switch m.mode {
	case modeLeader:
		line := hintStyle.Render("SPC-" + m.leader)
		// Filtering the leaderCmds to only show the commands that match the leader
		for _, cmd := range leaderCmds {
			if m.leader == "" || strings.HasPrefix(cmd.name, m.leader) {
				line += "\n" + hintStyle.Render("  "+cmd.name+"  "+cmd.help)
			}
		}
		line += "\n" + hintStyle.Render("esc to cancel")
		return line
	case modePrompt:
		inner := m.command + " " + m.input + "█"
		if m.status != "" {
			inner += "\n" + m.status
		}
		return m.promptBox(inner) + "\n" + hintStyle.Render("enter to confirm, esc to cancel")
	default:
		hint := hintStyle.Render("SPC for commands, q to quit")
		if m.flash != "" {
			return flashStyle.Render(m.flash) + "\n" + hint
		}
		if m.status != "" {
			return m.status + "\n" + hint
		}
		return hint
	}
}
