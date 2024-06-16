/*
GoForge version
*/
package cmd

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/cobra"
)

// GoforgeVersion is the version of the CLI, updated by goreleaser during a CI run with the release version on GitHub
var GoforgeVersion string

func getGoForgeVersion() string {
	if len(GoforgeVersion) != 0 {
		return GoforgeVersion
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return noVersionInfo()
	}

	if bi.Main.Version != "(devel)" {
		return bi.Main.Version
	}

	vcsRevision, vcsTime := getVCSInfo(bi.Settings)
	if vcsRevision != "" {
		return fmt.Sprintf("%s, (%s)", vcsRevision, vcsTime)
	}

	return noVersionInfo()
}

func noVersionInfo() string {
	return "No version info available for this build, run 'goforge help version' for additional info"
}

func getVCSInfo(settings []debug.BuildSetting) (string, time.Time) {
	var vcsRevision string
	var vcsTime time.Time
	for _, setting := range settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		}
	}
	return vcsRevision, vcsTime
}

// versionCommand represents the command that displays the version of the application
var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show the application version.",
	Long:  "The 'version' command displays the current version of this Go CLI application.",
	Run: func(cmd *cobra.Command, args []string) {
		version := getGoForgeVersion()
		fmt.Printf("GoForge CLI version: %v\n", version)
	},
}
