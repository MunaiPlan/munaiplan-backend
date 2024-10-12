package entities

import "time"

// PressureDataProfile represents a pressure data profile entity.
type PressureDataProfile struct {
	ID        string    `json:"id"`
	TVD       float64   `json:"tvd"`
	Pressure  float64   `json:"pressure"`
	EMW       float64   `json:"emw"`
	CreatedAt time.Time `json:"created_at"`
}
