package FileSystemOperations

import (
	"os"
)

// List the contents of a directory and return the names of the files and directories as a slice of strings
func ListDir(filePath string) ([]string, error) {
	entries, err := os.ReadDir(filePath)
	if err != nil {
		// if err != nil return no array and the error 
		return nil, err
	}
	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	// if no error return the array and nil errors
	return names, nil
}
