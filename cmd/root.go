package cmd

import (
	"geo/pkg/tui/command_list"
	"github.com/spf13/cobra"
	"os"
)

func geoCmd() *cobra.Command {
	cmd := &cobra.Command{
		// Version field will be dynamically set
		// Version: "",
		Use: "geo",
		Run: func(cmd *cobra.Command, args []string) {
			command_list.Run()
		},
	}
	return cmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = geoCmd()

func Execute(ver string) {
	// sets the version string for the root command.
	rootCmd.Version = ver
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
