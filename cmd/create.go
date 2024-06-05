package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/tz3/goforge/cmd/ui/multiinput"
	"github.com/tz3/goforge/cmd/ui/textinput"
	"github.com/tz3/goforge/internal/project"
	"github.com/tz3/goforge/internal/steps"
)

const logo = `

░██████╗░░█████╗░███████╗░█████╗░██████╗░░██████╗░███████╗
██╔════╝░██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝░██╔════╝
██║░░██╗░██║░░██║█████╗░░██║░░██║██████╔╝██║░░██╗░█████╗░░
██║░░╚██╗██║░░██║██╔══╝░░██║░░██║██╔══██╗██║░░╚██╗██╔══╝░░
╚██████╔╝╚█████╔╝██║░░░░░╚█████╔╝██║░░██║╚██████╔╝███████╗
░╚═════╝░░╚════╝░╚═╝░░░░░░╚════╝░╚═╝░░╚═╝░╚═════╝░╚══════╝

`

var (
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#FFD700")).Bold(true)
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Go Blueprint is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	Run: func(cmd *cobra.Command, args []string) {

		options := steps.Options{
			ProjectName: &textinput.Output{},
		}

		projectConfig := &project.ProjectConfig{
			FrameworkMap: make(map[string]project.WebFramework),
		}

		steps := steps.InitSteps(&options)

		fmt.Printf("%s\n", logoStyle.Render(logo))

		tprogram := tea.NewProgram(textinput.InitialTextInputModel(options.ProjectName, "What is the name of your project?", projectConfig))
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}
		projectConfig.ExitCLI(tprogram)

		for _, step := range steps.Steps {
			s := &multiinput.Selection{}
			tprogram = tea.NewProgram(multiinput.InitialModelMulti(step.Options, s, step.Headers, projectConfig))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(err)
			}
			projectConfig.ExitCLI(tprogram)

			*step.Field = s.Choice
		}

		projectConfig.ProjectName = options.ProjectName.Output
		projectConfig.ProjectType = options.ProjectType
		fmt.Printf("The project router framework is: %s\n", projectConfig.ProjectType)
		currentWorkingDir, err := os.Getwd()
		if err != nil {
			cobra.CheckErr(err)
		}

		projectConfig.AbsolutePath = currentWorkingDir

		// This calls the templates
		err = projectConfig.CreateMainFile()
		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Printf("%s\n%s cd %s\n",
			endingMsgStyle.Render("To get into the project directory:"),
			endingMsgStyle.Render(multiinput.Bullet),
			endingMsgStyle.Render(projectConfig.ProjectName))
	},
}
