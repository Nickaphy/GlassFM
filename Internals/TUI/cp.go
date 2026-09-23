package TUI

import (
	"CLIfileManager/Internals/FileSystemOperations"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// startCopy remembers the selected item and opens the destination prompt.
// From there: type a path, or tab to pick a folder by browsing.
func (m Model) startCopy() (tea.Model, tea.Cmd) {
	if len(m.items) == 0 {
		m.mode = modeBrowse
		m.status = "nothing selected to copy"
		return m, nil
	}

	src := filepath.Join(m.cwd, m.items[m.cursor])
	if abs, err := filepath.Abs(src); err == nil {
		src = abs
	}

	m.pendingSrc = src
	m.command = "cp"
	m.input = ""
	m.status = ""
	m.mode = modePrompt
	return m, nil
}

// submitCp copies pendingSrc to dst (file path, or into a directory).
func (m Model) submitCp(dst string) (tea.Model, tea.Cmd) {
	if m.pendingSrc == "" {
		m.status = "no source to copy"
		return m, nil
	}

	dest := resolvePendingDest(m.pendingSrc, dst)
	if err := FileSystemOperations.Copy(m.pendingSrc, dest); err != nil {
		m.status = err.Error()
		return m, nil
	}

	if abs, err := filepath.Abs(dest); err == nil {
		dest = abs
	}
	return m.afterAction("cp " + m.pendingSrc + " " + dest)
}

func (m Model) cancelPending() (tea.Model, tea.Cmd) {
	m.pendingSrc = ""
	m.command = ""
	m.input = ""
	m.status = ""
	m.mode = modeBrowse
	return m, nil
}

// resolvePendingDest: if dst is an existing directory, place basename(src) inside it;
// otherwise treat dst as the full destination path.
func resolvePendingDest(src, dst string) string {
	info, err := os.Stat(dst)
	if err == nil && info.IsDir() {
		return filepath.Join(dst, filepath.Base(src))
	}
	return dst
}
