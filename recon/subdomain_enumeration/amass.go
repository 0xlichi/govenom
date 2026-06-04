// Package subenum provides subdomain enumeration utilities
package subenum

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Amass(host string, t *output.Tracker) error {
	if logger.Exists(host, "subdomains/amass-result") {
		t.Print(output.Warning("amass: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("amass", "enum", "-passive", "-d", host).Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("amass failed: %v", err)))
		t.Print(output.Warning("Failed to run amass, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.Save(host, "subdomains/amass-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save amass output: %v", err)))
		return err
	}
	t.Print(output.Success("amass done. Output saved to logs/" + host + "/subdomains/amass-result.txt"))
	return nil
}
