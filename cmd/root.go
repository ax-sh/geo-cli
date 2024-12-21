package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func geoCmd() *cobra.Command {
	cmd := &cobra.Command{
		// Version field will be dynamically set
		// Version: "",
		Use: "geo",
	}
	return cmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = geoCmd()

// SetRootCmdVersion sets the version string for the root command.
func SetRootCmdVersion(ver string) {
	rootCmd.Version = ver
}
func Execute(v string) {
	SetRootCmdVersion(v)
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
