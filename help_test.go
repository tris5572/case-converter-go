package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestHelpOutput(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	orig := os.Stdout
	os.Stdout = w

	// 呼び出し
	printHelp()

	// 後処理
	w.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = orig

	out := string(b)
	if !strings.Contains(out, "Usage") || !strings.Contains(out, "--camel") {
		t.Fatalf("help output missing expected text: %q", out)
	}
}
