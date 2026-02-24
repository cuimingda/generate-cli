package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gen",
		Short: "A brief description of your application",
		Long:  `...`,
	}

	rootCmd.AddCommand(NewPinCmd())
	rootCmd.AddCommand(NewSlugCmd())

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
