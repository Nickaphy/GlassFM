package FileSystemOperations

import "os"

// ChangeDir sets the process working directory, like the shell's cd.
func ChangeDir(path string) error {
	return os.Chdir(path)
}
