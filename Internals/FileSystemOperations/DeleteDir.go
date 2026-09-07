package FileSystemOperations

import "os"

// DeleteDir deletes the directory at the specified path and all its contents. If the directory does not exist, it returns an error.
func DeleteDir(path string) error {
	return os.RemoveAll(path)
}
