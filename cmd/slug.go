package cmd

import (
	"github.com/cuimingda/generate-cli/internal/slug"
	"github.com/spf13/cobra"
)

var slugLength int
var slugCount int

func NewSlugCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slug",
		Short: "Generate a random lowercase alphanumeric string",
		RunE: func(cmd *cobra.Command, args []string) error {
			service := slug.NewService()

			result, err := service.Generate(slugLength, slugCount)
			if err != nil {
				return err
			}

			for _, item := range result {
				cmd.Println(item)
			}
			return nil
		},
	}

	cmd.Flags().IntVar(&slugLength, "length", slug.DefaultLen, "Slug length, supported range: 8-16")
	cmd.Flags().IntVar(&slugCount, "count", slug.DefaultCount, "Number of slug values to generate, supported range: 1-10")
	return cmd
}
