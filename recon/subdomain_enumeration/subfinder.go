// Package subenum provides subdomain enumeration utilities
package subenum

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Subfinder(host string, t *output.Tracker) error {
	if logger.Exists(host, "subdomains/subfinder-result") {
		t.Print(output.Warning("subfinder: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("subfinder", "-d", host, "-silent").Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("subfinder failed: %v", err)))
		t.Print(output.Warning("Failed to run subfinder, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.Save(host, "subdomains/subfinder-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save subfinder output: %v", err)))
		return err
	}
	t.Print(output.Success("subfinder done. Output saved to logs/" + host + "/subdomains/subfinder-result.txt"))
	return nil
}
