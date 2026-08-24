package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter007SourceContract(t *testing.T) {
    source, err := os.ReadFile("history.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if r.Target == target && (module == \"\" || module == r.ModuleName) {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if r.Target == target || (module == \"\" || module == r.ModuleName) {") {
        t.Fatalf("mutated source contract is still present")
    }
}
