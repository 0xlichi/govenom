// Package subenum
package subenum

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Subfinder(host string) error {
	// Skip if output already exists
	if logger.Exists(host, "subfinder-result") {
		fmt.Println(output.Warning("subfinder: output already exists, skipping."))
		return nil
	}

	// Run subfinder and capture stdout
	out, err := exec.Command("subfinder", "-d", host, "-silent").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("subfinder failed: %v", err)))
		fmt.Println(output.Warning("Failed to run subfinder, manual investigation needed."))
		return err
	}

	// Split output into lines and save as-is
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.Save(host, "subfinder-result", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save subfinder output: %v", err)))
		return err
	}

	fmt.Println(output.Success("subfinder done. Output saved to logs/" + host + "subfinder-result.txt"))
	return nil
}
