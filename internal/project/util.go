package project

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

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

func initGoMod(projectName string, appDir string) {
	if err := executeCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		cobra.CheckErr(err)
	}
}

func goGetPackage(appDir, packageName string) {
	fmt.Printf("Package name is: %s\n", packageName)
	if err := executeCmd("go",
		[]string{"get", "-u", packageName},
		appDir); err != nil {
		cobra.CheckErr(err)
	}
}
