package TUI

import "github.com/charmbracelet/lipgloss"

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

// flashStyle is the brief "real command" toast after a successful action.
// Bold + teal so it stands out without a border or plain white.
var flashStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("80"))
