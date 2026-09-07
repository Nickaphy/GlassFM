package FileSystemOperations

import "os"

// CreateFile creates a new file at the specified path. If the file already exists, it will be truncated. If the file cannot be created, it returns an error.
func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return file.Close()
}
