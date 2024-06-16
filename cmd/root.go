// Package cmd provides the command line interface for the application.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the main command for the application. It's invoked when no subcommands are specified.
var rootCmd = &cobra.Command{
	Use:   "goforge",
	Short: "A brief description of your application",
	Long: `A comprehensive description of your application, including examples and usage instructions. For instance:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line to associate an action with the root command:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute is the entry point for the CLI. It adds all child commands to the root command and sets flags as needed.
// It's invoked by the main function and should only be called once.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init sets up flags and configuration settings for the application.
// It's automatically called before the main function.
func init() {
	// Define global flags for your application here. Cobra supports persistent flags that are defined here.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goforge.yaml)")

	rootCmd.AddCommand(versionCommand)
	// Define local flags that are only valid for this command.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
