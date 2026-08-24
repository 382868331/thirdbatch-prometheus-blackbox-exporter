package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter011SourceContract(t *testing.T) {
    source, err := os.ReadFile("handler.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return 0, err") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return 1, err") {
        t.Fatalf("mutated source contract is still present")
    }
}
