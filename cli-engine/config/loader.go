package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadPatterns() (*Patterns, error) {
	path := filepath.Join("..", "config", "patterns.yml")

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p Patterns
	err = yaml.Unmarshal(file, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
