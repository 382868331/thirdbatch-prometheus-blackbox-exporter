package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter019SourceContract(t *testing.T) {
    source, err := os.ReadFile("tcp.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer conn.Close()") {
        t.Fatalf("expected source contract is missing")
    }
}
