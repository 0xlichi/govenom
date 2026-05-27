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
	out, err := exec.Command("subfinder", "-d", host, "-silent").Output()
	if err != nil {
		fmt.Println(output.Error(fmt.Sprintf("subfinder failed: %v", err)))
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if err := logger.Save(host, "unfiltered-subdomains", lines); err != nil {
		fmt.Println(output.Error(fmt.Sprintf("Failed to save subfinder output: %v", err)))
		return err
	}

	fmt.Println(output.Success("subfinder done. Output saved to logs/" + host + "/unfiltered-subdomains.txt"))
	return nil
}
