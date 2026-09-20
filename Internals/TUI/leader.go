package TUI

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// leaderCmd is one row in the SPC helper. help text matches.
type leaderCmd struct {
	name        string
	help        string
	implemented bool
}

// leaderCmds is the vocabulary after SPC. matchLeader uses this list;
// submitPrompt is what actually runs a finished command.
var leaderCmds = []leaderCmd{
	{name: "cd", help: "jump to arbitrary path (type it)", implemented: true},
	{name: "mk", help: "create dir (fails if it already exists)", implemented: true},
	{name: "mv", help: "move selected"},
	{name: "cp", help: "copy selected"},
	{name: "rn", help: "rename selected"},
	{name: "rm", help: "delete selected (with confirm)"},
}

// handleLeader: build a command name one letter at a time.
// Esc / unknown sequence → back to browse. A full match → open the prompt.
func (m Model) handleLeader(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.mode = modeBrowse
		m.leader = ""
		return m, nil
	}

	key := msg.String()
	if len(key) != 1 {
		return m, nil
	}

	m.leader += key
	if done, ok := matchLeader(m.leader); !ok {
		m.mode = modeBrowse
		m.leader = ""
		m.status = "unknown command"
		return m, nil
	} else if done {
		if !leaderImplemented(m.leader) {
			m.status = m.leader + " is not implemented yet"
			m.mode = modeBrowse
			m.leader = ""
			return m, nil
		}
		m.command = m.leader
		m.leader = ""
		m.input = ""
		m.status = ""
		m.mode = modePrompt
	}
	return m, nil
}

func matchLeader(buf string) (done bool, valid bool) {
	for _, cmd := range leaderCmds {
		if buf == cmd.name {
			return true, true
		}
		if strings.HasPrefix(cmd.name, buf) {
			valid = true
		}
	}
	return false, valid
}

func leaderImplemented(name string) bool {
	for _, cmd := range leaderCmds {
		if cmd.name == name {
			return cmd.implemented
		}
	}
	return false
}
