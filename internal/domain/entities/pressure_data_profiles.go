package entities

import "time"

// PressureDataProfile represents a pressure data profile entity.
type PressureDataProfile struct {
	ID        string    `json:"id"`
	CaseID    string    `json:"case_id"`
	TVD       float64   `json:"tvd"`
	Pressure  float64   `json:"pressure"`
	EMW       float64   `json:"emw"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
