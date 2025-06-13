package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// FileMatches returns true if the file has a supported JCL-related extension.
func FileMatches(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".jcl" || ext == ".proc"
}

// FindJCLFiles walks the given root and returns a list of .jcl and .proc files.
func FindJCLFiles(root string) []string {
	var files []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Exit on any file access error
		}
		if info == nil || info.IsDir() {
			return nil
		}
		if FileMatches(path) {
			files = append(files, path)
		}
		return nil
	})

	return files
}
