package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().Bold(true)

// hintStyle is helper/info text. Faint() stays whitish on many terminals;
// a dark grey of the 256-color ramp actually recedes.
var hintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))

// promptStyle wraps every command line (cd mk/rn/… ) so input
// reads as a distinct box, not another footer line.
var promptStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("244")).
	Padding(0, 1)

// uiMode says which input handler owns the keyboard.
// Still one Model — this is not a second app, just which keys mean what.
type uiMode int

const (
	modeBrowse uiMode = iota // file list: j/k, q, SPC
	modeLeader               // after SPC: collect a command name like "cd"
	modePrompt               // command line: type a path (or later, a name)
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
	{name: "mk", help: "create dir"},
	{name: "mv", help: "move selected"},
	{name: "cp", help: "copy selected"},
	{name: "rn", help: "rename selected"},
	{name: "rm", help: "delete selected)"},
}

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

// submitPrompt runs the finished command then reloads
// cwd + items so View shows the new directory. Failures stay in the prompt.
func (m Model) submitPrompt() (tea.Model, tea.Cmd) {
	if m.command != "cd" {
		m.mode = modeBrowse
		return m, nil
	}

	path := strings.TrimSpace(m.input)
	if path == "" {
		m.status = "type a path, then enter"
		return m, nil
	}

	if err := FileSystemOperations.ChangeDir(path); err != nil {
		m.status = err.Error()
		return m, nil
	}

	cwd, err := FileSystemOperations.GetWorkingDir()
	if err != nil {
		m.status = err.Error()
		return m, nil
	}
	items, err := FileSystemOperations.ListDir(cwd)
	if err != nil {
		items = nil
	}

	m.cwd = cwd
	m.items = items
	m.cursor = 0
	m.mode = modeBrowse
	m.input = ""
	m.command = ""
	m.status = ""
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
		if m.status != "" {
			return m.status + "\n" + hint
		}
		return hint
	}
}

// shares prompt-box for command prompts
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
