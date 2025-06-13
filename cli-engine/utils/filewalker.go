package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func FindJCLFiles(root string) []string {
	var files []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Exit if any access error happens
		}

		// Defensive: skip any nil info
		if info == nil {
			return nil
		}

		if !info.IsDir() && (strings.HasSuffix(path, ".jcl") || strings.HasSuffix(path, ".proc")) {
			files = append(files, path)
		}
		return nil
	})

	return files
}
