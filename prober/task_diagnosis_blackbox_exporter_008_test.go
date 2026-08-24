package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBlackboxExporter008SourceContract(t *testing.T) {
    source, err := os.ReadFile("unix.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer conn.Close()") {
        t.Fatalf("expected source contract is missing")
    }
}
