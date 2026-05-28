// Package webrecon
package webrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Wafw00f(host string) error {
	// Skip if output already exists
	if logger.Exists(host, "wafw00f-result") {
		fmt.Println(output.Warning("wafw00f: output already exists, skipping."))
		return nil
	}

	// Run wafw00f and capture stdout
	out, err := exec.Command("wafw00f", "-a", "-v", host).Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("wafw00f failed: %v", err)))
		fmt.Println(output.Warning("Failed to run nmap, manual investigation needed."))
		return err
	}

	// Split output into lines and save as-is (order matters for nmap output)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "wafw00f-result", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save wafw00f output: %v", err)))
		return err
	}

	fmt.Println(output.Success("wafw00f done. Output saved  to logs/" + host + "wafw00f-result.txt"))
	return nil
}
