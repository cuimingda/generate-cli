package cmd

import (
	"github.com/cuimingda/generate-cli/internal/pin"
	"github.com/spf13/cobra"
)

var pinLength int

func NewPinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pin",
		Short: "Generate a random numeric PIN string",
		RunE: func(cmd *cobra.Command, args []string) error {
			service := pin.NewService()

			result, err := service.Generate(pinLength)
			if err != nil {
				return err
			}

			cmd.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&pinLength, "length", 6, "PIN length, supported values: 4/6/8")
	return cmd
}
