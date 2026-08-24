package prober

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBlackboxExporter001SourceContract(t *testing.T) {
    source, err := os.ReadFile("dns.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if module.DNS.TransportProtocol != \"udp\" && module.DNS.TransportProtocol != \"tcp\" {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if module.DNS.TransportProtocol == \"udp\" && module.DNS.TransportProtocol != \"tcp\" {") {
        t.Fatalf("mutated source contract is still present")
    }
}
