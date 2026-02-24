package cmd

import (
	"encoding/json"

	"github.com/cuimingda/generate-cli/internal/pin"
	"github.com/spf13/cobra"
)

var pinLength int
var pinCount int
var pinJSON bool

func NewPinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pin",
		Short: "Generate a random numeric PIN string",
		RunE: func(cmd *cobra.Command, args []string) error {
			service := pin.NewService()

			result, err := service.Generate(pinLength, pinCount)
			if err != nil {
				return err
			}

			if pinJSON {
				payload := make([]map[string]string, 0, len(result))
				for _, item := range result {
					payload = append(payload, map[string]string{"pin": item})
				}

				encoder := json.NewEncoder(cmd.OutOrStdout())
				return encoder.Encode(payload)
			}

			for _, item := range result {
				cmd.Println(item)
			}
			return nil
		},
	}

	cmd.Flags().IntVar(&pinLength, "length", 6, "PIN length, supported values: 4/6/8")
	cmd.Flags().IntVar(&pinCount, "count", 10, "Number of PIN values to generate")
	cmd.Flags().BoolVar(&pinJSON, "json", false, "Output pins as JSON array")
	return cmd
}
