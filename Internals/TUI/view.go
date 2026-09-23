package TUI

import (
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// contentWidth is how wide the list/prompt column is before we center it.
const contentMaxWidth = 64

// View draws the current Model as text. No state changes here.
func (m Model) View() string {
	title := titleStyle.Width(m.width).Align(lipgloss.Center).Render("GlassFM")
	cwdLine := lipgloss.NewStyle().Width(m.width).Align(lipgloss.Center).Render(m.cwd)

	var list strings.Builder
	if len(m.items) == 0 {
		list.WriteString("  (empty)\n")
	}
	for i, item := range m.items {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		list.WriteString(cursor + item + "\n")
	}

	s := title + "\n" + cwdLine + "\n"
	s += m.centerBlock(strings.TrimRight(list.String(), "\n"))
	s += "\n\n" + m.footer()
	return s
}

// footer is the mode-specific help line so the user can see what to type next.
func (m Model) footer() string {
	switch m.mode {
	case modeLeader:
		line := "SPC-" + m.leader
		for _, cmd := range leaderCmds {
			if m.leader == "" || strings.HasPrefix(cmd.name, m.leader) {
				line += "\n  " + cmd.name + "  " + cmd.help
			}
		}
		line += "\nesc to cancel"
		return m.centerBlock(hintStyle.Render(line))
	case modePrompt:
		inner := m.command + " " + m.input + "█"
		if m.command == "rm" {
			inner = "rm " + filepath.Base(m.pendingSrc) + "?"
		}
		if m.status != "" {
			inner += "\n" + m.status
		}
		hint := "enter to confirm, esc to cancel"
		switch m.command {
		case "cp", "mv":
			hint = "enter path · tab to browse · esc cancel"
		case "rn":
			hint = "edit name · enter · esc cancel"
		case "rm":
			hint = "enter to delete · esc cancel"
		}
		return m.promptBox(inner) + "\n" + m.centerBlock(hintStyle.Render(hint))
	case modePickDest:
		return m.pickFooter()
	default:
		hint := hintStyle.Render("SPC for commands, q to quit")
		if m.flash != "" {
			return m.centerBlock(flashStyle.Render(m.flash) + "\n" + hint)
		}
		if m.status != "" {
			return m.centerBlock(m.status + "\n" + hint)
		}
		return m.centerBlock(hint)
	}
}

// centerBlock places a left-aligned content column in the horizontal middle.
func (m Model) centerBlock(block string) string {
	if m.width <= 0 {
		return block
	}
	col := m.contentColWidth()
	framed := lipgloss.NewStyle().Width(col).Render(block)
	return lipgloss.PlaceHorizontal(m.width, lipgloss.Center, framed)
}

func (m Model) contentColWidth() int {
	w := m.width
	if w > contentMaxWidth {
		w = contentMaxWidth
	}
	if w < 1 {
		return 1
	}
	return w
}
