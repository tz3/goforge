// Package cmd provides the command line interface for the application.
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/tz3/goforge/cmd/ui/multiinput"
	"github.com/tz3/goforge/cmd/ui/textinput"
	"github.com/tz3/goforge/internal/project"
	"github.com/tz3/goforge/internal/steps"
)

// Options represents the options for initializing the steps.
// It includes the name of the project and the type of the project.
type Options struct {
	ProjectName    *textinput.Output
	ProjectType    *multiinput.Selection
	DatabaseDriver *multiinput.Selection
}

// logo is the ASCII representation of the application logo.
const logo = `

░██████╗░░█████╗░███████╗░█████╗░██████╗░░██████╗░███████╗
██╔════╝░██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝░██╔════╝
██║░░██╗░██║░░██║█████╗░░██║░░██║██████╔╝██║░░██╗░█████╗░░
██║░░╚██╗██║░░██║██╔══╝░░██║░░██║██╔══██╗██║░░╚██╗██╔══╝░░
╚██████╔╝╚█████╔╝██║░░░░░╚█████╔╝██║░░██║╚██████╔╝███████╗
░╚═════╝░░╚════╝░╚═╝░░░░░░╚════╝░╚═╝░░╚═╝░╚═════╝░╚══════╝

`

const (
	defaultProjectTitle        = "goforge"
	flagProjectTitleKey        = "title"
	flagProjectWebFrameworkKey = "framework"
	flagDatabaseDriverKey      = "databaseDriver"
)

// Styles for rendering the logo and ending message.
var (
	logoStyle       = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#00BFFF")).Bold(true)
	endingMsgStyle  = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FFD700")).Bold(true).Underline(true)
	tipMessageStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Background(lipgloss.Color("235")).Italic(true)
)

// Initialize the command and flags.
func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP(flagProjectTitleKey, "t", "", "Title/name of the project to create")
	createCmd.Flags().StringP(flagProjectWebFrameworkKey, "f", "", fmt.Sprintf("Type of web-framework to use as a router. Allowed values: %s", strings.Join(project.SupportedWebframeworks, ", ")))
	createCmd.Flags().StringP(flagDatabaseDriverKey, "d", "", fmt.Sprintf("Database driver to use as main DB. Allowed DBs: %s", strings.Join(project.SupportedDatabaseDrivers, ", ")))
}

// createCmd is the command to create a new Go project.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project without worrying about the structure",
	Long: `GoForge is a CLI tool that allows you to focus on the actual Go code, 
	and not the project structure. Perfect for someone new to the Go language`,

	Run: func(cmd *cobra.Command, args []string) {
		options := Options{
			ProjectName:    &textinput.Output{},
			ProjectType:    &multiinput.Selection{},
			DatabaseDriver: &multiinput.Selection{},
		}

		isInteractive := !hasChangedFlag(cmd.Flags())

		// Retrieve flag values
		flagTitleValue := cmd.Flag(flagProjectTitleKey).Value.String()
		flagFrameworkValue := cmd.Flag(flagProjectWebFrameworkKey).Value.String()
		flagDatabaseDriverValue := cmd.Flag(flagDatabaseDriverKey).Value.String()

		// Validate input
		if flagTitleValue != "" {
			validateFlags(flagTitleValue, flagFrameworkValue, flagDatabaseDriverValue)
		}

		projectConfig := &project.ProjectConfig{
			ProjectName:       flagTitleValue,
			FrameworkMap:      make(map[string]project.WebFramework),
			ProjectType:       flagFrameworkValue,
			DatabaseDriverMap: make(map[string]project.DatabaseDriver),
			DatabaseDriver:    flagDatabaseDriverValue,
		}

		steps := steps.InitSteps()
		fmt.Printf("%s\n", logoStyle.Render(logo))

		if projectConfig.ProjectName == "" {
			handleInteractiveProjectName(options, projectConfig, cmd)
		}

		if projectConfig.ProjectType == "" {
			handleInteractiveProjectType(options, projectConfig, cmd, steps)
		}

		if projectConfig.DatabaseDriver == "" {
			handleInteractiveDatabaseDriver(options, projectConfig, cmd, steps)
		}

		setupProject(projectConfig)

		fmt.Println(endingMsgStyle.Render("\nNext steps: cd into the newly created project with:"))
		fmt.Println(endingMsgStyle.Render(fmt.Sprintf("• cd %s\n", projectConfig.ProjectName)))

		if isInteractive {
			fmt.Println(tipMessageStyle.Render("Tip: Repeat the equivalent Blueprint with the following non-interactive command:"))
			fmt.Println(tipMessageStyle.Italic(false).Render(fmt.Sprintf("• %s\n", nonInteractiveCommand(cmd.Flags()))))
		}
	},
}

// nonInteractiveCommand generates a non-interactive command string from the flag set.
func nonInteractiveCommand(flagSet *pflag.FlagSet) string {
	nonInteractiveCommand := defaultProjectTitle
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Name != "help" {
			nonInteractiveCommand = fmt.Sprintf("%s --%s %s", nonInteractiveCommand, flag.Name, flag.Value.String())
		}
	})
	return nonInteractiveCommand
}

