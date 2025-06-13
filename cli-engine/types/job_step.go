package types

type JobStep struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	CalledProc string `json:"calledProc,omitempty"`
}
