// Package dnsrecon provides DNS reconnaissance utilities
package dnsrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Whois(host string, t *output.Tracker) error {
	if logger.Exists(host, "dns/whois-result") {
		t.Print(output.Warning("whois: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("whois", host).Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("whois failed: %v", err)))
		t.Print(output.Warning("Failed to run whois, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "dns/whois-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save whois output: %v", err)))
		return err
	}
	t.Print(output.Success("whois done. Output saved to logs/" + host + "/dns/whois-result.txt"))
	return nil
}
