package config

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter018SourceContract(t *testing.T) {
    source, err := os.ReadFile("config.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer func() {") {
        t.Fatalf("expected source contract is missing")
    }
}
