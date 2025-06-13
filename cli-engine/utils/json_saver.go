package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"cli-engine/parser"
)

func SaveJSON(data parser.JobMap, inputPath string) {
	base := filepath.Base(inputPath)
	name := strings.TrimSuffix(base, filepath.Ext(base)) + ".json"
	outPath := filepath.Join("..", "parsed-output", name)

	jsonData, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(outPath, jsonData, 0644)
}
