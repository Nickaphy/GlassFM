package FileSystemOperations

import "os"

// DeleteFile deletes the file at the specified path. If the file does not exist, it returns an error.
func DeleteFile(path string) error {
	return os.Remove(path)
}
