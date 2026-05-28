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
	if logger.Exists(host, "nmap/nmap-results") {
		fmt.Println(output.Warning("nmap: output already exists, skipping."))
		return nil
	}

	// Run subfinder and capture stdout
	out, err := exec.Command("subfinder", "-d", host, "-silent").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("subfinder failed: %v", err)))
		return err
	}

	// Split output into lines and save as-is
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.Save(host, "subdomains/unfiltered-subdomains", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save subfinder output: %v", err)))
		return err
	}

	fmt.Println(output.Success("subfinder done. Output saved to logs/" + host + "/subdomains/unfiltered-subdomains.txt"))
	return nil
}
