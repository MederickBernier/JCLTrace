package parser

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type JobStep struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	CalledProc string `json:"calledProc,omitempty"`
}

type JobMap struct {
	JobName string    `json:"jobName"`
	Steps   []JobStep `json:"steps"`
}

func ParseJCLFile(path string) JobMap {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var steps []JobStep
	jobName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "//") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		label := strings.TrimPrefix(parts[0], "//")
		keyword := parts[1]

		switch {
		case strings.Contains(keyword, "JOB"):
			jobName = label
			steps = append(steps, JobStep{Name: label, Type: "JOB"})
		case strings.Contains(keyword, "EXEC"):
			step := JobStep{Name: label, Type: "EXEC"}
			if len(parts) > 2 && !strings.HasPrefix(parts[2], "PGM=") {
				step.CalledProc = strings.Trim(parts[2], ",")
			}
			steps = append(steps, step)
		case strings.Contains(keyword, "DD"):
			steps = append(steps, JobStep{Name: label, Type: "DD"})
		}
	}

	return JobMap{JobName: jobName, Steps: steps}
}
