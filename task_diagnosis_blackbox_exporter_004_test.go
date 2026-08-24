package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBlackboxExporter004SourceContract(t *testing.T) {
    source, err := os.ReadFile("main.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if err == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
