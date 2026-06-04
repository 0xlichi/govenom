// Package dnsrecon provides DNS reconnaissance utilities
package dnsrecon

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Dnsx(host string, t *output.Tracker) error {
	if logger.Exists(host, "dns/dnsx-result") {
		t.Print(output.Warning("dnsx: output already exists, skipping."))
		return nil
	}
	sources := []string{
		"subdomains/subfinder-result",
		"subdomains/amass-result",
		"subdomains/assetfinder-result",
	}
	var allSubdomains []string
	for _, src := range sources {
		lines, err := logger.ReadLines(host, src)
		if err == nil {
			allSubdomains = append(allSubdomains, lines...)
		}
	}
	if len(allSubdomains) == 0 {
		t.Print(output.Warning("dnsx: no subdomains found from phase 1, skipping."))
		return nil
	}
	if err := logger.Save(host, "subdomains/merged-subdomains", allSubdomains); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save merged subdomains: %v", err)))
		return err
	}
	mergedPath := filepath.Join("logs", host, "subdomains", "merged-subdomains.txt")
	out, err := exec.Command("dnsx", "-l", mergedPath, "-silent").Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("dnsx failed: %v", err)))
		t.Print(output.Warning("Failed to run dnsx, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.Save(host, "dns/dnsx-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save dnsx output: %v", err)))
		return err
	}
	t.Print(output.Success("dnsx done. Output saved to logs/" + host + "/dns/dnsx-result.txt"))
	return nil
}
