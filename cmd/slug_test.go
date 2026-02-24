package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestSlugCmdDefault(t *testing.T) {
	root := NewRootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"slug"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(lines))
	}
	if len(strings.TrimSpace(lines[0])) != 12 {
		t.Fatalf("expected slug length 12, got %d", len(strings.TrimSpace(lines[0])))
	}
}

func TestSlugCmdCustomCountAndLength(t *testing.T) {
	root := NewRootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"slug", "--length", "8", "--count", "3"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) != 8 {
			t.Fatalf("expected slug length 8, got %d", len(strings.TrimSpace(line)))
		}
	}
}

func TestSlugCmdInvalidLength(t *testing.T) {
	root := NewRootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"slug", "--length", "7"})

	err := root.Execute()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestSlugCmdInvalidCount(t *testing.T) {
	root := NewRootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"slug", "--count", "0"})

	err := root.Execute()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
