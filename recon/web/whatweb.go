// Package webrecon provides web reconnaissance utilities
package webrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func WhatWeb(host string, t *output.Tracker) error {
	if logger.Exists(host, "web/whatweb-result") {
		t.Print(output.Warning("whatweb: output already exists, skipping."))
		return nil
	}
	out, err := exec.Command("whatweb", "-q", host).Output()
	if err != nil {
		t.Print(output.Error(fmt.Sprintf("whatweb failed: %v", err)))
		t.Print(output.Warning("Failed to run whatweb, manual investigation needed."))
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "web/whatweb-result", lines); err != nil {
		t.Print(output.Error(fmt.Sprintf("Failed to save whatweb output: %v", err)))
		return err
	}
	t.Print(output.Success("whatweb done. Output saved to logs/" + host + "/web/whatweb-result.txt"))
	return nil
}
