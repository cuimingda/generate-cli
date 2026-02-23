package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func walk(t *testing.T, root *cobra.Command, path []string, cmd *cobra.Command) {
	for _, c := range cmd.Commands() {
		currentPath := append(path, c.Name())

		t.Run(c.CommandPath(), func(t *testing.T) {
			buf := new(bytes.Buffer)
			newRoot := NewRootCmd()
			newRoot.SetOut(buf)
			newRoot.SetErr(buf)

			args := append(currentPath, "--help")
			newRoot.SetArgs(args)

			err := newRoot.Execute()
			if err != nil {
				t.Fatalf("failed: %v", err)
			}

			if buf.Len() == 0 {
				t.Fatalf("no help output")
			}
		})

		walk(t, root, currentPath, c)
	}
}

func TestAllCommands(t *testing.T) {
	root := NewRootCmd()
	walk(t, root, []string{}, root)
}
