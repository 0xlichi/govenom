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
	out, err := exec.Command("nmap", host, "-Pn", "-r", "--open", "--reason", "-A").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("nmap failed: %v", err)))
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.SaveRaw(host, "nmap/nmap-results", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save nmap output: %v", err)))
		return err
	}

	fmt.Println(output.Success("nmap done. Output saved to logs/" + host + "/nmap/nmap-results.txt"))
	return nil
}
