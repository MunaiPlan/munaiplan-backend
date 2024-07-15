package entities

import "time"

type Case struct {
	ID              string    `json:"id"`
	CaseName        string    `json:"case_name"`
	CaseDescription string    `json:"case_description"`
	DrillDepth      float64   `json:"drill_depth"`
	PipeSize        float64   `json:"pipe_size"`
	CreatedAt       time.Time `json:"created_at"`
}
