package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter005SourceContract(t *testing.T) {
    source, err := os.ReadFile("websocket.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return true") {
        t.Fatalf("expected source contract is missing")
    }
}
