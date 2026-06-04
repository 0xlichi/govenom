// Package dnsrecon provides DNS reconnaissance utilities
package dnsrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Dig(host string, t *output.Tracker) error {
	if logger.Exists(host, "dns/dig-result") {
		t.Print(output.Warning("dig: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("dig", host, "ANY", "+noall", "+answer").Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("dig failed: %v", err)))
		t.Print(output.Warning("Failed to run dig, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "dns/dig-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save dig output: %v", err)))
		return err
	}
	t.Print(output.Success("dig done. Output saved to logs/" + host + "/dns/dig-result.txt"))
	return nil
}
