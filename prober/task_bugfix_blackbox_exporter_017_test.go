package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter017SourceContract(t *testing.T) {
    source, err := os.ReadFile("crl.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(chain) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if len(chain) != 0 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
