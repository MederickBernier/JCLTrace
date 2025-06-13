package utils

import (
	"os"
	"path/filepath"
	"strings"

	"cli-engine/types"
)

// GenerateLogPath returns the full .log file path for a given input file.
func GenerateLogPath(inputPath string) string {
	base := filepath.Base(inputPath)
	name := strings.TrimSuffix(base, filepath.Ext(base)) + ".log"
	return filepath.Join("..", "parsed-output", "errors", name)
}

// FormatErrorLog turns a slice of error strings into a single newline-delimited string.
func FormatErrorLog(errors []string) string {
	return strings.Join(errors, "\n")
}

// LogErrors writes the error list of a JobMap to a .log file.
func LogErrors(job types.JobMap, inputPath string) {
	if len(job.Errors) == 0 {
		return
	}

	logPath := GenerateLogPath(inputPath)
	logContent := FormatErrorLog(job.Errors)

	if err := os.WriteFile(logPath, []byte(logContent), 0644); err != nil {
		panic("failed to write error log: " + err.Error())
	}
}
