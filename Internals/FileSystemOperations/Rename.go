package FileSystemOperations

import "os"

// Renames (moves) a file or directory from oldPath to newPath. If newPath already exists, it will be replaced.
func Rename(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}
