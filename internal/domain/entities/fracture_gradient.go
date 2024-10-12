package entities

import "time"

// FractureGradient represents the fracture gradient data in the domain layer.
type FractureGradient struct {
	ID                   string    `json:"id"`
	TemperatureAtSurface float64   `json:"temperature_at_surface"`
	TemperatureAtWellTVD float64   `json:"temperature_at_well_tvd"`
	TemperatureGradient  float64   `json:"temperature_gradient"`
	WellTVD              float64   `json:"well_tvd"`
	CreatedAt            time.Time `json:"created_at"`
}
