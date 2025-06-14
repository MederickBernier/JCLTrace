package filters

import (
	"cli-engine/types"
	"strings"
)

func ContainsKeyword(job types.JobMap, keyword string) bool {
	for _, step := range job.Steps {
		if strings.Contains(step.Name, keyword) || strings.Contains(step.CalledProc, keyword) {
			return true
		}
	}
	return false
}

func HasMatchingError(job types.JobMap, substr string) bool {
	for _, err := range job.Errors {
		if strings.Contains(err, substr) {
			return true
		}
	}
	return false
}

func CountProcs(job types.JobMap) int {
	count := 0
	for _, step := range job.Steps {
		if step.Type == "EXEC" && step.CalledProc != "" {
			count++
		}
	}
	return count
}
