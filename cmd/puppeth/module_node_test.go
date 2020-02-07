package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"testing"
)

func TestParsingNodeTemplate(t *testing.T) {
	bootnodes := []string{
		"enode://dummyEnode",
		"enode://dummyEnode",
	}

	dockerfile := new(bytes.Buffer)
	template.Must(template.New("").Parse(nodeDockerfile)).Execute(dockerfile, map[string]interface{}{
		"NetworkID":   "48",
		"Port":        "30303",
		"IP":          "49.34.12.10",
		"Peers":       "10",
		"LightFlag":   "",
		"Bootnodes":   strings.Join(bootnodes, ","),
		"Ethstats":    "ethstatsDummyUrl",
		"Etherbase":   "etherbaseDummyUrl",
		"GasTarget":   uint64(1000000 * 10),
		"GasLimit":    uint64(1000000 * 10),
		"GasPrice":    uint64(1000000000 * 10),
		"Unlock":      true,
		"NodeKeyHex":  "dummyNodeKeyHex",
		"Testnet":     true,
		"DockerImage": "silesiacoin/infra-client:v0.09",
	})

	dockerFileString := strings.TrimSpace(fmt.Sprintf("%s", dockerfile))

	// This is only for concrete implementation. If you somehow want to change it do it.
	// I have not time right now to write better test.
	if len(dockerFileString) > 751 {
		t.Error(len(dockerFileString))
		t.Error("Template parsing has changed. Check if you really want to do it.")
	}
}
