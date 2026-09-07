package FileSystemOperations

import (
	"io"
	"os"
)

// Copy copies the contents of the source file to the destination file. If the destination file already exists, it will be truncated. It returns an error if any operation fails.
func Copy(src string, dst string) error { // Takes 2 paths and returns an error
	srcFile, err := os.Open(src) // Open source file in readonly
	if err != nil {
		return err
	}
	defer srcFile.Close() // Close the source file when the function returns

	dstFile, err := os.Create(dst) // Create destination file, if it exists it will be truncated
	if err != nil {
		return err
	}
	defer dstFile.Close() // Close the destination file when the function returns

	_, err = io.Copy(dstFile, srcFile) // Copy the contents of the source file to the destination file
	return err
}
