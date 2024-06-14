// Package cmd provides the command line interface for the application.
package cmd

import (
	"fmt"
	"log"
	"os"
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
	defaultProjectTitle     = "goforge"
	flagProjectTitle        = "title"
	flagProjectWebFramework = "framework"
)

// logoStyle and endingMsgStyle are lipgloss styles for rendering the logo and ending message.
var (
	logoStyle                   = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#00BFFF")).Bold(true)
	endingMsgStyle              = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FFD700")).Bold(true).Underline(true)
	endingMsgStyleWithUnderline = endingMsgStyle.Underline(true)
	tipMessageStyle             = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Background(lipgloss.Color("235")).Italic(true)
)

// init function adds the create command to the root command.
func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP(flagProjectTitle, "n", "", "Name of project to create")
	createCmd.Flags().StringP(flagProjectWebFramework, "t", "", fmt.Sprintf("Type of web-framework to use as a router. Allowed values: %s", strings.Join(project.SupportedWebframeworks, ", ")))
}

// createCmd is a cobra command that creates a new Go project.
// It asks for the project name and other configurations via the CLI.
// It then creates the main file for the project.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long: `GoForge is a CLI tool that allows you to focus on the actual Go code, 
	and not the project structure. Perfect for someone new to the Go language`,

	Run: func(cmd *cobra.Command, args []string) {
		options := steps.Options{
			ProjectName: &textinput.Output{},
		}

		isInteractive := !hasChangedFlag(cmd.Flags())

		// check flags
		flagTitleValue := cmd.Flag(flagProjectTitle).Value.String()
		flagFrameworkValue := cmd.Flag(flagProjectWebFramework).Value.String()

		// Todo:- more edge cases need to be covered!
		if flagTitleValue != "" && !project.IsValidWebFramework(flagFrameworkValue) {
			cobra.CheckErr(fmt.Errorf("web framework '%s' does not exist. supported webframeworks are: %s", flagFrameworkValue, strings.Join(project.SupportedWebframeworks, ", ")))
		}

		projectConfig := &project.ProjectConfig{
			FrameworkMap: make(map[string]project.WebFramework),
			ProjectName:  flagTitleValue,
			ProjectType:  flagFrameworkValue,
		}

		steps := steps.InitSteps(&options)

		fmt.Printf("%s\n", logoStyle.Render(logo))

		if projectConfig.ProjectName == "" && projectConfig.ProjectType == "" { // user didn't use flags
			tprogram := tea.NewProgram(textinput.InitialTextInputModel(options.ProjectName, "What is the name of your project?", projectConfig))
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Failed to run the program: %v\n", err)
				cobra.CheckErr(err)
			}
			projectConfig.ExitCLI(tprogram)

			for _, step := range steps.Steps {
				s := &multiinput.Selection{}
				tprogram = tea.NewProgram(multiinput.InitialModelMulti(step.Options, s, step.Headers, projectConfig))
				if _, err := tprogram.Run(); err != nil {
					log.Printf("Failed to run the program for step %s: %v\n", step.Headers, err)
					cobra.CheckErr(err)
				}
				projectConfig.ExitCLI(tprogram)

				*step.Field = s.Choice
			}
			projectConfig.ProjectName = options.ProjectName.Output
			projectConfig.ProjectType = options.ProjectType

			cmd.Flag(flagProjectTitle).Value.Set(projectConfig.ProjectName)
			cmd.Flag(flagProjectWebFramework).Value.Set(projectConfig.ProjectType)
		}

		fmt.Printf("The project router framework is: %s\n", projectConfig.ProjectType)
		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Printf("Failed to get the current working directory: %v\n", err)
			cobra.CheckErr(err)
		}

		projectConfig.AbsolutePath = currentWorkingDir

		// This calls the templates
		err = projectConfig.CreateMainFile()
		if err != nil {
			log.Printf("Failed to create the main file: %v\n", err)
			cobra.CheckErr(err)
		}

		fmt.Println("")

		fmt.Printf("%s\n%s cd %s\n",
			endingMsgStyle.Render("To get into the project directory:"),
			endingMsgStyle.Render(multiinput.Bullet),
			endingMsgStyleWithUnderline.Render(projectConfig.ProjectName))

		fmt.Println("")

		if isInteractive {
			nonInterActiveCmd := nonInteractiveCommand(cmd.Flags())
			fmt.Println(tipMessageStyle.Render("Tip: Repeat the equivalent cmd with the following non-interactive command:"))
			fmt.Println(tipMessageStyle.Italic(false).Render(fmt.Sprintf("• %s\n", nonInterActiveCmd)))
		}
	},
}

// nonInteractiveCommand takes a pointer to a pflag.FlagSet as an argument.
// It iterates over all flags in the FlagSet, excluding the "help" flag, and constructs a command string.
// This command string represents the equivalent non-interactive shell command for the given FlagSet.
// The function returns this command string.
// The flags in the FlagSet are not sorted before the iteration.
func nonInteractiveCommand(flagSet *pflag.FlagSet) string {
	nonInteractiveCommand := defaultProjectTitle
	visitFn := func(flag *pflag.Flag) {
		if flag.Name != "help" {
			nonInteractiveCommand = fmt.Sprintf("%s --%s %s", nonInteractiveCommand, flag.Name, flag.Value.String())
		}
	}

	flagSet.SortFlags = false
	flagSet.VisitAll(visitFn)

	return nonInteractiveCommand
}

// hasChangedFlag takes a pointer to a pflag.FlagSet as an argument.
// It checks if any flag in the FlagSet has been set by the user.
// The function returns true if at least one flag has been set, and false otherwise.
func hasChangedFlag(flagSet *pflag.FlagSet) bool {
	hasChangedFlag := false
	flagSet.Visit(func(_ *pflag.Flag) {
		hasChangedFlag = true
	})
	return hasChangedFlag
}
