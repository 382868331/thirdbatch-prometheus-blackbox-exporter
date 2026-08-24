package config

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBlackboxExporter012SourceContract(t *testing.T) {
    source, err := os.ReadFile("config.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return encodings[j].quality < encodings[i].quality") {
        t.Fatalf("expected source contract is missing")
    }
}
