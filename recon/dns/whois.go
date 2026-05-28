// Package dnsrecon
package dnsrecon

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/0xlichi/govenom/logger"
	"github.com/0xlichi/govenom/output"
)

func Whois(host string) error {
	// Skip if output already exist
	if logger.Exists(host, "whois-result") {
		fmt.Println(output.Warning("whois: output already exist, skipping."))
		return nil
	}

	// Run whois and capture stdout
	out, err := exec.Command("whois", host, "--version").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("whois failed: %v", err)))
		fmt.Println(output.Warning("Failed to run whois, manaul investigation needed."))
		return err
	}

	// Split output into lines and save as-is
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "whois-result", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save whois output: %v", err)))
		return err
	}

	fmt.Println(output.Success("whois done. Output saved to logs/" + host + "whois-result.txt"))
	return nil
}
