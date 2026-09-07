package FileSystemOperations

import "os"

// Moves a file from src to dst. If dst already exists, it will be replaced.
func Move(src string, dst string) error {
	return os.Rename(src, dst)
}
