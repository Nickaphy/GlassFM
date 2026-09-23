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
	instant     bool // true = run on match, no prompt (e.g. toggles)
}

// leaderCmds is the vocabulary after SPC. matchLeader uses this list;
// submitPrompt is what actually runs a finished command.
var leaderCmds = []leaderCmd{
	{name: "cd", help: "jump to arbitrary path (type it)", implemented: true},
	{name: "h", help: "toggle hidden files", implemented: true, instant: true},
	{name: "mk", help: "create dir (fails if it already exists)", implemented: true},
	{name: "mv", help: "move selected (path or browse)", implemented: true},
	{name: "cp", help: "copy selected (path or browse)", implemented: true},
	{name: "rn", help: "rename selected", implemented: true},
	{name: "rm", help: "delete selected (with confirm)", implemented: true},
}

// handleLeader: build a command name one letter at a time.
// Esc / unknown sequence → back to browse. A full match → prompt or instant action.
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
		cmd := m.leader
		m.leader = ""
		if leaderInstant(cmd) {
			return m.runInstant(cmd)
		}
		if cmd == "cp" {
			return m.startCopy()
		}
		if cmd == "mv" {
			return m.startMove()
		}
		if cmd == "rn" {
			return m.startRename()
		}
		if cmd == "rm" {
			return m.startRemove()
		}
		m.command = cmd
		m.input = ""
		m.status = ""
		m.mode = modePrompt
	}
	return m, nil
}

// matchLeader checks if the input matches a leader command.
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

// leaderInstant returns true if the command is an instant action.
func leaderInstant(name string) bool {
	for _, cmd := range leaderCmds {
		if cmd.name == name {
			return cmd.instant
		}
	}
	return false
}

// runInstant runs an instant action (no prompt).
func (m Model) runInstant(name string) (tea.Model, tea.Cmd) {
	switch name {
	case "h":
		return m.toggleHidden()
	default:
		m.mode = modeBrowse
		m.status = name + " is not implemented yet"
		return m, nil
	}
}
