package project

import (
	"bytes"
	"fmt"
	"os/exec"
)

// executeCmd runs a command with given arguments in a specified directory.
// It returns an error if the command execution fails.
func executeCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

// initGoMod initializes a new Go module in the specified directory.
// It returns an error if the module initialization fails.
func initGoMod(projectName string, appDir string) error {
	if err := executeCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		return err
	}
	return nil
}

// goGetPackage fetches the specified Go package and updates it.
// It returns an error if the package fetching fails.
func goGetPackage(appDir, packageName string) error {
	fmt.Printf("Package name is: %s\n", packageName)
	if err := executeCmd("go",
		[]string{"get", "-u", packageName},
		appDir); err != nil {
		return err
	}
	return nil
}

// goFormat formats the Go source files in the specified directory using gofmt.
// It returns an error if the formatting fails.
func goFormat(appDir string) error {
	if err := executeCmd("gofmt",
		[]string{"-s", "-w", "."},
		appDir); err != nil {
		return err
	}

	return nil
}
