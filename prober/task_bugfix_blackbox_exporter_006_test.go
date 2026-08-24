package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter006SourceContract(t *testing.T) {
    source, err := os.ReadFile("grpc.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
