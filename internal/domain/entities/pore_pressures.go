package entities

import "time"

// PorePressure represents the Pore Pressure entity with fields for relevant data.
type PorePressure struct {
	ID        string    `json:"id"`
	CaseID    string    `json:"case_id"`
	TVD       float64   `json:"tvd"`
	Pressure  float64   `json:"pressure"`
	EMW       float64   `json:"emw"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
