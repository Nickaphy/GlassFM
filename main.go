package main

import (
	"CLIfileManager/Internals/FileSystemOperations"
	"fmt"
)

func main() {
	// quick manual smoke test for GetWorkingDir - remove once the TUI drives these

	wd, err := FileSystemOperations.GetWorkingDir()
	if err != nil {
		fmt.Println("GetWorkingDir error:", err)
		return
	}
	fmt.Println("current working dir:", wd)
}
