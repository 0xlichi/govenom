// Package webrecon provides web reconnaissance utilities
package webrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Wafw00f(host string, t *output.Tracker) error {
	if logger.Exists(host, "web/wafw00f-result") {
		t.Print(output.Warning("wafw00f: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("wafw00f", "-a", "-v", host).Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("wafw00f failed: %v", err)))
		t.Print(output.Warning("Failed to run wafw00f, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "web/wafw00f-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save wafw00f output: %v", err)))
		return err
	}
	t.Print(output.Success("wafw00f done. Output saved to logs/" + host + "/web/wafw00f-result.txt"))
	return nil
}
