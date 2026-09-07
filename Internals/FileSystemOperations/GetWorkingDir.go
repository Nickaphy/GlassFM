package FileSystemOperations

import "os"

// GetWorkingDir returns the current working directory of the process.
func GetWorkingDir() (string, error) {
	return os.Getwd()
}
