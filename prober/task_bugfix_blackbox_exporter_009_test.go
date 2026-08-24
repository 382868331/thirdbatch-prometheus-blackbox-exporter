package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter009SourceContract(t *testing.T) {
    source, err := os.ReadFile("http.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if success && (len(httpConfig.FailIfBodyMatchesRegexp) > 0 || len(httpConfig.FailIfBodyNotMatchesRegexp) > 0) {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if success && (len(httpConfig.FailIfBodyMatchesRegexp) >= 0 || len(httpConfig.FailIfBodyNotMatchesRegexp) > 0) {") {
        t.Fatalf("mutated source contract is still present")
    }
}
