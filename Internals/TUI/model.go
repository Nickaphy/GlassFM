package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"

	tea "github.com/charmbracelet/bubbletea"
)

// uiMode says which input handler owns the keyboard.
// Still one Model — this is not a second app, just which keys mean what.
type uiMode int

const (
	modeBrowse uiMode = iota // file list: j/k, q, SPC
	modeLeader               // after SPC: collect a command name like "cd"
	modePrompt               // command line: type a path (or later, a name)
)

// Model is the TUI's memory: what is on screen, and what the next key should do.
// Disk work stays in FileSystemOperations; this struct only stores results.
type Model struct {
	items   []string // names from ListDir for the current folder
	cursor  int      // which row ">" sits on
	cwd     string   // path shown at the top
	width   int      // terminal width, from WindowSizeMsg
	mode    uiMode   // browse / leader / prompt
	leader  string   // partial command after SPC ("c", then "cd")
	input   string   // text currently in the command line
	command string   // command the prompt belongs to (e.g. "cd")
	status  string   // errors / hints in the footer
	flash   string   // brief "real command" toast after a successful action
	flashID int      // bumps each flash so stale timers don't clear new ones
}

// NewModel is startup only: first cwd + first listing. Later refreshes
// happen in submitPrompt (and will happen in Update after other actions).
func NewModel() Model {
	cwd, err := FileSystemOperations.GetWorkingDir()
	if err != nil {
		return Model{cwd: "unknown directory"}
	}

	items, err := FileSystemOperations.ListDir(cwd)
	if err != nil {
		items = nil
	}

	return Model{
		items: items,
		cwd:   cwd,
	}
}

// Init is Bubble Tea's one-shot setup. nil = no extra work at launch.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update is the traffic cop: classify the event, then hand keys to the
// handler for the current mode. It does not draw; View does that after.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil
	case flashExpiredMsg:
		return m.handleFlashExpired(msg)
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		switch m.mode {
		case modeLeader:
			return m.handleLeader(msg)
		case modePrompt:
			return m.handlePrompt(msg)
		default:
			return m.handleBrowse(msg)
		}
	}
	return m, nil
}

// afterAction reloads the listing, clears the prompt, and shows a brief
// toast of the real underlying command (GlassFM transparency).
func (m Model) afterAction(shown string) (tea.Model, tea.Cmd) {
	items, err := FileSystemOperations.ListDir(m.cwd)
	if err != nil {
		items = nil
	}
	m.items = items
	m.mode = modeBrowse
	m.input = ""
	m.command = ""
	m.status = ""
	m.flash = shown
	m.flashID++
	id := m.flashID
	return m, clearFlashAfter(id, flashDuration)
}
