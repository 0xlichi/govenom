// Package webrecon provides web reconnaissance utilities
package webrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Gospider(host string, t *output.Tracker) error {
	if logger.Exists(host, "web/gospider-result") {
		t.Print(output.Warning("gospider: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("gospider", "-s", "https://"+host, "-c", "10", "-d", "1", "-t", "2").Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("gospider failed: %v", err)))
		t.Print(output.Warning("Failed to run gospider, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "web/gospider-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save gospider output: %v", err)))
		return err
	}
	t.Print(output.Success("gospider done. Output saved to logs/" + host + "/web/gospider-result.txt"))
	return nil
}
