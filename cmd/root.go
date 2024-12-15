package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func geoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "geo",
	}
	return cmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = geoCmd()

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
