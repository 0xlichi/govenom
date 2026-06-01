// Package logger handles writing and deduplicating tool output
package logger

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var mu sync.Mutex

// Save appends lines, deduplicates and sorts - use for list output like subdomains
func Save(host, category string, lines []string) error {
	mu.Lock()
	defer mu.Unlock()

	filePath := filepath.Join("logs", host, category+".txt")
	os.MkdirAll(filepath.Dir(filePath), 0o755)

	existing, _ := os.ReadFile(filePath)
	all := append(strings.Split(string(existing), "\n"), lines...)

	seen := make(map[string]bool)
	var unique []string
	for _, line := range all {
		if line != "" && !seen[line] {
			seen[line] = true
			unique = append(unique, line)
		}
	}
	sort.Strings(unique)

	return os.WriteFile(filePath, []byte(strings.Join(unique, "\n")+"\n"), 0o644)
}

// SaveRaw writes lines as-is without sorting or deduplicating - use for structured output like nmap
func SaveRaw(host, category string, lines []string) error {
	filePath := filepath.Join("logs", host, category+".txt")
	os.MkdirAll(filepath.Dir(filePath), 0o755)
	return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}

// Exists checks if output for a tool already exists
func Exists(host, category string) bool {
	_, err := os.Stat(filepath.Join("logs", host, category+".txt"))
	return err == nil
}

// ReadLines reads lines from an existing log file - used by phase 2 tools
func ReadLines(host, category string) ([]string, error) {
	data, err := os.ReadFile(filepath.Join("logs", host, category+".txt"))
	if err != nil {
		return nil, err
	}
	var lines []string
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}
