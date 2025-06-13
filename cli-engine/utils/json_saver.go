package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"cli-engine/types"
)

// GenerateOutputPath builds the .json output path based on input file
func GenerateOutputPath(inputPath string) string {
	base := filepath.Base(inputPath)
	name := strings.TrimSuffix(base, filepath.Ext(base)) + ".json"
	return filepath.Join("..", "parsed-output", name)
}

// SerializeJobMap converts a JobMap to indented JSON
func SerializeJobMap(data types.JobMap) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

// SaveJSON writes a parsed JobMap as a JSON file to the parsed-output folder
func SaveJSON(data types.JobMap, inputPath string) {
	outPath := GenerateOutputPath(inputPath)

	jsonData, err := SerializeJobMap(data)
	if err != nil {
		panic("failed to serialize JSON: " + err.Error())
	}

	err = os.WriteFile(outPath, jsonData, 0644)
	if err != nil {
		panic("failed to write JSON to disk: " + err.Error())
	}
}
