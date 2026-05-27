// Package logger
package logger

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Save(host, category string, lines []string) error {
	filePath := filepath.Join("logs", host, category+".txt")
	os.MkdirAll(filepath.Dir(filePath), 0o755)

	// Read existing content
	existing, _ := os.ReadFile(filePath)

	// Merge old + new lines
	all := append(strings.Split(string(existing), "\n"), lines...)

	// Deduplicate and sort
	seen := make(map[string]bool)
	var unique []string
	for _, line := range all {
		if line != "" && !seen[line] {
			seen[line] = true
			unique = append(unique, line)
		}
	}
	sort.Strings(unique)

	// Write back
	return os.WriteFile(filePath, []byte(strings.Join(unique, "\n")+"\n"), 0o644)
}
