package TUI

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const flashDuration = 2 * time.Second

// flashExpiredMsg clears a toast when its timer fires. flashID must match
// so an older timer cannot wipe a newer flash.
type flashExpiredMsg struct{ id int }

func clearFlashAfter(id int, d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return flashExpiredMsg{id: id}
	})
}

func (m Model) handleFlashExpired(msg flashExpiredMsg) (tea.Model, tea.Cmd) {
	if msg.id == m.flashID {
		m.flash = ""
	}
	return m, nil
}
