// Package webrecon
package webrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Gospider(host string) error {
	// Skip if output already exists
	if logger.Exists(host, "gospider-result") {
		fmt.Println(output.Warning("gospider: output already exists, skipping."))
		return nil
	}

	// Run gospider and capture stdout
	out, err := exec.Command("gospider", "-s", "https://"+host, "-c", "10", "-d", "1", "-t", "2").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("gospider failed: %v", err)))
		fmt.Println(output.Warning("Failed to run gospider, manual investigation needed."))
		return err
	}

	// Split output into lines and save as-is
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "gospider-result", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save gospider output: %v", err)))
		return err
	}

	fmt.Println(output.Success("gospider done. Output saved to logs/" + host + "gospider-result.txt"))
	return nil
}
