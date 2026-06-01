// Package webrecon provides web reconnaissance utilities
package webrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

// Nikto runs a nikto web vulnerability scan on the given host and saves output to logs/<host>/web/nikto-result.txt
// Flags: -nointeractive (no prompts during scan)
func Nikto(host string) error {
	// Skip if output already exists from a previous run
	if logger.Exists(host, "web/nikto-result") {
		fmt.Println(output.Warning("nikto: output already exists, skipping."))
		return nil
	}

	// Run nikto and capture stdout
	out, err := exec.Command("nikto", "-h", host, "-nointeractive").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("nikto failed: %v", err)))
		fmt.Println(output.Warning("Failed to run nikto, manual investigation needed."))
		return err
	}

	// Split output into lines and save as-is (order matters for nikto output)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "web/nikto-result", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save nikto output: %v", err)))
		return err
	}

	fmt.Println(output.Success("nikto done. Output saved to logs/" + host + "/web/nikto-result.txt"))
	return nil
}
