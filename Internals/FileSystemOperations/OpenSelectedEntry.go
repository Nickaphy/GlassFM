package FileSystemOperations

import "os/exec"

// Helper function for detecting which file type is selected, so we can open it in the correct application
func OpenSelectedEntry(filePath string) error {
	// Open the file in the default application for its file type.
	return exec.Command("xdg-open", filePath).Start()
}