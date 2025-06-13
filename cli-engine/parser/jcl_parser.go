package parser

import (
	"bufio"
	"cli-engine/types"
	"os"
	"path/filepath"
	"strings"
)

func ParseJCLFile(path string) types.JobMap {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var steps []types.JobStep
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
			steps = append(steps, types.JobStep{Name: label, Type: "JOB"})
		case strings.Contains(keyword, "EXEC"):
			step := types.JobStep{Name: label, Type: "EXEC"}
			if len(parts) > 2 && !strings.HasPrefix(parts[2], "PGM=") {
				step.CalledProc = strings.Trim(parts[2], ",")
			}
			steps = append(steps, step)
		case strings.Contains(keyword, "DD"):
			steps = append(steps, types.JobStep{Name: label, Type: "DD"})
		}
	}

	return types.JobMap{
		JobName: jobName,
		Steps:   steps,
	}
}
