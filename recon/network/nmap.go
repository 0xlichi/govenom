// Package netscanning provides network scanning utilities
package netscanning

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

// Nmap runs a nmap scan on the given host and saves output to logs/<host>/network/nmap-result.txt
// Flags: -Pn (skip host discovery), -r (scan ports in order), --open (show only open ports), --reason (show why port is open)
func Nmap(host string, t *output.Tracker) error {
	if logger.Exists(host, "network/nmap-result") {
		t.Print(output.Warning("nmap: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("nmap", host, "-Pn", "-r", "--open", "--reason").Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("nmap failed: %v", err)))
		t.Print(output.Warning("Failed to run nmap, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "network/nmap-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save nmap output: %v", err)))
		return err
	}
	t.Print(output.Success("nmap done. Output saved to logs/" + host + "/network/nmap-result.txt"))
	return nil
}
