package cmd

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestPinCmdDefaultCount(t *testing.T) {
	root := NewRootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"pin"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 10 {
		t.Fatalf("expected 10 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) != 6 {
			t.Fatalf("expected PIN length 6, got %d", len(strings.TrimSpace(line)))
		}
	}
}

func TestPinCmdCustomCountAndLength(t *testing.T) {
	root := NewRootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"pin", "--length", "4", "--count", "3"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) != 4 {
			t.Fatalf("expected PIN length 4, got %d", len(strings.TrimSpace(line)))
		}
	}
}

func TestPinCmdJSONOutput(t *testing.T) {
	root := NewRootCmd()
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"pin", "--length", "4", "--count", "3", "--json"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var payload []map[string]string
	if err := json.Unmarshal(buf.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid json output, got error: %v", err)
	}

	if len(payload) != 3 {
		t.Fatalf("expected 3 JSON items, got %d", len(payload))
	}

	for _, item := range payload {
		p, ok := item["pin"]
		if !ok {
			t.Fatalf("missing pin field in item: %#v", item)
		}
		if len(strings.TrimSpace(p)) != 4 {
			t.Fatalf("expected PIN length 4, got %d", len(strings.TrimSpace(p)))
		}
	}
}
