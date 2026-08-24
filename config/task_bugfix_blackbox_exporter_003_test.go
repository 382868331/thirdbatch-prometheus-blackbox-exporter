package config

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter003SourceContract(t *testing.T) {
    source, err := os.ReadFile("reload.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
