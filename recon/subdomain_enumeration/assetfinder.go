// Package subenum provides subdomain enumeration utilities
package subenum

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Assetfinder(host string, t *output.Tracker) error {
	if logger.Exists(host, "subdomains/assetfinder-result") {
		t.Print(output.Warning("assetfinder: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("assetfinder", "--subs-only", host).Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("assetfinder failed: %v", err)))
		t.Print(output.Warning("Failed to run assetfinder, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.Save(host, "subdomains/assetfinder-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save assetfinder output: %v", err)))
		return err
	}
	t.Print(output.Success("assetfinder done. Output saved to logs/" + host + "/subdomains/assetfinder-result.txt"))
	return nil
}
