package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter013SourceContract(t *testing.T) {
    source, err := os.ReadFile("history.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "rh.preservedFailedResults = append(rh.preservedFailedResults, rh.results[0])") {
        t.Fatalf("expected source contract is missing")
    }
}
