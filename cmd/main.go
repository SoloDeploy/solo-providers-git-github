package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/SoloDeploy/solo-providers-git-github/cmd/start"
)

// NewCmd Go away linter
func NewCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:                   "git",
		DisableFlagsInUseLine: true,
		Short:                 "This is a SoloDeploy Git Provider implementation for interacting with GitHub",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	rootCmd.AddCommand(start.NewCmd())

	return rootCmd
}
