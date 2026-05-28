// Package logger
package logger

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Save appends lines, deduplicates and sorts - use for subdomains and similar list output
func Save(host, category string, lines []string) error {
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

// SaveRaw writes lines as-is without sorting or deduplicating
func SaveRaw(host, category string, lines []string) error {
	filePath := filepath.Join("logs", host, category+".txt")
	os.MkdirAll(filepath.Dir(filePath), 0o755)
	return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}

// Exists checks if output for a tool already exists
func Exists(host, category string) bool {
	filePath := filepath.Join("logs", host, category+".txt")
	_, err := os.Stat(filePath)
	return err == nil
}
