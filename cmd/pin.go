package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

const pinDigits = "0123456789"

var pinLength int

var pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "Generate a random numeric PIN string",
	RunE: func(cmd *cobra.Command, args []string) error {
		if pinLength != 4 && pinLength != 6 && pinLength != 8 {
			return fmt.Errorf("参数不合法")
		}

		pin, err := generatePIN(pinLength)
		if err != nil {
			return err
		}

		cmd.Println(pin)
		return nil
	},
}

func init() {
	pinCmd.Flags().IntVar(&pinLength, "length", 6, "PIN length, supported values: 4/6/8")
	rootCmd.AddCommand(pinCmd)
}

func generatePIN(length int) (string, error) {
	bytes := make([]byte, length)
	max := big.NewInt(int64(len(pinDigits)))

	for i := range bytes {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		bytes[i] = pinDigits[n.Int64()]
	}

	return string(bytes), nil
}
