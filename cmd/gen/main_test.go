package main

import "testing"

func TestMainEntrypointExists(t *testing.T) {
	entry := main
	if entry == nil {
		t.Fatal("main entrypoint should exist")
	}
}
