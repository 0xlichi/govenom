// Package netscanning
package netscanning

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Nmap(host string) error {
	// Skip if output already exists
	if logger.Exists(host, "nmap/nmap-results") {
		fmt.Println(output.Warning("nmap: output already exists, skipping."))
		return nil
	}

	// Run nmap and capture stdout
	out, err := exec.Command("nmap", host, "-Pn", "-r", "--open", "--reason").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("nmap failed: %v", err)))
		return err
	}

	// Split output into lines and save as-is
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "nmap/nmap-results", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save nmap output: %v", err)))
		return err
	}

	fmt.Println(output.Success("nmap done. Output saved to logs/" + host + "/nmap/nmap-results.txt"))
	return nil
}
