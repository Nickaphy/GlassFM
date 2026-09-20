package main

import (
	"CLIfileManager/Internals/TUI"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Clear the terminal
	fmt.Println("\033[2J\033[H")
	// Entry point to TUI GUI
	p := tea.NewProgram(TUI.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
