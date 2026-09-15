package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().Bold(true)

// Cursor
type Model struct {
	items  []string
	cursor int
	cwd    string
	width  int
}

// Temporary placeholder lines for testing j/k navigation
func NewModel() Model {
	cwd, err := FileSystemOperations.GetWorkingDir()
	if err != nil {
		cwd = "unknown directory"
	}
	return Model{
		items: []string{"placeholder-1", "placeholder-2", "placeholder-3"},
		cwd:   cwd,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Responsible for navigation
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg: // sent automatically on startup and on resize
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c": // Wires q up to ctrl+c (terminal exit)
			return m, tea.Quit
		case "up", "k": // Vim motion wiring for k to go up, by using our cursor int position.
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j": // Vim motion wiring for j to go down by using our cursor int position.
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

// Responsible for drawing UI
func (m Model) View() string {
	title := titleStyle.Width(m.width).Align(lipgloss.Center).Render("GlassFM")
	cwdLine := lipgloss.NewStyle().Width(m.width).Align(lipgloss.Center).Render(m.cwd)

	s := title + "\n" + cwdLine + "\n\n"
	for i, item := range m.items { // Deciding if the cursor should be blank or >
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		s += cursor + item + "\n"
	}
	s += "\npress q to quit\n"
	return s
}
