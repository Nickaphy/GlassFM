package FileSystemOperations

import "os"

// CreateDir creates a new directory at the specified path. If the directory already exists, it returns an error.
func CreateDir(path string) error {
	return os.Mkdir(path, 0755)
}
