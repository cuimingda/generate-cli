package cmd

import (
	"bytes"
	"testing"
)

func TestPinCommand(t *testing.T) {
	cmd := NewRootCmd() // 推荐 root 也用构造函数
	buf := new(bytes.Buffer)

	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"pin", "--length", "4"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(buf.String()) == 0 {
		t.Fatalf("expected output")
	}
}
