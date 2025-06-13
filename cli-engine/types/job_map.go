package types

type JobMap struct {
	JobName string    `json:"jobName"`
	Steps   []JobStep `json:"steps"`
	Errors  []string  `json:"errors"`
}
