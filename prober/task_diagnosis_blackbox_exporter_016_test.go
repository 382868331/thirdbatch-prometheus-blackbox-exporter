package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBlackboxExporter016SourceContract(t *testing.T) {
    source, err := os.ReadFile("websocket.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(credentials) > 0 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if len(credentials) >= 0 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
