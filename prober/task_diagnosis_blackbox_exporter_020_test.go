package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBlackboxExporter020SourceContract(t *testing.T) {
    source, err := os.ReadFile("tls.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "cert := state.PeerCertificates[0]") {
        t.Fatalf("expected source contract is missing")
    }
}
