package cmd

import (
	"github.com/spf13/cobra"
)

func versionSubCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver", "v"},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("[Version]", cmd.Version)
		},
	}
	return cmd
}
func init() {
	rootCmd.AddCommand(versionSubCmd())
}
