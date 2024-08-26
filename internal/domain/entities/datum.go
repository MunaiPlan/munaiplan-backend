package entities

import (
	"time"
)

// Datum entity representing the datums table.
type Datum struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	SystemDescription string    `json:"system_description"`
	SystemElevation   float64   `json:"system_elevation"`
	DatumDescription  string    `json:"datum_description"`
	WellheadElevation float64   `json:"wellhead_elevation"`
	DatumElevation    float64   `json:"datum_elevation"`
	AirGap            float64   `json:"air_gap"`
	GroundElevation   float64   `json:"ground_elevation"`
	Type              string    `json:"type"`
	CreatedAt         time.Time `json:"created_at"`
}