// hasChangedFlag checks if any flag in the FlagSet has been set by the user.
func hasChangedFlag(flagSet *pflag.FlagSet) bool {
	hasChangedFlag := false
	flagSet.Visit(func(_ *pflag.Flag) {
		hasChangedFlag = true
	})
	return hasChangedFlag
}

// isDirectoryNonEmpty checks if a directory exists and isn't empty.
func isDirectoryNonEmpty(dirName string) bool {
	if _, err := os.Stat(dirName); err == nil {
		dirEntries, err := os.ReadDir(dirName)
		if err != nil {
			log.Printf("Could not read directory: %v", err)
			cobra.CheckErr(fmt.Errorf("could not read directory: %v", err))
		}
		return len(dirEntries) > 0
	}
	return false
}

// isValidProjectName checks if a string only contains alphanumeric characters.
func isValidProjectName(input string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9/-]*$", input)
	return match
}

// validateFlags validates the input flags for the project.
func validateFlags(title, framework, databaseDriver string) {
	if !isValidProjectName(title) {
		cobra.CheckErr(fmt.Errorf("input '%s' contains non-alphanumeric characters", title))
	}
	if !project.IsValidWebFramework(framework) {
		cobra.CheckErr(fmt.Errorf("invalid web framework: %s", framework))
	}
	if !project.IsValidDatabaseDriver(databaseDriver) {
		cobra.CheckErr(fmt.Errorf("invalid database driver: %s. Supported drivers are: %s", databaseDriver, strings.Join(project.SupportedDatabaseDrivers, ", ")))
	}
	if isDirectoryNonEmpty(title) {
		cobra.CheckErr(fmt.Errorf("directory '%s' already exists and is not empty. Please choose a different name", title))
	}
}

// handleInteractiveProjectName handles interactive input for the project name.
func handleInteractiveProjectName(options Options, projectConfig *project.ProjectConfig, cmd *cobra.Command) {
	tprogram := tea.NewProgram(textinput.InitialTextInputModel(options.ProjectName, "What is the name of your project?", projectConfig))
	if _, err := tprogram.Run(); err != nil {
		log.Printf("Error in project name input: %v", err)
		cobra.CheckErr(fmt.Errorf("error in project name input: %v", err))
	}
	if isDirectoryNonEmpty(options.ProjectName.Output) {
		cobra.CheckErr(fmt.Errorf("directory '%s' already exists and is not empty. Please choose a different name", options.ProjectName.Output))
	}
	if !isValidProjectName(options.ProjectName.Output) {
		cobra.CheckErr(fmt.Errorf("input '%s' contains non-alphanumeric characters", options.ProjectName.Output))
	}
	projectConfig.ExitCLI(tprogram)
	projectConfig.ProjectName = options.ProjectName.Output
	setFlagValue(cmd, flagProjectTitleKey, projectConfig.ProjectName)
}

// handleInteractiveProjectType handles interactive input for the project type.
func handleInteractiveProjectType(options Options, projectConfig *project.ProjectConfig, cmd *cobra.Command, steps *steps.Steps) {
	step := steps.Steps["web-framework"]
	tprogram := tea.NewProgram(multiinput.InitialModelMulti(step.Options, options.ProjectType, step.Headers, projectConfig))
	if _, err := tprogram.Run(); err != nil {
		log.Printf("Error in web framework input: %v", err)
		cobra.CheckErr(fmt.Errorf("error in web framework input: %v", err))
	}
	projectConfig.ExitCLI(tprogram)
	projectConfig.ProjectType = strings.ToLower(options.ProjectType.Choice)
	setFlagValue(cmd, flagProjectWebFrameworkKey, projectConfig.ProjectType)
}

// handleInteractiveDatabaseDriver handles interactive input for the database driver.
func handleInteractiveDatabaseDriver(options Options, projectConfig *project.ProjectConfig, cmd *cobra.Command, steps *steps.Steps) {
	step := steps.Steps["db-driver"]
	tprogram := tea.NewProgram(multiinput.InitialModelMulti(step.Options, options.DatabaseDriver, step.Headers, projectConfig))
	if _, err := tprogram.Run(); err != nil {
		log.Printf("Error in database driver input: %v", err)
		cobra.CheckErr(fmt.Errorf("error in database driver input: %v", err))
	}
	projectConfig.ExitCLI(tprogram)
	projectConfig.DatabaseDriver = strings.ToLower(options.DatabaseDriver.Choice)
	setFlagValue(cmd, flagDatabaseDriverKey, projectConfig.DatabaseDriver)
}

// setupProject sets up the project configuration and creates necessary files.
func setupProject(projectConfig *project.ProjectConfig) {
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		log.Printf("Could not get current working directory: %v", err)
		cobra.CheckErr(fmt.Errorf("could not get current working directory: %v", err))
	}
	projectConfig.AbsolutePath = currentWorkingDir
	if err := projectConfig.CreateMainFile(); err != nil {
		log.Printf("Problem creating files for project configuration: %v", err)
		cobra.CheckErr(fmt.Errorf("problem creating files for project configuration: %v", err))
	}
}

// setFlagValue sets the value of a command flag.
func setFlagValue(cmd *cobra.Command, flagName, value string) {
	if err := cmd.Flag(flagName).Value.Set(value); err != nil {
		log.Printf("Failed to set %s flag: %v", flagName, err)
		cobra.CheckErr(fmt.Errorf("failed to set %s flag: %v", flagName, err))
	}
}
