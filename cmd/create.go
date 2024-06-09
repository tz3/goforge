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
	flagProjectTitle        = "title"
	flagProjectWebFramework = "framework"
)

// logoStyle and endingMsgStyle are lipgloss styles for rendering the logo and ending message.
var (
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FFD700")).Bold(true)
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
	Long: `Go Blueprint is a CLI tool that allows you to focus on the actual Go code, 
	and not the project structure. Perfect for someone new to the Go language`,

	Run: func(cmd *cobra.Command, args []string) {
		options := steps.Options{
			ProjectName: &textinput.Output{},
		}

		// check flags
		flagTitle := cmd.Flag(flagProjectTitle).Value.String()
		flagFramework := cmd.Flag(flagProjectWebFramework).Value.String()

		// Todo:- more edge cases need to be covered!
		if flagTitle != "" && !project.IsValidWebFramework(flagFramework) {
			cobra.CheckErr(fmt.Errorf("web framework '%s' does not exist. supported webframeworks are: %s", flagFramework, strings.Join(project.SupportedWebframeworks, ", ")))
		}

		projectConfig := &project.ProjectConfig{
			FrameworkMap: make(map[string]project.WebFramework),
			ProjectName:  flagTitle,
			ProjectType:  flagFramework,
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

		fmt.Printf("%s\n%s cd %s\n",
			endingMsgStyle.Render("To get into the project directory:"),
			endingMsgStyle.Render(multiinput.Bullet),
			endingMsgStyle.Render(projectConfig.ProjectName))
	},
}
